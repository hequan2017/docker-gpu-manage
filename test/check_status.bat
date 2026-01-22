@echo off
chcp 65001 >nul
echo ========================================
echo     端口转发状态检查
echo ========================================
echo.

echo [1] 检查9999端口 (Echo服务器)
powershell -Command "if (Get-NetTCPConnection -LocalPort 9999 -ErrorAction SilentlyContinue) { Write-Host '✅ Echo服务器运行中 (端口9999)' -ForegroundColor Green } else { Write-Host '❌ Echo服务器未运行' -ForegroundColor Red }"
echo.

echo [2] 检查8081端口 (端口转发)
powershell -Command "if (Get-NetTCPConnection -LocalPort 8081 -ErrorAction SilentlyContinue) { Write-Host '✅ 8081端口已监听' -ForegroundColor Yellow; Write-Host '  如果是后端服务，说明端口转发规则未创建' -ForegroundColor Cyan } else { Write-Host '❌ 8081端口未监听' -ForegroundColor Red }"
echo.

echo [3] 端口转发检查清单
echo.
echo   创建端口转发规则:
echo   1. 打开浏览器: http://localhost:8080
echo   2. 登录系统
echo   3. 进入"端口转发"菜单
echo   4. 点击"新建"
echo   5. 填写配置:
echo      源IP: 0.0.0.0
echo      源端口: 8081
echo      协议: TCP
echo      目标IP: 127.0.0.1
echo      目标端口: 9999
echo      状态: 启用 ✓
echo   6. 点击"确定"
echo.

echo [4] 创建规则后测试
echo.
echo   运行测试命令:
echo   cd D:\devops\test-2025\docker-gpu-manage\test
echo   go run test_client.go -server="127.0.0.1:8081" -count=5
echo.
echo   或双击运行: run_test.bat
echo.

echo ========================================
pause
