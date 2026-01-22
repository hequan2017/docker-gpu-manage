# 端口转发测试工具集

## 📁 文件清单

```
test/
├── README.md                          # 本文件 - 总体说明
├── 端口转发测试指南.md                 # 完整的测试教程
├── 删除端口转发规则-快速指南.md        # 权限问题解决方案
├── echo_server.go                     # Echo测试服务器
├── test_client.go                     # TCP测试客户端
├── test_port_forwarding.ps1           # 自动化测试脚本
├── 删除端口转发规则.ps1                # PowerShell删除脚本
└── 删除端口转发规则.bat                # 批处理删除脚本
```

## 🎯 测试目标

验证Windows端口转发功能：
- **源端口**: 8081（外部访问端口）
- **目标端口**: 9999（Echo服务器监听端口）
- **测试路径**: 客户端 → 8081 → 转发 → 9999 → Echo服务器

## 🚀 快速开始

### ⭐ 推荐方式：使用自动化脚本

1. **以管理员身份打开PowerShell**

2. **运行自动化测试**
   ```bash
   cd D:\devops\test-2025\docker-gpu-manage\test
   .\test_port_forwarding.ps1
   ```

3. **脚本会自动完成**:
   - ✅ 检查端口占用
   - ✅ 配置端口转发
   - ✅ 启动Echo服务器（需手动）
   - ✅ 运行测试客户端
   - ✅ 验证数据传输

### 手动测试流程

#### 步骤1：启动Echo服务器
```bash
cd D:\devops\test-2025\docker-gpu-manage\test
go run echo_server.go 9999
```

#### 步骤2：配置端口转发（需要管理员权限）
```bash
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=9999 connectaddress=127.0.0.1
```

#### 步骤3：运行测试客户端
```bash
go run test_client.go -server=127.0.0.1:8081 -count=5
```

#### 步骤4：清理端口转发规则
```bash
# 方法1：使用批处理文件（双击运行）
.\删除端口转发规则.bat

# 方法2：使用PowerShell命令（需要管理员权限）
netsh interface portproxy delete v4tov4 listenport=8081 listenaddress=0.0.0.0
```

## 📚 详细文档

| 文档 | 说明 | 适用场景 |
|------|------|----------|
| [端口转发测试指南.md](./端口转发测试指南.md) | 完整的测试教程 | 首次测试、深入学习 |
| [删除端口转发规则-快速指南.md](./删除端口转发规则-快速指南.md) | 权限问题解决方案 | 删除规则时遇到权限错误 |

## 🛠️ 工具说明

### echo_server.go - Echo测试服务器

**功能**:
- 监听指定端口
- 接收客户端消息
- 原样返回消息（添加时间戳）
- 支持多客户端并发连接

**使用方法**:
```bash
# 使用默认端口9999
go run echo_server.go

# 指定端口
go run echo_server.go 8888
```

**输出示例**:
```
✅ Echo服务器已启动！
📡 监听地址: :9999
⏰ 启动时间: 2025-01-22 15:30:00
--------------------------------------------------
🔗 新连接来自: 127.0.0.1:54321
📨 收到消息: Hello Test
📤 发送响应: Echo: Hello Test [时间: 15:30:15]
```

### test_client.go - TCP测试客户端

**功能**:
- 连接到指定服务器
- 发送测试消息
- 接收并验证响应
- 支持多次发送和统计

**参数**:
- `-server`: 服务器地址（默认: 127.0.0.1:8081）
- `-msg`: 发送消息（默认: "Hello Port Forwarding!"）
- `-count`: 发送次数（默认: 1）
- `-interval`: 发送间隔毫秒（默认: 1000）

**使用方法**:
```bash
# 基本测试
go run test_client.go

# 自定义测试
go run test_client.go -server=127.0.0.1:8081 -msg="Test Message" -count=10

# 压力测试
go run test_client.go -count=1000 -interval=10
```

### test_port_forwarding.ps1 - 自动化测试脚本

**功能**:
- 自动检查端口占用
- 配置Windows端口转发
- 运行完整测试流程
- 生成测试报告

**使用方法**:
```powershell
# 以管理员身份运行
.\test_port_forwarding.ps1
```

### 删除端口转发规则工具

**PowerShell脚本** (`删除端口转发规则.ps1`):
```powershell
# 以管理员身份运行
.\删除端口转发规则.ps1
```

**批处理文件** (`删除端口转发规则.bat`):
```batch
# 双击运行（自动请求管理员权限）
.\删除端口转发规则.bat
```

## 🔧 Windows端口转发命令

### 查看所有规则
```bash
netsh interface portproxy show all
```

### 添加规则
```bash
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=9999 connectaddress=127.0.0.1
```

### 删除指定规则
```bash
netsh interface portproxy delete v4tov4 listenport=8081 listenaddress=0.0.0.0
```

### 删除所有规则
```bash
netsh interface portproxy reset
```

## ⚠️ 重要提示

### 1. 管理员权限
所有 `netsh interface portproxy` 命令都需要管理员权限。

