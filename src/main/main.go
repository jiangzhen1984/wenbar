

package main

import (
    "gotom"
    "main/handlers"
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
     DebugMode      : true,
     Port           : ":8080", 
     Tpldir         : "./view/",
     Mapping        : []*gotom.Mapping {
                           {"/hot_list",      handlers.HotListHandler},
                           {"/login",         handlers.LoginHandler},
                           {"/personal",      handlers.PersonalHandler},
                      },
     TplMapping     : map[string]*gotom.GTTemplateMapping {
                           "/hot_list" : {
                                             Uri  : "/hot_list",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "hot_list" : {Name : "hot_list" , Path : "view/hot_list.html"},
                                                    },
                                         }  ,
                           "/login"    : {
                                             Uri  : "/login",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "login" : {Name : "login"    , Path : "view/login.html"},
                                                    },
                                         }  ,
                           "/personal" : {
                                             Uri  : "/personal",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "personal" : {Name : "login"    , Path : "view/personal.html"},
                                                    },
                                         }  ,
                      },
}


func main() {
    gotom.InitServer(conf)
}
