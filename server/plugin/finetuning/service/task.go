package service

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	finetuningModel "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model"
	finetuningRequest "github.com/flipped-aurora/gin-vue-admin/server/plugin/finetuning/model/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type FinetuningTaskService struct{}

// CreateFinetuningTask 创建微调任务
func (s *FinetuningTaskService) CreateFinetuningTask(ctx context.Context, task *finetuningModel.FinetuningTask) (err error) {
	// 设置初始状态
	task.Status = finetuningModel.TaskStatusPending
	task.Progress = new(float64)
	*task.Progress = 0

	// 生成日志文件路径
	logDir := filepath.Join(global.GVA_CONFIG.Server.SavePath, "finetuning_logs")
	if err = os.MkdirAll(logDir, 0755); err != nil {
		return errors.Wrap(err, "创建日志目录失败")
	}
	logFileName := fmt.Sprintf("task_%d_%s.log", time.Now().Unix(), task.Name)
	logPath := filepath.Join(logDir, logFileName)
	task.LogPath = &logPath

	// 构建执行命令
	if task.Command == nil || *task.Command == "" {
		command, err := s.buildCommand(task)
		if err != nil {
			return errors.Wrap(err, "构建命令失败")
		}
		task.Command = &command
	}

	// 保存任务到数据库
	if err = global.GVA_DB.Create(task).Error; err != nil {
		return errors.Wrap(err, "创建任务失败")
	}

	// 异步执行任务
	go s.executeTask(task.ID)

	return nil
}

