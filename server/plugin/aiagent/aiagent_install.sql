-- =====================================================
-- AI Agent 插件 - 初始化 SQL 脚本
-- =====================================================
-- 执行说明：
-- 1. 本脚本根据 gin-vue-admin 实际表结构调整
-- 2. 包含数据表、菜单、API、权限的完整初始化
-- 3. 支持智谱 GLM-4.7 模型
-- =====================================================

-- =====================================================
-- 第1步：创建数据表
-- =====================================================

-- 1.1 创建会话表
CREATE TABLE IF NOT EXISTS `gva_aiagent_conversations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(200) NOT NULL COMMENT '会话标题',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '所属用户ID',
  `model` varchar(50) NOT NULL DEFAULT 'glm-4-plus' COMMENT '使用的模型',
  `system_prompt` text COMMENT '系统提示词',
  `temperature` double DEFAULT 0.7 COMMENT '温度参数',
  `max_tokens` int DEFAULT 4096 COMMENT '最大token数',
  `is_active` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否激活',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_is_active` (`is_active`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='AI对话会话表';

-- 1.2 创建消息表
CREATE TABLE IF NOT EXISTS `gva_aiagent_messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `conversation_id` bigint unsigned NOT NULL COMMENT '会话ID',
  `role` varchar(20) NOT NULL COMMENT '角色(user/assistant/system)',
  `content` text NOT NULL COMMENT '消息内容',
  `token_count` int DEFAULT NULL COMMENT 'token数量',
  `metadata` json COMMENT '元数据(如finish_reason等)',
  PRIMARY KEY (`id`),
  KEY `idx_conversation_id` (`conversation_id`),
  KEY `idx_role` (`role`),
  KEY `idx_deleted_at` (`deleted_at`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='AI对话消息表';

-- 1.3 创建配置表
CREATE TABLE IF NOT EXISTS `gva_aiagent_configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '配置名称',
  `api_key` varchar(200) NOT NULL COMMENT 'API Key',
  `base_url` varchar(500) NOT NULL DEFAULT 'https://open.bigmodel.cn/api/paas/v4/' COMMENT 'API基础URL',
  `model` varchar(50) NOT NULL DEFAULT 'glm-4-plus' COMMENT '默认模型',
  `temperature` double NOT NULL DEFAULT 0.7 COMMENT '默认温度',
  `max_tokens` int NOT NULL DEFAULT 4096 COMMENT '默认最大token数',
  `is_active` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_is_active` (`is_active`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='AI Agent配置表';

-- =====================================================
-- 第2步：创建菜单
-- =====================================================

-- 2.1 创建AI Agent主菜单（作为系统功能的子菜单）
-- 注意：如果系统功能的ID不是24，请根据实际情况调整
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), 24, 'aiagent', 'aiagent', 0, 'view/routerHolder.vue', 6, 0, 'AI Agent', 'chat-dot-square');

-- 获取主菜单ID
SET @aiagent_menu_id = LAST_INSERT_ID();

-- 2.2 创建子菜单
-- AI 对话
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @aiagent_menu_id, 'chat', 'aiagentChat', 0, 'plugin/aiagent/view/chat.vue', 1, 0, 'AI 对话', 'chat-line-round');

SET @chat_menu_id = LAST_INSERT_ID();

-- AI 配置
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @aiagent_menu_id, 'config', 'aiagentConfig', 0, 'plugin/aiagent/view/config.vue', 2, 0, 'AI 配置', 'setting');

SET @config_menu_id = LAST_INSERT_ID();

-- =====================================================
-- 第3步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
-- 会话相关API
(NOW(), NOW(), '/conversation/createConversation', '创建会话', 'AI Agent', 'POST'),
(NOW(), NOW(), '/conversation/deleteConversation', '删除会话', 'AI Agent', 'DELETE'),
(NOW(), NOW(), '/conversation/updateConversation', '更新会话', 'AI Agent', 'PUT'),
(NOW(), NOW(), '/conversation/findConversation', '根据ID获取会话', 'AI Agent', 'GET'),
(NOW(), NOW(), '/conversation/getConversationList', '获取会话列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/conversation/setActive', '设置会话激活状态', 'AI Agent', 'POST'),
(NOW(), NOW(), '/conversation/getActive', '获取激活的会话', 'AI Agent', 'GET'),
-- 消息相关API
(NOW(), NOW(), '/message/getMessageList', '获取消息列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/message/deleteMessage', '删除消息', 'AI Agent', 'DELETE'),
-- 聊天相关API
(NOW(), NOW(), '/chat/sendMessage', '发送消息', 'AI Agent', 'POST'),
-- 配置相关API
(NOW(), NOW(), '/config/createConfig', '创建AI配置', 'AI Agent', 'POST'),
(NOW(), NOW(), '/config/deleteConfig', '删除AI配置', 'AI Agent', 'DELETE'),
(NOW(), NOW(), '/config/updateConfig', '更新AI配置', 'AI Agent', 'PUT'),
(NOW(), NOW(), '/config/findConfig', '根据ID获取AI配置', 'AI Agent', 'GET'),
(NOW(), NOW(), '/config/getConfigList', '获取AI配置列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/config/setActive', '设置AI配置激活状态', 'AI Agent', 'POST'),
(NOW(), NOW(), '/config/getActive', '获取激活的AI配置', 'AI Agent', 'GET');

-- =====================================================
-- 第4步：为管理员角色授权菜单（authority_id = 888）
-- =====================================================

-- 插入菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES
(888, @aiagent_menu_id),
(888, @chat_menu_id),
(888, @config_menu_id);

-- =====================================================
-- 第5步：为管理员角色授权API（使用 sys_casbin 表）
-- =====================================================

INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
-- 会话API权限
(NULL, 'p', '888', '/conversation/createConversation', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/conversation/deleteConversation', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/conversation/updateConversation', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/conversation/findConversation', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/conversation/getConversationList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/conversation/setActive', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/conversation/getActive', 'GET', '', '', '', ''),
-- 消息API权限
(NULL, 'p', '888', '/message/getMessageList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/message/deleteMessage', 'DELETE', '', '', '', ''),
-- 聊天API权限
(NULL, 'p', '888', '/chat/sendMessage', 'POST', '', '', '', ''),
-- 配置API权限
(NULL, 'p', '888', '/config/createConfig', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/config/deleteConfig', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/config/updateConfig', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/config/findConfig', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/config/getConfigList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/config/setActive', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/config/getActive', 'GET', '', '', '', '');

-- =====================================================
-- 第6步：创建默认配置（可选）
-- =====================================================

-- 插入示例配置（用户需要修改API Key）
INSERT INTO gva_aiagent_configs (created_at, updated_at, name, api_key, base_url, model, temperature, max_tokens, is_active)
VALUES (NOW(), NOW(), '默认GLM配置', 'your-api-key-here', 'https://open.bigmodel.cn/api/paas/v4/', 'glm-4-plus', 0.7, 4096, 1);

-- =====================================================
-- 第7步：验证安装
-- =====================================================

SELECT '✅ AI Agent 插件 SQL 执行完成！' as status;
SELECT COUNT(*) as '数据表数量' FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name LIKE 'gva_aiagent%';
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name LIKE 'aiagent%';
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group = 'AI Agent';
SELECT COUNT(*) as '菜单权限数' FROM sys_authority_menus WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name LIKE 'aiagent%'
);
SELECT COUNT(*) as 'API权限数' FROM sys_casbin WHERE v1 LIKE '/conversation%' OR v1 LIKE '/message%' OR v1 LIKE '/chat%' OR v1 LIKE '/config%';

-- =====================================================
-- 使用说明
-- =====================================================

SELECT '📝 使用说明：' as info;
SELECT '1. 首次使用请先到【AI Agent】->【AI 配置】页面添加智谱AI的API Key' as step1;
SELECT '2. API Key可以从智谱AI开放平台获取：https://open.bigmodel.cn/' as step2;
SELECT '3. 配置完成后，到【AI Agent】->【AI 对话】页面开始对话' as step3;
SELECT '4. 支持的模型：GLM-4-Plus、GLM-4-Air、GLM-4-Flash、GLM-3-Turbo' as step4;

-- =====================================================
-- 回滚SQL（如需删除，请谨慎使用）
-- =====================================================

-- 删除API权限
-- DELETE FROM sys_casbin WHERE v1 LIKE '/conversation%' OR v1 LIKE '/message%' OR v1 LIKE '/chat%' OR v1 LIKE '/config%' AND v0 = '888';

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'aiagent%'
-- );

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE name LIKE 'aiagent%';

-- 删除API
-- DELETE FROM sys_apis WHERE api_group = 'AI Agent';

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS gva_aiagent_messages;
-- DROP TABLE IF EXISTS gva_aiagent_conversations;
-- DROP TABLE IF EXISTS gva_aiagent_configs;
