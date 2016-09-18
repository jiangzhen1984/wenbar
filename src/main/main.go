

package main

import (
    "gotom"
)


func testHandler1(resp gotom.GTResponse, req * gotom.GTRequest) {
     sess := req.GetSession()
     if sess == nil {
           sess = req.CreateSession(resp)
           gotom.LD("==create session =%s\n", sess)
     }
     gotom.LD("===%s\n", sess)
}

var conf *gotom.GTConfig = &gotom.GTConfig { 
        Port    : ":8080", 
        Tpldir  : "./view/",
        Mapping : []*gotom.Mapping{{"/test1", testHandler1}, {"/test1", testHandler1}},
}


func main() {
    gotom.InitServer(conf)
}
