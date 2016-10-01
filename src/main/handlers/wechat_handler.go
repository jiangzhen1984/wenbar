
package handlers

import (
    "gotom"
    "main/service/wechat"
    "main/service/vo"
)

func WeChatHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
          st := req.P("state")
          gotom.LD("===we chat parameter state %s\n", st)
          sess := req.GetSession()
          uobj := sess.GetAttribute("user")
          wuobj := sess.GetAttribute("wechatuser")
          gotom.LD("==== wenbar user ===>%s \n", uobj)
          gotom.LD("==== wechat user ===>%s \n", wuobj)

          user := uobj.(*vo.User)
          if wuobj != nil {
                wechatuser := wuobj.(*wechat.WeChatUser)
                wechatuser.AuthToken(user)
          }
     }

     return nil, nil, nil
}
