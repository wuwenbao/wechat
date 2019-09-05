# 微信支付

目前仅实现将要用到的接口，剩下的需要慢慢实现

```go
package main

import (
    "fmt"

    "github.com/wuwenbao/wechat/payment"
    "github.com/wuwenbao/wechat/payment/order"
)

conf := payment.Conf("", "", "")

app := payment.New(conf)

//统一下单
jsapiOrder := order.JSAPI()
jsapiOrder.各种参数 = 值
jsapiOrder.[] = []

res, err := app.Order().Unify(jsapiOrder)
if err != nil {
    panic(err)
}
fmt.Println(res)

```