-- =====================================================
-- 端口转发模块 - 快速创建菜单和API
-- =====================================================
-- 请立即执行此SQL脚本！
-- =====================================================

-- 第1步：创建菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level)
VALUES (NOW(), NOW(), 0, '/portForward', 'portForward', 0, 'view/routerHolder.vue', 10, 0);

-- 获取父菜单ID（假设是新的ID）
SET @parent_id = LAST_INSERT_ID();

-- 第2步：创建子菜单
INSERT INTO sys_base_menus (created_at, updated_at, parent_id, path, name, hidden, component, sort, menu_level)
VALUES (NOW(), NOW(), @parent_id, 'portForwardRules', 'portForwardRules', 0, 'plugin/portforward/view/portForward.vue', 1, 1);

-- 第3步：创建API
INSERT INTO sys_apis (created_at, updated_at, path, description, api_group, method)
VALUES
    (NOW(), NOW(), '/portForward/createPortForward', '创建端口转发规则', '端口转发', 'POST'),
    (NOW(), NOW(), '/portForward/deletePortForward', '删除端口转发规则', '端口转发', 'DELETE'),
    (NOW(), NOW(), '/portForward/deletePortForwardByIds', '批量删除端口转发规则', '端口转发', 'DELETE'),
    (NOW(), NOW(), '/portForward/updatePortForward', '更新端口转发规则', '端口转发', 'PUT'),
    (NOW(), NOW(), '/portForward/updatePortForwardStatus', '更新端口转发规则状态', '端口转发', 'PUT'),
    (NOW(), NOW(), '/portForward/findPortForward', '根据ID获取端口转发规则', '端口转发', 'GET'),
    (NOW(), NOW(), '/portForward/getPortForwardList', '获取端口转发规则列表', '端口转发', 'GET');

-- 第4步：为管理员角色授权（假设角色ID是888）
-- 先获取菜单ID
SET @menu1_id = (SELECT id FROM sys_base_menus WHERE name = 'portForward' AND parent_id = 0);
SET @menu2_id = (SELECT id FROM sys_base_menus WHERE name = 'portForwardRules' AND parent_id != 0);

-- 授权菜单
INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id, created_at, updated_at)
VALUES
    (888, @menu1_id, NOW(), NOW()),
    (888, @menu2_id, NOW(), NOW());

-- 授权API
INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
VALUES
    (NULL, 'p', '888', '/portForward/createPortForward', 'POST', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/deletePortForward', 'DELETE', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/deletePortForwardByIds', 'DELETE', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/updatePortForward', 'PUT', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/updatePortForwardStatus', 'PUT', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/findPortForward', 'GET', '', '', '', ''),
    (NULL, 'p', '888', '/portForward/getPortForwardList', 'GET', '', '', '', '');

-- 验证
SELECT '✅ 创建完成！' as status;
SELECT COUNT(*) as '菜单数量' FROM sys_base_menus WHERE name LIKE 'portForward%';
SELECT COUNT(*) as 'API数量' FROM sys_apis WHERE api_group = '端口转发';
