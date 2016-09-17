
package gotom


import (
      "net/http"
      "sync"
)


var ctxmu  * sync.Mutex = &sync.Mutex{}

func InitContext() *GTServerContext {
    ctxmu.Lock()
    defer ctxmu.Unlock()
    if SerCtx == nil {
         SerCtx = new(GTServerContext)
    }

    return SerCtx
}


func InitMapping(mappings ...*Mapping) {

    if SerCtx == nil {
       LE("Server Context doesn't initalize yet")
       return
    }
    var ret EnumRet
    for _ , mp := range mappings {
        ret = SerCtx.AddMapping(mp) 
        LD("====> Add Mapping ret %d  Mapping:%s\n", ret, mp)
    }
}

func InitMappings(mappings []*Mapping) {

    if SerCtx == nil {
       GoTomLogger.Fatal("Server Context doesn't initalize yet")
       return
    }
    var ret EnumRet
    for _ , mp := range mappings {
        ret = SerCtx.AddMapping(mp) 
        LD("====> Add Mapping ret %d Mapping:%s\n", ret, *mp)
    }
}



func InitServer(conf * GTConfig) {

   InitContext()

   if conf == nil {
         conf = & GTConfig{port : ":8080", tpldir : "./view/"}
   }

   LI(" Server Config :%s\n", conf)
   http.Handle("/", http.FileServer(http.Dir(conf.tpldir)))
   for key, value := range SerCtx.mapping {
        http.HandleFunc(key, value.hld.OnHandler)
   }
   http.ListenAndServe(conf.port, nil)
}
   


