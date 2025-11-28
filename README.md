# Go Cert Checker 🔐

一个用 **Go** 写的轻量级 SSL/TLS 证书检查工具，可以快速查看证书信息、有效期和剩余天数，并判断证书状态（有效 / 已过期 / 尚未生效）。

---

## 功能

- ✅ 查看证书 CN（通用名称）和 SANs（域名列表）  
- ✅ 显示签发机构、有效期、序列号  
- ✅ 剩余天数计算  
- ✅ 判断证书状态（有效 / 已过期 / 尚未生效）  
- ✅ 支持自定义端口和连接超时  

---

## 安装

```bash
go install github.com/wantnotshould/go-cert-checker@latest
````

---

## 使用方法

```bash
# 检查默认 443 端口
go-cert-checker -domain=example.com

# 指定端口
go-cert-checker -domain=example.com -port=8443

# 指定连接超时（秒）
go-cert-checker -domain=example.com -timeout=10
```

---

## 输出示例

```
🌐 域名: musei.cn
🔑 证书 CN: musei.cn
📜 SANs: [musei.cn]
🏛️ 签发机构: CN=E8,O=Let's Encrypt,C=US
📅 生效时间: 2025-10-16 20:22:53
📅 到期时间: 2026-01-14 20:22:52
🆔 序列号: 571427202731249133172236069215885298715265
⏳ 剩余天数: 47 天
✅ 证书状态: 有效
```

---

## 注意事项

* 程序会忽略证书有效性验证，只用于查看信息，不保证安全通信
* 端口范围必须在 `1-65535`
* 如果证书过期或尚未生效，会在输出中明确提示

---

## 许可证

MIT License
