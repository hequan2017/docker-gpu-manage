-- =====================================================
-- 端口转发权限修复 - 快速执行SQL
-- =====================================================
-- 使用方法：
-- 1. 替换 your_database_name 为你的数据库名
-- 2. 在MySQL中执行此脚本
-- 3. 重新登录前端
-- =====================================================

USE your_database_name;  -- ⚠️ 替换为你的数据库名

-- =====================================================
-- 第1步：删除旧的端口转发API和权限（避免重复）
-- =====================================================
DELETE FROM casbin_rule WHERE v1 LIKE '/portForward%';
DELETE FROM sys_apis WHERE api_group = '端口转发';

-- =====================================================
-- 第2步：创建端口转发API
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
-- 第3步：为所有角色分配端口转发API权限
-- =====================================================
INSERT INTO casbin_rule (ptype, v0, v1, v2)
SELECT 'p', sa.authority_id, CONCAT(sapi.path, sapi.method), sapi.method
FROM sys_authorities sa
CROSS JOIN sys_apis sapi
WHERE sapi.api_group = '端口转发';

-- =====================================================
-- 第4步：验证修复结果
-- =====================================================
SELECT '端口转发API数量:' as 检查项, COUNT(*) as 数量 FROM sys_apis WHERE api_group = '端口转发'
UNION ALL
SELECT '端口转发权限数量:' as 检查项, COUNT(*) as 数量 FROM casbin_rule WHERE v1 LIKE '/portForward%';

-- =====================================================
-- 执行完成后：
-- 1. 退出前端登录
-- 2. 清除浏览器缓存 (Ctrl + Shift + Delete)
-- 3. 重新登录
-- 4. 尝试新建端口转发规则
-- =====================================================
