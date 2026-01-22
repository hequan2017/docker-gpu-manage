# 端口转发测试脚本
# PowerShell版本

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "     端口转发测试工具" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 配置参数
$SERVER = "127.0.0.1:8081"
$COUNT = 5
$MESSAGE = "Hello Port Forwarding!"

Write-Host "📋 测试配置:" -ForegroundColor Yellow
Write-Host "  目标服务器: $SERVER"
Write-Host "  测试次数: $COUNT"
Write-Host "  测试消息: $MESSAGE"
Write-Host ""

Write-Host "🚀 开始编译和运行测试客户端..." -ForegroundColor Green
Write-Host ""

# 编译并运行
go run test_client.go -server="$SERVER" -count=$COUNT -msg="$MESSAGE"
