# Token V1

## 概要描述
单点登录态，有状态的 token。通过生成 refresh token 和 access token，使用 access token 进行访问。存储在 Redis 进行管理，可对用户进行下线、token 撤销、token 刷新。

### 优势
1. **可控性和可撤销性：** 使用 Redis 存储 token 可以提供更好的可控性和撤销性。通过在服务器端存储 token，可以对其进行管理、撤销或设置过期时间。
2. **灵活性：** 可以根据需要存储更多的用户信息和状态信息，并可以在需要时更新或删除。

### 适用场景
1. **会话管理：** 适用于需要跟踪用户会话和状态的场景，例如需要记录用户的登录状态、权限或其他信息时。
2. **令牌刷新和撤销：** 如果需要更细粒度地控制令牌的刷新、撤销和管理，可以使用 Redis 存储。

## 功能列表
- [x] 生成 token
  - [x] refresh token + access token，不支持单独刷新 refresh token
  - [x] access token
- [x] 撤销 token
  - [x] 撤销 refresh token + access token
  - [x] 撤销 access token
- [x] 刷新 token
  - [x] 刷新 refresh token + access token
  - [x] 刷新 access token

## 可实现的功能
- [x] 带有状态的用户单点登录
- [x] 撤销用户登录态
- [x] 无感续租
