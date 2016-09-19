
package handlers

import (
    "gotom"
    "net/http"
)


func LogoutHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     sess := req.GetSession()
     if sess != nil {
          sess.Invalidate()
     }

     http.Redirect(*resp.Resp, req.Req, "/login", 302) 
     return nil, nil, nil
}



