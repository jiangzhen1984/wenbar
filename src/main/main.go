

package main

import (
    "gotom"
    "main/service"
    "main/service/wechat"
)



func main() {
    var  confs []* wechat.WeChatConfig =  wechat.InitWeChatConfig("")
    for _, val := range confs {
           val.AuthServer()
    }
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
