


package handlers

import (
    "gotom"
    "main/service/vo"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LI("==== start call\n")

     if req.Req.Method == "GET" {
         from := req.Req.FormValue("from")
         gotom.LI("====> from :%s\n", from)
         
         return tpls.Tpls["login"], &vo.LoginHtml{From : from}, nil
     } else if req.Req.Method == "POST" {
         //TODO update for wexin auth
         sess := req.CreateSession(resp)
         sess.SetAttribute("user", &vo.User{Name : "test", Title :"test"})
         from := req.Req.FormValue("from")
         gotom.LI("====from :%s\n", from)
         n := len(from)
         if n <= 0 {
              from = "/hot_list"
         }
         Redirect(resp, req, from) 
     } else {
         Redirect(resp, req, "/hot_list") 
         return nil, nil, nil
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}



