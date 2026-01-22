# 端口转发测试脚本
# 用于测试端口转发功能

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "     端口转发测试工具" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 配置参数
$SOURCE_PORT = 8081        # 源端口（转发端口）
$TARGET_IP = "127.0.0.1"   # 目标IP（本机）
$TARGET_PORT = 9999        # 目标端口（Echo服务器监听端口）

Write-Host "📋 测试配置:" -ForegroundColor Yellow
Write-Host "  源端口: $SOURCE_PORT"
Write-Host "  目标地址: $TARGET_IP`:$TARGET_PORT"
Write-Host ""

# ============================================
# 步骤1: 检查端口占用
# ============================================
Write-Host "📌 步骤1: 检查端口占用状态" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

$sourcePortInUse = Get-NetTCPConnection -LocalPort $SOURCE_PORT -ErrorAction SilentlyContinue
$targetPortInUse = Get-NetTCPConnection -LocalPort $TARGET_PORT -ErrorAction SilentlyContinue

if ($sourcePortInUse) {
    Write-Host "⚠️  端口 $SOURCE_PORT 已被占用" -ForegroundColor Red
    Write-Host "   进程PID: $($sourcePortInUse.OwningProcess)"
    $process = Get-Process -Id $sourcePortInUse.OwningProcess -ErrorAction SilentlyContinue
    if ($process) {
        Write-Host "   进程名称: $($process.ProcessName)"
    }
} else {
    Write-Host "✅ 端口 $SOURCE_PORT 可用" -ForegroundColor Green
}

if ($targetPortInUse) {
    Write-Host "⚠️  端口 $TARGET_PORT 已被占用" -ForegroundColor Red
    Write-Host "   进程PID: $($targetPortInUse.OwningProcess)"
    $process = Get-Process -Id $targetPortInUse.OwningProcess -ErrorAction SilentlyContinue
    if ($process) {
        Write-Host "   进程名称: $($process.ProcessName)"
    }
} else {
    Write-Host "✅ 端口 $TARGET_PORT 可用" -ForegroundColor Green
}

Write-Host ""

# ============================================
# 步骤2: 启动Echo服务器
# ============================================
Write-Host "📌 步骤2: 启动Echo测试服务器" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "请先在新终端窗口中运行Echo服务器:" -ForegroundColor Yellow
Write-Host "  cd D:\devops\test-2025\docker-gpu-manage\test" -ForegroundColor White
Write-Host "  go run echo_server.go $TARGET_PORT" -ForegroundColor Cyan
Write-Host ""

$response = Read-Host "Echo服务器已启动? (y/n)"
if ($response -ne "y") {
    Write-Host "❌ 请先启动Echo服务器再继续测试" -ForegroundColor Red
    exit 1
}

Write-Host ""

# ============================================
# 步骤3: 配置端口转发
# ============================================
Write-Host "📌 步骤3: 配置Windows端口转发" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "添加端口转发规则:" -ForegroundColor Yellow
$command = "netsh interface portproxy add v4tov4 listenport=$SOURCE_PORT listenaddress=0.0.0.0 connectport=$TARGET_PORT connectaddress=$TARGET_IP"
Write-Host "  $command" -ForegroundColor Cyan

try {
    $result = Invoke-Expression $command 2>&1
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ 端口转发规则添加成功" -ForegroundColor Green
    } else {
        Write-Host "⚠️  可能已存在同名规则，继续测试..." -ForegroundColor Yellow
    }
} catch {
    Write-Host "⚠️  添加规则时出现警告: $_" -ForegroundColor Yellow
}

Write-Host ""

# ============================================
# 步骤4: 验证端口转发配置
# ============================================
Write-Host "📌 步骤4: 验证端口转发配置" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

$proxyRules = netsh interface portproxy show all
Write-Host "当前端口转发规则:" -ForegroundColor Yellow
Write-Host $proxyRules
Write-Host ""

# ============================================
# 步骤5: 测试端口连接
# ============================================
Write-Host "📌 步骤5: 测试端口连接" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "测试连接到 localhost:$SOURCE_PORT ..." -ForegroundColor Yellow
try {
    $tcpClient = New-Object System.Net.Sockets.TcpClient
    $tcpClient.ReceiveTimeout = 5000
    $tcpClient.Connect("127.0.0.1", $SOURCE_PORT)
    Write-Host "✅ 成功连接到端口 $SOURCE_PORT" -ForegroundColor Green
    $tcpClient.Close()
} catch {
    Write-Host "❌ 连接失败: $_" -ForegroundColor Red
    Write-Host ""
    Write-Host "可能的原因:" -ForegroundColor Yellow
    Write-Host "  1. Echo服务器未启动" -ForegroundColor White
    Write-Host "  2. 端口转发规则未正确配置" -ForegroundColor White
    Write-Host "  3. 防火墙阻止了连接" -ForegroundColor White
    exit 1
}

Write-Host ""

# ============================================
# 步骤6: 测试数据传输
# ============================================
Write-Host "📌 步骤6: 测试数据传输" -ForegroundColor Green
Write-Host "----------------------------------------" -ForegroundColor Gray

Write-Host "发送测试消息..." -ForegroundColor Yellow
try {
    $tcpClient = New-Object System.Net.Sockets.TcpClient
    $tcpClient.ReceiveTimeout = 5000
    $tcpClient.Connect("127.0.0.1", $SOURCE_PORT)

    $stream = $tcpClient.GetStream()
    $writer = New-Object System.IO.StreamWriter($stream)
    $reader = New-Object System.IO.StreamReader($stream)

    # 发送测试消息
    $testMessage = "Hello Port Forwarding!`n"
    $writer.Write($testMessage)
    $writer.Flush()

    Write-Host "📤 发送: $testMessage" -ForegroundColor Cyan

    # 接收响应
    $response = $reader.ReadLine()
    if ($response) {
        Write-Host "📨 接收: $response" -ForegroundColor Green
        Write-Host "✅ 端口转发测试成功！" -ForegroundColor Green
    } else {
        Write-Host "⚠️  未收到响应" -ForegroundColor Yellow
    }

    $writer.Close()
    $reader.Close()
    $stream.Close()
    $tcpClient.Close()

} catch {
    Write-Host "❌ 数据传输失败: $_" -ForegroundColor Red
}

Write-Host ""

# ============================================
# 完成
# ============================================
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  测试完成" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "查看端口转发规则:" -ForegroundColor Yellow
Write-Host "  netsh interface portproxy show all" -ForegroundColor Cyan
Write-Host ""
Write-Host "删除端口转发规则:" -ForegroundColor Yellow
Write-Host "  netsh interface portproxy delete v4tov4 listenport=$SOURCE_PORT listenaddress=0.0.0.0" -ForegroundColor Cyan
Write-Host ""
Write-Host "删除所有端口转发规则:" -ForegroundColor Yellow
Write-Host "  netsh interface portproxy reset" -ForegroundColor Cyan
Write-Host ""
