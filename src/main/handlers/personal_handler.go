



package handlers

import (
    "gotom"
    "net/http"
    "main/vo"
)


func PersonalHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     gotom.LD("=============>")
     sess := req.GetSession()
     if sess == nil || sess.GetAttribute("user") == nil {
         gotom.LE("====  not login session  %s  \n", sess)
         http.Redirect(*resp.Resp, req.Req, "/login", 302) 
         return nil, nil, nil
     }

     obj := sess.GetAttribute("user")
     user, ok := obj.(*vo.User)
     if ok == false {
         gotom.LE("==== type mismatch %s  \n", user)
         http.Redirect(*resp.Resp, req.Req, "/login", 302) 
         return nil, nil, nil
     }

     if user.Personal == nil {
           //TODO load user data
     }

     data := vo.PersonalHtml{Name : user.Name,  Title : user.Title}

     return &gotom.GTTemplate{tpls.Tpls["personal"]}, data, nil
}



