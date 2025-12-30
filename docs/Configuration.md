# ⚙️ 配置说明

本页面详细介绍天启算力管理平台的配置项。

## 后端配置

后端配置文件位于 `server/config.yaml`。

### 系统配置

```yaml
system:
  db-type: mysql              # 数据库类型：mysql/pgsql/sqlite/mssql/oracle
  addr: 8888                  # 服务监听端口
  use-redis: false            # 是否使用 Redis（建议生产环境开启）
  oss-type: local             # 对象存储类型
  use-multipoint: false       # 多点登录拦截
```

### SSH 跳板机配置

```yaml
jumpbox:
  enabled: true               # 是否启用 SSH 跳板机
  port: 2026                  # SSH 监听端口（默认 2026）
  server-ip: "192.168.112.148"  # 服务器 IP（用于显示连接命令）
  host-key: ""                # SSH 主机密钥路径（可选，不设置则自动生成）
  banner: "欢迎使用SSH跳板机服务\r\n"  # SSH 欢迎信息
```

### JWT 配置

```yaml
jwt:
  signing-key: "your-key"     # JWT 签名密钥（生产环境请修改！）
  expires-time: 7d            # Token 过期时间
  buffer-time: 1d             # Token 缓冲时间
  issuer: gva                 # 签发者
```

### MySQL 配置

```yaml
mysql:
  path: 127.0.0.1             # 数据库地址
  port: 3306                  # 端口
  config: charset=utf8mb4&parseTime=True&loc=Local
  db-name: gva                # 数据库名
  username: root              # 用户名
  password: ""                # 密码
  prefix: ""                  # 表前缀
  singular: false             # 是否使用单数表名
  engine: ""                  # 数据库引擎（默认 InnoDB）
  max-idle-conns: 10          # 最大空闲连接数
  max-open-conns: 100         # 最大打开连接数
  log-mode: ""                # 日志模式
  log-zap: false              # 是否使用 zap 记录日志
```

### PostgreSQL 配置

```yaml
pgsql:
  path: ""
  port: 5432
  config: sslmode=disable TimeZone=Asia/Shanghai
  db-name: ""
  username: ""
  password: ""
  prefix: ""
  singular: false
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: ""
  log-zap: false
```

### Redis 配置

```yaml
redis:
  addr: 127.0.0.1:6379        # Redis 地址
  password: ""                # 密码
  db: 0                       # 数据库索引
  useCluster: false           # 是否使用集群
  clusterAddrs: []            # 集群地址列表
```

### Zap 日志配置

```yaml
zap:
  level: info                 # 日志级别：debug/info/warn/error
  prefix: "[GVA]"             # 日志前缀
  format: console             # 日志格式：console/json
  director: log               # 日志目录
  encode-level: LowercaseColorLevelEncoder  # 日志编码器
  stacktrace-key: stacktrace  # 栈跟踪键名
  max-age: 0                  # 日志保留天数（0 表示不删除）
  show-line: true             # 是否显示行号
  log-in-console: true        # 是否输出到控制台
```

### CORS 跨域配置

```yaml
cors:
  mode: whitelist             # 模式：whitelist/allow-all/strict-whitelist
  whitelist:
    - allow-origin: example.com
      allow-headers: Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id
      allow-methods: POST,GET,OPTIONS,DELETE,PUT
      expose-headers: Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type
      allow-credentials: true
```

---

## 前端配置

前端通过环境变量控制后端地址与代理，配置文件位于 `web/.env.*`。

### 开发环境 (.env.development)

```env
# 前端请求前缀（开发模式由 Vite 代理到后端）
VITE_BASE_API=/api

# 后端地址与端口
VITE_BASE_PATH=http://127.0.0.1
VITE_SERVER_PORT=8888

# 前端开发端口
VITE_CLI_PORT=8080

# 静态/上传文件基础路径
VITE_FILE_API=/uploads/file
```

### 生产环境 (.env.production)

```env
VITE_BASE_API=/api
VITE_BASE_PATH=http://your-production-server.com
VITE_SERVER_PORT=8888
VITE_CLI_PORT=80
VITE_FILE_API=/uploads/file
```

### 代理说明

- **开发模式**：`/api` 将被 Vite 代理到 `${VITE_BASE_PATH}:${VITE_SERVER_PORT}/`（详见 `web/vite.config.js`）
- **生产部署**：建议由 Nginx/网关将 `/api` 反向代理到后端服务
- **代码实现**：`web/src/utils/request.js` 使用 `import.meta.env.VITE_BASE_API` 作为 Axios 的 baseURL

---

## Nginx 配置示例

生产环境推荐使用 Nginx 作为反向代理：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 前端静态资源
    location / {
        root /path/to/web/dist;
        index index.html;
        try_files $uri $uri/ /index.html;
    }

    # API 反向代理
    location /api/ {
        proxy_pass http://127.0.0.1:8888/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
        # WebSocket 支持（Web 终端需要）
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }

    # 静态文件
    location /uploads/ {
        proxy_pass http://127.0.0.1:8888/uploads/;
    }
}
```

---

## 生产环境建议

### 安全性

- ✅ 修改默认管理员密码
- ✅ 修改 JWT 签名密钥（`jwt.signing-key`）
- ✅ 启用 HTTPS
- ✅ 配置防火墙规则
- ✅ 启用 TLS 连接 Docker

### 性能优化

- ✅ 启用 Redis 缓存（`system.use-redis: true`）
- ✅ 配置数据库连接池
- ✅ 使用 Nginx 反向代理
- ✅ 启用 Gzip 压缩

### 监控和日志

- ✅ 配置日志轮转（`zap.max-age`）
- ✅ 设置合适的日志级别
- ✅ 配置监控告警
- ✅ 定期备份数据库

