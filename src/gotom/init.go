
package gotom


import (
)

func InitContext() {
    SerCtx = new(GTServerContext)
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
