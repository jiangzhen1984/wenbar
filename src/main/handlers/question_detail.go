



package handlers

import (
    "gotom"
    "fmt"
)


func QuestionDetailHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         return tpls.Tpls["questiondetail"], nil, nil
     }
    
     return nil, nil, fmt.Errorf("Not support method %s", req.Req.Method)
}



