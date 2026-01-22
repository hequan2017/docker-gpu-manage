-- =====================================================
-- 端口转发模块 - 手动初始化SQL脚本
-- =====================================================
-- 说明：如果自动初始化失败，使用此脚本手动创建菜单和API
-- 执行此脚本后，需要为角色分配权限
-- =====================================================

-- =====================================================
-- 第一步：创建菜单
-- =====================================================

-- 1. 创建顶级菜单：端口转发
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level)
VALUES (NOW(), NOW(), 0, '/portForward', 'portForward', 0, 'view/routerHolder.vue', 10, 0);

-- 获取父菜单ID
SET @parent_menu_id = LAST_INSERT_ID();

-- 2. 创建子菜单：转发规则
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level)
VALUES (NOW(), NOW(), @parent_menu_id, 'portForwardRules', 'portForwardRules', 0, 'plugin/portforward/view/portForward.vue', 1, 1);

-- 获取子菜单ID
SET @child_menu_id = LAST_INSERT_ID();

-- 验证菜单创建
SELECT '菜单创建完成' as status;
SELECT id, parent_id, title, path, name FROM sys_base_menus WHERE name LIKE 'portForward%';


-- =====================================================
-- 第二步：创建API
-- =====================================================

INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
    (NOW(), NOW(), '/portForward/createPortForward', '创建端口转发规则', '端口转发', 'POST'),
    (NOW(), NOW(), '/portForward/deletePortForward', '删除端口转发规则', '端口转发', 'DELETE'),
    (NOW(), NOW(), '/portForward/deletePortForwardByIds', '批量删除端口转发规则', '端口转发', 'DELETE'),
    (NOW(), NOW(), '/portForward/updatePortForward', '更新端口转发规则', '端口转发', 'PUT'),
    (NOW(), NOW(), '/portForward/updatePortForwardStatus', '更新端口转发规则状态', '端口转发', 'PUT'),
    (NOW(), NOW(), '/portForward/findPortForward', '根据ID获取端口转发规则', '端口转发', 'GET'),
    (NOW(), NOW(), '/portForward/getPortForwardList', '获取端口转发规则列表', '端口转发', 'GET');

-- 验证API创建
SELECT 'API创建完成' as status;
SELECT id, path, method, description, api_group FROM sys_apis WHERE api_group = '端口转发';


-- =====================================================
-- 第三步：为管理员角色（ID=888）授权
-- =====================================================
-- 注意：如果您的管理员角色ID不是888，请修改下面的值

-- 3.1 为角色分配菜单权限
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id, created_at, updated_at)
VALUES
    (888, @parent_menu_id, NOW(), NOW()),  -- 端口转发父菜单
    (888, @child_menu_id, NOW(), NOW());  -- 转发规则子菜单

-- 3.2 为角色分配API权限（Casbin规则）
INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
    (NULL, 'p', '888', '/portForward/createPortForward', 'POST', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/deletePortForward', 'DELETE', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/deletePortForwardByIds', 'DELETE', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/updatePortForward', 'PUT', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/updatePortForwardStatus', 'PUT', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/findPortForward', 'GET', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/getPortForwardList', 'GET', '', '', '', '');


-- =====================================================
-- 第四步：验证安装
-- =====================================================

-- 4.1 验证菜单
SELECT
    '=== 菜单验证 ===' as info;
SELECT
    id,
    parent_id,
    title,
    path,
    name
FROM sys_base_menus
WHERE name LIKE 'portForward%'
ORDER BY parent_id, sort;

-- 4.2 验证API
SELECT
    '=== API验证 ===' as info;
SELECT
    id,
    path,
    method,
    description
FROM sys_apis
WHERE api_group = '端口转发'
ORDER BY id;

-- 4.3 验证角色权限（菜单）
SELECT
    '=== 角色菜单权限验证 ===' as info;
SELECT
    m.id,
    m.parent_id,
    m.title
