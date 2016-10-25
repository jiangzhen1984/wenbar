

package main

import (
    "gotom"
    "main/handlers"
    "main/service"
    "main/service/wechat"
    "simulation"
    "os"
    "time"
    "encoding/json"
    "io/ioutil"
)


func printHelp() {
    var str string

    str = " ./main [-s] [-w] [-p] [-d] [-hc] [-h] \n"
    str +="        -s  use simulation wechat server for authorization\n"
    str +="        -w  wechat configuration file path \n"
    str +="        -p  listen port \n"
    str +="        -hc host configuration file path \n"
    str +="        -h  print this \n"
    str +="   E.G.\n"
    str +="        bin/main -s \n"
    str +="        bin/main -s  -p 8082\n"
    str +="        bin/main -s  -p 8082 -w conf/conf.json\n"
    str +="        bin/main -s  -p 8082 -w conf/conf.json  -hc host config \n"
    str +="\n\n"

    gotom.LI(str)
}



func main() {
    
    var simuflag bool = false
    var wechatConfigPath string
    var port string
    var debugmode bool = false
    var hostconf string
    for idx, arg := range os.Args {
         switch {
            case arg == "-s":
                simuflag = true
            case arg == "-w":
                if len(os.Args) - 1 <= idx {
                     gotom.LP(" argument incorrect -w filepath ")
                }
                wechatConfigPath = os.Args[idx + 1]
            case arg == "-p":
                if len(os.Args) - 1 <= idx {
                     gotom.LP(" argument incorrect -p port ")
                }
                port = os.Args[idx + 1]
            case arg == "-d":
                debugmode = true
            case arg == "-hc":
                if len(os.Args) - 1 <= idx {
                     gotom.LP(" argument incorrect -hc host_conf ")
                }
                hostconf = os.Args[idx + 1]
            case arg == "-h":
                printHelp();
                return
         }
    }


    if hostconf != "" {
         gotom.LD("===> read host config%s\n", hostconf)
         if data, err := ioutil.ReadFile(hostconf); err == nil {
              var hc handlers.HostConfig
              if err := json.Unmarshal(data, &hc); err == nil {
                   handlers.HostConf = &hc
                   gotom.LD("===%s   %s\n", hc, err)
              } else {
                   gotom.LE("===%s   %s\n", hc, err)
              }
          } else {
               gotom.LE("===> read host config%s\n", err)
          }
    } else {
          handlers.HostConf = new(handlers.HostConfig)
          gotom.LW("No host configuration ")
    }

    if port == "" {
         gotom.LW("Using default port 8081 for listening ")
         port = "8081"
    } else {
         gotom.LW("Using  port %s for listening ", port)
    }
    conf.Port = ":" + port
    conf.DebugMode = debugmode
    if simuflag == true {
         gotom.LI("Use local wechat simuation server")
         go sim.InitSIMWeChatServer()
         time.Sleep(3000)
    }
    if wechatConfigPath == "" {
         gotom.LW("No wechat configuration\n") 
    } else {
         wechat.WeChatConfs = wechat.InitWeChatConfig(wechatConfigPath)
         if wechat.WeChatConfs == nil || len(wechat.WeChatConfs) <= 0{
              gotom.LP("Initalize wechat config failed \n")
              return
         }
    }
    for _, conf := range wechat.WeChatConfs {
        conf.AuthServer()
        conf.AuthJS()
    }
    ws.InitDB(ws.DBConfiguration{DBUrl:"localhost"})
    gotom.InitServer(conf)
}
