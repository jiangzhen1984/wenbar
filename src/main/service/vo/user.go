
package vo

import (
    "strconv"
    "gotom"
    "time"
    "main/service/wechat"
)


type User struct {

     Uid       Wid     `json:"id" bson:"_id,omitempty"`
     
     Name      string 

     Title     string

     Avatar1   string
  
     Avatar2   string

     NativeId  uint64

     OutId     string

     Personal  *UserPersonal
  
     WeChat    *PersonalWeChat
}


func (u User) GetNativeID() string {
     return strconv.FormatUint(u.NativeId, 10)
}



type UserPersonal struct {

     Incoming     float32

     Revenue      float32

     QuesFee      float32 

     Ans          int
 
     BeViewed     int

     UnAns        int
}


type PersonalWeChat struct {

     OpenId        string

     NickName      string
  
     Sex           string

     Province      string

     City          string

     Country       string

     Unionid       string

     Avatar        string

     Token         string

     TokenTime     int64

     TokenExpired  int
}



func (u User) OnResponse(t int, us * wechat.WeChatUser, ret bool, data interface{}) {
     if !ret {
           gotom.LD(" Handle wechatuser auth failed \n")
           return
     }

     switch t {
          case wechat.RESPONSE_TYPE_USER_AUTH:
                u.getUserTokenInfo(us, ret, data)
          case wechat.RESPONSE_TYPE_GET_USER_INFO:
                u.getUserInfo(us, ret, data)
          default:
                gotom.LW("Unkown type of response %d\n", t)
     }

}


func (u User) getUserTokenInfo(user * wechat.WeChatUser, ret bool, data interface{}) {
     var pwc *PersonalWeChat
     ar := data.(*wechat.AuthResponse)
     if u.WeChat == nil {
         pwc = new(PersonalWeChat)
         u.WeChat = pwc
     } else {
         pwc = u.WeChat 
     }
     pwc.OpenId        = ar.Openid
     pwc.Token         = ar.Access_token
     pwc.TokenTime     = time.Now().Unix()
     pwc.TokenExpired  = ar.Expires_in
    
     gotom.LI("====> got token user %s\n", ar)
 /*    gobj := gotom.Object(&u)
     if _, err := ws.DoService(ws.UpdateUserWeChat, &gobj); err != nil {
          gotom.LE("update wechat token failed %s\n", err)
     }
*/
}

func (u User) getUserInfo(user * wechat.WeChatUser, ret bool, data interface{}) {
     var pwc *PersonalWeChat
     ar := data.(*wechat.UserInfoResp)
     if u.WeChat == nil {
         pwc = new(PersonalWeChat)
         u.WeChat = pwc
     } else {
         pwc = u.WeChat 
     }
     pwc.NickName    = ar.NickName
     pwc.City        = ar.City
     pwc.Unionid     = ar.UnionId
     pwc.Avatar      = ar.Headimgurl
     pwc.Sex         = ar.Sex
     pwc.Country     = ar.Country
     u.Avatar1       = pwc.Avatar

     gotom.LI("====> get user info user %s\n", ar)
/*     gobj := gotom.Object(&u)
     if _, err := ws.DoService(ws.UpdateUserWeChat, &gobj); err != nil {
          gotom.LE("update wechat information failed %s\n", err)
     }
*/
}

