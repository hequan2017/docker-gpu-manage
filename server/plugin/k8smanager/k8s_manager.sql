-- ========================================
-- K8s 管理插件 - SQL 脚本
-- ========================================

-- 注意：执行前请根据实际情况调整 menu_id 和 parent_id
-- 以下 SQL 假设 K8s 管理作为顶级菜单插入

-- ========================================
-- 1. 菜单数据
-- ========================================

-- 1.1 K8s 管理父菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1000, NOW(), NOW(), 0, 'k8s', 'k8s', 0, 'view/routerHolder.vue', 8, '{"title": "K8s管理", "icon": "cpu-line"}');

-- 1.2 集群管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1001, NOW(), NOW(), 1000, 'k8sCluster', 'k8sCluster', 0, 'plugin/k8smanager/view/cluster.vue', 1, '{"title": "集群管理", "icon": "server-line"}');

-- 1.3 Pod 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1002, NOW(), NOW(), 1000, 'k8sPod', 'k8sPod', 0, 'plugin/k8smanager/view/pod.vue', 2, '{"title": "Pod管理", "icon": "apps-line"}');

-- 1.4 Deployment 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1003, NOW(), NOW(), 1000, 'k8sDeployment', 'k8sDeployment', 0, 'plugin/k8smanager/view/deployment.vue', 3, '{"title": "Deployment管理", "icon": "stack-line"}');

-- 1.5 Service 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1004, NOW(), NOW(), 1000, 'k8sService', 'k8sService', 0, 'plugin/k8smanager/view/service.vue', 4, '{"title": "Service管理", "icon": "links-line"}');

-- 1.6 Namespace 管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1005, NOW(), NOW(), 1000, 'k8sNamespace', 'k8sNamespace', 0, 'plugin/k8smanager/view/namespace.vue', 5, '{"title": "Namespace管理", "icon": "folder-line"}');

-- 1.7 事件管理子菜单
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `meta`)
VALUES
(1006, NOW(), NOW(), 1000, 'k8sEvent', 'k8sEvent', 0, 'plugin/k8smanager/view/event.vue', 6, '{"title": "事件管理", "icon": "notification-line"}');

-- ========================================
-- 2. API 权限数据
-- ========================================

-- 2.1 K8s集群相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2001, NOW(), NOW(), '/k8s/cluster/create', '创建K8s集群', 'K8s集群', 'POST'),
(2002, NOW(), NOW(), '/k8s/cluster/delete', '删除K8s集群', 'K8s集群', 'DELETE'),
(2003, NOW(), NOW(), '/k8s/cluster/deleteByIds', '批量删除K8s集群', 'K8s集群', 'DELETE'),
(2004, NOW(), NOW(), '/k8s/cluster/update', '更新K8s集群', 'K8s集群', 'PUT'),
(2005, NOW(), NOW(), '/k8s/cluster/get', '获取K8s集群详情', 'K8s集群', 'GET'),
(2006, NOW(), NOW(), '/k8s/cluster/list', '获取K8s集群列表', 'K8s集群', 'GET'),
(2007, NOW(), NOW(), '/k8s/cluster/refresh', '刷新K8s集群状态', 'K8s集群', 'POST'),
(2008, NOW(), NOW(), '/k8s/cluster/all', '获取所有K8s集群', 'K8s集群', 'GET');

-- 2.2 Pod相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2009, NOW(), NOW(), '/k8s/pod/list', '获取Pod列表', 'K8s Pod', 'GET'),
(2010, NOW(), NOW(), '/k8s/pod/get', '获取Pod详情', 'K8s Pod', 'GET'),
(2011, NOW(), NOW(), '/k8s/pod/delete', '删除Pod', 'K8s Pod', 'DELETE'),
(2012, NOW(), NOW(), '/k8s/pod/log', '获取Pod日志', 'K8s Pod', 'POST'),
(2013, NOW(), NOW(), '/k8s/pod/containers', '获取Pod容器列表', 'K8s Pod', 'GET'),
(2014, NOW(), NOW(), '/k8s/pod/events', '获取Pod事件', 'K8s Pod', 'GET');

-- 2.3 Deployment相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2015, NOW(), NOW(), '/k8s/deployment/list', '获取Deployment列表', 'K8s Deployment', 'GET'),
(2016, NOW(), NOW(), '/k8s/deployment/get', '获取Deployment详情', 'K8s Deployment', 'GET'),
(2017, NOW(), NOW(), '/k8s/deployment/scale', '扩缩容Deployment', 'K8s Deployment', 'POST'),
(2018, NOW(), NOW(), '/k8s/deployment/restart', '重启Deployment', 'K8s Deployment', 'POST'),
(2019, NOW(), NOW(), '/k8s/deployment/delete', '删除Deployment', 'K8s Deployment', 'DELETE'),
(2020, NOW(), NOW(), '/k8s/deployment/pods', '获取Deployment关联的Pods', 'K8s Deployment', 'GET');

