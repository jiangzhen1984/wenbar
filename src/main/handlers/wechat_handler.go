
package handlers

import (
    "gotom"
    "main/service/wechat"
    "main/service/vo"
    "main/service"
    "strconv"
    "time"
)

type weChatRespHdr struct {
     
     user    * vo.User
     wechat  * wechat.WeChatUser
}


func (h * weChatRespHdr) OnResponse(t int, us * wechat.WeChatUser, ret bool, data interface{}) {
     if !ret {
           gotom.LD(" Handle wechat response  failed :%b \n", ret)
           return
     }

     switch t {
          case wechat.RESPONSE_TYPE_USER_AUTH:
                h.getUserTokenInfo(us, ret, data)
          case wechat.RESPONSE_TYPE_GET_USER_INFO:
                h.getUserInfo(us, ret, data)
          default:
                gotom.LW("Unkown type of response %d\n", t)
     }

}


func (h * weChatRespHdr) getUserTokenInfo(user * wechat.WeChatUser, ret bool, data interface{}) {
     var pwc *vo.PersonalWeChat
     ar := data.(*wechat.AuthResponse)
     if h.user.WeChat == nil {
         pwc = new(vo.PersonalWeChat)
         h.user.WeChat = pwc
     } else {
         pwc = h.user.WeChat 
     }
     pwc.OpenId        = ar.Openid
     pwc.Token         = ar.Access_token
     pwc.TokenTime     = time.Now().Unix()
     pwc.TokenExpired  = ar.Expires_in
    
     gotom.LI("====> got token user %s\n", h.user.WeChat)
     if _, err := ws.DoService(ws.UpdateUserWeChat, h.user); err != nil {
          gotom.LE("update wechat token failed %s\n", err)
     }
     gotom.LI("====> User updated  %s\n", h.user)
     // after token get user personal information
     go user.GetUserInfoFromServer(wechat.WeChatRespHandler(h))
     
}

func (h * weChatRespHdr) getUserInfo(user * wechat.WeChatUser, ret bool, data interface{}) {
     var pwc *vo.PersonalWeChat
     ar := data.(*wechat.UserInfoResp)
     if h.user.WeChat == nil {
         pwc = new(vo.PersonalWeChat)
         h.user.WeChat = pwc
     } else {
         pwc = h.user.WeChat 
     }
     pwc.NickName    = ar.NickName
     pwc.City        = ar.City
     pwc.Unionid     = ar.UnionId
     pwc.Avatar      = ar.Headimgurl
     pwc.Sex         = strconv.Itoa(ar.Sex)
     pwc.Country     = ar.Country
     h.user.Avatar1  = pwc.Avatar

     gotom.LI("====> get user info user %s\n", h.user.WeChat)
     if _, err := ws.DoService(ws.UpdateUserWeChat, h.user); err != nil {
          gotom.LE("update wechat information failed %s\n", err)
     }
}


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
         wechatuser.AuthToken(wechat.WeChatRespHandler(&weChatRespHdr{user : user, wechat : wechatuser}))
         sess.SetAttribute("wechatuser", wechatuser)
         Redirect(resp, req, "/hot_list")
     }

     return nil, nil, nil
}
