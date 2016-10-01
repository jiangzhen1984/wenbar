

package main

import (
    "gotom"
    "main/service"
    "main/service/wechat"
)



func main() {
    wechat.WeChatConfs = wechat.InitWeChatConfig("conf/conf.json")
    if wechat.WeChatConfs == nil || len(wechat.WeChatConfs) <= 0{
           gotom.LP("Initalize wechat config failed \n")
           return
    }
    wechat.DC().AuthServer()
    gotom.LI("[wechat] auth initalized ==>%s\n", WechatConf)
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
