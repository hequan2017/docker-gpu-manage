-- =====================================================
-- 端口转发模块 - 角色授权示例SQL
-- =====================================================
-- 说明：此脚本演示如何为角色分配端口转发模块的权限
-- 使用前请根据实际情况修改角色ID和权限ID
-- =====================================================

-- =====================================================
-- 第一步：查看现有角色和菜单
-- =====================================================

-- 查看所有角色
-- SELECT authority_id, authority_name FROM sys_authorities;

-- 查看端口转发相关菜单
-- SELECT id, parent_id, name, title FROM sys_base_menus WHERE name LIKE 'portForward%';

-- 查看端口转发相关API
-- SELECT id, path, description, api_group FROM sys_apis WHERE api_group = '端口转发';


-- =====================================================
-- 第二步：为角色分配菜单权限
-- =====================================================

-- 示例：为管理员角色（ID=888）分配端口转发菜单权限
-- 注意：实际使用时需要查询获取正确的菜单ID

-- 假设端口转发父菜单ID为 1001，转发规则子菜单ID为 1002
-- INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id, created_at, updated_at)
-- VALUES
--     (888, 1001, NOW(), NOW()),  -- 端口转发父菜单
--     (888, 1002, NOW(), NOW());  -- 转发规则子菜单


-- =====================================================
-- 第三步：为角色分配API权限（使用Casbin）
-- =====================================================

-- 示例：为管理员角色（authority_id=888）分配所有端口转发API权限
-- Casbin规则格式： p, 角色ID, 路径, 方法

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
-- 第四步：查询现有权限配置
-- =====================================================

-- 查询某个角色拥有的菜单权限
-- SELECT m.id, m.parent_id, m.title
-- FROM sys_base_menus m
-- INNER JOIN sys_authority_menus am ON m.id = am.sys_base_menu_id
-- WHERE am.sys_authority_authority_id = 888
-- ORDER BY m.parent_id, m.sort;

-- 查询某个角色拥有的API权限（Casbin）
-- SELECT v0 as authority_id, v1 as path, v2 as method
-- FROM sys_casbin
-- WHERE ptype = 'p' AND v0 = '888' AND v1 LIKE '/portForward%';


-- =====================================================
-- 第五步：清理权限配置（如需要）
-- =====================================================

-- 删除角色的菜单权限
-- DELETE FROM sys_authority_menus
-- WHERE sys_authority_authority_id = 888
-- AND sys_base_menu_id IN (
--     SELECT id FROM sys_base_menus WHERE name LIKE 'portForward%'
-- );

-- 删除角色的API权限
-- DELETE FROM sys_casbin
-- WHERE ptype = 'p'
-- AND v0 = '888'
-- AND v1 LIKE '/portForward%';


-- =====================================================
-- 实际使用建议
-- =====================================================

/*
建议使用系统的后台管理界面进行权限配置：

1. 登录系统
2. 进入"超级管理员" → "角色管理"
3. 选择需要配置的角色
4. 点击"设置权限"按钮
5. 在权限树中勾选"端口转发"相关权限
6. 保存

使用界面操作的优势：
- 自动处理菜单和API的关联
- 自动生成Casbin规则
- 避免手动SQL操作的风险
- 操作更直观、更安全
*/


-- =====================================================
-- 验证配置是否成功
-- =====================================================

-- 验证菜单是否创建成功
SELECT
    id,
    parent_id,
    title,
    path,
    name,
    icon,
    sort
FROM sys_base_menus
WHERE name LIKE 'portForward%'
ORDER BY parent_id, sort;

-- 验证API是否创建成功
SELECT
    id,
    path,
    method,
    description,
    api_group
FROM sys_apis
WHERE api_group = '端口转发'
ORDER BY id;

-- 统计信息
SELECT
    '菜单数量' as type,
    COUNT(*) as count
FROM sys_base_menus
WHERE name LIKE 'portForward%'
UNION ALL
SELECT
    'API数量' as type,
    COUNT(*) as count
FROM sys_apis
WHERE api_group = '端口转发';
