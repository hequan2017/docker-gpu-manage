# 端口转发完整测试脚本
# 自动化测试流程

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  端口转发完整测试工具" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查端口占用
Write-Host "📌 步骤1: 检查端口占用" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

$port8081 = Get-NetTCPConnection -LocalPort 8081 -ErrorAction SilentlyContinue
$port9999 = Get-NetTCPConnection -LocalPort 9999 -ErrorAction SilentlyContinue

if ($port8081) {
    Write-Host "⚠️  端口8081已被占用" -ForegroundColor Red
    Write-Host "   进程: $($port8081.OwningProcess)"
} else {
    Write-Host "✅ 端口8081可用" -ForegroundColor Green
}

if ($port9999) {
    Write-Host "⚠️  端口9999已被占用" -ForegroundColor Red
    Write-Host "   进程: $($port9999.OwningProcess)"
} else {
    Write-Host "✅ 端口9999可用" -ForegroundColor Green
}

Write-Host ""

# Echo服务器检查
Write-Host "📌 步骤2: 检查Echo服务器" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

$echoRunning = Get-NetTCPConnection -LocalPort 9999 -ErrorAction SilentlyContinue
if ($echoRunning) {
    Write-Host "✅ Echo服务器正在运行" -ForegroundColor Green
} else {
    Write-Host "❌ Echo服务器未运行" -ForegroundColor Red
    Write-Host ""
    Write-Host "请先在新窗口启动Echo服务器:" -ForegroundColor Yellow
    Write-Host "  cd D:\devops\test-2025\docker-gpu-manage\test" -ForegroundColor White
    Write-Host "  go run echo_server.go 9999" -ForegroundColor Cyan
    Write-Host ""

    $response = Read-Host "Echo服务器已启动? (y/n)"
    if ($response -ne "y") {
        Write-Host "❌ 请先启动Echo服务器" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""

# 后端服务检查
Write-Host "📌 步骤3: 检查后端服务" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "请确认后端服务已启动 (端口8080)" -ForegroundColor Yellow
Write-Host "  cd D:\devops\test-2025\docker-gpu-manage\server" -ForegroundColor White
Write-Host "  go run main.go" -ForegroundColor Cyan
Write-Host ""

$response = Read-Host "后端服务已启动? (y/n)"
if ($response -ne "y") {
    Write-Host "❌ 请先启动后端服务" -ForegroundColor Red
    exit 1
}

Write-Host ""

# 检查端口转发规则
Write-Host "📌 步骤4: 检查端口转发规则" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "请确认已创建端口转发规则:" -ForegroundColor Yellow
Write-Host "  源地址: 0.0.0.0:8081" -ForegroundColor White
Write-Host "  目标地址: 127.0.0.1:9999" -ForegroundColor White
Write-Host "  协议: TCP" -ForegroundColor White
Write-Host "  状态: 启用" -ForegroundColor White
Write-Host ""

$response = Read-Host "端口转发规则已创建? (y/n)"
if ($response -ne "y") {
    Write-Host ""
    Write-Host "请通过前端创建规则:" -ForegroundColor Yellow
    Write-Host "  1. 打开浏览器访问: http://localhost:8080" -ForegroundColor White
    Write-Host "  2. 登录系统" -ForegroundColor White
    Write-Host "  3. 进入端口转发管理" -ForegroundColor White
    Write-Host "  4. 点击新建，创建规则" -ForegroundColor White
    Write-Host ""

    $response = Read-Host "规则创建完成? (y/n)"
    if ($response -ne "y") {
        Write-Host "❌ 请先创建端口转发规则" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""

# 运行测试
Write-Host "📌 步骤5: 运行测试" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray
Write-Host ""

Write-Host "🚀 开始测试..." -ForegroundColor Cyan
Write-Host ""

# 设置参数
$SERVER = "127.0.0.1:8081"
$COUNT = 5
$MESSAGE = "Hello Port Forwarding!"

Write-Host "配置:" -ForegroundColor Yellow
Write-Host "  目标: $SERVER"
Write-Host "  次数: $COUNT"
Write-Host "  消息: $MESSAGE"
Write-Host ""

# 运行测试客户端
& {
    $ErrorActionPreference = "Continue"
    go run test_client.go -server="$SERVER" -count=$COUNT -msg="$MESSAGE"
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  测试完成" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
