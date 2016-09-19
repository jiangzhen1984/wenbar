

package main

import (
    "gotom"
    "main/handlers"
    "html/template"
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
     Port           : ":8080", 
     Tpldir         : "./view/",
     Mapping        : []*gotom.Mapping {
                           {"/hot_list", handlers.HotListHandler},
                           {"/login", handlers.LoginHandler},
                      },
     TplMapping     : map[string]*gotom.GTTemplateMapping {
                           "/hot_list" : {
                                             Uri  : "/test1",
                                             Tpls : map[string]*template.Template{
                                                     "hot_list" : template.Must(template.ParseFiles("view/hot_list.html")),
                                                    },
                                         }  ,
                           "/login"    : {
                                             Uri  : "/login",
                                             Tpls : map[string]*template.Template{
                                                     "login" : template.Must(template.ParseFiles("view/login.html")),
                                                    },
                                         }  ,
                      },
}


func main() {
    gotom.InitServer(conf)
}
