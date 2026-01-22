@echo off
chcp 65001 >nul
echo ========================================
echo     检查端口状态
echo ========================================
echo.

echo [1] 检查8082端口
powershell -Command "if (Get-NetTCPConnection -LocalPort 8082 -ErrorAction SilentlyContinue) { Write-Host '❌ 8082端口已被占用' -ForegroundColor Red } else { Write-Host '✅ 8082端口可用，可以使用' -ForegroundColor Green }"
echo.

echo [2] 检查8888端口
powershell -Command "if (Get-NetTCPConnection -LocalPort 8888 -ErrorAction SilentlyContinue) { Write-Host '❌ 8888端口已被占用' -ForegroundColor Red } else { Write-Host '✅ 8888端口可用，可以使用' -ForegroundColor Green }"
echo.

echo [3] 检查9999端口 (Echo服务器)
powershell -Command "if (Get-NetTCPConnection -LocalPort 9999 -ErrorAction SilentlyContinue) { Write-Host '✅ Echo服务器运行中 (端口9999)' -ForegroundColor Green } else { Write-Host '❌ Echo服务器未运行' -ForegroundColor Red }"
echo.

echo ========================================
echo.
echo 推荐使用可用端口创建规则
echo ========================================
pause