FROM sys_base_menus m
INNER JOIN sys_authority_menus am ON m.id = am.sys_base_menu_id
WHERE am.sys_authority_authority_id = 888
  AND m.name LIKE 'portForward%';

-- 4.4 验证角色权限（API）
SELECT
    '=== 角色API权限验证 ===' as info;
SELECT
    v1 as path,
    v2 as method
FROM sys_casbin
WHERE ptype = 'p'
  AND v0 = '888'
  AND v1 LIKE '/portForward%'
ORDER BY v1;

-- 4.5 统计信息
SELECT
    '=== 统计信息 ===' as info;
SELECT
    '菜单总数' as type,
    COUNT(*) as count
FROM sys_base_menus
WHERE name LIKE 'portForward%'
UNION ALL
SELECT
    'API总数' as type,
    COUNT(*) as count
FROM sys_apis
WHERE api_group = '端口转发'
UNION ALL
SELECT
    '已授权菜单数（角色888）' as type,
    COUNT(*) as count
FROM sys_authority_menus am
INNER JOIN sys_base_menus m ON am.sys_base_menu_id = m.id
WHERE am.sys_authority_authority_id = 888
  AND m.name LIKE 'portForward%'
UNION ALL
SELECT
    '已授权API数（角色888）' as type,
    COUNT(*) as count
FROM sys_casbin
WHERE ptype = 'p'
  AND v0 = '888'
  AND v1 LIKE '/portForward%';


-- =====================================================
-- 第五步：创建数据库表（如果还没有创建）
-- =====================================================

-- 检查表是否存在
-- SELECT COUNT(*) FROM information_schema.tables
-- WHERE table_schema = DATABASE()
-- AND table_name = 'gva_port_forward';

-- 如果表不存在，后端启动时会自动创建
-- 或者执行以下SQL手动创建：

/*
CREATE TABLE IF NOT EXISTS `gva_port_forward` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `source_ip` varchar(64) NOT NULL COMMENT '源IP地址',
  `source_port` int NOT NULL COMMENT '源端口',
  `protocol` varchar(10) NOT NULL DEFAULT 'tcp' COMMENT '协议类型: tcp/udp',
  `target_ip` varchar(64) NOT NULL COMMENT '目标IP地址',
  `target_port` int NOT NULL COMMENT '目标端口',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态: 1-启用, 0-禁用',
  `description` varchar(255) DEFAULT '' COMMENT '规则描述',
  PRIMARY KEY (`id`),
  KEY `idx_source` (`source_ip`, `source_port`),
  KEY `idx_protocol` (`protocol`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='端口转发规则表';
*/


-- =====================================================
-- 安装完成提示
-- =====================================================

SELECT
    '========================================' as '';
SELECT
    '✅ 端口转发模块安装完成！' as status;
SELECT
    '========================================' as '';
SELECT
    '接下来的步骤：' as info;
SELECT
    '1. 清除浏览器缓存' as step1;
SELECT
    '2. 退出登录' as step2;
SELECT
    '3. 重新登录' as step3;
SELECT
    '4. 在左侧菜单查看"端口转发"' as step4;
SELECT
    '========================================' as '';


-- =====================================================
-- 如果需要重新安装，先执行清理
-- =====================================================

/*
⚠️ 警告：以下操作会删除所有数据，请谨慎使用！

-- 删除角色权限
DELETE FROM sys_authority_menus
WHERE sys_base_menu_id IN (
    SELECT id FROM sys_base_menus WHERE name LIKE 'portForward%'
);

DELETE FROM sys_casbin
WHERE v1 LIKE '/portForward%';

-- 删除菜单
DELETE FROM sys_base_menus WHERE name LIKE 'portForward%';

-- 删除API
DELETE FROM sys_apis WHERE api_group = '端口转发';

-- 删除数据表（可选）
DROP TABLE IF EXISTS gva_port_forward;

执行完清理后，重新运行本脚本
*/
