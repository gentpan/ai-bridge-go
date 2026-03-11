# AI Bridge Go

自托管 AI API 反向代理网关，专为解决中国大陆及香港地区无法直接访问 OpenAI、Claude、Google Gemini 等海外 AI 服务而设计。

## 项目背景

由于网络环境和政策限制，中国大陆和香港的服务器无法稳定访问 OpenAI、Anthropic (Claude)、Google Gemini 等海外 AI 服务的 API。虽然本地开发时可以借助代理工具，但对于线上运行的 WordPress 站点或其他服务端应用，直接从国内/港区服务器调用这些 AI API 几乎不可行。

AI Bridge 通过在海外部署一台轻量级代理网关来解决这一问题 —— 国内服务器将 AI 请求发送到你的海外网关，再由网关转发至 AI 服务商，实现稳定、安全的 AI API 访问。

> **⚠️ 部署要求：** 本网关必须部署在能够正常访问海外 AI 服务 API 的服务器上，如美国、日本、新加坡等地区的 VPS。**请勿部署在中国大陆或香港服务器上**，否则仍然无法访问目标 API。

**安全保障：** API Key 仅在你自己的服务器上流转，不经过任何第三方平台，杜绝密钥泄露风险。

> **相关仓库**
>
> - WordPress 插件：[gentpan/global-ai-bridge](https://github.com/gentpan/global-ai-bridge)
> - PHP 后端：[gentpan/ai-bridge-php](https://github.com/gentpan/ai-bridge-php)

## 工作原理

本网关是一个纯粹的 **反向代理（Reverse Proxy）**，不存储任何 API Key，不缓存任何对话内容。每个请求的处理流程：

```
┌─────────────┐         ┌─────────────────┐         ┌─────────────────┐
│  WordPress  │  POST   │   AI Bridge Go  │  POST   │   AI 服务商      │
│  (国内服务器) │ ──────→ │  (海外 VPS)     │ ──────→ │  (OpenAI 等)    │
│             │         │                 │         │                 │
│  插件发送    │         │  1. 读取 provider│         │                 │
│  AI 请求    │         │  2. 查内置地址表  │         │                 │
│             │         │  3. 用你的 API   │         │                 │
│             │ ←────── │     Key 转发请求 │ ←────── │  返回 AI 响应    │
│  收到响应    │  JSON   │  4. 原样回传结果  │  JSON   │                 │
└─────────────┘         └─────────────────┘         └─────────────────┘
```

- **插件端**（WordPress）：发送请求时附带 `provider`（如 openai）、`model`（如 gpt-4.1-mini）、消息内容和你的 API Key
- **网关端**（本项目）：根据 `provider` 查出对应的上游 API 地址（如 `https://api.openai.com/v1`），用你的 API Key 构造请求发给 AI 服务商，再把响应原样返回
- **网关不做任何 AI 处理**，纯网络中继，解决的是国内服务器无法直连海外 API 的网络问题

## 功能特性

- 支持 OpenAI、Claude、Google Gemini、DeepSeek 及任何 OpenAI 兼容 API
- 无需鉴权，直接使用 AI 服务商 API Key
- 内置速率限制、请求指标统计
- Docker 一键部署

## 快速开始

```bash
docker run -d \
  --name ai-bridge \
  -p 8080:8080 \
  ghcr.io/gentpan/ai-bridge-go:latest
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

| 变量                      | 默认值                                             | 说明                                                       |
| ------------------------- | -------------------------------------------------- | ---------------------------------------------------------- |
| `LISTEN_ADDR`             | `:8080`                                            | 监听地址                                                   |
| `DEFAULT_PROVIDER`        | `openai`                                           | 默认 AI 提供商                                             |
| `DEFAULT_MODEL`           | `gpt-4.1-mini`                                     | 默认模型                                                   |
| `REQUEST_TIMEOUT_SECONDS` | `60`                                               | 请求超时（秒）                                             |
| `RATE_LIMIT_PER_MINUTE`   | `120`                                              | 每分钟请求限制                                             |
| `NODE_NAME`               | `ai-bridge-node`                                   | 节点名称                                                   |
| `NODE_TRAFFIC_MODE`       | `outbound`                                         | 流量方向：`outbound`（中国→海外）或 `inbound`（海外→中国） |
| `OPENAI_BASE_URL`         | `https://api.openai.com/v1`                        | OpenAI API 地址                                            |
| `ANTHROPIC_BASE_URL`      | `https://api.anthropic.com/v1`                     | Claude API 地址                                            |
| `GOOGLE_BASE_URL`         | `https://generativelanguage.googleapis.com/v1beta` | Google Gemini API 地址                                     |
| `DEEPSEEK_BASE_URL`       | `https://api.deepseek.com/v1`                      | DeepSeek API 地址                                          |
| `QWEN_BASE_URL`           | —                                                  | 通义千问（配置后启用）                                     |
| `DOUBAO_BASE_URL`         | —                                                  | 豆包（配置后启用）                                         |
| `KIMI_BASE_URL`           | —                                                  | Kimi（配置后启用）                                         |

每个提供商还支持 `_ENABLED`（`true`/`false`）、`_DEFAULT_MODEL`、`_API_VERSION` 后缀变量。

## API 端点

| 方法   | 路径                   | 说明       |
| ------ | ---------------------- | ---------- |
| `GET`  | `/healthz`             | 健康检查   |
| `GET`  | `/metrics`             | 请求指标   |
| `POST` | `/v1/chat/completions` | 聊天完成   |
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

## 更多信息

- [PHP 后端](https://github.com/gentpan/ai-bridge-php)
- [WordPress 插件](https://github.com/gentpan/global-ai-bridge)

## License

MIT
