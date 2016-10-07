

package main

import (
    "gotom"
    "main/service"
    "main/service/wechat"
 //   "simulation"
)



func main() {
//    go sim.InitSIMWeChatServer()
    wechat.WeChatConfs = wechat.InitWeChatConfig("conf/conf.json")
    if wechat.WeChatConfs == nil || len(wechat.WeChatConfs) <= 0{
           gotom.LP("Initalize wechat config failed \n")
           return
    }
    wechat.DC().AuthServer()
    wechat.DC().AuthJS()
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
