# 模型微调

## 概述

模型微调功能提供大语言模型微调任务的全生命周期管理，支持 LLaMA、ChatGLM、Qwen 等主流开源模型。通过可视化界面创建训练任务、实时监控训练进度、管理 GPU 资源，实现高效的模型定制化训练。

## 功能特性

### 核心功能

- ✅ **多模型支持**：支持 LLaMA、ChatGLM、Qwen 等开源模型
- ✅ **任务管理**：创建、删除、停止微调任务
- ✅ **实时监控**：查看训练进度和日志输出
- ✅ **GPU 管理**：配置 GPU 设备和显存使用
- ✅ **参数配置**：灵活配置训练参数
- ✅ **异步执行**：后台异步执行训练任务

## 数据模型

### 微调任务 (FinetuningTask)

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| 任务名称 | string | ✅ | 任务显示名称 |
| 任务描述 | string | | 任务详细描述 |
| 所属用户 | int | | 创建用户ID |
| 任务状态 | string | ✅ | pending/running/completed/failed/stopped |
| 任务进度 | float64 | | 0-100 |
| 基础模型 | string | ✅ | 基础模型路径或名称 |
| 数据集路径 | string | ✅ | 训练数据集路径 |
| 输出路径 | string | | 输出模型保存路径 |
| 训练参数 | json | | 训练超参数配置 |
| GPU配置 | json | | GPU 设备配置 |
| 执行命令 | string | | 完整的执行命令 |
| 日志文件路径 | string | | 训练日志文件路径 |
| 错误信息 | text | | 失败时的错误信息 |
| 开始时间 | int64 | | 任务开始时间戳 |
| 结束时间 | int64 | | 任务结束时间戳 |
| 进程ID | int | | 训练进程ID |
| 训练指标 | json | | loss、accuracy 等指标 |

## 使用指南

### 创建微调任务

1. 进入【算法微调】→【微调任务】页面
2. 点击"新建任务"按钮
3. 填写任务信息：

#### 基本信息区
- **任务名称**：任务的唯一标识名称
- **任务描述**：任务的详细说明（可选）

#### 模型配置区
- **基础模型**：选择或输入基础模型路径
  - 支持的模型：
    - LLaMA 2/3 系列
    - ChatGLM2/3 系列
    - Qwen1.5/2 系列
- **数据集路径**：训练数据集所在的目录路径
- **输出路径**：微调后模型的保存目录

#### 训练参数区
点击"展开高级参数"配置详细训练参数：

| 参数 | 说明 | 默认值 |
|------|------|--------|
| num_train_epochs | 训练轮数 | 3 |
| per_device_train_batch_size | 每设备批大小 | 4 |
| gradient_accumulation_steps | 梯度累积步数 | 1 |
| learning_rate | 学习率 | 2e-4 |
| warmup_steps | 预热步数 | 100 |
| logging_steps | 日志记录步数 | 10 |
| save_steps | 模型保存步数 | 500 |
| max_seq_length | 最大序列长度 | 512 |

#### GPU 配置区
- **CUDA 设备**：指定使用的 GPU（如 "0,1" 或 "0"）
- **显存优化**：启用显存优化技术

4. 点击"确定"创建任务
5. 任务将自动开始执行

### 监控训练进度

**查看任务列表：**
- 任务列表显示所有微调任务
- 状态标识：
  - 🟡 pending - 待执行
  - 🟢 running - 执行中
  - ✅ completed - 已完成
  - ❌ failed - 失败
  - ⏹️ stopped - 已停止

**查看实时进度：**
- 点击任务名称进入详情页
- 查看当前进度百分比
- 查看已用时间和预计剩余时间

**查看训练日志：**
1. 在任务详情页点击"查看日志"
2. 实时显示训练输出
3. 支持日志刷新和滚动

### 停止微调任务

1. 在任务列表中找到正在运行的任务
2. 点击"停止"按钮
3. 确认停止操作
4. 系统会优雅地终止训练进程

**注意**：停止后的任务无法恢复，如需继续请重新创建任务。

### 删除微调任务

1. 在任务列表中找到要删除的任务
2. 点击"删除"按钮
3. 确认删除操作

**注意**：
- 删除任务会同时删除训练日志文件
- 已输出的模型文件不会被删除
- 运行中的任务会先被停止

## 支持的模型

### LLaMA 系列

| 模型 | 参数量 | 显存需求 | 特点 |
|------|--------|----------|------|
| LLaMA 2-7B | 7B | ~16GB | 基础版本 |
| LLaMA 2-13B | 13B | ~24GB | 中等规模 |
| LLaMA 2-70B | 70B | ~140GB | 大规模 |
| LLaMA 3-8B | 8B | ~16GB | 新一代 |
| LLaMA 3-70B | 70B | ~140GB | 新一代大模型 |

### ChatGLM 系列

| 模型 | 参数量 | 显存需求 | 特点 |
|------|--------|----------|------|
| ChatGLM2-6B | 6B | ~14GB | 中文优化 |
| ChatGLM3-6B | 6B | ~14GB | 功能增强 |
| GLM-4-9B | 9B | ~20GB | 最新版本 |

### Qwen 系列

| 模型 | 参数量 | 显存需求 | 特点 |
|------|--------|----------|------|
| Qwen1.5-7B | 7B | ~16GB | 多语言 |
| Qwen1.5-14B | 14B | ~24GB | 中等规模 |
| Qwen2-7B | 7B | ~16GB | 第二代 |
| Qwen2-72B | 72B | ~144GB | 大规模 |

