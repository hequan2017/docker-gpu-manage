-- ====================================================================
-- PortForward 插件 Casbin 权限修复 SQL
-- 修复 getAllForwarderStatus 等 API 权限不足问题
-- ====================================================================

-- 为管理员角色 (authority_id = 888) 添加 portForward API 权限
INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`) VALUES
('p', '888', '/portForward/createPortForward', 'POST'),
('p', '888', '/portForward/deletePortForward', 'DELETE'),
('p', '888', '/portForward/deletePortForwardByIds', 'DELETE'),
('p', '888', '/portForward/updatePortForward', 'PUT'),
('p', '888', '/portForward/updatePortForwardStatus', 'PUT'),
('p', '888', '/portForward/findPortForward', 'GET'),
('p', '888', '/portForward/getPortForwardList', 'GET'),
('p', '888', '/portForward/getServerIP', 'GET'),
('p', '888', '/portForward/getForwarderStatus', 'GET'),
('p', '888', '/portForward/getAllForwarderStatus', 'GET')
ON DUPLICATE KEY UPDATE `v2` = VALUES(`v2`);

-- 验证插入结果
SELECT * FROM `casbin_rule` WHERE `v1` LIKE '/portForward%' AND `v0` = '888';
