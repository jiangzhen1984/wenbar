



package handlers

import (
    "gotom"
    "main/service/vo"
)


func PersonalHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     sess := req.GetSession()
     if sess == nil || sess.GetAttribute("user") == nil {
         gotom.LE("====  not login session  %s  \n", sess)
         Redirect(resp, req, "/login")
         return nil, nil, nil
     }

     obj := sess.GetAttribute("user")
     user, ok := obj.(*vo.User)
     if ok == false {
         gotom.LE("==== type mismatch %s  \n", user)
         Redirect(resp, req, "/login") 
         return nil, nil, nil
     }

     if user.Personal == nil {
           //TODO load user data
     }

     data := vo.PersonalHtml{Name : user.Name,  Title : user.Title}

     return tpls.Tpls["personal"], data, nil
}



