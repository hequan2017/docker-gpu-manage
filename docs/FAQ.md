# ❓ 常见问题

本页面汇总了天启算力管理平台的常见问题与解决方案。

## 部署相关

### 1. 数据库连接失败

**问题描述**：启动后端服务时提示数据库连接失败。

**解决方案**：
- 检查数据库服务是否启动
- 检查 `config.yaml` 中的数据库配置是否正确
- 检查数据库用户是否有创建数据库的权限
- 检查网络连接和防火墙设置

```yaml
# 检查配置项
mysql:
  path: 127.0.0.1        # 确认地址正确
  port: 3306             # 确认端口正确
  db-name: gva           # 确认数据库名
  username: root         # 确认用户名
  password: "yourpass"   # 确认密码
```

### 2. 前端无法连接后端

**问题描述**：前端页面加载正常，但 API 请求失败。

**解决方案**：
- 检查后端服务是否正常启动（默认端口 8888）
- 检查 `.env.development` 中的配置是否正确：

```env
VITE_BASE_PATH=http://127.0.0.1
VITE_SERVER_PORT=8888
```

- 检查浏览器控制台的网络请求，确认请求地址
- 检查 `vite.config.js` 中的代理配置是否生效
- 检查防火墙和端口是否开放

### 3. 初始化页面无法访问

**问题描述**：访问前端地址后页面空白或报错。

**解决方案**：
- 确保 Node.js 版本 >= 20
- 删除 `node_modules` 并重新安装依赖：

```bash
cd web
rm -rf node_modules package-lock.json
npm install
npm run dev
```

---

## Docker 容器相关

### 4. Docker 容器无法连接

**问题描述**：创建算力节点后，Docker 连接状态显示「连接失败」。

**解决方案**：

1. **检查 Docker 服务是否运行**
```bash
systemctl status docker
```

2. **检查 Docker API 是否开放远程访问**
```bash
# 编辑 /etc/docker/daemon.json
{
  "hosts": ["unix:///var/run/docker.sock", "tcp://0.0.0.0:2376"],
  "tls": true,
  "tlscacert": "/etc/docker/ca.pem",
  "tlscert": "/etc/docker/server-cert.pem",
  "tlskey": "/etc/docker/server-key.pem"
}
```

3. **检查 TLS 证书配置**
   - 确保 CA 证书、客户端证书、客户端私钥正确配置
   - 证书格式应为 PEM 格式

4. **检查防火墙**
```bash
# 开放 Docker 端口
firewall-cmd --permanent --add-port=2376/tcp
firewall-cmd --reload
```

### 5. 容器创建失败

**问题描述**：创建实例时提示容器创建失败。

**解决方案**：

1. **检查镜像是否存在**
   - 确认镜像地址正确
   - 在算力节点上手动拉取镜像测试：
   ```bash
   docker pull <镜像地址>
   ```

2. **检查资源是否充足**
   - 确认算力节点的 CPU、内存、磁盘资源充足
   - 检查 GPU 是否被占用

3. **检查 overlay2.size 支持**
   - 如果存储驱动不支持 `overlay2.size`，系统会自动跳过该参数重试

4. **查看后端日志**
   - 检查 `server/log/` 目录下的日志文件获取详细错误信息

### 6. GPU 不可用

**问题描述**：容器启动后无法使用 GPU。

**解决方案**：

1. **检查 NVIDIA 驱动**
```bash
nvidia-smi
```

2. **检查 NVIDIA Container Toolkit**
```bash
# 安装 NVIDIA Container Toolkit
distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add -
curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | sudo tee /etc/apt/sources.list.d/nvidia-docker.list
sudo apt-get update && sudo apt-get install -y nvidia-container-toolkit
sudo systemctl restart docker
```

3. **测试 GPU 容器**
```bash
docker run --rm --gpus all nvidia/cuda:11.0-base nvidia-smi
```

---

## SSH 跳板机相关

### 7. SSH 跳板机无法连接

**问题描述**：无法通过 SSH 连接到跳板机服务。

**解决方案**：

1. **检查跳板机是否启用**
```yaml
# config.yaml
jumpbox:
  enabled: true    # 确保为 true
  port: 2026       # 确认端口
```

