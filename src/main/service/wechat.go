

package ws

import (
     "time"
     "gotom"
)


var WeChatConfs  []WeChatConfig


var weChatURL = "https://api.weixin.qq.com"
var requestUri []string = []string{"/sns/oauth2/access_token"} 
var WeChatUserAuthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?"


const (
    WE_REQ_URI_USER_TOKEN = 0
)

type WeChatConfig struct {

    AppId        string

    Secret       string

    GrantType    string

    Token        string

    Experies     int

    Timestamp    time.Time

    ResponseURL  string

    requestUri   []string

}


func InitWeChatConfig(path string) []*WeChatConfig {
     return nil
}


func (wcc * WeChatConfig) IsValid() bool {
      return false
}

func (wcc * WeChatConfig) IsExperied() int {
      return -1
}


func (wcc * WeChatConfig) AuthServer() bool {
      return false
}


func (wcc * WeChatConfig) getRequestUrl(ty int) string{
 
     switch {
          case ty == WE_REQ_URI_USER_TOKEN:
               return weChatURL + requestUri[WE_REQ_URI_USER_TOKEN] +"?appid=" + wcc.AppId+"&secret=" + wcc.Secret+"&grant_type=authorization_code"
     }
     return ""
}



const (
      WC_USER_INIT           = iota
      WC_USER_CODE
      WC_USER_TOKEN
      WC_USER_TOKEN_REFRESH
      WC_USER_EXPIRED
    
)


type WeChatUser struct {

     Code          string

     Token         string

     RefreshToken  string

     Timestamp     string

     Expired       int

     OpenId        string

     UnionId       string

     State         int 

     WebAuthUrl    string

     Conf          *WeChatConfig
}



func InitWeChatUser(wcc * WeChatConfig) (*WeChatUser, error) {
      if wcc == nil || wcc.IsValid() == false {
             return nil, gotom.ErrorMsg("WeChat Config doesn't initalize yet ==>%s\n", wcc)
      }

      return &WeChatUser{State : WC_USER_INIT, Conf : wcc}, nil
}


func (wcu * WeChatUser) BuildAuthUrl(wcc * WeChatConfig, redirectUrl string) {
      if wcu == nil {
            gotom.LE(" we chat user is nil\n")
            return
      }
      if wcc == nil {
            gotom.LE(" we chat config is nil\n")
            return
      }

      var redirect = redirectUrl
      if len(redirect) <=0 {
            redirect = wcc.ResponseURL
            gotom.LI("Use default redirect url %s\n", redirect)
      }
      url :=  WeChatUserAuthURL +  
              "appid=" + wcc.AppId +
              "&response_type=code" +
              "&scope=snsapi_userinfo" +
              "&response_type=code" +
              "&redirect_uri=" + redirect

      gotom.LD("user auth url:%s\n", url)

      wcu.WebAuthUrl = url
}


func (wc * WeChatUser) UpdateAuthCode(code string) {
       if wc == nil {
           gotom.LE(" illegal state for we chat user ")
           return
       } 

       if wc.State != WC_USER_INIT {
           gotom.LW(" we chat user illegal state %d  %s\n", wc.State, wc)
       }

       wc.State = WC_USER_CODE
       wc.Code = code

       //TODO trigger get user token and information
}


func (wc * WeChatUser) authToken(handler WeChatUserAuthRespHandler) {
    authTokenUrl := wc.Conf.getRequestUrl(WE_REQ_URI_USER_TOKEN)+"&code=" + wc.Code
    gotom.LD("get token url %s\n", authTokenUrl)

    //TODO send request and get url


    //TODO parse response
    handler.OnUserAuthResponse(wc, true, nil)
}


type  WeChatUserAuthRespHandler interface {
      OnUserAuthResponse(user * WeChatUser, ret bool, data interface{})
}
