-- =====================================================
-- 端口转发模块 - 快速验证和授权脚本
-- =====================================================
-- 说明：此脚本用于验证端口转发模块是否正确安装
-- 并为管理员角色授权
-- =====================================================

-- =====================================================
-- 第一步：验证菜单是否创建成功
-- =====================================================

-- 查看端口转发菜单
SELECT
    '=== 菜单验证 ===' as info;
SELECT
    id,
    parent_id,
    title,
    path,
    name,
    component,
    sort,
    hidden
FROM sys_base_menus
WHERE name LIKE 'portForward%'
ORDER BY parent_id, sort;

-- 期望结果：2条记录
-- 1. 端口转发（父菜单，parent_id=0）
-- 2. 转发规则（子菜单，parent_id=父菜单ID）


-- =====================================================
-- 第二步：验证API是否创建成功
-- =====================================================

SELECT
    '=== API验证 ===' as info;
SELECT
    id,
    path,
    method,
    description,
    api_group
FROM sys_apis
WHERE api_group = '端口转发'
ORDER BY id;

-- 期望结果：7条记录
-- 所有端口转发相关的API


-- =====================================================
-- 第三步：查看当前所有角色
-- =====================================================

SELECT
    '=== 角色列表 ===' as info;
SELECT
    authority_id,
    authority_name,
    created_at
FROM sys_authorities
ORDER BY authority_id;


-- =====================================================
-- 第四步：为管理员角色（ID=888）授权
-- =====================================================
-- 说明：如果您的管理员角色ID不是888，请先执行上面的查询获取正确ID
-- 然后修改下面的SQL中的角色ID

SELECT
    '=== 开始授权（管理员角色ID=888）===' as info;

-- 4.1 获取菜单ID
-- 注意：这些SELECT语句用于获取菜单ID，请记录返回的ID值用于后续操作

-- 获取端口转发父菜单ID
-- SELECT id FROM sys_base_menus WHERE name = 'portForward' AND parent_id = 0;

-- 获取转发规则子菜单ID
-- SELECT id FROM sys_base_menus WHERE name = 'portForwardRules' AND parent_id != 0;


-- 4.2 为角色分配菜单权限
-- 请将下方 1001 和 1002 替换为上面查询到的实际菜单ID

-- 检查是否已存在菜单权限
SELECT
    '=== 检查现有菜单权限 ===' as info;
SELECT
    am.sys_authority_authority_id,
    am.sys_base_menu_id,
    m.title
FROM sys_authority_menus am
LEFT JOIN sys_base_menus m ON am.sys_base_menu_id = m.id
WHERE am.sys_authority_authority_id = 888
  AND m.name LIKE 'portForward%';

-- 如果上面的查询没有结果，执行下面的INSERT语句
-- 注意：替换 1001 和 1002 为实际菜单ID

-- INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id, created_at, updated_at)
-- VALUES
--     (888, 1001, NOW(), NOW()),  -- 端口转发父菜单
--     (888, 1002, NOW(), NOW());  -- 转发规则子菜单


-- 4.3 为角色分配API权限（Casbin规则）
SELECT
    '=== 检查现有API权限 ===' as info;
SELECT
    v0 as authority_id,
    v1 as path,
    v2 as method
FROM sys_casbin
WHERE ptype = 'p'
  AND v0 = '888'
  AND v1 LIKE '/portForward%';

