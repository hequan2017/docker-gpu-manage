#!/bin/bash
# 端口转发自动化测试脚本

echo "========================================"
echo "     端口转发自动化测试"
echo "========================================"
echo ""

# 检查端口占用
echo "[1] 检查端口占用..."
if netstat -tuln 2>/dev/null | grep -q ":9999"; then
    echo "⚠️  端口9999已被占用"
else
    echo "✅ 端口9999可用"
fi

if netstat -tuln 2>/dev/null | grep -q ":8081"; then
    echo "⚠️  端口8081已被占用"
else
    echo "✅ 端口8081可用"
fi
echo ""

# 启动Echo服务器
echo "[2] 启动Echo服务器..."
cd /d/devops/test-2025/docker-gpu-manage/test
nohup go run echo_server.go 9999 > echo_server.log 2>&1 &
ECHO_PID=$!
echo "Echo服务器已启动 (PID: $ECHO_PID)"
sleep 2
echo ""

# 检查Echo服务器是否正常启动
if ps -p $ECHO_PID > /dev/null 2>&1; then
    echo "✅ Echo服务器运行正常"
else
    echo "❌ Echo服务器启动失败"
    cat echo_server.log
    exit 1
fi
echo ""

# 测试Echo服务器
echo "[3] 测试Echo服务器直接连接..."
echo "Test Message" | timeout 3 nc localhost 9999 2>&1 || echo "nc命令不可用，跳过直接测试"
echo ""

# 提示用户创建端口转发规则
echo "[4] 端口转发规则检查"
echo "⚠️  请确认已在通过前端创建端口转发规则："
echo "   源地址: 0.0.0.0:8081"
echo "   目标地址: 127.0.0.1:9999"
echo "   协议: TCP"
echo "   状态: 启用"
echo ""

read -p "端口转发规则已创建? (y/n): " created

if [ "$created" != "y" ]; then
    echo ""
    echo "请先创建端口转发规则，然后重新运行此脚本"
    echo ""
    echo "创建步骤："
    echo "1. 打开浏览器: http://localhost:8080"
    echo "2. 登录系统"
    echo "3. 进入'端口转发'菜单"
    echo "4. 点击'新建'创建规则"
    echo "5. 填写上述配置并启用"
    exit 0
fi
echo ""

# 运行端口转发测试
echo "[5] 运行端口转发测试..."
echo "目标: 127.0.0.1:8081"
echo "次数: 5"
echo ""

go run test_client.go -server="127.0.0.1:8081" -count=5 -msg="Hello Test"
TEST_RESULT=$?

echo ""
echo "========================================"
echo "  测试完成"
echo "========================================"
echo ""

# 清理
echo "[6] 清理..."
kill $ECHO_PID 2>/dev/null
echo "Echo服务器已停止"
echo ""

if [ $TEST_RESULT -eq 0 ]; then
    echo "🎉 测试成功！端口转发工作正常！"
else
    echo "⚠️  测试失败，请检查配置"
fi

exit $TEST_RESULT
