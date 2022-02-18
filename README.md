# wecom-go-3rd-app-demo

使用 [go-zero](https://github.com/zeromicro/go-zero) 框架编写的企业微信第三方应用接收消息回调API,只做参数的解析验证不涉及逻辑处理

后台使用时数据回调URL填写：

    https://your.host.name/callback/data

指令回调URL填写：

    https://your.host.name/callback/cmd

请将配置中的`your.host.name`替换成真实域名地址

## 注意事项

- 服务器部署协议`http` or `https`
- 前端加`nginx`或其它代理到端口`8888`;端口号可在`etc/zero-api.yaml`里配置
- 注意修改`etc/suite-api.yaml`里的

    - CorpId
    - SuiteId
    - SuiteToken
    - SuiteEncodingAesKey
    - SuiteSecret
- post回调返回给微信端必需是小写的 `success`