

package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
    "main/service/wechat"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LI("==== start call Login \n")
     gotom.LD("==== Cookies %s \n", req.Req.Cookies())

     if req.Req.Method == "GET" {
         from := req.Req.FormValue("from")
         gotom.LI("====> from :%s\n", from)
         if req.GetSession() != nil {
                sess := req.GetSession()
                gotom.LD("====> %s\n", sess)
                gotom.LD("====> %s\n", sess.GetAttribute("user"))
                gotom.LD("====> %s\n", sess.GetAttribute("wechatuser"))
                //TODO check user information and get from database
                Redirect(resp, req, "/hot_list") 
         }
         
         return tpls.Tpls["login"], &vo.LoginHtml{From : from}, nil
     } else if req.Req.Method == "POST" {
         if req.P("type") == "cellphone" {
              sess := req.CreateSession(resp)
              user := new(vo.User)
              user.Personal = new(vo.UserPersonal)
              ws.DoService(ws.CreateUser, user)
              sess.SetAttribute("user", user)
              
              Redirect(resp, req, "/hot_list") 
         } else {
              wechatuser := wechat.DC().InitWeChatUser()
              wechatuser.BuildAuthUrl("")
              Redirect(resp, req, wechatuser.WebAuthUrl) 
         }
     } else {
         Redirect(resp, req, "/hot_list") 
         return nil, nil, nil
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}



