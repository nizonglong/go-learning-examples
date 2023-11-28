## 项目名称：JWT 学习项目 V1

### 项目描述：

JWT 学习项目旨在提供对 JSON Web Token（JWT）的全面了解和应用实践。JWT 是一种用于无状态身份验证和授权的令牌，本项目旨在探索其特性、优势和适用场景。

### 主要特性：

- **轻量级和自包含性：** JWT 是一种轻量级的令牌，具有自包含性，无需额外存储状态信息。
- **无状态性：** 无需在服务器端存储 token 信息，减轻服务器负担。
- **跨平台和跨语言性：** 基于标准的 JWT 可在多种平台和语言之间使用。

### 适用场景：

- **分布式系统：** 适用于多个服务之间的身份验证和授权，特别是在分布式系统中。
- **前后端分离应用：** 可在前后端分离的应用中使用 JWT 进行身份验证和授权。

### 功能列表：

- **生成 token：** 实现了生成 JWT token 的功能。
- **校验 token：** 实现了校验 JWT token 的功能。

### 可实现的功能：

- **无状态的请求访问：** 适合于简单无场景的请求访问。

### 注意事项：

- **安全性和敏感信息存储：** JWT 应谨慎存储敏感信息，避免直接将敏感数据放置在 payload 中。