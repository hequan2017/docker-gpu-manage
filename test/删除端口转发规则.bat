@echo off
:: 批处理文件 - 自动请求管理员权限并删除端口转发规则

echo ========================================
echo      删除端口转发规则工具
echo ========================================
echo.

:: 检查是否有管理员权限
net session >nul 2>&1
if %errorLevel% == 0 (
    echo [√] 已获取管理员权限
    echo.
    goto :menu
) else (
    echo [!] 需要管理员权限运行此脚本
    echo [*] 正在请求管理员权限...
    echo.
    powershell -Command "Start-Process '%~f0' -Verb RunAs"
    exit /b
)

:menu
echo 当前端口转发规则:
echo ----------------------------------------
netsh interface portproxy show all
echo ----------------------------------------
echo.
echo 请选择操作:
echo   1. 删除指定端口转发规则
echo   2. 删除所有端口转发规则
echo   3. 退出
echo.

set /p choice="请输入选项 (1/2/3): "

if "%choice%"=="1" goto :delete_one
if "%choice%"=="2" goto :delete_all
if "%choice%"=="3" goto :end

echo.
echo [×] 无效的选项
goto :end

:delete_one
echo.
set /p port="请输入要删除的端口号 (例如: 8081): "
set /p address="请输入监听地址 (默认: 0.0.0.0，直接回车使用默认值): "

if "%address%"=="" set address=0.0.0.0

echo.
echo [*] 删除端口转发规则: %address%:%port%
netsh interface portproxy delete v4tov4 listenport=%port% listenaddress=%address%

if %errorLevel%==0 (
    echo [√] 端口转发规则删除成功
) else (
    echo [×] 删除失败
)
echo.
goto :show_remaining

:delete_all
echo.
set /p confirm="[!] 确认要删除所有端口转发规则? (yes/no): "

if "%confirm%"=="yes" (
    echo [*] 删除所有端口转发规则...
    netsh interface portproxy reset
    echo [√] 所有端口转发规则已删除
    echo.
) else (
    echo [*] 已取消操作
    echo.
)
goto :show_remaining

:show_remaining
echo ----------------------------------------
echo 剩余的端口转发规则:
echo ----------------------------------------
netsh interface portproxy show all
echo.
echo 操作完成！
echo.

:end
pause
