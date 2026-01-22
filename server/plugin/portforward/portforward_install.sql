-- =====================================================
-- 端口转发插件 - 初始化 SQL 脚本
-- =====================================================
-- 执行说明：
-- 1. 本脚本根据 gin-vue-admin 实际表结构调整
-- 2. 包含数据表、菜单、API、权限的完整初始化
-- 3. 支持TCP/UDP端口转发功能
-- =====================================================

-- =====================================================
-- 第1步：创建数据表
-- =====================================================

-- 1.1 创建端口转发规则表
CREATE TABLE IF NOT EXISTS `gva_port_forward` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `source_ip` varchar(64) NOT NULL COMMENT '源IP地址',
  `source_port` int NOT NULL COMMENT '源端口',
  `protocol` varchar(10) NOT NULL DEFAULT 'tcp' COMMENT '协议类型: tcp/udp',
  `target_ip` varchar(64) NOT NULL COMMENT '目标IP地址',
  `target_port` int NOT NULL COMMENT '目标端口',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态: 1-启用, 0-禁用',
  `description` varchar(255) DEFAULT NULL COMMENT '规则描述',
  PRIMARY KEY (`id`),
  KEY `idx_source_ip` (`source_ip`),
  KEY `idx_source_port` (`source_port`),
  KEY `idx_protocol` (`protocol`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='端口转发规则表';

-- =====================================================
-- 第2步：创建菜单
-- =====================================================

-- 2.1 创建端口转发主菜单（作为顶级菜单）
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), 0, '/portForward', 'portForward', 0, 'view/routerHolder.vue', 10, 0, '端口转发', 'position');

-- 获取主菜单ID
SET @portforward_menu_id = LAST_INSERT_ID();

-- 2.2 创建子菜单
-- 转发规则管理
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level, title, icon)
VALUES (NOW(), NOW(), @portforward_menu_id, 'portForwardRules', 'portForwardRules', 0, 'plugin/portforward/view/portForward.vue', 1, 0, '转发规则', 'list');

SET @rules_menu_id = LAST_INSERT_ID();

-- =====================================================
-- 第3步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
-- 端口转发规则管理API
(NOW(), NOW(), '/portForward/createPortForward', '创建端口转发规则', 'PortForward', 'POST'),
(NOW(), NOW(), '/portForward/deletePortForward', '删除端口转发规则', 'PortForward', 'DELETE'),
(NOW(), NOW(), '/portForward/deletePortForwardByIds', '批量删除端口转发规则', 'PortForward', 'DELETE'),
(NOW(), NOW(), '/portForward/updatePortForward', '更新端口转发规则', 'PortForward', 'PUT'),
(NOW(), NOW(), '/portForward/updatePortForwardStatus', '更新端口转发规则状态', 'PortForward', 'PUT'),
(NOW(), NOW(), '/portForward/findPortForward', '根据ID获取端口转发规则', 'PortForward', 'GET'),
(NOW(), NOW(), '/portForward/getPortForwardList', '获取端口转发规则列表', 'PortForward', 'GET'),
(NOW(), NOW(), '/portForward/getServerIP', '获取服务器IP地址', 'PortForward', 'GET'),
(NOW(), NOW(), '/portForward/getForwarderStatus', '获取端口转发运行状态', 'PortForward', 'GET'),
(NOW(), NOW(), '/portForward/getAllForwarderStatus', '获取所有端口转发运行状态', 'PortForward', 'GET');

-- =====================================================
-- 第4步：为管理员角色授权菜单（authority_id = 888）
-- =====================================================

-- 插入菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id)
VALUES
(888, @portforward_menu_id),
(888, @rules_menu_id);

-- =====================================================
-- 第5步：为管理员角色授权API（使用 sys_casbin 表）
-- =====================================================

INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
-- 端口转发规则管理API权限
(NULL, 'p', '888', '/portForward/createPortForward', 'POST', '', '', '', ''),
(NULL, 'p', '888', '/portForward/deletePortForward', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/portForward/deletePortForwardByIds', 'DELETE', '', '', '', ''),
(NULL, 'p', '888', '/portForward/updatePortForward', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/portForward/updatePortForwardStatus', 'PUT', '', '', '', ''),
(NULL, 'p', '888', '/portForward/findPortForward', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/portForward/getPortForwardList', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/portForward/getServerIP', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/portForward/getForwarderStatus', 'GET', '', '', '', ''),
(NULL, 'p', '888', '/portForward/getAllForwarderStatus', 'GET', '', '', '', '');

-- =====================================================
-- 第6步：验证安装
-- =====================================================

SELECT '✅ 端口转发插件 SQL 执行完成！' as status;
SELECT COUNT(*) as '数据表数量' FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name LIKE 'gva_port_forward';
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name LIKE 'portforward%';
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group = 'PortForward';
SELECT COUNT(*) as '菜单权限数' FROM sys_authority_menus WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name LIKE 'portforward%'
);
SELECT COUNT(*) as 'API权限数' FROM sys_casbin WHERE v1 LIKE '/portForward%' AND v0 = '888';

-- =====================================================
-- 使用说明
-- =====================================================

SELECT '📝 使用说明：' as info;
SELECT '1. 进入【端口转发】->【转发规则】页面管理转发规则' as step1;
SELECT '2. 支持TCP和UDP协议的端口转发' as step2;
SELECT '3. 点击"获取服务器IP"可以快速获取本机可用的IP地址' as step3;
SELECT '4. 支持批量删除和状态快速切换' as step4;

-- =====================================================
-- 回滚SQL（如需删除，请谨慎使用）
-- =====================================================

-- 删除API权限
-- DELETE FROM sys_casbin WHERE v1 LIKE '/portForward%' AND v0 = '888';

-- 删除菜单权限
-- DELETE FROM sys_authority_menus WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'portforward%'
-- );

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE name LIKE 'portforward%';

-- 删除API
-- DELETE FROM sys_apis WHERE api_group = 'PortForward';

-- 删除数据表（慎用）
-- DROP TABLE IF EXISTS gva_port_forward;
