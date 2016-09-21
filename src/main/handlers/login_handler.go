


package handlers

import (
    "gotom"
    "net/http"
    "fmt"
    "main/service/vo"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         return tpls.Tpls["login"], nil, nil
     } else if req.Req.Method == "POST" {
         //TODO update for wexin auth
         sess := req.CreateSession(resp)
         sess.SetAttribute("user", &vo.User{Name : "test", Title :"test"})
         http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
     } else {
         http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
         return nil, nil, nil
     }
    
     return nil, nil, fmt.Errorf("Not support method %s", req.Req.Method)
}



