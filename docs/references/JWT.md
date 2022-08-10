#### JWT

分别由三部分组成：Header + Payload + Signature，之间用`.`进行分隔。

##### Header

标识头的信息：描述的是 JWT 的元数据，`alg`表示使用什么签名算法, `typ`表示是什么类型的。

- HS256：HMAC SHA256（默认的）

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

##### Payload

用来存放实际需要传递的数据, 标准包含了 7 个部分

```
- iss：签发者 (issuer)
- sub：主题 (subject)
- aud：接收者/受众 (audience)
- exp：过期时间 (expiration time)
- nbf：生效时间 (Not Before)
- iat：签发时间 (Issued At)
- jti：编号 (JWT ID)
```

除了以上字段外，还可以支持自定义字段。

```json
{
  "iss": "joe",
  "exp": 1300819380,
  "is_root": true
}
```

> JWT 默认是不加密的，任何人都可以读到，所以不要把秘密信息放在这个部分。这个 JSON 对象是使用 Base64URL 算法转成字符串

##### Signature

防止数据篡改，Signature 用来验证数据的完整性，防止数据被篡改的 (这个密钥只有服务器知道)，你可以理解为某种加盐的行为

> 这个部分才是保障数据安全的重要部分

##### 加密规则

```bash
HMACSHA256 (
    base64urlencode(Header) + "." +
    base64urlencode(Payload),
    Secret
)
```
