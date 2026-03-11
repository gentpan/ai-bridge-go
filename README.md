# AI Bridge Go

自托管 AI 代理网关 —— 部署在你自己的服务器上，安全代理 AI API 请求。

API Key 只在你自己的服务器上流转，不经过任何第三方，避免泄露风险。

> **相关仓库**
> - WordPress 插件：[gentpan/global-ai-bridge](https://github.com/gentpan/global-ai-bridge)
> - PHP 后端：[gentpan/ai-bridge-php](https://github.com/gentpan/ai-bridge-php)

## 功能特性

- 支持 OpenAI、Claude、Google Gemini、DeepSeek 及任何 OpenAI 兼容 API
- 无需鉴权，直接使用 AI 服务商 API Key
- 内置速率限制、请求指标统计
- Docker / 二进制 / systemd 多种部署方式

## 快速开始

### Docker 一键部署（推荐）

```bash
docker run -d \
  --name ai-bridge \
  -p 8080:8080 \
  -e OPENAI_BASE_URL=https://api.openai.com/v1 \
  ghcr.io/gentpan/ai-bridge-go:latest
```

### Docker Compose

```bash
curl -O https://raw.githubusercontent.com/gentpan/ai-bridge-go/main/docker-compose.selfhosted.yml
docker compose -f docker-compose.selfhosted.yml up -d
```

### 二进制部署

```bash
wget https://github.com/gentpan/ai-bridge-go/releases/latest/download/ai-bridge-linux-amd64
chmod +x ai-bridge-linux-amd64
./ai-bridge-linux-amd64
```

### 验证

```bash
curl http://localhost:8080/healthz
# {"ok":true, "mode":"Self-Hosted", ...}
```

## WordPress 插件配置

1. 安装 [AI Bridge 插件](https://github.com/gentpan/global-ai-bridge)
2. 进入 WordPress 后台 → 工具 → AI Bridge
3. 连接方式选择「使用自己的服务器（自托管）」
4. 填入后端地址：`https://your-domain.com/v1/chat/completions`
5. **AI Bridge 访问令牌**：留空
6. **模型 API Token**：填入你的 OpenAI / Claude 等 API Key
7. 保存后点击「测速当前节点」验证

## 环境变量

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `LISTEN_ADDR` | `:8080` | 监听地址 |
| `DEFAULT_PROVIDER` | `openai` | 默认 AI 提供商 |
| `DEFAULT_MODEL` | `gpt-4.1-mini` | 默认模型 |
| `REQUEST_TIMEOUT_SECONDS` | `60` | 请求超时（秒） |
| `RATE_LIMIT_PER_MINUTE` | `120` | 每分钟请求限制 |
| `NODE_NAME` | `ai-bridge-node` | 节点名称 |
| `NODE_TRAFFIC_MODE` | `outbound` | 流量方向：`outbound`（中国→海外）或 `inbound`（海外→中国） |
| `OPENAI_BASE_URL` | `https://api.openai.com/v1` | OpenAI API 地址 |
| `ANTHROPIC_BASE_URL` | `https://api.anthropic.com/v1` | Claude API 地址 |
| `GOOGLE_BASE_URL` | `https://generativelanguage.googleapis.com/v1beta` | Google Gemini API 地址 |
| `DEEPSEEK_BASE_URL` | `https://api.deepseek.com/v1` | DeepSeek API 地址 |
| `QWEN_BASE_URL` | — | 通义千问（配置后启用） |
| `DOUBAO_BASE_URL` | — | 豆包（配置后启用） |
| `KIMI_BASE_URL` | — | Kimi（配置后启用） |

每个提供商还支持 `_ENABLED`（`true`/`false`）、`_DEFAULT_MODEL`、`_API_VERSION` 后缀变量。

## API 端点

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/healthz` | 健康检查 |
| `GET` | `/metrics` | 请求指标 |
| `POST` | `/v1/chat/completions` | 聊天完成 |
| `POST` | `/v1/connectors/proxy` | 代理连接器 |

## 请求示例

```bash
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "X-AIBRIDGE-PROVIDER-TOKEN: sk-your-openai-key" \
  -d '{
    "provider": "openai",
    "model": "gpt-4.1-mini",
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
- [部署检查清单](DEPLOY-CHECKLIST.md)
- [PHP 后端](https://github.com/gentpan/ai-bridge-php)

## License

MIT
