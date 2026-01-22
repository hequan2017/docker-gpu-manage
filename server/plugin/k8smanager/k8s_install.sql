-- =====================================================
-- K8s 管理插件 - 正确的 SQL 脚本
-- =====================================================
-- 执行说明：
-- 1. 本脚本已根据 gin-vue-admin 实际表结构调整
-- 2. meta 字段是嵌入式，需要分别存储 title 和 icon
-- 3. API 权限使用 sys_casbin 表，不是 sys_authority_apis
-- =====================================================

-- =====================================================
-- 第1步：创建数据表
-- =====================================================

CREATE TABLE IF NOT EXISTS `k8s_clusters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '集群名称',
  `kube_config` longtext COMMENT 'kubeconfig配置内容',
  `endpoint` varchar(500) DEFAULT NULL COMMENT 'API Server地址',
  `version` varchar(50) DEFAULT NULL COMMENT 'K8s版本',
  `status` varchar(20) DEFAULT 'unknown' COMMENT '集群状态',
  `description` varchar(500) DEFAULT NULL COMMENT '集群描述',
  `region` varchar(100) DEFAULT NULL COMMENT '区域',
  `provider` varchar(50) DEFAULT NULL COMMENT '云服务商',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认集群',
  `node_count` int DEFAULT '0' COMMENT '节点数量',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='K8s集群配置表';

-- =====================================================
-- 第2步：创建菜单
-- =====================================================

-- 2.1 创建 K8s 管理父菜单（menu_level = 0 表示顶级菜单）
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), 0, 'k8s', 'k8s', 0, 'view/routerHolder.vue', 8, 0, 'K8s管理', 'cpu-line');

-- 获取父菜单ID
SET @k8s_parent_id = LAST_INSERT_ID();

-- 2.2 创建集群管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sCluster', 'k8sCluster', 0, 'plugin/k8smanager/view/cluster.vue', 1, 1, '集群管理', 'server-line');

-- 2.3 创建 Pod 管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sPod', 'k8sPod', 0, 'plugin/k8smanager/view/pod.vue', 2, 1, 'Pod管理', 'apps-line');

-- 2.4 创建 Deployment 管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sDeployment', 'k8sDeployment', 0, 'plugin/k8smanager/view/deployment.vue', 3, 1, 'Deployment管理', 'stack-line');

-- 2.5 创建 Service 管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sService', 'k8sService', 0, 'plugin/k8smanager/view/service.vue', 4, 1, 'Service管理', 'links-line');

-- 2.6 创建 Namespace 管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sNamespace', 'k8sNamespace', 0, 'plugin/k8smanager/view/namespace.vue', 5, 1, 'Namespace管理', 'folder-line');

-- 2.7 创建事件管理子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @k8s_parent_id, 'k8sEvent', 'k8sEvent', 0, 'plugin/k8smanager/view/event.vue', 6, 1, '事件管理', 'notification-line');

-- =====================================================
-- 第3步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
-- K8s集群相关API (8个)
(NOW(), NOW(), '/k8s/cluster/create', '创建K8s集群', 'K8s集群', 'POST'),
(NOW(), NOW(), '/k8s/cluster/delete', '删除K8s集群', 'K8s集群', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/deleteByIds', '批量删除K8s集群', 'K8s集群', 'DELETE'),
(NOW(), NOW(), '/k8s/cluster/update', '更新K8s集群', 'K8s集群', 'PUT'),
(NOW(), NOW(), '/k8s/cluster/get', '获取K8s集群详情', 'K8s集群', 'GET'),
(NOW(), NOW(), '/k8s/cluster/list', '获取K8s集群列表', 'K8s集群', 'GET'),
(NOW(), NOW(), '/k8s/cluster/refresh', '刷新K8s集群状态', 'K8s集群', 'POST'),
(NOW(), NOW(), '/k8s/cluster/all', '获取所有K8s集群', 'K8s集群', 'GET'),

-- Pod相关API (6个)
(NOW(), NOW(), '/k8s/pod/list', '获取Pod列表', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/get', '获取Pod详情', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/delete', '删除Pod', 'K8s Pod', 'DELETE'),
(NOW(), NOW(), '/k8s/pod/log', '获取Pod日志', 'K8s Pod', 'POST'),
(NOW(), NOW(), '/k8s/pod/containers', '获取Pod容器列表', 'K8s Pod', 'GET'),
(NOW(), NOW(), '/k8s/pod/events', '获取Pod事件', 'K8s Pod', 'GET'),

-- Deployment相关API (6个)
(NOW(), NOW(), '/k8s/deployment/list', '获取Deployment列表', 'K8s Deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/get', '获取Deployment详情', 'K8s Deployment', 'GET'),
(NOW(), NOW(), '/k8s/deployment/scale', '扩缩容Deployment', 'K8s Deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/restart', '重启Deployment', 'K8s Deployment', 'POST'),
(NOW(), NOW(), '/k8s/deployment/delete', '删除Deployment', 'K8s Deployment', 'DELETE'),
(NOW(), NOW(), '/k8s/deployment/pods', '获取Deployment关联的Pods', 'K8s Deployment', 'GET'),

-- Service相关API (4个)
(NOW(), NOW(), '/k8s/service/list', '获取Service列表', 'K8s Service', 'GET'),
(NOW(), NOW(), '/k8s/service/get', '获取Service详情', 'K8s Service', 'GET'),
(NOW(), NOW(), '/k8s/service/delete', '删除Service', 'K8s Service', 'DELETE'),
(NOW(), NOW(), '/k8s/service/endpoints', '获取Service的Endpoints', 'K8s Service', 'GET'),

-- Namespace相关API (4个)
(NOW(), NOW(), '/k8s/namespace/list', '获取Namespace列表', 'K8s Namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/get', '获取Namespace详情', 'K8s Namespace', 'GET'),
(NOW(), NOW(), '/k8s/namespace/create', '创建Namespace', 'K8s Namespace', 'POST'),
(NOW(), NOW(), '/k8s/namespace/delete', '删除Namespace', 'K8s Namespace', 'DELETE'),

-- Event相关API (1个)
(NOW(), NOW(), '/k8s/event/list', '获取Event列表', 'K8s Event', 'POST');

-- =====================================================
-- 第4步：为管理员角色授权菜单（authority_id = 888）
-- =====================================================

-- 先获取所有K8s相关菜单的ID
SET @menu1 = (SELECT id FROM sys_base_menus WHERE name = 'k8s' AND parent_id = 0);
SET @menu2 = (SELECT id FROM sys_base_menus WHERE name = 'k8sCluster' AND parent_id != 0);
SET @menu3 = (SELECT id FROM sys_base_menus WHERE name = 'k8sPod' AND parent_id != 0);
SET @menu4 = (SELECT id FROM sys_base_menus WHERE name = 'k8sDeployment' AND parent_id != 0);
SET @menu5 = (SELECT id FROM sys_base_menus WHERE name = 'k8sService' AND parent_id != 0);
SET @menu6 = (SELECT id FROM sys_base_menus WHERE name = 'k8sNamespace' AND parent_id != 0);
SET @menu7 = (SELECT id FROM sys_base_menus WHERE name = 'k8sEvent' AND parent_id != 0);

-- 插入菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES
(888, @menu1),
(888, @menu2),
(888, @menu3),
(888, @menu4),
(888, @menu5),
(888, @menu6),
(888, @menu7);

-- =====================================================
-- 第5步：为管理员角色授权API（使用 sys_casbin 表）
-- =====================================================

INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
-- K8s集群 API
(NULL, 'p', '888', '/k8s/cluster/create', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/delete', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/deleteByIds', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/update', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/get', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/list', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/refresh', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/cluster/all', 'GET', '', '', '', ''),

-- Pod API
(NULL, 'p', '888', '/k8s/pod/list', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/pod/get', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/pod/delete', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/k8s/pod/log', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/pod/containers', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/pod/events', 'GET', '', '', '', ''),

-- Deployment API
(NULL, 'p', '888', '/k8s/deployment/list', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/deployment/get', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/deployment/scale', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/deployment/restart', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/deployment/delete', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/k8s/deployment/pods', 'GET', '', '', '', ''),

-- Service API
(NULL, 'p', '888', '/k8s/service/list', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/service/get', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/service/delete', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/k8s/service/endpoints', 'GET', '', '', '', ''),

-- Namespace API
(NULL, 'p', '888', '/k8s/namespace/list', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/namespace/get', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/k8s/namespace/create', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/k8s/namespace/delete', 'DELETE', '', '', '', ''),

-- Event API
(NULL, 'p', '888', '/k8s/event/list', 'POST', '', '', '', '');

-- =====================================================
-- 第6步：验证安装
-- =====================================================

SELECT '✅ K8s 管理插件 SQL 执行完成！' as status;
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name LIKE 'k8s%';
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group LIKE 'K8s%';
SELECT COUNT(*) as '菜单权限数' FROM sys_authority_menus WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name LIKE 'k8s%'
);
SELECT COUNT(*) as 'API权限数' FROM sys_casbin WHERE v1 LIKE '/k8s%' AND v0 = '888';

-- =====================================================
-- 回滚SQL（如需删除，请谨慎使用）
-- =====================================================

-- 删除API权限
-- DELETE FROM sys_casbin WHERE v1 LIKE '/k8s%' AND v0 = '888';

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'k8s%'
-- );

-- 删除API
-- DELETE FROM sys_apis WHERE api_group LIKE 'K8s%';

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE name LIKE 'k8s%';

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS k8s_clusters;