-- 2.4 Service相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2021, NOW(), NOW(), '/k8s/service/list', '获取Service列表', 'K8s Service', 'GET'),
(2022, NOW(), NOW(), '/k8s/service/get', '获取Service详情', 'K8s Service', 'GET'),
(2023, NOW(), NOW(), '/k8s/service/delete', '删除Service', 'K8s Service', 'DELETE'),
(2024, NOW(), NOW(), '/k8s/service/endpoints', '获取Service的Endpoints', 'K8s Service', 'GET');

-- 2.5 Namespace相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2025, NOW(), NOW(), '/k8s/namespace/list', '获取Namespace列表', 'K8s Namespace', 'GET'),
(2026, NOW(), NOW(), '/k8s/namespace/get', '获取Namespace详情', 'K8s Namespace', 'GET'),
(2027, NOW(), NOW(), '/k8s/namespace/create', '创建Namespace', 'K8s Namespace', 'POST'),
(2028, NOW(), NOW(), '/k8s/namespace/delete', '删除Namespace', 'K8s Namespace', 'DELETE');

-- 2.6 Event相关API
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `path`, `description`, `api_group`, `method`)
VALUES
(2029, NOW(), NOW(), '/k8s/event/list', '获取Event列表', 'K8s Event', 'POST');

-- ========================================
-- 3. 菜单与API关联（sys_authority_menus）
-- ========================================
-- 注意：以下 SQL 需要根据实际的 sys_authorities 表中的 authority_id 进行调整
-- 假设 888 为管理员角色ID

-- 3.1 为管理员角色添加 K8s 管理菜单权限
INSERT INTO `sys_authority_menus` (`created_at`, `updated_at`, `menu_id`, `authority_id`)
VALUES
(NOW(), NOW(), 1000, 888),  -- K8s管理父菜单
(NOW(), NOW(), 1001, 888),  -- 集群管理
(NOW(), NOW(), 1002, 888),  -- Pod管理
(NOW(), NOW(), 1003, 888),  -- Deployment管理
(NOW(), NOW(), 1004, 888),  -- Service管理
(NOW(), NOW(), 1005, 888),  -- Namespace管理
(NOW(), NOW(), 1006, 888);  -- 事件管理

-- 3.2 为管理员角色添加所有 API 权限
INSERT INTO `sys_authority_apis` (`created_at`, `updated_at`, `api_id`, `authority_id`)
VALUES
-- K8s集群 API
(NOW(), NOW(), 2001, 888),
(NOW(), NOW(), 2002, 888),
(NOW(), NOW(), 2003, 888),
(NOW(), NOW(), 2004, 888),
(NOW(), NOW(), 2005, 888),
(NOW(), NOW(), 2006, 888),
(NOW(), NOW(), 2007, 888),
(NOW(), NOW(), 2008, 888),
-- Pod API
(NOW(), NOW(), 2009, 888),
(NOW(), NOW(), 2010, 888),
(NOW(), NOW(), 2011, 888),
(NOW(), NOW(), 2012, 888),
(NOW(), NOW(), 2013, 888),
(NOW(), NOW(), 2014, 888),
-- Deployment API
(NOW(), NOW(), 2015, 888),
(NOW(), NOW(), 2016, 888),
(NOW(), NOW(), 2017, 888),
(NOW(), NOW(), 2018, 888),
(NOW(), NOW(), 2019, 888),
(NOW(), NOW(), 2020, 888),
-- Service API
(NOW(), NOW(), 2021, 888),
(NOW(), NOW(), 2022, 888),
(NOW(), NOW(), 2023, 888),
(NOW(), NOW(), 2024, 888),
-- Namespace API
(NOW(), NOW(), 2025, 888),
(NOW(), NOW(), 2026, 888),
(NOW(), NOW(), 2027, 888),
(NOW(), NOW(), 2028, 888),
-- Event API
(NOW(), NOW(), 2029, 888);

-- ========================================
-- 4. 数据表结构（如果需要手动创建）
-- ========================================

CREATE TABLE IF NOT EXISTS `k8s_clusters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
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
  KEY `idx_provider` (`provider`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='K8s集群配置表';

-- ========================================
-- 注意事项
-- ========================================
-- 1. 菜单 ID (1000-1006) 如果与现有菜单冲突，需要调整
-- 2. API ID (2001-2029) 如果与现有API冲突，需要调整
-- 3. authority_id = 888 假设为管理员角色，实际使用时需要根据系统调整
-- 4. parent_id = 0 表示顶级菜单，如需挂载到其他父菜单下，请修改
-- 5. 菜单的 component 路径需要与前端实际文件路径匹配
-- ========================================
