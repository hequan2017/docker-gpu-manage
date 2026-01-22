-- =====================================================
-- 天启算力管理平台 - 菜单和API权限SQL
-- =====================================================
-- 说明：此SQL包含"项目说明"菜单的完整配置
-- 使用方法：在数据库中直接执行此SQL即可
-- =====================================================

-- =====================================================
-- 第一部分：菜单数据
-- =====================================================

-- 1. 添加"项目说明"父级菜单
INSERT INTO `sys_base_menus` (
    `created_at`,
    `updated_at`,
    `menu_level`,
    `parent_id`,
    `path`,
    `name`,
    `hidden`,
    `component`,
    `sort`,
    `title`
) VALUES (
    NOW(),
    NOW(),
    0,
    0,
    'project',
    'project',
    0,
    'view/about/project.vue',
    2,
    '项目说明'
);

-- 获取刚才插入的菜单ID（用于后续权限配置）
-- SET @project_menu_id = LAST_INSERT_ID();

-- =====================================================
-- 第二部分：为角色分配菜单权限
-- =====================================================

-- 1. 为超级管理员(role_id = 888)分配"项目说明"菜单权限
-- 注意：超级管理员默认拥有所有权限，通常不需要单独配置
-- 如果需要，可以执行以下SQL：
-- INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
-- SELECT 888, id, NOW(), NOW() FROM `sys_base_menus` WHERE `name` = 'project';

-- 2. 为普通用户(role_id = 8881)分配"项目说明"菜单权限
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
SELECT 8881, id, NOW(), NOW()
FROM `sys_base_menus`
WHERE `name` = 'project'
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus`
    WHERE `sys_authority_authority_id` = 8881
    AND `sys_base_menu_id` = (SELECT id FROM `sys_base_menus` WHERE `name` = 'project')
);

-- 3. 为测试角色(role_id = 9528)分配"项目说明"菜单权限
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`, `sys_base_menu_id`, `created_at`, `updated_at`)
SELECT 9528, id, NOW(), NOW()
FROM `sys_base_menus`
WHERE `name` = 'project'
AND NOT EXISTS (
    SELECT 1 FROM `sys_authority_menus`
    WHERE `sys_authority_authority_id` = 9528
    AND `sys_base_menu_id` = (SELECT id FROM `sys_base_menus` WHERE `name` = 'project')
);

-- =====================================================
-- 第三部分：验证查询（可选）
-- =====================================================

-- 查询"项目说明"菜单是否添加成功
-- SELECT * FROM `sys_base_menus` WHERE `name` = 'project';

-- 查询各角色的菜单权限
-- SELECT
--     a.authority_id,
--     a.authority_name,
--     m.title,
--     m.path
-- FROM `sys_authority_menus` am
-- JOIN `sys_authorities` a ON am.sys_authority_authority_id = a.authority_id
-- JOIN `sys_base_menus` m ON am.sys_base_menu_id = m.id
-- WHERE m.name = 'project'
-- ORDER BY a.authority_id;

-- =====================================================
-- 第四部分：回滚SQL（如果需要删除）
-- =====================================================

-- 删除角色菜单关联
-- DELETE FROM `sys_authority_menus`
-- WHERE `sys_base_menu_id` = (SELECT id FROM `sys_base_menus` WHERE `name` = 'project');

-- 删除"项目说明"菜单
-- DELETE FROM `sys_base_menus` WHERE `name` = 'project';