-- 如果上面的查询没有结果，执行下面的INSERT语句
-- INSERT INTO sys_casbin (id, ptype, v0, v1, v2, v3, v4, v5)
-- VALUES
--     (NULL, 'p', '888', '/portForward/createPortForward', 'POST', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/deletePortForward', 'DELETE', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/deletePortForwardByIds', 'DELETE', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/updatePortForward', 'PUT', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/updatePortForwardStatus', 'PUT', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/findPortForward', 'GET', '', '', '', ''),
--     (NULL, 'p', '888', '/portForward/getPortForwardList', 'GET', '', '', '', '');


-- =====================================================
-- 第五步：验证授权是否成功
-- =====================================================

SELECT
    '=== 验证菜单授权 ===' as info;
SELECT
    m.id,
    m.parent_id,
    m.title,
    m.path
FROM sys_base_menus m
INNER JOIN sys_authority_menus am ON m.id = am.sys_base_menu_id
WHERE am.sys_authority_authority_id = 888
  AND m.name LIKE 'portForward%'
ORDER BY m.parent_id, m.sort;

SELECT
    '=== 验证API授权 ===' as info;
SELECT
    v1 as path,
    v2 as method
FROM sys_casbin
WHERE ptype = 'p'
  AND v0 = '888'
  AND v1 LIKE '/portForward%'
ORDER BY v1;


-- =====================================================
-- 第六步：统计信息
-- =====================================================

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
-- 故障排查SQL
-- =====================================================

-- 如果看不到菜单，检查以下内容：

-- 1. 检查父菜单是否存在
SELECT
    '=== 检查父菜单 ===' as info;
SELECT
    id,
    parent_id,
    title,
    path
FROM sys_base_menus
WHERE name = 'portForward';

-- 2. 检查子菜单是否存在
SELECT
    '=== 检查子菜单 ===' as info;
SELECT
    id,
    parent_id,
    title,
    path
FROM sys_base_menus
WHERE name = 'portForwardRules';

-- 3. 检查当前登录用户的角色
-- SELECT * FROM sys_users WHERE nick_name = '当前用户名';

-- 4. 检查该用户是否有菜单权限
-- SELECT u.nick_name, m.title
-- FROM sys_users u
-- INNER JOIN sys_user_authorities ua ON u.uuid = ua.sys_user_authority_id
-- INNER JOIN sys_authorities a ON ua.sys_authority_authority_id = a.authority_id
-- INNER JOIN sys_authority_menus am ON a.authority_id = am.sys_authority_authority_id
-- INNER JOIN sys_base_menus m ON am.sys_base_menu_id = m.id
-- WHERE u.nick_name = '当前用户名'
--   AND m.name LIKE 'portForward%';


-- =====================================================
-- 快速修复命令（如果菜单或API不存在）
-- =====================================================

-- 如果菜单不存在，说明插件未正确初始化
-- 解决方法：
-- 1. 检查 server/main.go 中是否注册了插件
-- 2. 检查 server/initialize/plugin_biz_v2.go 中是否有 portforward.Plugin
-- 3. 重启服务器
-- 4. 菜单会自动创建


-- =====================================================
-- 手动清理（如需重新初始化）
-- =====================================================

-- ⚠️ 警告：以下操作会删除所有端口转发相关的数据，请谨慎使用！

-- 删除菜单权限
-- DELETE FROM sys_authority_menus
-- WHERE sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'portForward%'
-- );

-- 删除API权限
-- DELETE FROM sys_casbin
-- WHERE v1 LIKE '/portForward%';

-- 删除菜单
-- DELETE FROM sys_base_menus WHERE name LIKE 'portForward%';

-- 删除API
-- DELETE FROM sys_apis WHERE api_group = '端口转发';

-- 删除数据表
-- DROP TABLE IF EXISTS gva_port_forward;

-- 执行完后，重启服务器，所有数据会重新创建


-- =====================================================
-- 使用说明
-- =====================================================

/*
1. 验证阶段：
   - 执行"第一步"到"第三步"的查询
   - 确认菜单和API已创建
   - 记录需要授权的角色ID

2. 授权阶段：
   - 将SQL中的角色ID（888）替换为实际角色ID
   - 执行"第四步"的INSERT语句
   - 或使用系统后台界面进行授权

3. 验证授权：
   - 执行"第五步"的验证查询
   - 确认权限已正确分配

4. 前端验证：
   - 清除浏览器缓存
   - 退出登录
   - 重新登录
   - 在左侧菜单查找"端口转发"

5. 如果仍然看不到菜单：
   - 检查服务是否重启
   - 查看服务器日志
   - 执行"故障排查SQL"
*/
