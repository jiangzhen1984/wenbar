


package handlers

import (
    "gotom"
    "net/http"
    "fmt"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         return &gotom.GTTemplate{tpls.Tpls["login"]}, nil, nil
     } else if req.Req.Method == "POST" {
         http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
         return nil, nil, nil
     }
    
     return nil, nil, fmt.Errorf("Not support method %s", req.Req.Method)
}



