-- =====================================================
-- 补充缺失的菜单和API权限 SQL
-- 生成时间: 自动生成
-- 说明: 补充 K8s Manager 缺失的监控指标和审计日志菜单/API
-- =====================================================

-- =====================================================
-- 一、K8s Manager - 补充缺失的菜单
-- =====================================================

-- 获取K8s父菜单ID
SET @k8s_parent_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'k8s' AND `parent_id` = 0 LIMIT 1);

-- 1. 创建监控指标菜单 (如果不存在)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
SELECT NOW(), NOW(), @k8s_parent_id, 'metrics', 'k8sMetrics', 0, 'plugin/k8smanager/view/metrics.vue', 7,
'{"title": "监控指标", "icon": "line-chart-line", "type": 0}'
WHERE NOT EXISTS (
    SELECT 1 FROM `sys_base_menus` WHERE `path` = 'metrics' AND `parent_id` = @k8s_parent_id
);

SET @metrics_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'metrics' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- 2. 创建审计日志菜单 (如果不存在)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
SELECT NOW(), NOW(), @k8s_parent_id, 'audit', 'k8sAudit', 0, 'plugin/k8smanager/view/audit.vue', 8,
'{"title": "审计日志", "icon": "file-list-line", "type": 0}'
WHERE NOT EXISTS (
    SELECT 1 FROM `sys_base_menus` WHERE `path` = 'audit' AND `parent_id` = @k8s_parent_id
);

SET @audit_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'audit' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- =====================================================
-- 二、K8s Manager - 补充缺失的API
-- =====================================================

-- 监控指标相关API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/cluster', '获取集群指标', 'k8s-metrics', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/cluster' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/cluster/refresh', '刷新集群指标', 'k8s-metrics', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/cluster/refresh' AND `method` = 'POST');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/nodes', '获取节点指标', 'k8s-metrics', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/nodes' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/pods', '获取Pod指标', 'k8s-metrics', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/pods' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/summary', '获取指标汇总', 'k8s-metrics', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/summary' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/collector/start', '启动指标收集器', 'k8s-metrics', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/collector/start' AND `method` = 'POST');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/metrics/collector/stop', '停止指标收集器', 'k8s-metrics', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/metrics/collector/stop' AND `method` = 'POST');

-- 审计日志相关API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/audit/list', '获取审计日志列表', 'k8s-audit', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/audit/list' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/audit/stats', '获取审计统计信息', 'k8s-audit', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/audit/stats' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/audit/client-stats', '获取客户端统计信息', 'k8s-audit', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/audit/client-stats' AND `method` = 'GET');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/audit/cleanup', '清理审计日志', 'k8s-audit', 'DELETE'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/audit/cleanup' AND `method` = 'DELETE');

INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
SELECT NOW(), NOW(), '/k8s/audit/export', '导出审计日志', 'k8s-audit', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_apis` WHERE `path` = '/k8s/audit/export' AND `method` = 'GET');

-- =====================================================
-- 三、K8s Manager - 补充API权限 (sys_casbin)
-- =====================================================

-- 监控指标API权限
INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/cluster', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/cluster' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/cluster/refresh', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/cluster/refresh' AND `v2` = 'POST');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/nodes', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/nodes' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/pods', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/pods' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/summary', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/summary' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/collector/start', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/collector/start' AND `v2` = 'POST');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/metrics/collector/stop', 'POST'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/metrics/collector/stop' AND `v2` = 'POST');

-- 审计日志API权限
INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/audit/list', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/audit/list' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/audit/stats', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/audit/stats' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/audit/client-stats', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/audit/client-stats' AND `v2` = 'GET');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/audit/cleanup', 'DELETE'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/audit/cleanup' AND `v2` = 'DELETE');

INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
SELECT 'p', '888', '/k8s/audit/export', 'GET'
WHERE NOT EXISTS (SELECT 1 FROM `sys_casbin` WHERE `v0` = '888' AND `v1` = '/k8s/audit/export' AND `v2` = 'GET');

-- =====================================================
-- 四、补充菜单权限 (sys_authority_menus)
-- =====================================================

-- 监控指标菜单权限
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`)
SELECT 888, @metrics_menu_id
WHERE @metrics_menu_id IS NOT NULL
  AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus`
    WHERE `sys_authority_authority_id` = 888 AND `sys_base_menu_id` = @metrics_menu_id
);

-- 审计日志菜单权限
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`)
SELECT 888, @audit_menu_id
WHERE @audit_menu_id IS NOT NULL
  AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus`
    WHERE `sys_authority_authority_id` = 888 AND `sys_base_menu_id` = @audit_menu_id
);

-- =====================================================
-- 五、AI Agent - 修复API路径（添加 /aiagent 前缀）
-- =====================================================

-- 删除错误的API（路径缺少 /aiagent 前缀）
DELETE FROM `sys_apis` WHERE `api_group` = 'AI Agent' AND `path` NOT LIKE '/aiagent/%';

-- 重新创建正确的API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
-- 会话相关API
(NOW(), NOW(), '/aiagent/conversation/createConversation', '创建会话', 'AI Agent', 'POST'),
(NOW(), NOW(), '/aiagent/conversation/deleteConversation', '删除会话', 'AI Agent', 'DELETE'),
(NOW(), NOW(), '/aiagent/conversation/updateConversation', '更新会话', 'AI Agent', 'PUT'),
(NOW(), NOW(), '/aiagent/conversation/findConversation', '根据ID获取会话', 'AI Agent', 'GET'),
(NOW(), NOW(), '/aiagent/conversation/getConversationList', '获取会话列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/aiagent/conversation/setActive', '设置会话激活状态', 'AI Agent', 'POST'),
(NOW(), NOW(), '/aiagent/conversation/getActive', '获取激活的会话', 'AI Agent', 'GET'),
-- 消息相关API
(NOW(), NOW(), '/aiagent/message/getMessageList', '获取消息列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/aiagent/message/deleteMessage', '删除消息', 'AI Agent', 'DELETE'),
-- 聊天相关API
(NOW(), NOW(), '/aiagent/chat/sendMessage', '发送消息', 'AI Agent', 'POST'),
-- 配置相关API
(NOW(), NOW(), '/aiagent/config/createConfig', '创建AI配置', 'AI Agent', 'POST'),
(NOW(), NOW(), '/aiagent/config/deleteConfig', '删除AI配置', 'AI Agent', 'DELETE'),
(NOW(), NOW(), '/aiagent/config/updateConfig', '更新AI配置', 'AI Agent', 'PUT'),
(NOW(), NOW(), '/aiagent/config/findConfig', '根据ID获取AI配置', 'AI Agent', 'GET'),
(NOW(), NOW(), '/aiagent/config/getConfigList', '获取AI配置列表', 'AI Agent', 'GET'),
(NOW(), NOW(), '/aiagent/config/setActive', '设置AI配置激活状态', 'AI Agent', 'POST'),
(NOW(), NOW(), '/aiagent/config/getActive', '获取激活的AI配置', 'AI Agent', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 删除错误的API权限
DELETE FROM `sys_casbin` WHERE `v0` = '888' AND (
    (`v1` LIKE '/conversation%' OR `v1` LIKE '/message%' OR `v1` LIKE '/chat%' OR `v1` LIKE '/config%')
    AND `v1` NOT LIKE '/aiagent/%'
);

-- 重新创建正确的API权限
INSERT INTO `sys_casbin` (`ptype`, `v0`, `v1`, `v2`)
VALUES
-- 会话API权限
('p', '888', '/aiagent/conversation/createConversation', 'POST'),
('p', '888', '/aiagent/conversation/deleteConversation', 'DELETE'),
('p', '888', '/aiagent/conversation/updateConversation', 'PUT'),
('p', '888', '/aiagent/conversation/findConversation', 'GET'),
('p', '888', '/aiagent/conversation/getConversationList', 'GET'),
('p', '888', '/aiagent/conversation/setActive', 'POST'),
('p', '888', '/aiagent/conversation/getActive', 'GET'),
-- 消息API权限
('p', '888', '/aiagent/message/getMessageList', 'GET'),
('p', '888', '/aiagent/message/deleteMessage', 'DELETE'),
-- 聊天API权限
('p', '888', '/aiagent/chat/sendMessage', 'POST'),
-- 配置API权限
('p', '888', '/aiagent/config/createConfig', 'POST'),
('p', '888', '/aiagent/config/deleteConfig', 'DELETE'),
('p', '888', '/aiagent/config/updateConfig', 'PUT'),
('p', '888', '/aiagent/config/findConfig', 'GET'),
('p', '888', '/aiagent/config/getConfigList', 'GET'),
('p', '888', '/aiagent/config/setActive', 'POST'),
('p', '888', '/aiagent/config/getActive', 'GET')
ON DUPLICATE KEY UPDATE `v1` = `v1`;

-- =====================================================
-- 验证结果
-- =====================================================

SELECT '✅ 权限同步SQL执行完成！' as status;

SELECT 'K8s监控指标菜单' as item, COUNT(*) as count
FROM sys_base_menus WHERE path = 'metrics' AND name = 'k8sMetrics'
UNION ALL
SELECT 'K8s审计日志菜单', COUNT(*)
FROM sys_base_menus WHERE path = 'audit' AND name = 'k8sAudit'
UNION ALL
SELECT 'K8s监控指标API数', COUNT(*)
FROM sys_apis WHERE api_group = 'k8s-metrics'
UNION ALL
SELECT 'K8s审计日志API数', COUNT(*)
FROM sys_apis WHERE api_group = 'k8s-audit'
UNION ALL
SELECT 'AI Agent正确路径API数', COUNT(*)
FROM sys_apis WHERE api_group = 'AI Agent' AND path LIKE '/aiagent/%'
UNION ALL
SELECT '监控指标API权限数', COUNT(*)
FROM sys_casbin WHERE v0 = '888' AND v1 LIKE '/k8s/metrics/%'
UNION ALL
SELECT '审计日志API权限数', COUNT(*)
FROM sys_casbin WHERE v0 = '888' AND v1 LIKE '/k8s/audit/%'
UNION ALL
SELECT 'AI Agent API权限数', COUNT(*)
FROM sys_casbin WHERE v0 = '888' AND v1 LIKE '/aiagent/%';

-- =====================================================
-- SQL 执行完成
-- =====================================================
