-- =====================================================
-- 算法微调插件 - 初始化 SQL 脚本
-- =====================================================
-- 执行说明：
-- 1. 本脚本根据 gin-vue-admin 实际表结构调整
-- 2. 包含数据表、菜单、API、权限的完整初始化
-- 3. 支持 LLaMA、ChatGLM 等大语言模型微调
-- =====================================================

-- =====================================================
-- 第1步：创建数据表
-- =====================================================

-- 1.1 创建微调任务表
CREATE TABLE IF NOT EXISTS `gva_finetuning_tasks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(200) NOT NULL COMMENT '任务名称',
  `description` text COMMENT '任务描述',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '所属用户ID',
  `status` varchar(20) NOT NULL DEFAULT 'pending' COMMENT '任务状态: pending, running, completed, failed, stopped',
  `progress` double DEFAULT 0 COMMENT '任务进度 0-100',
  `base_model` varchar(200) NOT NULL COMMENT '基础模型路径或名称',
  `dataset_path` varchar(500) NOT NULL COMMENT '数据集路径',
  `output_path` varchar(500) DEFAULT NULL COMMENT '输出模型路径',
  `training_args` json COMMENT '训练参数JSON配置',
  `gpu_config` json COMMENT 'GPU配置JSON配置',
  `command` text COMMENT '执行的完整命令',
  `log_path` varchar(500) DEFAULT NULL COMMENT '日志文件路径',
  `error_message` text COMMENT '错误信息',
  `started_at` bigint DEFAULT NULL COMMENT '开始时间戳',
  `finished_at` bigint DEFAULT NULL COMMENT '结束时间戳',
  `pid` int DEFAULT NULL COMMENT '进程ID',
  `metrics` json COMMENT '训练指标JSON配置',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='算法微调任务表';

-- =====================================================
-- 第2步：创建菜单
-- =====================================================

-- 2.1 创建算法微调主菜单（作为系统功能的子菜单）
-- 注意：如果系统功能的ID不是24，请根据实际情况调整
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), 24, 'finetuning', 'finetuning', 0, 'view/routerHolder.vue', 7, 0, '算法微调', 'cpu');

-- 获取主菜单ID
SET @finetuning_menu_id = LAST_INSERT_ID();

-- 2.2 创建子菜单
-- 微调任务列表
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @finetuning_menu_id, 'taskList', 'finetuningTaskList', 0, 'plugin/finetuning/view/taskList.vue', 1, 0, '微调任务', 'list');

SET @task_list_menu_id = LAST_INSERT_ID();

-- 任务详情（隐藏菜单）
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @finetuning_menu_id, 'taskDetail', 'finetuningTaskDetail', 1, 'plugin/finetuning/view/taskDetail.vue', 2, 0, '任务详情', 'document');

SET @task_detail_menu_id = LAST_INSERT_ID();

-- =====================================================
-- 第3步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
-- 任务管理API
(NOW(), NOW(), '/finetuning/createTask', '创建微调任务', 'Finetuning', 'POST'),
(NOW(), NOW(), '/finetuning/deleteTask', '删除微调任务', 'Finetuning', 'DELETE'),
(NOW(), NOW(), '/finetuning/stopTask', '停止微调任务', 'Finetuning', 'POST'),
(NOW(), NOW(), '/finetuning/getTask', '根据ID获取微调任务', 'Finetuning', 'GET'),
(NOW(), NOW(), '/finetuning/getTaskList', '获取微调任务列表', 'Finetuning', 'GET'),
(NOW(), NOW(), '/finetuning/getTaskLog', '获取微调任务日志', 'Finetuning', 'GET');

-- =====================================================
-- 第4步：为管理员角色授权菜单（authority_id = 888）
-- =====================================================

-- 插入菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES
(888, @finetuning_menu_id),
(888, @task_list_menu_id),
(888, @task_detail_menu_id);

-- =====================================================
-- 第5步：为管理员角色授权API（使用 sys_casbin 表）
-- =====================================================

INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
-- 任务管理API权限
(NULL, 'p', '888', '/finetuning/createTask', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/finetuning/deleteTask', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/finetuning/stopTask', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/finetuning/getTask', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/finetuning/getTaskList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/finetuning/getTaskLog', 'GET', '', '', '', '');

-- =====================================================
-- 第6步：验证安装
-- =====================================================

SELECT '✅ 算法微调插件 SQL 执行完成！' as status;
SELECT COUNT(*) as '数据表数量' FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name LIKE 'gva_finetuning%';
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name LIKE 'finetuning%';
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group = 'Finetuning';
SELECT COUNT(*) as '菜单权限数' FROM sys_authority_menus WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name LIKE 'finetuning%'
);
SELECT COUNT(*) as 'API权限数' FROM sys_casbin WHERE v1 LIKE '/finetuning%' AND v0 = '888';

-- =====================================================
-- 使用说明
-- =====================================================

SELECT '📝 使用说明：' as info;
SELECT '1. 创建微调任务前，请确保服务器已安装 Python 和 PyTorch' as step1;
SELECT '2. 准备好基础模型（支持本地路径或HuggingFace模型）' as step2;
SELECT '3. 准备训练数据集' as step3;
SELECT '4. 进入【算法微调】->【微调任务】页面创建任务' as step4;
SELECT '5. 支持的模型：LLaMA、ChatGLM、Qwen 等主流大语言模型' as step5;

-- =====================================================
-- 回滚SQL（如需删除，请谨慎使用）
-- =====================================================

-- 删除API权限
-- DELETE FROM sys_casbin WHERE v1 LIKE '/finetuning%' AND v0 = '888';

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'finetuning%'
-- );

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE name LIKE 'finetuning%';

-- 删除API
-- DELETE FROM sys_apis WHERE api_group = 'Finetuning';

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS gva_finetuning_tasks;
