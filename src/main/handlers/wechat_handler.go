
package handlers

import (
    "gotom"
    "main/service/wechat"
    "main/service/vo"
)

func WeChatHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     gotom.LD("==== Cookies %s \n", req.Req.Cookies)
     if req.Req.Method == "GET" {
         st := req.P("state")
         code := req.P("code")
         gotom.LD("===we chat parameter state %s   code:%s\n", st, code)
         sess := req.CreateSession(resp)
         user := new(vo.User)
         sess.SetAttribute("user", user)

         wechatuser := wechat.DC().InitWeChatUser()
         wechatuser.UpdateAuthCode(code)
         wechatuser.AuthToken(user)
         sess.SetAttribute("wechatuser", wechatuser)
         //TODO get user info
         Redirect(resp, req, "/hot_list")

     }

     return nil, nil, nil
}
