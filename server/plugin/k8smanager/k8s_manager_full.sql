-- ========================================
-- K8s 管理插件 - 完整 SQL 脚本
-- ========================================
-- 执行说明：
-- 1. 建议先备份数据库
-- 2. 根据实际情况调整ID（避免冲突）
-- 3. authority_id 需要根据实际系统角色表调整
-- ========================================

-- ========================================
-- 第一部分：创建数据表
-- ========================================

-- 删除已存在的表（慎用）
-- DROP TABLE IF EXISTS `k8s_clusters`;

-- 创建 K8s 集群配置表
CREATE TABLE IF NOT EXISTS `k8s_clusters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(100) NOT NULL COMMENT '集群名称',
  `kube_config` longtext COMMENT 'kubeconfig配置内容',
  `endpoint` varchar(500) DEFAULT NULL COMMENT 'API Server地址',
  `version` varchar(50) DEFAULT NULL COMMENT 'K8s版本',
  `status` varchar(20) DEFAULT 'unknown' COMMENT '集群状态: online, offline, unknown',
  `description` varchar(500) DEFAULT NULL COMMENT '集群描述',
  `region` varchar(100) DEFAULT NULL COMMENT '区域',
  `provider` varchar(50) DEFAULT NULL COMMENT '云服务商',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认集群',
  `node_count` int DEFAULT '0' COMMENT '节点数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_status` (`status`),
  KEY `idx_provider` (`provider`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='K8s集群配置表';

-- ========================================
-- 第二部分：插入菜单数据
-- ========================================

-- 注意：菜单ID从5000开始，避免与现有菜单冲突
-- 如需调整，请批量替换所有ID

-- 2.1 K8s 管理父菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5000, NOW(), NOW(), 0, 'k8s', 'k8s', 0, 'view/routerHolder.vue', 8, '{"title": "K8s管理", "icon": "cpu-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.2 集群管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5001, NOW(), NOW(), 5000, 'k8sCluster', 'k8sCluster', 0, 'plugin/k8smanager/view/cluster.vue', 1, '{"title": "集群管理", "icon": "server-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.3 Pod 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5002, NOW(), NOW(), 5000, 'k8sPod', 'k8sPod', 0, 'plugin/k8smanager/view/pod.vue', 2, '{"title": "Pod管理", "icon": "apps-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.4 Deployment 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5003, NOW(), NOW(), 5000, 'k8sDeployment', 'k8sDeployment', 0, 'plugin/k8smanager/view/deployment.vue', 3, '{"title": "Deployment管理", "icon": "stack-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.5 Service 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5004, NOW(), NOW(), 5000, 'k8sService', 'k8sService', 0, 'plugin/k8smanager/view/service.vue', 4, '{"title": "Service管理", "icon": "links-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.6 Namespace 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5005, NOW(), NOW(), 5000, 'k8sNamespace', 'k8sNamespace', 0, 'plugin/k8smanager/view/namespace.vue', 5, '{"title": "Namespace管理", "icon": "folder-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 2.7 事件管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(5006, NOW(), NOW(), 5000, 'k8sEvent', 'k8sEvent', 0, 'plugin/k8smanager/view/event.vue', 6, '{"title": "事件管理", "icon": "notification-line"}')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ========================================
-- 第三部分：插入 API 权限数据
-- ========================================

-- 注意：API ID从6000开始，避免与现有API冲突

-- 3.1 K8s集群相关API (8个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6001, NOW(), NOW(), '/k8s/cluster/create', '创建K8s集群', 'K8s集群', 'POST'),
(6002, NOW(), NOW(), '/k8s/cluster/delete', '删除K8s集群', 'K8s集群', 'DELETE'),
(6003, NOW(), NOW(), '/k8s/cluster/deleteByIds', '批量删除K8s集群', 'K8s集群', 'DELETE'),
(6004, NOW(), NOW(), '/k8s/cluster/update', '更新K8s集群', 'K8s集群', 'PUT'),
(6005, NOW(), NOW(), '/k8s/cluster/get', '获取K8s集群详情', 'K8s集群', 'GET'),
(6006, NOW(), NOW(), '/k8s/cluster/list', '获取K8s集群列表', 'K8s集群', 'GET'),
(6007, NOW(), NOW(), '/k8s/cluster/refresh', '刷新K8s集群状态', 'K8s集群', 'POST'),
(6008, NOW(), NOW(), '/k8s/cluster/all', '获取所有K8s集群', 'K8s集群', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3.2 Pod相关API (6个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6009, NOW(), NOW(), '/k8s/pod/list', '获取Pod列表', 'K8s Pod', 'GET'),
(6010, NOW(), NOW(), '/k8s/pod/get', '获取Pod详情', 'K8s Pod', 'GET'),
(6011, NOW(), NOW(), '/k8s/pod/delete', '删除Pod', 'K8s Pod', 'DELETE'),
(6012, NOW(), NOW(), '/k8s/pod/log', '获取Pod日志', 'K8s Pod', 'POST'),
(6013, NOW(), NOW(), '/k8s/pod/containers', '获取Pod容器列表', 'K8s Pod', 'GET'),
(6014, NOW(), NOW(), '/k8s/pod/events', '获取Pod事件', 'K8s Pod', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3.3 Deployment相关API (6个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6015, NOW(), NOW(), '/k8s/deployment/list', '获取Deployment列表', 'K8s Deployment', 'GET'),
(6016, NOW(), NOW(), '/k8s/deployment/get', '获取Deployment详情', 'K8s Deployment', 'GET'),
(6017, NOW(), NOW(), '/k8s/deployment/scale', '扩缩容Deployment', 'K8s Deployment', 'POST'),
(6018, NOW(), NOW(), '/k8s/deployment/restart', '重启Deployment', 'K8s Deployment', 'POST'),
(6019, NOW(), NOW(), '/k8s/deployment/delete', '删除Deployment', 'K8s Deployment', 'DELETE'),
(6020, NOW(), NOW(), '/k8s/deployment/pods', '获取Deployment关联的Pods', 'K8s Deployment', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3.4 Service相关API (4个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6021, NOW(), NOW(), '/k8s/service/list', '获取Service列表', 'K8s Service', 'GET'),
(6022, NOW(), NOW(), '/k8s/service/get', '获取Service详情', 'K8s Service', 'GET'),
(6023, NOW(), NOW(), '/k8s/service/delete', '删除Service', 'K8s Service', 'DELETE'),
(6024, NOW(), NOW(), '/k8s/service/endpoints', '获取Service的Endpoints', 'K8s Service', 'GET')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3.5 Namespace相关API (4个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6025, NOW(), NOW(), '/k8s/namespace/list', '获取Namespace列表', 'K8s Namespace', 'GET'),
(6026, NOW(), NOW(), '/k8s/namespace/get', '获取Namespace详情', 'K8s Namespace', 'GET'),
(6027, NOW(), NOW(), '/k8s/namespace/create', '创建Namespace', 'K8s Namespace', 'POST'),
(6028, NOW(), NOW(), '/k8s/namespace/delete', '删除Namespace', 'K8s Namespace', 'DELETE')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- 3.6 Event相关API (1个)
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(6029, NOW(), NOW(), '/k8s/event/list', '获取Event列表', 'K8s Event', 'POST')
ON DUPLICATE KEY UPDATE `updated_at` = NOW();

-- ========================================
-- 第四部分：为管理员角色添加菜单权限
-- ========================================

-- 注意：authority_id = 888 为gin-vue-admin默认管理员角色ID
-- 如果您的系统管理员角色ID不同，请替换下面的 888

-- 4.1 添加菜单权限（7个菜单）
INSERT IGNORE INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`)
VALUES
(5000, 888),  -- K8s管理父菜单
(5001, 888),  -- 集群管理
(5002, 888),  -- Pod管理
(5003, 888),  -- Deployment管理
(5004, 888),  -- Service管理
(5005, 888),  -- Namespace管理
(5006, 888);  -- 事件管理

-- ========================================
-- 第五部分：为管理员角色添加 API 权限
-- ========================================

-- 5.1 添加API权限（29个API）
INSERT IGNORE INTO `sys_authority_apis` (`sys_api_id`, `sys_authority_authority_id`)
VALUES
-- K8s集群 API (6001-6008)
(6001, 888),
(6002, 888),
(6003, 888),
(6004, 888),
(6005, 888),
(6006, 888),
(6007, 888),
(6008, 888),
-- Pod API (6009-6014)
(6009, 888),
(6010, 888),
(6011, 888),
(6012, 888),
(6013, 888),
(6014, 888),
-- Deployment API (6015-6020)
(6015, 888),
(6016, 888),
(6017, 888),
(6018, 888),
(6019, 888),
(6020, 888),
-- Service API (6021-6024)
(6021, 888),
(6022, 888),
(6023, 888),
(6024, 888),
-- Namespace API (6025-6028)
(6025, 888),
(6026, 888),
(6027, 888),
(6028, 888),
-- Event API (6029)
(6029, 888);

-- ========================================
-- 第六部分：验证SQL（可选）
-- ========================================

-- 查询已创建的菜单
-- SELECT id, name, title FROM sys_base_menus WHERE id BETWEEN 5000 AND 5006 ORDER BY id;

-- 查询已创建的API
-- SELECT id, path, description, api_group FROM sys_apis WHERE id BETWEEN 6001 AND 6029 ORDER BY id;

-- 查询菜单权限
-- SELECT * FROM sys_authority_menus WHERE sys_base_menu_id BETWEEN 5000 AND 5006;

-- 查询API权限
-- SELECT * FROM sys_authority_apis WHERE sys_api_id BETWEEN 6001 AND 6029;

-- ========================================
-- 回滚SQL（如需删除，请谨慎使用）
-- ========================================

-- 删除API权限
-- DELETE FROM sys_authority_apis WHERE sys_api_id BETWEEN 6001 AND 6029;

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id BETWEEN 5000 AND 5006;

-- 删除API
-- DELETE FROM sys_apis WHERE id BETWEEN 6001 AND 6029;

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE id BETWEEN 5000 AND 5006;

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS k8s_clusters;

-- ========================================
-- 执行完成提示
-- ========================================

-- SQL执行完成后，请检查：
-- 1. k8s_clusters 表是否创建成功
-- 2. sys_base_menus 表中是否有7条新记录（ID 5000-5006）
-- 3. sys_apis 表中是否有29条新记录（ID 6001-6029）
-- 4. sys_authority_menus 表中是否有7条新记录
-- 5. sys_authority_apis 表中是否有29条新记录

-- ========================================