**获取管理员权限**:
- 方法1: 右键PowerShell → "以管理员身份运行"
- 方法2: 使用批处理文件（自动请求权限）

### 2. 防火墙设置
如果从外部网络访问，需要允许端口通过防火墙：

```bash
# 添加防火墙规则
netsh advfirewall firewall add rule name="Port8081" dir=in action=allow protocol=TCP localport=8081

# 删除防火墙规则
netsh advfirewall firewall delete rule name="Port8081"
```

### 3. 端口冲突
确保端口未被占用：

```bash
# 检查端口占用
netstat -ano | findstr :8081
netstat -ano | findstr :9999
```

## 🧪 测试场景

### 场景1：基本功能测试
```bash
# 1. 启动Echo服务器
go run echo_server.go 9999

# 2. 配置端口转发
netsh interface portproxy add v4tov4 listenport=8081 listenaddress=0.0.0.0 connectport=9999 connectaddress=127.0.0.1

# 3. 运行测试客户端
go run test_client.go -server=127.0.0.1:8081 -count=5

# 4. 清理规则
netsh interface portproxy delete v4tov4 listenport=8081 listenaddress=0.0.0.0
```

### 场景2：并发测试
```bash
# 在多个终端窗口中同时运行
go run test_client.go -server=127.0.0.1:8081 -count=10 -interval=100
```

### 场景3：压力测试
```bash
# 发送1000条消息
go run test_client.go -count=1000 -interval=10
```

### 场景4：外部网络访问
```bash
# 从另一台机器访问（替换IP为本机局域网IP）
go run test_client.go -server=192.168.1.100:8081
```

## 📊 测试结果示例

### 成功的测试输出
```
========================================
     端口转发测试客户端
========================================
目标服务器: 127.0.0.1:8081
测试消息: Hello Test
发送次数: 5
发送间隔: 1000 ms
========================================

【第 1/5 次测试】
📡 已连接到服务器: 127.0.0.1:8081
⏰ 时间: 15:30:15.123
📤 发送: [1] Hello Test
📥 接收: Echo: [1] Hello Test [时间: 15:30:15]
✅ 测试通过

【第 2/5 次测试】
📡 已连接到服务器: 127.0.0.1:8081
⏰ 时间: 15:30:16.234
📤 发送: [2] Hello Test
📥 接收: Echo: [2] Hello Test [时间: 15:30:16]
✅ 测试通过

...

========================================
           测试结果统计
========================================
总测试次数: 5
✅ 成功: 5
❌ 失败: 0

🎉 所有测试通过！端口转发工作正常！
========================================
```

## 🔍 故障排查

| 问题 | 可能原因 | 解决方案 |
|------|----------|----------|
| 连接被拒绝 | Echo服务器未启动 | 启动echo_server.go |
| 连接超时 | 防火墙阻止 | 添加防火墙规则 |
| 权限不足 | 未使用管理员权限 | 以管理员身份运行 |
| 端口占用 | 端口已被使用 | 检查并释放端口 |
| 删除失败 | 权限不足或参数错误 | 使用管理员权限+检查参数 |

详细的故障排查请参考 [端口转发测试指南.md](./端口转发测试指南.md)

## 🎓 学习要点

### 端口转发原理
```
客户端 → [源IP:源端口] → 转发规则 → [目标IP:目标端口] → 目标服务器
        (例如 0.0.0.0:8081)    (127.0.0.1:9999)
```

### Windows netsh
- Windows自带的端口转发工具
- 在内核层拦截并转发TCP连接
- 对客户端和服务器透明
- 支持v4tov4、v4tov6、v6tov4、v6tov6

### 应用场景
- 内网穿透
- 端口映射
- 开发测试
- 服务迁移

## 📞 常见问题

### Q: 如何确认端口转发是否工作？
A: 运行 `test_client.go`，如果收到Echo响应则说明转发正常。

### Q: 为什么需要管理员权限？
A: Windows安全策略要求，修改网络配置需要管理员权限。

### Q: 端口转发会持续多久？
A: 重启计算机后会失效，需要重新配置。可创建启动脚本自动配置。

### Q: 如何查看端口转发的日志？
A: Windows端口转发本身不生成日志，需要查看应用程序日志。

### Q: 可以转发UDP端口吗？
A: `netsh interface portproxy` 只支持TCP。UDP转发需要其他工具。

## 📝 测试检查清单

完成测试前，请确认：

- [ ] Echo服务器成功启动
- [ ] 端口转发规则已配置
- [ ] `netsh interface portproxy show all` 显示正确规则
- [ ] 测试客户端成功连接
- [ ] Echo服务器收到转发的消息
- [ ] 客户端收到Echo响应
- [ ] 多次测试都成功
- [ ] 测试完成后已清理规则

## 🎉 完成测试

测试完成后，记得清理：

1. 停止Echo服务器（Ctrl+C）
2. 删除端口转发规则：
   ```bash
   .\删除端口转发规则.bat
   ```
3. 验证规则已清除：
   ```bash
   netsh interface portproxy show all
   ```

---

**准备开始测试？** 查看 [端口转发测试指南.md](./端口转发测试指南.md) 获取详细说明！🚀
