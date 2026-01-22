@echo off
REM 端口转发测试脚本 - 批处理版本

echo ========================================
echo      端口转发测试工具
echo ========================================
echo.

echo 配置:
echo   目标服务器: 127.0.0.1:8081
echo   测试次数: 5
echo   测试消息: Hello Port Forwarding!
echo.

echo 开始测试...
echo.

go run test_client.go -server=127.0.0.1:8081 -count=5 -msg="Hello Port Forwarding!"

echo.
echo 测试完成！
pause
