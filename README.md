<div align="center">

# AI Bridge Go

**高性能 Go AI API 反向代理网关 · Docker 一键部署**

<p>
  <img src="https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Docker-ready-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/license-MIT-brightgreen?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/github/v/release/gentpan/ai-bridge-go?style=for-the-badge" alt="Release">
  <img src="https://img.shields.io/github/stars/gentpan/ai-bridge-go?style=for-the-badge" alt="Stars">
  <img src="https://img.shields.io/github/go-mod/go-version/gentpan/ai-bridge-go?style=for-the-badge" alt="Go Version">
</p>

<p>
  <a href="https://github.com/gentpan/ai-bridge-php">PHP 版</a> ·
  <a href="https://github.com/gentpan/ai-bridge">WordPress 插件</a>
</p>

</div>

---

## 📖 概述

AI Bridge Go 是一个高性能的自托管 AI API 反向代理网关，专为解决中国大陆及香港地区无法直接访问 OpenAI、Claude、Google Gemini 等海外 AI 服务而设计。

**推荐生产环境使用。** 相比 PHP 版，Go 版具有更高的并发性能、更低的资源占用，并支持 Caddy 自动 HTTPS、Docker 一键部署。

> ⚠️ **部署要求：** 必须部署在能正常访问海外 AI 服务 API 的服务器上（美国、日本、新加坡等地区的 VPS）。**请勿部署在中国大陆或香港服务器上。**

---

## 🔧 工作原理

本网关是一个纯粹的**反向代理**，不存储任何 API Key，不缓存任何对话内容：

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

---

## ✨ 特性

- ⚡ **高性能** — Go 原生并发，低内存占用
- 🐳 **Docker 一键部署** — 一条命令启动
- 🤖 **多平台支持** — OpenAI、Claude、Google Gemini、DeepSeek 及任何 OpenAI 兼容 API
- 🔑 **绝对安全** — API Key 仅在你的服务器上流转
- 🌐 **Caddy 自动 HTTPS** — 内建 TLS 支持
- 🔗 **兼容任何应用** — 不仅仅是 WordPress，任何 HTTP 客户端均可使用

---

## 🚀 快速开始

### Docker 部署（推荐）

```bash
docker run -d \
  --name ai-bridge \
  -p 9260:9260 \
  ghcr.io/gentpan/ai-bridge-go:latest
```

### Docker Compose

```bash
git clone https://github.com/gentpan/ai-bridge-go.git
cd ai-bridge-go
cp .env.example .env
# 编辑 .env 配置
docker compose up -d
```

### 验证

```bash
curl http://localhost:9260/healthz
# {"ok":true, "mode":"Self-Hosted", ...}
```

---

## 🔗 WordPress 插件配置

1. 安装 [AI Bridge 插件](https://github.com/gentpan/ai-bridge)
2. 进入 WordPress 后台 → 工具 → AI Bridge
3. 连接方式选择「使用自己的服务器（自托管）」
4. 填入后端地址：`https://your-domain.com/v1/chat/completions`
5. **AI Bridge 访问令牌**：留空
6. **模型 API Token**：填入你的 OpenAI / Claude 等 API Key
7. 保存后点击「测速当前节点」验证

---

## 🔗 关联项目

| 仓库 | 说明 |
|------|------|
| [ai-bridge](https://github.com/gentpan/ai-bridge) | WordPress 插件 + Go 后端组合包 |
| [ai-bridge-php](https://github.com/gentpan/ai-bridge-php) | PHP 单文件版（适合共享主机） |

---

## 📄 License

MIT