// GetFinetuningTaskList 获取任务列表
func (s *FinetuningTaskService) GetFinetuningTaskList(info *finetuningRequest.FinetuningTaskSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 构建查询
	db := global.GVA_DB.Model(&finetuningModel.FinetuningTask{})

	// 添加搜索条件
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if info.BaseModel != "" {
		db = db.Where("base_model LIKE ?", "%"+info.BaseModel+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 排序
	orderKey := info.OrderKey
	if orderKey == "" {
		orderKey = "created_at"
	}
	if info.Desc {
		orderKey += " DESC"
	}

	// 分页查询
	var taskList []finetuningModel.FinetuningTask
	err = db.Limit(limit).Offset(offset).Order(orderKey).Find(&taskList).Error
	return taskList, total, err
}

// GetFinetuningTaskById 获取任务详情
func (s *FinetuningTaskService) GetFinetuningTaskById(id uint) (task finetuningModel.FinetuningTask, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&task).Error
	return
}

// UpdateFinetuningTask 更新任务
func (s *FinetuningTaskService) UpdateFinetuningTask(task *finetuningModel.FinetuningTask) (err error) {
	err = global.GVA_DB.Save(task).Error
	return err
}

// DeleteFinetuningTask 删除任务
func (s *FinetuningTaskService) DeleteFinetuningTask(id uint) (err error) {
	// 先获取任务信息
	task, err := s.GetFinetuningTaskById(id)
	if err != nil {
		return errors.Wrap(err, "获取任务信息失败")
	}

	// 如果任务正在运行，先停止
	if task.Status == finetuningModel.TaskStatusRunning {
		if task.Pid != nil {
			if err = s.killProcess(*task.Pid); err != nil {
				global.GVA_LOG.Warn("停止任务进程失败", zap.Uint("task_id", id), zap.Error(err))
			}
		}
	}

	// 删除日志文件
	if task.LogPath != nil {
		if err = os.Remove(*task.LogPath); err != nil && !os.IsNotExist(err) {
			global.GVA_LOG.Warn("删除日志文件失败", zap.String("path", *task.LogPath), zap.Error(err))
		}
	}

	// 删除数据库记录
	err = global.GVA_DB.Delete(&finetuningModel.FinetuningTask{}, id).Error
	return err
}

// StopFinetuningTask 停止任务
func (s *FinetuningTaskService) StopFinetuningTask(id uint) (err error) {
	// 获取任务信息
	task, err := s.GetFinetuningTaskById(id)
	if err != nil {
		return errors.Wrap(err, "获取任务信息失败")
	}

	// 检查任务状态
	if task.Status != finetuningModel.TaskStatusRunning {
		return errors.New("任务未在运行中")
	}

	// 停止进程
	if task.Pid != nil {
		if err = s.killProcess(*task.Pid); err != nil {
			return errors.Wrap(err, "停止进程失败")
		}
	}

	// 更新任务状态
	now := int64(time.Now().Unix())
	task.Status = finetuningModel.TaskStatusStopped
	task.FinishedAt = &now
	err = s.UpdateFinetuningTask(&task)

	return err
}

// GetTaskLog 获取任务日志
func (s *FinetuningTaskService) GetTaskLog(id uint, lines, offset *int) (logContent string, err error) {
	// 获取任务信息
	task, err := s.GetFinetuningTaskById(id)
	if err != nil {
		return "", errors.Wrap(err, "获取任务信息失败")
	}

	// 检查日志文件
	if task.LogPath == nil {
		return "", errors.New("任务没有日志文件")
	}

	// 读取日志文件
	logContent, err = s.readLogFile(*task.LogPath, lines, offset)
	if err != nil {
		return "", errors.Wrap(err, "读取日志文件失败")
	}

	return logContent, nil
}

// executeTask 执行微调任务（异步）
func (s *FinetuningTaskService) executeTask(id uint) {
	// 获取任务信息
	task, err := s.GetFinetuningTaskById(id)
	if err != nil {
		global.GVA_LOG.Error("获取任务信息失败", zap.Uint("task_id", id), zap.Error(err))
		return
	}

	// 更新任务状态为运行中
	task.Status = finetuningModel.TaskStatusRunning
	now := int64(time.Now().Unix())
	task.StartedAt = &now
	if err = s.UpdateFinetuningTask(&task); err != nil {
		global.GVA_LOG.Error("更新任务状态失败", zap.Uint("task_id", id), zap.Error(err))
		return
	}

	// 打开日志文件
	logFile, err := os.OpenFile(*task.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		global.GVA_LOG.Error("打开日志文件失败", zap.String("path", *task.LogPath), zap.Error(err))
		s.updateTaskStatus(id, finetuningModel.TaskStatusFailed, err.Error())
		return
	}
	defer logFile.Close()

	// 解析命令
	cmdParts := strings.Fields(*task.Command)
	if len(cmdParts) == 0 {
		err = errors.New("命令为空")
		global.GVA_LOG.Error("命令解析失败", zap.Uint("task_id", id), zap.Error(err))
		s.updateTaskStatus(id, finetuningModel.TaskStatusFailed, err.Error())
		return
	}

	// 创建命令
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	// 设置环境变量
	if task.GPUConfig != nil {
		var gpuConfig map[string]interface{}
		if err = json.Unmarshal([]byte(*task.GPUConfig), &gpuConfig); err == nil {
			if cudaVisibleDevices, ok := gpuConfig["cuda_visible_devices"]; ok {
				cmd.Env = append(os.Environ(), fmt.Sprintf("CUDA_VISIBLE_DEVICES=%v", cudaVisibleDevices))
			}
		}
	}

	// 启动命令
	if err = cmd.Start(); err != nil {
		global.GVA_LOG.Error("启动命令失败", zap.Uint("task_id", id), zap.Error(err))
		s.updateTaskStatus(id, finetuningModel.TaskStatusFailed, err.Error())
		return
	}

	// 保存进程ID
	pid := cmd.Process.Pid
	task.Pid = &pid
	if err = s.UpdateFinetuningTask(&task); err != nil {
		global.GVA_LOG.Error("保存进程ID失败", zap.Uint("task_id", id), zap.Error(err))
	}

	// 等待命令完成
	err = cmd.Wait()

	// 更新任务状态
	finishedAt := int64(time.Now().Unix())
	task.FinishedAt = &finishedAt

	if err != nil {
		global.GVA_LOG.Error("任务执行失败", zap.Uint("task_id", id), zap.Error(err))
		errorMsg := err.Error()
		task.ErrorMessage = &errorMsg
		task.Status = finetuningModel.TaskStatusFailed
	} else {
		task.Status = finetuningModel.TaskStatusCompleted
		progress := 100.0
		task.Progress = &progress
	}

	if err = s.UpdateFinetuningTask(&task); err != nil {
		global.GVA_LOG.Error("更新任务状态失败", zap.Uint("task_id", id), zap.Error(err))
	}
}

// buildCommand 构建执行命令
func (s *FinetuningTaskService) buildCommand(task *finetuningModel.FinetuningTask) (string, error) {
	// 基础命令模板
	baseCmd := "python train.py"

	// 构建参数
	args := []string{
		fmt.Sprintf("--base_model=%s", task.BaseModel),
		fmt.Sprintf("--data_path=%s", task.DatasetPath),
	}

	// 添加输出路径
	if task.OutputPath != nil {
		args = append(args, fmt.Sprintf("--output_dir=%s", *task.OutputPath))
	}

	// 解析训练参数
	if task.TrainingArgs != nil {
		var trainingArgs map[string]interface{}
		if err := json.Unmarshal([]byte(*task.TrainingArgs), &trainingArgs); err == nil {
			for key, value := range trainingArgs {
				args = append(args, fmt.Sprintf("--%s=%v", key, value))
			}
		}
	}

	// 组合命令
	return fmt.Sprintf("%s %s", baseCmd, strings.Join(args, " ")), nil
}

// readLogFile 读取日志文件
func (s *FinetuningTaskService) readLogFile(logPath string, lines, offset *int) (string, error) {
	file, err := os.Open(logPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	lineCount := 0
	skipLines := 0

	if offset != nil {
		skipLines = *offset
	}

	for scanner.Scan() {
		if lineCount < skipLines {
			lineCount++
			continue
		}
		result = append(result, scanner.Text())
		lineCount++
	}

	// 如果指定了行数，从末尾截取
	if lines != nil && *lines > 0 && len(result) > *lines {
		result = result[len(result)-*lines:]
	}

	return strings.Join(result, "\n"), nil
}

// killProcess 杀死进程
func (s *FinetuningTaskService) killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	// 先尝试优雅退出（SIGTERM）
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		// 如果失败，强制杀死（SIGKILL）
		return process.Signal(syscall.SIGKILL)
	}

	// 等待进程退出
	done := make(chan error, 1)
	go func() {
		_, err := process.Wait()
		done <- err
	}()

	select {
	case <-done:
		return nil
	case <-time.After(10 * time.Second):
		// 超时后强制杀死
		return process.Signal(syscall.SIGKILL)
	}
}

// updateTaskStatus 更新任务状态
func (s *FinetuningTaskService) updateTaskStatus(id uint, status string, errorMsg string) {
	task, err := s.GetFinetuningTaskById(id)
	if err != nil {
		global.GVA_LOG.Error("获取任务信息失败", zap.Uint("task_id", id), zap.Error(err))
		return
	}

	task.Status = status
	task.ErrorMessage = &errorMsg
	finishedAt := int64(time.Now().Unix())
	task.FinishedAt = &finishedAt

	if err = s.UpdateFinetuningTask(&task); err != nil {
		global.GVA_LOG.Error("更新任务状态失败", zap.Uint("task_id", id), zap.Error(err))
	}
}
