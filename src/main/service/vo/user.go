
package vo

import (
    "strconv"
    "gotom"
    "time"
    "main/service/wechat"
)


type User struct {

     Uid       Wid
     
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
  
     Sex           int

     Province      string

     City          string

     Country       string

     Unionid       string

     Avatar        string

     Token         string

     TokenTime     int64

     TokenExpired  int
}



func (u User) HandleWeChatResponse(t int, us * wechat.WeChatUser, ret bool, data interface{}) {
     if !ret {
           gotom.LD(" Handle wechatuser auth failed \n")
           return
     }

     switch t {
          case wechat.RESPONSE_TYPE_USER_AUTH:
                u.getUserTokenInfo(us, ret, data)
          default:
                gotom.LW("Unkown type of response %d\n", t)
     }

}


func (u User) getUserTokenInfo(user * wechat.WeChatUser, ret bool, data interface{}) {
     ar := data.(wechat.AuthResponse)
     pwc := new(PersonalWeChat)
     pwc.OpenId        = ar.Openid
     pwc.Token         = ar.Access_token
     pwc.TokenTime     = time.Now().Unix()
     pwc.TokenExpired  = ar.Expires_in
    
     //TODO save user information to database 
}

