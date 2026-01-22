# 端口转发自动化测试脚本 - PowerShell版本

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "     端口转发自动化测试" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 配置
$ECHO_PORT = 9999
$FORWARD_PORT = 8081
$TEST_COUNT = 5
$TEST_MESSAGE = "Hello Test"

# 步骤1: 检查端口占用
Write-Host "[步骤1] 检查端口占用" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

$port9999 = Get-NetTCPConnection -LocalPort $ECHO_PORT -ErrorAction SilentlyContinue
$port8081 = Get-NetTCPConnection -LocalPort $FORWARD_PORT -ErrorAction SilentlyContinue

if ($port9999) {
    Write-Host "⚠️  端口 $ECHO_PORT 已被占用 (PID: $($port9999.OwningProcess))" -ForegroundColor Red
    $stop = Read-Host "是否停止占用进程? (y/n)"
    if ($stop -eq "y") {
        Stop-Process -Id $port9999.OwningProcess -Force
        Write-Host "✅ 进程已停止" -ForegroundColor Green
    }
} else {
    Write-Host "✅ 端口 $ECHO_PORT 可用" -ForegroundColor Green
}

if ($port8081) {
    Write-Host "⚠️  端口 $FORWARD_PORT 已被占用 (PID: $($port8081.OwningProcess))" -ForegroundColor Yellow
    Write-Host "   这可能是后端服务或端口转发器" -ForegroundColor Cyan
} else {
    Write-Host "✅ 端口 $FORWARD_PORT 可用" -ForegroundColor Green
}

Write-Host ""

# 步骤2: 启动Echo服务器
Write-Host "[步骤2] 启动Echo服务器" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "启动Echo服务器 (端口 $ECHO_PORT)..." -ForegroundColor Cyan

$echoProcess = Start-Process -FilePath "go" -ArgumentList "run", "echo_server.go", $ECHO_PORT -PassThru -WindowStyle Hidden -WorkingDirectory "D:\devops\test-2025\docker-gpu-manage\test"

Start-Sleep -Seconds 2

# 检查Echo服务器是否运行
$echoRunning = Get-NetTCPConnection -LocalPort $ECHO_PORT -ErrorAction SilentlyContinue
if ($echoRunning) {
    Write-Host "✅ Echo服务器已启动 (PID: $($echoProcess.Id))" -ForegroundColor Green
} else {
    Write-Host "❌ Echo服务器启动失败" -ForegroundColor Red
    exit 1
}

Write-Host ""

# 步骤3: 测试Echo服务器
Write-Host "[步骤3] 测试Echo服务器" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "测试直接连接到Echo服务器 (端口 $ECHO_PORT)..." -ForegroundColor Cyan

$testResult = & go run test_client.go -server="127.0.0.1:$ECHO_PORT" -count=1 -msg="Direct Test" 2>&1

if ($LASTEXITCODE -eq 0 -and $testResult -match "✅ 测试通过") {
    Write-Host "✅ Echo服务器工作正常" -ForegroundColor Green
} else {
    Write-Host "⚠️  Echo服务器测试失败，但继续测试" -ForegroundColor Yellow
}

Write-Host ""

# 步骤4: 检查端口转发规则
Write-Host "[步骤4] 检查端口转发规则" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "请确认已创建以下端口转发规则:" -ForegroundColor Yellow
Write-Host "  源地址: 0.0.0.0:$FORWARD_PORT" -ForegroundColor White
Write-Host "  目标地址: 127.0.0.1:$ECHO_PORT" -ForegroundColor White
Write-Host "  协议: TCP" -ForegroundColor White
Write-Host "  状态: 启用" -ForegroundColor White
Write-Host ""

$created = Read-Host "端口转发规则已创建? (y/n)"

if ($created -ne "y") {
    Write-Host ""
    Write-Host "请先创建端口转发规则:" -ForegroundColor Yellow
    Write-Host "1. 打开浏览器访问: http://localhost:8080" -ForegroundColor White
    Write-Host "2. 登录系统" -ForegroundColor White
    Write-Host "3. 进入'端口转发'菜单" -ForegroundColor White
    Write-Host "4. 点击'新建'按钮" -ForegroundColor White
    Write-Host "5. 填写规则配置并启用" -ForegroundColor White
    Write-Host ""

    # 检查是否需要权限修复
    $needPermission = Read-Host "是否遇到权限不足问题? (y/n)"
    if ($needPermission -eq "y") {
        Write-Host ""
        Write-Host "请执行以下SQL修复权限:" -ForegroundColor Yellow
        Write-Host "1. 打开 MySQL 客户端" -ForegroundColor White
        Write-Host "2. 执行 D:\devops\test-2025\docker-gpu-manage\fix_portforward_permissions.sql" -ForegroundColor White
        Write-Host "3. 退出登录并重新登录" -ForegroundColor White
    }

    Write-Host ""
    Write-Host "创建规则后，请重新运行此脚本" -ForegroundColor Cyan

    # 清理
    Stop-Process -Id $echoProcess.Id -Force 2>$null
    exit 0
}

Write-Host ""

# 步骤5: 运行端口转发测试
Write-Host "[步骤5] 运行端口转发测试" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray
Write-Host "配置:" -ForegroundColor Yellow
Write-Host "  测试目标: 127.0.0.1:$FORWARD_PORT" -ForegroundColor White
Write-Host "  测试次数: $TEST_COUNT" -ForegroundColor White
Write-Host "  测试消息: $TEST_MESSAGE" -ForegroundColor White
Write-Host ""

Write-Host "开始测试..." -ForegroundColor Cyan
Write-Host ""

& {
    $ErrorActionPreference = "Continue"
    go run test_client.go -server="127.0.0.1:$FORWARD_PORT" -count=$TEST_COUNT -msg=$TEST_MESSAGE
}

$testExitCode = $LASTEXITCODE

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  测试完成" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 步骤6: 清理
Write-Host "[步骤6] 清理资源" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Stop-Process -Id $echoProcess.Id -Force 2>$null
Write-Host "✅ Echo服务器已停止" -ForegroundColor Green

Write-Host ""

# 结果总结
if ($testExitCode -eq 0) {
    Write-Host "🎉 测试成功！端口转发工作正常！" -ForegroundColor Green
    Write-Host ""
    Write-Host "端口转发功能已成功实现并验证！" -ForegroundColor Green
} else {
    Write-Host "⚠️  测试失败" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "请检查:" -ForegroundColor Yellow
    Write-Host "1. 端口转发规则是否创建" -ForegroundColor White
    Write-Host "2. 规则状态是否为'启用'" -ForegroundColor White
    Write-Host "3. 后端日志是否有错误" -ForegroundColor White
    Write-Host "4. 防火墙是否阻止连接" -ForegroundColor White
}

Write-Host ""
pause