## 训练参数说明

### 基础参数

| 参数 | 说明 | 推荐值 |
|------|------|--------|
| learning_rate | 学习率 | 2e-4 ~ 5e-5 |
| num_train_epochs | 训练轮数 | 3 ~ 10 |
| batch_size | 批大小 | 1 ~ 8 |
| max_seq_length | 序列长度 | 256 ~ 2048 |

### 高级参数

| 参数 | 说明 | 推荐值 |
|------|------|--------|
| warmup_ratio | 预热比例 | 0.03 ~ 0.1 |
| weight_decay | 权重衰减 | 0.01 |
| gradient_checkpointing | 梯度检查点 | true |
| lora_r | LoRA 秩 | 8 ~ 64 |
| lora_alpha | LoRA alpha | 16 ~ 128 |
| lora_dropout | LoRA dropout | 0.05 ~ 0.1 |

## 数据集准备

### 数据格式

支持 JSONL 格式的训练数据：

```jsonl
{"instruction": "问题内容", "input": "", "output": "回答内容"}
{"instruction": "任务描述", "input": "额外信息", "output": "期望输出"}
```

### 数据集目录结构

```
dataset/
├── train.jsonl      # 训练集
├── valid.jsonl      # 验证集（可选）
└── test.jsonl       # 测试集（可选）
```

### 数据质量要求

1. **数据量**：建议至少 1000 条样本
2. **数据质量**：确保输入输出准确无误
3. **数据多样性**：覆盖不同场景和问题类型
4. **数据去重**：去除重复样本
5. **数据清洗**：过滤低质量数据

## GPU 资源管理

### 单卡训练

适用于小模型（7B 及以下）：

```bash
CUDA_VISIBLE_DEVICES=0 python train.py ...
```

### 多卡训练

适用于大模型（13B 及以上）：

```bash
CUDA_VISIBLE_DEVICES=0,1,2,3 python train.py ...
```

### 显存优化

当显存不足时，可以：
1. 减小 batch_size
2. 启用 gradient_checkpointing
3. 使用量化技术（4bit/8bit）
4. 使用 LoRA 等参数高效微调方法

## 常见问题

### Q1: 如何选择合适的微调方法？

A:
- **全量微调**：适用于有充足资源，追求最佳效果
- **LoRA 微调**：推荐方法，资源需求低，效果好
- **QLoRA 微调**：显存受限时使用，4bit 量化

### Q2: 训练显存不足怎么办？

A:
1. 减小 `per_device_train_batch_size`
2. 增加 `gradient_accumulation_steps`
3. 启用 `gradient_checkpointing`
4. 使用 LoRA 或 QLoRA
5. 使用多卡并行训练

### Q3: 如何判断训练是否收敛？

A:
1. 观察 loss 曲线是否下降并趋于稳定
2. 查看验证集性能指标
3. 检查是否发生过拟合
4. 合理设置早停策略

### Q4: 微调后的模型如何部署？

A:
1. 训练完成后，模型保存在输出目录
2. 可以直接加载使用
3. 或转换为其他格式（如 GGUF）部署

### Q5: 任务失败如何排查？

A:
1. 查看任务日志了解错误信息
2. 检查数据集格式是否正确
3. 验证模型路径是否存在
4. 确认 GPU 资源是否充足
5. 检查训练参数配置

## API 接口

### 任务管理

```
POST /finetuning/createTask       # 创建微调任务
DELETE /finetuning/deleteTask     # 删除微调任务
POST /finetuning/stopTask         # 停止微调任务
GET /finetuning/getTask           # 获取任务详情
GET /finetuning/getTaskList       # 获取任务列表
```

### 日志接口

```
GET /finetuning/getTaskLog        # 获取任务日志
```

### 请求示例

```json
{
  "name": "my_finetuning_task",
  "description": "ChatGLM3 医疗问答微调",
  "baseModel": "THUDM/chatglm3-6b",
  "datasetPath": "/data/medicalqa",
  "outputPath": "/output/medical_chatglm3",
  "trainingArgs": {
    "num_train_epochs": 5,
    "per_device_train_batch_size": 4,
    "gradient_accumulation_steps": 4,
    "learning_rate": 2e-4,
    "max_seq_length": 512
  },
  "gpuConfig": {
    "cuda_visible_devices": "0,1"
  }
}
```

## 最佳实践

### 数据准备

1. **数据质量优先于数量**
   - 确保 1000+ 高质量样本
   - 去除重复和错误数据

2. **数据集划分**
   - 训练集：80%
   - 验证集：15%
   - 测试集：5%

### 参数调优

1. **学习率**
   - 小模型：2e-4
   - 大模型：5e-5

2. **Batch Size**
   - 根据显存调整
   - 配合梯度累积使用

3. **训练轮数**
   - 简单任务：3-5 轮
   - 复杂任务：5-10 轮

### 资源优化

1. **使用 LoRA**
   - 显著降低资源需求
   - 保持良好的微调效果

2. **多卡并行**
   - 提高训练速度
   - 支持更大模型

### 效果评估

1. **定量评估**
   - 使用验证集计算指标
   - 与基线模型对比

2. **定性评估**
   - 人工测试生成效果
   - 检查回答质量

## 相关文档

- [AI 智能助手](AI-Agent.md)
- [算力节点管理](Compute-Nodes.md)

---

**最后更新**：2025-01-22
