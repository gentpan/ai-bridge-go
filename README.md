# AI Bridge

High-performance Go backend gateway for AI Bridge — routes, authenticates, and proxies AI API requests across providers.

> **相关仓库**
> - WordPress 插件：[gentpan/global-ai-bridge](https://github.com/gentpan/global-ai-bridge)
> - PHP 后端：[gentpan/ai-bridge-php](https://github.com/gentpan/ai-bridge-php)

## 功能特性

- 支持 OpenAI、Claude、Google Gemini、DeepSeek 及任何 OpenAI 兼容 API
- 两种运行模式：托管模式（需 Token 鉴权）和自托管模式（免鉴权）
- 自动识别模式：配置了 `SITE_TOKEN` 就是托管模式，未配置就是自托管模式
- 内置速率限制、请求指标统计
- Token 申请与管理（托管模式）
- 邮件通知（托管模式）
- Docker / 二进制 / systemd 多种部署方式

## 快速开始

### 自托管模式（推荐个人用户）

不需要 Token，直接用你自己的 AI 服务商 API Key。

**Docker 一键部署：**

```bash
docker run -d \
  --name ai-bridge \
  -p 8080:8080 \
  -e OPENAI_BASE_URL=https://api.openai.com/v1 \
  ghcr.io/gentpan/ai-bridge-go:latest
```

**Docker Compose：**

```bash
curl -O https://raw.githubusercontent.com/gentpan/ai-bridge/main/docker-compose.selfhosted.yml
docker compose -f docker-compose.selfhosted.yml up -d
```

**二进制部署：**

```bash
# 从 Releases 下载对应平台的二进制
wget https://github.com/gentpan/ai-bridge-go/releases/latest/download/ai-bridge-linux-amd64
chmod +x ai-bridge-linux-amd64
./ai-bridge-linux-amd64
```

**验证：**

```bash
curl http://localhost:8080/healthz
# 返回 {"ok":true, "mode":"Self-Hosted", ...}
```

### 托管模式（多用户共享）

需要配置 `SITE_TOKEN`，用户通过邮件申请访问 Token。

```bash
docker run -d \
  --name ai-bridge \
  -p 8080:8080 \
  -e SITE_TOKEN=your-admin-token \
  -e EMAIL_PROVIDER=sendflare \
  -e EMAIL_API_KEY=your-key \
  -e OPENAI_BASE_URL=https://api.openai.com/v1 \
  ghcr.io/gentpan/ai-bridge-go:latest
```

## WordPress 插件配置

1. 安装 [AI Bridge 插件](https://github.com/gentpan/global-ai-bridge)
2. 进入 WordPress 后台 → 工具 → AI Bridge
3. 配置后端地址：`http://your-server:8080/v1/chat/completions`
4. **自托管模式**：AI Bridge 访问令牌留空，填入 AI 服务商 API Key
5. **托管模式**：填写申请到的访问令牌和 AI 服务商 API Key
6. 保存后点击「测速当前节点」验证连接

## 环境变量

### 基础配置

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `LISTEN_ADDR` | `:8080` | 监听地址 |
| `DEFAULT_PROVIDER` | `openai` | 默认 AI 提供商 |
| `DEFAULT_MODEL` | `gpt-4.1-mini` | 默认模型 |
| `REQUEST_TIMEOUT_SECONDS` | `60` | 请求超时（秒） |
| `RATE_LIMIT_PER_MINUTE` | `120` | 每分钟请求限制 |
| `NODE_NAME` | `ai-bridge-node` | 节点名称 |
| `NODE_TRAFFIC_MODE` | `outbound` | 流量方向：`outbound`（中国→海外）或 `inbound`（海外→中国） |
| `DATA_DIR` | `./data` | 数据目录（Token 存储等） |

### AI 提供商配置

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `OPENAI_BASE_URL` | `https://api.openai.com/v1` | OpenAI API 地址 |
| `ANTHROPIC_BASE_URL` | `https://api.anthropic.com/v1` | Claude API 地址 |
| `GOOGLE_BASE_URL` | `https://generativelanguage.googleapis.com/v1beta` | Google Gemini API 地址 |
| `DEEPSEEK_BASE_URL` | `https://api.deepseek.com/v1` | DeepSeek API 地址 |
| `QWEN_BASE_URL` | — | 通义千问 API 地址（配置后启用） |
| `DOUBAO_BASE_URL` | — | 豆包 API 地址（配置后启用） |
| `KIMI_BASE_URL` | — | Kimi API 地址（配置后启用） |

每个提供商还支持 `_ENABLED`（`true`/`false`）、`_DEFAULT_MODEL`、`_API_VERSION` 后缀变量。

### 托管模式配置（可选）

| 变量 | 说明 |
|------|------|
| `SITE_TOKEN` | 管理员 Token（配置后启用托管模式） |
| `SITE_TOKENS` | 多个 Token，逗号分隔 |
| `METRICS_TOKEN` | Metrics 接口鉴权 Token |
| `EMAIL_PROVIDER` | 邮件服务商：`sendflare` 或 `resend` |
| `EMAIL_API_KEY` | 邮件服务 API Key |
| `EMAIL_FROM_ADDR` | 发件人地址 |
| `EMAIL_FROM_NAME` | 发件人名称 |

## API 端点

| 方法 | 路径 | 说明 | 模式 |
|------|------|------|------|
| `GET` | `/healthz` | 健康检查 | 所有 |
| `GET` | `/metrics` | 请求指标 | 所有 |
| `POST` | `/v1/chat/completions` | 聊天完成 | 所有 |
| `POST` | `/v1/connectors/proxy` | 代理连接器 | 所有 |
| `POST` | `/api/apply-token` | 申请 Token | 托管 |
| `GET` | `/api/tokens` | Token 列表 | 托管 |
| `GET` | `/api/tokens/stats` | 使用统计 | 托管 |
| `POST` | `/api/tokens/revoke` | 吊销 Token | 托管 |

## 请求示例

```bash
# 自托管模式（无需 Authorization header）
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "X-AIBRIDGE-PROVIDER-TOKEN: sk-your-openai-key" \
  -d '{
    "provider": "openai",
    "model": "gpt-4.1-mini",
    "messages": [{"role": "user", "content": "Hello"}]
  }'

# 托管模式（需要 Authorization header）
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-site-token" \
  -H "X-AIBRIDGE-PROVIDER-TOKEN: sk-your-openai-key" \
  -d '{
    "provider": "claude",
    "model": "claude-sonnet-4-20250514",
    "messages": [{"role": "user", "content": "Hello"}]
  }'
```

## 构建

```bash
make build          # 编译二进制
make docker-build   # 构建 Docker 镜像
make docker-push    # 推送到 GHCR
```

## 更多文档

- [自托管部署指南](SELFHOSTED.md)
- [托管部署指南](DEPLOY.md)
- [部署检查清单](DEPLOY-CHECKLIST.md)
- [PHP 后端](https://github.com/gentpan/ai-bridge-php)

## License

MIT