2. **检查端口是否被占用**
```bash
netstat -tlnp | grep 2026
```

3. **检查防火墙**
```bash
firewall-cmd --permanent --add-port=2026/tcp
firewall-cmd --reload
```

4. **查看后端日志**
   - 检查是否有 SSH 服务启动相关的错误日志

### 8. SSH 连接后无法选择容器

**问题描述**：SSH 登录成功，但容器列表为空。

**解决方案**：
- 确认当前用户已创建实例
- 普通用户只能看到自己创建的容器
- 管理员（authorityId=888）可以看到所有容器
- 检查实例的容器状态是否为 `running`

---

## 显存切分相关

### 9. 显存切分功能不生效

**问题描述**：创建支持显存切分的实例后，显存限制不生效。

**解决方案**：

1. **检查 HAMi-core 部署**
   - 确认在算力节点上正确编译了 HAMi-core
   - 参考：https://github.com/Project-HAMi/HAMi-core

2. **检查算力节点配置**
   - 确认「HAMi-core 目录」字段填写正确（如 `/root/HAMi-core/build`）

3. **检查镜像配置**
   - 确认镜像的「是否支持显存切分」已勾选

4. **检查产品规格配置**
   - 确认产品规格的「是否支持显存切分」已勾选
   - 确认「显存容量」字段已正确填写

5. **验证容器环境变量**
```bash
# 进入容器检查环境变量
docker exec -it <容器ID> env | grep CUDA
# 应该看到：
# LD_PRELOAD=/libvgpu/build/libvgpu.so
# CUDA_DEVICE_MEMORY_LIMIT=<显存大小>g
# CUDA_DEVICE_SM_LIMIT=<SM限制>
```

---

## 权限相关

### 10. 普通用户无法操作实例

**问题描述**：普通用户无法看到或操作实例。

**解决方案**：
- 普通用户只能操作自己创建的实例
- 确认实例的创建用户是否为当前登录用户
- 检查用户的角色权限配置

### 11. API 权限不足

**问题描述**：调用某些 API 时提示权限不足。

**解决方案**：
- 在「超级管理员」->「菜单管理」中添加对应的 API 权限
- 在「超级管理员」->「角色管理」中为角色分配 API 权限
- 刷新页面或重新登录

---

## 性能相关

### 12. 系统响应缓慢

**问题描述**：系统整体响应速度较慢。

**解决方案**：

1. **启用 Redis 缓存**
```yaml
system:
  use-redis: true

redis:
  addr: 127.0.0.1:6379
```

2. **优化数据库连接池**
```yaml
mysql:
  max-idle-conns: 10
  max-open-conns: 100
```

3. **调整日志级别**
```yaml
zap:
  level: warn    # 生产环境建议使用 warn 或 error
```

### 13. 定时任务影响性能

**问题描述**：定时任务执行时系统变慢。

**解决方案**：
- 定时任务默认每 30 秒执行一次
- 如果节点和实例数量较多，可以适当调整检查间隔
- 定时任务日志已优化，仅保留必要的 Error/Warn 日志

---

## 其他问题

### 14. 如何重置管理员密码？

**解决方案**：

直接在数据库中更新密码：
```sql
-- 默认密码 123456 的加密值
UPDATE sys_users SET password = 'e10adc3949ba59abbe56e057f20f883e' WHERE username = 'admin';
```

### 15. 如何查看系统日志？

**解决方案**：

后端日志位于 `server/log/` 目录下，按日期分目录存储：
```bash
ls server/log/
# 2025-12-30/
# ...

cat server/log/2025-12-30/server.log
```

### 16. 如何备份数据库？

**解决方案**：

```bash
# MySQL 备份
mysqldump -u root -p gva > gva_backup_$(date +%Y%m%d).sql

# 恢复
mysql -u root -p gva < gva_backup_20251230.sql
```

---

## 获取帮助

如果以上方案无法解决您的问题，请：

1. 查看后端日志获取详细错误信息
2. 在 [GitHub Issues](https://github.com/hequan2017/docker-gpu-manage/issues) 提交问题
3. 提供以下信息以便排查：
   - 系统环境（OS、Go 版本、Node.js 版本）
   - 错误日志截图或文本
   - 复现步骤

