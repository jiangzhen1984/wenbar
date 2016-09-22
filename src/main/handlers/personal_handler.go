



package handlers

import (
    "gotom"
    "main/service/vo"
)


func PersonalHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     user := GetLoggedUser(req)
     if user == nil {
         Redirect(resp, req, "/login?from=/personal")
         return nil, nil, nil
     }

     if user.Personal == nil {
           //TODO load user data
     }

     data := vo.PersonalHtml{Name : user.Name,  Title : user.Title}

     return tpls.Tpls["personal"], data, nil
}



