-- =====================================================
-- 戴尔资产管理插件 - 初始化 SQL 脚本
-- =====================================================
-- 执行说明：
-- 1. 本脚本根据 gin-vue-admin 实际表结构调整
-- 2. 包含数据表、菜单、API、权限的完整初始化
-- =====================================================

-- =====================================================
-- 第1步：创建数据表
-- =====================================================

CREATE TABLE IF NOT EXISTS `gva_dell_asset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `host_name` varchar(100) NOT NULL COMMENT '主机名',
  `service_tag` varchar(20) NOT NULL COMMENT '服务标签(戴尔唯一标识)',
  `asset_number` varchar(50) DEFAULT NULL COMMENT '资产编号',
  `model` varchar(100) DEFAULT NULL COMMENT '型号',
  `serial_number` varchar(50) DEFAULT NULL COMMENT '序列号',
  `cpu_model` varchar(100) DEFAULT NULL COMMENT 'CPU型号',
  `cpu_cores` int DEFAULT 0 COMMENT 'CPU核心数',
  `cpu_threads` int DEFAULT 0 COMMENT 'CPU线程数',
  `memory_capacity` int DEFAULT 0 COMMENT '内存容量(GB)',
  `disk_info` varchar(500) DEFAULT NULL COMMENT '磁盘信息',
  `network_info` varchar(500) DEFAULT NULL COMMENT '网卡信息',
  `ip_address` varchar(50) DEFAULT NULL COMMENT 'IP地址',
  `mac_address` varchar(50) DEFAULT NULL COMMENT 'MAC地址',
  `cabinet` varchar(50) DEFAULT NULL COMMENT '机柜位置',
  `rack_position` varchar(20) DEFAULT NULL COMMENT '机架位置(U位)',
  `power_status` varchar(20) DEFAULT 'offline' COMMENT '电源状态',
  `purchase_date` varchar(20) DEFAULT NULL COMMENT '购买日期',
  `warranty_expiry` varchar(20) DEFAULT NULL COMMENT '保修到期日',
  `os` varchar(100) DEFAULT NULL COMMENT '操作系统',
  `department` varchar(100) DEFAULT NULL COMMENT '所属部门',
  `manager` varchar(50) DEFAULT NULL COMMENT '负责人',
  `status` varchar(20) DEFAULT 'offline' COMMENT '状态',
  `remarks` text COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_service_tag` (`service_tag`),
  KEY `idx_host_name` (`host_name`),
  KEY `idx_ip_address` (`ip_address`),
  KEY `idx_status` (`status`),
  KEY `idx_department` (`department`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='戴尔服务器资产表';

-- =====================================================
-- 第2步：创建菜单
-- =====================================================

-- 2.1 创建戴尔资产管理菜单（作为资源管理的子菜单，parent_id=3 表示资源管理）
-- 注意：如果资源管理的ID不是3，请根据实际情况调整
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), 3, 'dellAsset', 'dellAsset', 0, 'plugin/dellasset/view/dellAsset.vue', 10, 1, '戴尔资产管理', 'cpu');

-- 获取菜单ID（用于后续权限配置）
SET @dell_asset_menu_id = LAST_INSERT_ID();

-- =====================================================
-- 第3步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
-- 戴尔资产基础API
(NOW(), NOW(), '/dellAsset/createDellAsset', '创建戴尔服务器资产', '戴尔资产', 'POST'),
(NOW(), NOW(), '/dellAsset/deleteDellAsset', '删除戴尔服务器资产', '戴尔资产', 'DELETE'),
(NOW(), NOW(), '/dellAsset/deleteDellAssetByIds', '批量删除戴尔服务器资产', '戴尔资产', 'DELETE'),
(NOW(), NOW(), '/dellAsset/updateDellAsset', '更新戴尔服务器资产', '戴尔资产', 'PUT'),
(NOW(), NOW(), '/dellAsset/findDellAsset', '查询戴尔服务器资产', '戴尔资产', 'GET'),
(NOW(), NOW(), '/dellAsset/getDellAssetList', '获取戴尔服务器资产列表', '戴尔资产', 'GET'),
(NOW(), NOW(), '/dellAsset/getStatistics', '获取资产统计信息', '戴尔资产', 'GET');

