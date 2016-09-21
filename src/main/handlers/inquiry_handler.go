



package handlers

import (
    "gotom"
    "net/http"
)


func InquiryHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         return tpls.Tpls["inquiry"], nil, nil
     } else if req.Req.Method == "POST" {
         http.Redirect(*resp.Resp, req.Req, "/my_inquiry", 301) 
     } else {
         http.Redirect(*resp.Resp, req.Req, "/hot_list", 301) 
         return nil, nil, nil
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}



