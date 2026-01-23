-- ====================================================================
-- K8s Manager 和 PortForward 插件菜单和权限初始化 SQL
-- 生成时间: 2026-01-23
-- 说明: 包含 K8s 管理和端口转发功能的所有菜单、按钮和 API 权限
-- ====================================================================

-- ====================================================================
-- 一、K8s 管理主菜单 (父菜单 ID: 5000)
-- ====================================================================

-- 插入 K8s 管理父菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), 0, 'k8s', 'k8s', 'view/routerHolder.vue', 8, '{"title": "K8s管理", "icon": "cpu-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 获取父菜单ID (假设为5000，实际环境中需要查询)
SET @k8s_parent_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'k8s' AND `parent_id` = 0 LIMIT 1);

-- ====================================================================
-- 二、K8s 管理子菜单
-- ====================================================================

-- 1. 集群管理 (ID: 5001)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'cluster', 'k8sCluster', 'plugin/k8smanager/view/cluster.vue', 1, '{"title": "集群管理", "icon": "server-2-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @cluster_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'cluster' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- 集群管理按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @cluster_menu_id, 'create', '创建集群', 1),
(NOW(), NOW(), @cluster_menu_id, 'update', '更新集群', 1),
(NOW(), NOW(), @cluster_menu_id, 'delete', '删除集群', 1),
(NOW(), NOW(), @cluster_menu_id, 'refresh', '刷新状态', 1),
(NOW(), NOW(), @cluster_menu_id, 'testConnection', '测试连接', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2. Pod管理 (ID: 5002)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'pod', 'k8sPod', 'plugin/k8smanager/view/pod.vue', 2, '{"title": "Pod管理", "icon": "apps-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @pod_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'pod' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- Pod管理按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @pod_menu_id, 'view', '查看详情', 1),
(NOW(), NOW(), @pod_menu_id, 'log', '查看日志', 1),
(NOW(), NOW(), @pod_menu_id, 'terminal', '打开终端', 1),
(NOW(), NOW(), @pod_menu_id, 'delete', '删除Pod', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3. Deployment管理 (ID: 5003)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'deployment', 'k8sDeployment', 'plugin/k8smanager/view/deployment.vue', 3, '{"title": "Deployment管理", "icon": "layout-grid-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @deployment_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'deployment' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- Deployment管理按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @deployment_menu_id, 'view', '查看详情', 1),
(NOW(), NOW(), @deployment_menu_id, 'scale', '扩缩容', 1),
(NOW(), NOW(), @deployment_menu_id, 'restart', '重启', 1),
(NOW(), NOW(), @deployment_menu_id, 'delete', '删除Deployment', 1),
(NOW(), NOW(), @deployment_menu_id, 'viewPods', '查看Pods', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 4. Service管理 (ID: 5004)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'service', 'k8sService', 'plugin/k8smanager/view/service.vue', 4, '{"title": "Service管理", "icon": "links-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @service_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'service' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- Service管理按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @service_menu_id, 'view', '查看详情', 1),
(NOW(), NOW(), @service_menu_id, 'viewEndpoints', '查看Endpoints', 1),
(NOW(), NOW(), @service_menu_id, 'delete', '删除Service', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 5. Namespace管理 (ID: 5005)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'namespace', 'k8sNamespace', 'plugin/k8smanager/view/namespace.vue', 5, '{"title": "Namespace管理", "icon": "folder-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @namespace_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'namespace' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- Namespace管理按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @namespace_menu_id, 'create', '创建Namespace', 1),
(NOW(), NOW(), @namespace_menu_id, 'view', '查看详情', 1),
(NOW(), NOW(), @namespace_menu_id, 'delete', '删除Namespace', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 6. 事件管理 (ID: 5006)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'event', 'k8sEvent', 'plugin/k8smanager/view/event.vue', 6, '{"title": "事件管理", "icon": "alert-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @event_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'event' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- 7. 监控指标 (ID: 5007)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'metrics', 'k8sMetrics', 'plugin/k8smanager/view/metrics.vue', 7, '{"title": "监控指标", "icon": "line-chart-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @metrics_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'metrics' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- 监控指标按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @metrics_menu_id, 'refresh', '刷新指标', 1),
(NOW(), NOW(), @metrics_menu_id, 'export', '导出数据', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 8. 审计日志 (ID: 5008)
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @k8s_parent_id, 'audit', 'k8sAudit', 'plugin/k8smanager/view/audit.vue', 8, '{"title": "审计日志", "icon": "file-list-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @audit_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'audit' AND `parent_id` = @k8s_parent_id LIMIT 1);

-- 审计日志按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @audit_menu_id, 'view', '查看详情', 1),
(NOW(), NOW(), @audit_menu_id, 'export', '导出日志', 1),
(NOW(), NOW(), @audit_menu_id, 'cleanup', '清理日志', 1),
(NOW(), NOW(), @audit_menu_id, 'stats', '统计信息', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ====================================================================
-- 三、端口转发管理菜单 (父菜单 ID: 4000)
-- ====================================================================

-- 插入端口转发父菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), 0, 'portForward', 'portForward', 'view/routerHolder.vue', 9, '{"title": "端口转发", "icon": "position-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @portforward_parent_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'portForward' AND `parent_id` = 0 LIMIT 1);

-- 端口转发规则管理子菜单
INSERT INTO `sys_base_menus` (`created_at`, `updated_at`, `parent_id`, `path`, `name`, `component`, `sort`, `meta`)
VALUES (NOW(), NOW(), @portforward_parent_id, 'rules', 'portForwardRules', 'plugin/portforward/view/portForward.vue', 1, '{"title": "转发规则", "icon": "route-line", "type": 0}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

SET @portforward_menu_id = (SELECT `id` FROM `sys_base_menus` WHERE `path` = 'rules' AND `parent_id` = @portforward_parent_id LIMIT 1);

-- 端口转发按钮权限
INSERT INTO `sys_base_menu_btns` (`created_at`, `updated_at`, `menu_id`, `name`, `desc`, `default_status`)
VALUES
(NOW(), NOW(), @portforward_menu_id, 'create', '创建规则', 1),
(NOW(), NOW(), @portforward_menu_id, 'update', '更新规则', 1),
(NOW(), NOW(), @portforward_menu_id, 'delete', '删除规则', 1),
(NOW(), NOW(), @portforward_menu_id, 'start', '启动转发', 1),
(NOW(), NOW(), @portforward_menu_id, 'stop', '停止转发', 1),
(NOW(), NOW(), @portforward_menu_id, 'viewStatus', '查看状态', 1)
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ====================================================================
-- 四、API 权限定义
-- ====================================================================

-- 插入 K8s 管理相关 API (假设从 ID 6000 开始)
-- 注意: 这些 path 需要与实际 API 路径匹配，权限标识通常格式为 "k8s:cluster:create" 等

-- 1. 集群管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/cluster/create', '创建K8s集群', 'k8s-cluster', 'POST'),
(NOW(), NOW(), '/k8s/cluster/delete', '删除K8s集群', 'k8s-cluster', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/deleteByIds', '批量删除K8s集群', 'k8s-cluster', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/update', '更新K8s集群', 'k8s-cluster', 'PUT'),
(NOW(), NOW(), '/k8s/cluster/get', '获取K8s集群详情', 'k8s-cluster', 'GET'),
(NOW(), NOW(), '/k8s/cluster/list', '获取K8s集群列表', 'k8s-cluster', 'GET'),
(NOW(), NOW(), '/k8s/cluster/refresh', '刷新集群状态', 'k8s-cluster', 'POST'),
(NOW(), NOW(), '/k8s/cluster/all', '获取所有集群', 'k8s-cluster', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2. Pod 管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/pod/list', '获取Pod列表', 'k8s-pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/get', '获取Pod详情', 'k8s-pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/delete', '删除Pod', 'k8s-pod', 'DELETE'),
(NOW(), NOW(), '/k8s/pod/log', '获取Pod日志', 'k8s-pod', 'POST'),
(NOW(), NOW(), '/k8s/pod/containers', '获取Pod容器列表', 'k8s-pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/events', '获取Pod事件', 'k8s-pod', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3. Deployment 管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/deployment/list', '获取Deployment列表', 'k8s-deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/get', '获取Deployment详情', 'k8s-deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/scale', '扩缩容Deployment', 'k8s-deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/restart', '重启Deployment', 'k8s-deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/delete', '删除Deployment', 'k8s-deployment', 'DELETE'),
(NOW(), NOW(), '/k8s/deployment/pods', '获取Deployment关联的Pods', 'k8s-deployment', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 4. Service 管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/service/list', '获取Service列表', 'k8s-service', 'GET'),
(NOW(), NOW(), '/k8s/service/get', '获取Service详情', 'k8s-service', 'GET'),
(NOW(), NOW(), '/k8s/service/delete', '删除Service', 'k8s-service', 'DELETE'),
(NOW(), NOW(), '/k8s/service/endpoints', '获取Service的Endpoints', 'k8s-service', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 5. Namespace 管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/namespace/list', '获取Namespace列表', 'k8s-namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/get', '获取Namespace详情', 'k8s-namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/create', '创建Namespace', 'k8s-namespace', 'POST'),
(NOW(), NOW(), '/k8s/namespace/delete', '删除Namespace', 'k8s-namespace', 'DELETE')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 6. Event 管理 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/event/list', '获取Event列表', 'k8s-event', 'POST')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 7. 审计日志 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/audit/list', '获取审计日志列表', 'k8s-audit', 'GET'),
(NOW(), NOW(), '/k8s/audit/stats', '获取审计统计信息', 'k8s-audit', 'GET'),
(NOW(), NOW(), '/k8s/audit/client-stats', '获取客户端统计信息', 'k8s-audit', 'GET'),
(NOW(), NOW(), '/k8s/audit/cleanup', '清理审计日志', 'k8s-audit', 'DELETE'),
(NOW(), NOW(), '/k8s/audit/export', '导出审计日志', 'k8s-audit', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 8. 监控指标 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/k8s/metrics/cluster', '获取集群指标', 'k8s-metrics', 'GET'),
(NOW(), NOW(), '/k8s/metrics/cluster/refresh', '刷新集群指标', 'k8s-metrics', 'POST'),
(NOW(), NOW(), '/k8s/metrics/nodes', '获取节点指标', 'k8s-metrics', 'GET'),
(NOW(), NOW(), '/k8s/metrics/pods', '获取Pod指标', 'k8s-metrics', 'GET'),
(NOW(), NOW(), '/k8s/metrics/summary', '获取指标汇总', 'k8s-metrics', 'GET'),
(NOW(), NOW(), '/k8s/metrics/collector/start', '启动指标收集器', 'k8s-metrics', 'POST'),
(NOW(), NOW(), '/k8s/metrics/collector/stop', '停止指标收集器', 'k8s-metrics', 'POST')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 9. 端口转发 API
INSERT INTO `sys_apis` (`created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(NOW(), NOW(), '/portForward/createPortForward', '创建端口转发规则', 'portForward', 'POST'),
(NOW(), NOW(), '/portForward/deletePortForward', '删除端口转发规则', 'portForward', 'DELETE'),
(NOW(), NOW(), '/portForward/deletePortForwardByIds', '批量删除端口转发规则', 'portForward', 'DELETE'),
(NOW(), NOW(), '/portForward/updatePortForward', '更新端口转发规则', 'portForward', 'PUT'),
(NOW(), NOW(), '/portForward/findPortForward', '根据ID获取端口转发规则', 'portForward', 'GET'),
(NOW(), NOW(), '/portForward/getPortForwardList', '获取端口转发规则列表', 'portForward', 'GET'),
(NOW(), NOW(), '/portForward/updatePortForwardStatus', '更新端口转发规则状态', 'portForward', 'PUT'),
(NOW(), NOW(), '/portForward/getServerIP', '获取服务器IP地址', 'portForward', 'GET'),
(NOW(), NOW(), '/portForward/getForwarderStatus', '获取端口转发运行状态', 'portForward', 'GET'),
(NOW(), NOW(), '/portForward/getAllForwarderStatus', '获取所有端口转发运行状态', 'portForward', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ====================================================================
-- 五、关联菜单按钮与 API 权限
-- ====================================================================

-- 为管理员角色 (authority_id = 888) 分配所有权限
-- 注意: 实际使用时需要根据您的 authority_id 调整

-- 1. 集群管理按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/cluster/create' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @cluster_menu_id AND `name` = 'create'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/cluster/update' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @cluster_menu_id AND `name` = 'update'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/cluster/delete' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @cluster_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2. Pod管理按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/pod/get' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @pod_menu_id AND `name` = 'view'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/pod/log' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @pod_menu_id AND `name` = 'log'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/pod/delete' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @pod_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3. Deployment管理按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/deployment/get' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @deployment_menu_id AND `name` = 'view'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/deployment/scale' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @deployment_menu_id AND `name` = 'scale'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/deployment/restart' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @deployment_menu_id AND `name` = 'restart'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/deployment/delete' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @deployment_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/deployment/pods' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @deployment_menu_id AND `name` = 'viewPods'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 4. Service管理按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/service/get' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @service_menu_id AND `name` = 'view'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/service/endpoints' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @service_menu_id AND `name` = 'viewEndpoints'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/service/delete' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @service_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 5. Namespace管理按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/namespace/create' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @namespace_menu_id AND `name` = 'create'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/namespace/get' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @namespace_menu_id AND `name` = 'view'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/k8s/namespace/delete' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @namespace_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 6. 端口转发按钮与API关联
INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/portForward/createPortForward' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @portforward_menu_id AND `name` = 'create'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/portForward/updatePortForward' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @portforward_menu_id AND `name` = 'update'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/portForward/deletePortForward' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @portforward_menu_id AND `name` = 'delete'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/portForward/updatePortForwardStatus' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @portforward_menu_id AND `name` IN ('start', 'stop')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

INSERT INTO `sys_authority_btn_api` (`created_at`, `updated_at`, `authority_id`, `menu_id`, `btn_id`, `api_id`)
SELECT NOW(), NOW(), 888, menu_id, id, (SELECT `id` FROM `sys_apis` WHERE `path` = '/portForward/getForwarderStatus' LIMIT 1)
FROM `sys_base_menu_btns` WHERE `menu_id` = @portforward_menu_id AND `name` = 'viewStatus'
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ====================================================================
-- 六、为管理员角色分配菜单访问权限
-- ====================================================================

-- 假设管理员角色 authority_id = 888，根据实际情况调整
-- 关联角色与菜单

INSERT INTO `sys_authority_menus` (`created_at`, `updated_at`, `menu_id`, `authority_id`, `authority_id_menu_id`)
SELECT NOW(), NOW(), `id`, 888, CONCAT('888-', `id`)
FROM `sys_base_menus`
WHERE `path` IN ('k8s', 'cluster', 'pod', 'deployment', 'service', 'namespace', 'event', 'metrics', 'audit', 'portForward', 'rules')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ====================================================================
-- SQL 执行完成
-- ====================================================================
