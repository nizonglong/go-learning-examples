# Token V2

## 1. 概要描述
多端多设备登录。相较于 token v1，新增了根据设备类型和设备 ID 管控 token 数量，实现对设备登录的限制。

这是一个有状态的、带有设备限制的 token 方案。通过设备类型和设备 ID 生成 refresh token 和 access token，使用 access token 进行访问。存储在 Redis 进行管理，可对用户进行下线、token 撤销、token 刷新、设备限制和下线设备等操作。

### 1.1 优势
1. **设备管理：** 有效管理用户在不同设备上的登录情况，限制登录设备数量，提高账户的安全性。
2. **安全性提升：** 控制设备数量降低账户被恶意盗用的风险。达到登录设备数量上限时，新设备的登录会被阻止，降低了账户被未经授权访问的可能性。
3. **灵活性和可控性：** 根据业务需求灵活限制不同用户在不同设备上的登录次数，提供更细粒度的控制。

### 1.2 适用场景
1. **安全敏感应用：** 特别适用于安全性要求较高的应用场景，例如金融应用、医疗保健应用等。
2. **账户安全控制：** 需要控制和保护用户账户，防止未授权的设备访问。
3. **用户体验和安全平衡：** 在提供良好用户体验的同时，确保账户的安全性，设备数量控制提供了一种平衡。

## 2. 功能列表
- [x] 生成 token，带有设备数量限制
  - [x] refresh token + access token，不支持单独刷新 refresh token
  - [x] access token
- [x] 撤销 token
  - [x] 撤销 refresh token + access token
  - [x] 撤销 access token
- [x] 刷新 token
  - [x] 刷新 refresh token + access token
  - [x] 刷新 access token

## 3. 可实现的功能
- [x] 多端多设备登录且带有数量管控
- [x] 撤销用户设备登录态
- [x] 无感续租
