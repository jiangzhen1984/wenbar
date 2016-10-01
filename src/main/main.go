

package main

import (
    "gotom"
    "main/service"
    "main/service/wechat"
)



func main() {
    wechat.WeChatConfs = []*wechat.WeChatConfig{WechatConf}
    WechatConf.AuthServer()
    gotom.LI("[wechat] auth initalized ==>%s\n", WechatConf)
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