-- =====================================================
-- 第4步：为管理员角色授权菜单（authority_id = 888）
-- =====================================================

-- 插入菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES (888, @dell_asset_menu_id);

-- =====================================================
-- 第5步：为管理员角色授权API（使用 sys_casbin 表）
-- =====================================================

INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
-- 戴尔资产API权限
(NULL, 'p', '888', '/dellAsset/createDellAsset', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/deleteDellAsset', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/deleteDellAssetByIds', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/updateDellAsset', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/findDellAsset', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/getDellAssetList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/dellAsset/getStatistics', 'GET', '', '', '', '');

-- =====================================================
-- 第6步：为菜单按钮添加API权限关联
-- =====================================================

-- 获取API ID
SET @api_create = (SELECT id FROM sys_apis WHERE path = '/dellAsset/createDellAsset' AND method = 'POST');
SET @api_update = (SELECT id FROM sys_apis WHERE path = '/dellAsset/updateDellAsset' AND method = 'PUT');
SET @api_delete = (SELECT id FROM sys_apis WHERE path = '/dellAsset/deleteDellAsset' AND method = 'DELETE');

-- 为管理员角色添加按钮API权限（通过sys_authority_menus的menu_id关联到API）
-- 注意：gin-vue-admin中按钮权限通常通过菜单名称关联API
-- 这里创建按钮菜单项
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES
(NOW(), NOW(), @dell_asset_menu_id, '', 'create', 1, '', 1, 2, '新增', ''),
(NOW(), NOW(), @dell_asset_menu_id, '', 'update', 1, '', 2, 2, '编辑', ''),
(NOW(), NOW(), @dell_asset_menu_id, '', 'delete', 1, '', 3, 2, '删除', '');

-- 获取按钮菜单ID并授权
SET @btn_create = (SELECT id FROM sys_base_menus WHERE name = 'create' AND parent_id = @dell_asset_menu_id);
SET @btn_update = (SELECT id FROM sys_base_menus WHERE name = 'update' AND parent_id = @dell_asset_menu_id);
SET @btn_delete = (SELECT id FROM sys_base_menus WHERE name = 'delete' AND parent_id = @dell_asset_menu_id);

-- 为按钮添加菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES
(888, @btn_create),
(888, @btn_update),
(888, @btn_delete);

-- =====================================================
-- 第7步：验证安装
-- =====================================================

SELECT '✅ 戴尔资产管理插件 SQL 执行完成！' as status;
SELECT COUNT(*) as '数据表数量' FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'gva_dell_asset';
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name = 'dellAsset' OR parent_id IN (SELECT id FROM sys_base_menus WHERE name = 'dellAsset');
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group = '戴尔资产';
SELECT COUNT(*) as '菜单权限数' FROM sys_authority_menus WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name = 'dellAsset' OR parent_id IN (SELECT id FROM sys_base_menus WHERE name = 'dellAsset')
);
SELECT COUNT(*) as 'API权限数' FROM sys_casbin WHERE v1 LIKE '/dellAsset%' AND v0 = '888';

-- =====================================================
-- 回滚SQL（如需删除，请谨慎使用）
-- =====================================================

-- 删除API权限
-- DELETE FROM sys_casbin WHERE v1 LIKE '/dellAsset%' AND v0 = '888';

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name = 'dellAsset' OR parent_id IN (SELECT id FROM sys_base_menus WHERE name = 'dellAsset')
-- );

-- 删除按钮菜单
-- DELETE FROM sys_base_menus WHERE parent_id IN (SELECT id FROM sys_base_menus WHERE name = 'dellAsset') AND name IN ('create', 'update', 'delete');

-- 删除主菜单
-- DELETE FROM sys_base_menus WHERE name = 'dellAsset';

-- 删除API
-- DELETE FROM sys_apis WHERE api_group = '戴尔资产';

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS gva_dell_asset;
