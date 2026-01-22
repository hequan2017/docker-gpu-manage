@echo off
REM ========================================
REM 端口转发权限快速修复脚本
REM ========================================

echo ========================================
echo   端口转发权限快速修复工具
echo ========================================
echo.

echo [步骤1] 检查数据库连接信息...
echo.

REM 检查config.yaml是否存在
if not exist "..\server\config.yaml" (
    echo [错误] 找不到配置文件: ..\server\config.yaml
    echo 请确认当前目录是否在 test/ 目录下
    pause
    exit /b 1
)

echo [√] 配置文件存在
echo.
echo [步骤2] 请手动执行以下SQL修复权限:
echo.
echo ----------------------------------------
echo USE your_database_name;
echo.
echo -- 删除旧的端口转发API和权限
echo DELETE FROM casbin_rule WHERE v1 LIKE '/portForward%%';
echo DELETE FROM sys_apis WHERE api_group = '端口转发';
echo.
echo -- 插入端口转发API
echo INSERT INTO `sys_apis` (`path`, `description`, `api_group`, `method`, `created_at`, `updated_at`) VALUES
echo ('/portForward/createPortForward', '新建端口转发规则', '端口转发', 'POST', NOW(), NOW()),
echo ('/portForward/deletePortForward', '删除端口转发规则', '端口转发', 'DELETE', NOW(), NOW()),
echo ('/portForward/deletePortForwardByIds', '批量删除端口转发规则', '端口转发', 'DELETE', NOW(), NOW()),
echo ('/portForward/updatePortForward', '更新端口转发规则', '端口转发', 'PUT', NOW(), NOW()),
echo ('/portForward/updatePortForwardStatus', '更新端口转发规则状态', '端口转发', 'PUT', NOW(), NOW()),
echo ('/portForward/findPortForward', '根据ID获取端口转发规则', '端口转发', 'GET', NOW(), NOW()),
echo ('/portForward/getPortForwardList', '获取端口转发规则列表', '端口转发', 'GET', NOW(), NOW()),
echo ('/portForward/getServerIP', '获取服务器IP地址', '端口转发', 'GET', NOW(), NOW()),
echo ('/portForward/getForwarderStatus', '获取端口转发运行状态', '端口转发', 'GET', NOW(), NOW()),
echo ('/portForward/getAllForwarderStatus', '获取所有端口转发运行状态', '端口转发', 'GET', NOW(), NOW());
echo.
echo -- 为所有角色添加权限
echo INSERT INTO casbin_rule (ptype, v0, v1, v2)
echo SELECT 'p', sa.authority_id, CONCAT(sapi.path, sapi.method), sapi.method
echo FROM sys_authorities sa
echo CROSS JOIN sys_apis sapi
echo WHERE sapi.api_group = '端口转发';
echo ----------------------------------------
echo.
echo [步骤3] 执行完SQL后:
echo   1. 退出登录前端
echo   2. 清除浏览器缓存 (Ctrl+Shift+Delete)
echo   3. 重新登录
echo   4. 尝试新建端口转发规则
echo.
pause
