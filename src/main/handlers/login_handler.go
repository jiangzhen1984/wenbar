


package handlers

import (
    "gotom"
    "net/http"
    "fmt"
    "main/service/vo"
    "main/service"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         return tpls.Tpls["login"], nil, nil
     } else if req.Req.Method == "POST" {
         ws.DoService(ws.GetUserWS)
         if req.Req.FormValue("phoneNumber") == "13811962467" {
              sess := req.CreateSession(resp)
              sess.SetAttribute("user", &vo.User{Name : "test", Title :"test"})
              http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
         } else {
                return tpls.Tpls["login"], vo.LoginHtml{ PhoneNumber: "1ee", ErrMsg : "error" }, nil
         }
     } else {
         http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
         return nil, nil, nil
     }
    
     return nil, nil, fmt.Errorf("Not support method %s", req.Req.Method)
}



