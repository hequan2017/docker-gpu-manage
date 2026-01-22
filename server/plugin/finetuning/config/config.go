package config

// TrainingConfig 训练配置
type TrainingConfig struct {
	// DefaultPythonCommand 默认Python命令
	DefaultPythonCommand string
	// DefaultTrainScript 默认训练脚本
	DefaultTrainScript string
	// DefaultLogDir 默认日志目录
	DefaultLogDir string
	// DefaultOutputDir 默认输出目录
	DefaultOutputDir string
	// ProcessStopTimeout 进程停止超时时间（秒）
	ProcessStopTimeout int
	// LogAutoRefreshInterval 日志自动刷新间隔（秒）
	LogAutoRefreshInterval int
}

// DefaultConfig 默认配置
var DefaultConfig = TrainingConfig{
	DefaultPythonCommand:  "python",
	DefaultTrainScript:    "train.py",
	DefaultLogDir:         "finetuning_logs",
	DefaultOutputDir:      "finetuning_outputs",
	ProcessStopTimeout:    10,
	LogAutoRefreshInterval: 5,
}

// CommandTemplate 命令模板
type CommandTemplate struct {
	Name        string
	Description string
	Template    string
	// RequiredParams 必需参数
	RequiredParams []string
	// OptionalParams 可选参数
	OptionalParams map[string]string
}

// BuiltInTemplates 内置命令模板
var BuiltInTemplates = map[string]CommandTemplate{
	"python_train": {
		Name:        "Python训练脚本",
		Description: "使用Python训练脚本进行微调",
		Template:    "{{.PythonCmd}} {{.Script}} --base_model {{.BaseModel}} --data_path {{.DataPath}} {{if .OutputDir}}--output_dir {{.OutputDir}}{{end}} {{.TrainingArgs}}",
		RequiredParams: []string{"base_model", "data_path"},
		OptionalParams: map[string]string{
			"output_dir":     "--output_dir",
			"learning_rate":  "--learning_rate",
			"batch_size":     "--batch_size",
			"num_epochs":     "--num_train_epochs",
			"max_steps":      "--max_steps",
			"warmup_steps":   "--warmup_steps",
			"logging_steps":  "--logging_steps",
			"save_steps":     "--save_steps",
		},
	},
	"bash_script": {
		Name:        "Bash脚本",
		Description: "使用Bash脚本执行训练",
		Template:    "bash {{.Script}} {{.Args}}",
		RequiredParams: []string{"script"},
		OptionalParams: map[string]string{},
	},
	"docker_train": {
		Name:        "Docker容器训练",
		Description: "在Docker容器中执行训练",
		Template:    "docker run --gpus {{.GPUs}} -v {{.DataDir}}:/data -v {{.OutputDir}}:/output {{.Image}} {{.Command}}",
		RequiredParams: []string{"image", "command"},
		OptionalParams: map[string]string{
			"gpus":      "all",
			"data_dir":  "/data",
			"output_dir": "/output",
		},
	},
}

// StatusLabels 状态标签映射
var StatusLabels = map[string]string{
	"pending":   "待执行",
	"running":   "执行中",
	"completed": "已完成",
	"failed":    "失败",
	"stopped":   "已停止",
}

// StatusTypes 状态类型映射（用于前端）
var StatusTypes = map[string]string{
	"pending":   "info",
	"running":   "primary",
	"completed": "success",
	"failed":    "danger",
	"stopped":   "warning",
}

// TrainingPresets 训练预设配置
var TrainingPresets = map[string]map[string]interface{}{
	"quick_test": {
		"learning_rate":     0.0001,
		"batch_size":        16,
		"num_train_epochs":  1,
		"max_steps":         100,
		"warmup_steps":      10,
		"logging_steps":     10,
		"save_steps":        50,
	},
	"standard": {
		"learning_rate":     0.0002,
		"batch_size":        32,
		"num_train_epochs":  3,
		"max_steps":         -1,
		"warmup_steps":      100,
		"logging_steps":     50,
		"save_steps":        500,
	},
	"full_finetune": {
		"learning_rate":     0.00005,
		"batch_size":        64,
		"num_train_epochs":  10,
		"max_steps":         -1,
		"warmup_steps":      500,
		"logging_steps":     100,
		"save_steps":        1000,
	},
}
