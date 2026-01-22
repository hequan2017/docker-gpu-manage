# 删除端口转发规则脚本
# 需要管理员权限运行

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "     删除端口转发规则工具" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查管理员权限
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if (-not $isAdmin) {
    Write-Host "❌ 错误: 此脚本需要管理员权限运行！" -ForegroundColor Red
    Write-Host ""
    Write-Host "请按以下步骤操作:" -ForegroundColor Yellow
    Write-Host "  1. 右键点击 PowerShell 图标" -ForegroundColor White
    Write-Host "  2. 选择 '以管理员身份运行'" -ForegroundColor White
    Write-Host "  3. 重新运行此脚本" -ForegroundColor White
    Write-Host ""
    Write-Host "或者使用以下命令自动提升权限:" -ForegroundColor Yellow
    Write-Host "  Start-Process powershell -Verb runAs -ArgumentList '-NoExit', '-Command', 'cd D:\devops\test-2025\docker-gpu-manage\test; .\删除端口转发规则.ps1'" -ForegroundColor Cyan
    Write-Host ""
    pause
    exit 1
}

Write-Host "✅ 已获取管理员权限" -ForegroundColor Green
Write-Host ""

# 显示当前所有端口转发规则
Write-Host "📌 当前端口转发规则:" -ForegroundColor Yellow
Write-Host "----------------------------------------" -ForegroundColor Gray

$rules = netsh interface portproxy show all 2>&1

if ($rules -match "v4tov4") {
    Write-Host $rules -ForegroundColor White
    Write-Host ""

    # 询问用户操作
    Write-Host "请选择操作:" -ForegroundColor Yellow
    Write-Host "  1. 删除指定端口转发规则" -ForegroundColor White
    Write-Host "  2. 删除所有端口转发规则" -ForegroundColor White
    Write-Host "  3. 取消" -ForegroundColor White
    Write-Host ""

    $choice = Read-Host "请输入选项 (1/2/3)"

    switch ($choice) {
        "1" {
            Write-Host ""
            $port = Read-Host "请输入要删除的端口号 (例如: 8081)"
            $address = Read-Host "请输入监听地址 (默认: 0.0.0.0，直接回车使用默认值)"

            if ([string]::IsNullOrWhiteSpace($address)) {
                $address = "0.0.0.0"
            }

            Write-Host ""
            Write-Host "删除端口转发规则: $address`:$port" -ForegroundColor Yellow

            $command = "netsh interface portproxy delete v4tov4 listenport=$port listenaddress=$address"
            $result = Invoke-Expression $command 2>&1

            if ($LASTEXITCODE -eq 0) {
                Write-Host "✅ 端口转发规则删除成功" -ForegroundColor Green
            } else {
                Write-Host "❌ 删除失败: $result" -ForegroundColor Red
            }
        }

        "2" {
            Write-Host ""
            $confirm = Read-Host "⚠️  确认要删除所有端口转发规则? (yes/no)"

            if ($confirm -eq "yes") {
                Write-Host "删除所有端口转发规则..." -ForegroundColor Yellow
                $result = netsh interface portproxy reset 2>&1

                Write-Host "✅ 所有端口转发规则已删除" -ForegroundColor Green
            } else {
                Write-Host "已取消操作" -ForegroundColor Yellow
            }
        }

        "3" {
            Write-Host "已取消操作" -ForegroundColor Yellow
        }

        default {
            Write-Host "❌ 无效的选项" -ForegroundColor Red
        }
    }
} else {
    Write-Host "✅ 当前没有端口转发规则" -ForegroundColor Green
}

Write-Host ""
Write-Host "----------------------------------------" -ForegroundColor Gray
Write-Host "剩余的端口转发规则:" -ForegroundColor Yellow
netsh interface portproxy show all
Write-Host ""
Write-Host "操作完成！" -ForegroundColor Green
Write-Host ""
