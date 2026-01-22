@echo off
chcp 65001 >nul
echo ========================================
echo     端口转发测试 - 8082端口
echo ========================================
echo.

cd /d/devops/test-2025\docker-gpu-manage/test
echo 测试目标: 127.0.0.1:8082
echo 测试次数: 5
echo.
echo 开始测试...
echo.

go run test_client.go -server="127.0.0.1:8082" -count=5 -msg="Hello Test"

echo.
echo ========================================
echo 测试完成！
echo ========================================
pause
