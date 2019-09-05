# 企业微信内部开发api

目前仅实现将要用到的接口，剩下的需要慢慢实现

```go
package main

import (
    "fmt"

    "github.com/wuwenbao/wechat/work"
    "github.com/wuwenbao/wechat/work/message"
)

conf := work.Conf("corpId", "corpSecret", nil)

app := work.New(conf)

//推送消息
msg := app.Message(1000002)
res, err:= msg.ToUser("snowman").Send(message.Text("test")
if err != nil {
    panic(err)
}
fmt.Println(res)

```