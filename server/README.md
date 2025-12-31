# Docker GPU Manage（服务端）说明

本项目基于 tianqi，提供面向 GPU 工作负载的 Docker 实例生命周期管理与监控能力，支持远程 Docker 主机（含 TLS 双向认证）、GPU 分配、显存分割（HAMi）、系统/数据盘配置、容器日志与运行指标采集等。

后端核心容器管理逻辑位于：backend/service/instance/docker.go

- 关键结构体与服务
  - DockerService：容器连接、创建、启停、删除、日志与指标采集
  - ContainerConfig：创建容器时的规格描述（CPU/内存/磁盘/GPU/显存分割等）

## 功能特性

- 远程 Docker 连接
  - 支持 DOCKER_HOST=unix:///var/run/docker.sock 或 tcp://host:port
  - 支持 TLS 双向认证：CA、Client Cert、Client Key 通过数据库字段内联存储与加载
  - 提供 TestDockerConnection 用于连通性自检（cli.Ping）

- 容器生命周期管理
  - CreateContainer：根据镜像与产品规格创建并启动容器
  - StartContainer/StopContainer/RestartContainer：启停与重启
  - DeleteContainer：删除容器并清理其命名数据卷（/data 挂载的卷）

- 资源与规格配置
  - CPU：--cpus（通过 HostConfig.NanoCPUs 实现）
  - 内存：--memory=NGB
  - 系统盘：overlay2.size=NG（若后端不支持将自动回退为不限制）
  - 数据盘：创建命名卷 <name>-data 并挂载到 /data
  - GPU：--gpus N（通过 DeviceRequests，nvidia 驱动）

- 显存分割（HAMi）
  - 支持在产品规格开启 SupportMemorySplit 时，自动挂载 HAMi 库目录并设置环境变量：
    - LD_PRELOAD=/libvgpu/build/libvgpu.so
    - CUDA_DEVICE_MEMORY_LIMIT=<显存容量>g
    - CUDA_DEVICE_SM_LIMIT=<显存容量/单卡总容量>
  - HamiCore 路径默认 /root/HAMi-core/build，可在计算节点上通过字段配置覆盖

- 运行状态与日志
  - GetContainerStatus：读取容器状态（running/exited/…）
  - SyncContainerStatus：将状态同步回数据库字段 container_status
  - GetContainerLogs：获取容器标准输出/错误日志，支持 tail 与 timestamps

- 指标采集（CPU/内存/PIDs/GPU）
  - 优先通过 docker stats --no-stream --format {{json .}} 获取 CPU/内存/PIDs
    - 远程 TLS 模式会自动临时写入 ca.pem/cert.pem/key.pem 并设置 DOCKER_HOST、DOCKER_CERT_PATH、DOCKER_TLS_VERIFY
  - 额外采集 GPU 显存大小与利用率：
    - 首选容器内 nvidia-smi（兼容多 GPU、多行、带/不带单位）
    - 若容器内不可用，尝试宿主机 nvidia-smi + 环境变量 NVIDIA_VISIBLE_DEVICES 兜底
    - 再兜底读取 CUDA_DEVICE_MEMORY_LIMIT 环境变量，仅返回总显存大小
  - 指标结果缓存 30 秒（减少系统开销）

## 关键数据模型字段（节选）

- 计算节点 computenode.ComputeNode
  - DockerAddress：Docker 连接地址（如 tcp://10.0.0.5:2376）
  - UseTls：是否启用 TLS
  - CaCert / ClientCert / ClientKey：TLS 所需证书内容（PEM）
  - HamiCore：HAMi 核心库目录（默认 /root/HAMi-core/build）
  - MemoryCapacity / GpuCount：用于计算单卡显存容量（用于 CUDA_DEVICE_SM_LIMIT）

- 镜像 imageregistry.ImageRegistry
  - Address：镜像地址（如 registry.example.com/ns/image:tag）

- 规格 product.ProductSpec
  - CpuCores / MemoryGb / SystemDiskGb / DataDiskGb / GpuCount
  - SupportMemorySplit / MemoryCapacity（显存分割相关）

## 使用示例（伪代码）

```go
svc := &instance.DockerService{}
ctx := context.TODO()

// 1) 测试 Docker 连接
ok, msg := svc.TestDockerConnection(ctx, &computenode.ComputeNode{
    DockerAddress: ptr("tcp://10.0.0.5:2376"),
    UseTls:        ptr(true),
    CaCert:        ptr(caPEM),
    ClientCert:    ptr(cliCertPEM),
    ClientKey:     ptr(cliKeyPEM),
})
if !ok { log.Fatal(msg) }

// 2) 由镜像与产品规格构建容器配置
cfg := svc.BuildContainerConfig(image, spec, node, svc.GenerateInstanceName("demo", 123))

// 3) 创建并启动容器
id, err := svc.CreateContainer(ctx, node, cfg)
if err != nil { log.Fatal(err) }

// 4) 获取日志与指标
logs, _ := svc.GetContainerLogs(ctx, node, id, "100")
stats, _ := svc.GetContainerStats(ctx, node, id)

// 5) 删除容器（并清理命名数据卷）
_ = svc.DeleteContainer(ctx, node, id, cfg.Name)
```

提示：GetContainerLogs 的 tail 参数可设置为 "100"、"200" 等；指标采集需要本机可执行 docker（CLI），以及容器或宿主具备 nvidia-smi。

## 注意事项

- overlay2.size 对底层存储/文件系统有依赖，若启动失败会自动回退为不设置该参数重试
- GPU 指标采集需满足以下至少一项：
  - 容器内已安装 nvidia-smi；或
  - 服务与 Docker 宿主同机且宿主安装 nvidia-smi；并且容器配置了 NVIDIA_VISIBLE_DEVICES
- 启用 TLS 时需确保证书正确匹配目标 Docker 守护进程
- 删除容器会尝试清理挂载的命名卷，请谨慎操作

## server 项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── packfile
├── resource
│   ├── excel
│   ├── page
│   └── template
├── router
├── service
├── source
└── utils
    ├── timer
    └── upload
```

- 目录说明
  - api：接口层
  - config：配置结构体，与 config.yaml 对应
  - core：核心组件初始化（zap、viper、server）
  - docs：swagger 文档
  - global：全局对象
  - initialize：router、redis、gorm、validator、timer 等初始化
    - internal：仅供 initialize 调用的内部实现（如 gorm logger）
  - middleware：gin 中间件
  - model：数据库模型
    - request/response：请求与响应结构体
  - packfile：静态文件打包
  - resource：静态资源（excel、page、template）
  - router：路由
  - service：业务逻辑（本 README 重点的 DockerService 位于 backend/service/instance/docker.go）
  - source：初始化数据
  - utils：工具封装（timer、upload 等）

---

如需进一步完善文档（部署步骤、前后端联调、配置样例等），请告知具体需求。