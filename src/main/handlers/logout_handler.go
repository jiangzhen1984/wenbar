
package handlers

import (
    "gotom"
)


func LogoutHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     sess := req.GetSession()
     if sess != nil {
          sess.Invalidate()
     }

     Redirect(resp, req, "/login") 
     return nil, nil, nil
}



