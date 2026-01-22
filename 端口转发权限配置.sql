-- =====================================================
-- 端口转发模块 - 权限配置SQL脚本
-- 解决"权限不足"问题
-- =====================================================

-- 步骤1: 确认端口转发API已创建（如果初始化代码未执行，手动执行）

-- 检查API是否存在
SELECT COUNT(*) as api_count FROM sys_apis WHERE path LIKE '/portForward%';

-- 如果上面的查询结果为0，说明API未创建，执行以下INSERT

-- =====================================================
-- 端口转发API定义
-- =====================================================
INSERT INTO `sys_apis` (`path`, `description`, `api_group`, `method`, `created_at`, `updated_at`) VALUES
('/portForward/createPortForward', '新建端口转发规则', '端口转发', 'POST', NOW(), NOW()),
('/portForward/deletePortForward', '删除端口转发规则', '端口转发', 'DELETE', NOW(), NOW()),
('/portForward/deletePortForwardByIds', '批量删除端口转发规则', '端口转发', 'DELETE', NOW(), NOW()),
('/portForward/updatePortForward', '更新端口转发规则', '端口转发', 'PUT', NOW(), NOW()),
('/portForward/updatePortForwardStatus', '更新端口转发规则状态', '端口转发', 'PUT', NOW(), NOW()),
('/portForward/findPortForward', '根据ID获取端口转发规则', '端口转发', 'GET', NOW(), NOW()),
('/portForward/getPortForwardList', '获取端口转发规则列表', '端口转发', 'GET', NOW(), NOW()),
('/portForward/getServerIP', '获取服务器IP地址', '端口转发', 'GET', NOW(), NOW()),
('/portForward/getForwarderStatus', '获取端口转发运行状态', '端口转发', 'GET', NOW(), NOW()),
('/portForward/getAllForwarderStatus', '获取所有端口转发运行状态', '端口转发', 'GET', NOW(), NOW());

-- =====================================================
-- 步骤2: 给管理员角色（通常ID为888）分配所有端口转发API权限
-- =====================================================

-- 2.1 首先查询管理员角色的authority_id
SELECT authority_id, authority_name, authority_type FROM casbin_rule WHERE ptype = 'g' AND v0 LIKE 'admin%';

-- 2.2 查询端口转发的所有API ID
SELECT id, path, method FROM sys_apis WHERE api_group = '端口转发';

-- 2.3 为管理员角色（假设authority_id为888）添加所有端口转发API权限
-- 注意：如果你的管理员角色ID不是888，请替换为实际的authority_id

-- 获取管理员角色ID（通常情况）
SET @admin_authority_id = 888;

-- 为管理员角色添加端口转发API权限
INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`)
SELECT
    'p',
    @admin_authority_id,
    CONCAT(path, [method]),
    method
FROM sys_apis
WHERE api_group = '端口转发';

-- =====================================================
-- 步骤3: 为超级管理员角色（通常ID为999）分配权限
-- =====================================================

SET @super_admin_authority_id = 999;

INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`)
SELECT
    'p',
    @super_admin_authority_id,
    CONCAT(path, [method]),
    method
FROM sys_apis
WHERE api_group = '端口转发';

-- =====================================================
-- 步骤4: 验证权限是否添加成功
-- =====================================================

-- 查看端口转发的API
SELECT id, path, method, description FROM sys_apis WHERE api_group = '端口转发';

-- 查看管理员角色的端口转发权限
SELECT * FROM casbin_rule
WHERE ptype = 'p'
  AND v0 IN (@admin_authority_id, @super_admin_authority_id)
  AND v1 LIKE '/portForward%';

-- =====================================================
-- 备选方案：如果上面的方法不适用，使用直接API路径方式
-- =====================================================

-- 查询当前所有API
-- SELECT * FROM sys_apis ORDER BY id DESC LIMIT 20;

-- 如果需要删除并重新创建API权限
-- DELETE FROM casbin_rule WHERE v1 LIKE '/portForward%';

-- =====================================================
-- 故障排查查询
-- =====================================================

-- 查看当前登录用户的角色
-- SELECT * FROM sys_users WHERE nick_name = '当前用户名';

-- 查看用户角色关联
-- SELECT ur.*, sa.authority_name
-- FROM sys_user_authorities ur
-- JOIN sys_authorities sa ON ur.authority_id = sa.authority_id
-- WHERE ur.user_uuid = '你的用户UUID';

-- 查看角色的所有API权限
-- SELECT * FROM casbin_rule WHERE ptype = 'p' AND v0 = '你的角色ID';

-- =====================================================
-- 完成提示
-- =====================================================

-- 执行完以上SQL后：
-- 1. 退出登录
-- 2. 清除浏览器缓存（Ctrl+F5）
-- 3. 重新登录
-- 4. 尝试点击"新建"按钮

-- 如果仍然提示权限不足：
-- 1. 检查浏览器控制台的错误信息（F12）
-- 2. 检查后端日志中的权限验证信息
-- 3. 确认当前用户的角色ID是否正确
