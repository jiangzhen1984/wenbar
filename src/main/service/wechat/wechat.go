

package wechat

import (
     "encoding/json"
     "io/ioutil"
     "net/http"
     "time"
     "gotom"
     "net/url"
     "crypto/sha1"
     "math/rand"
     "strconv"
     "fmt"
     "os"
)


var WeChatConfs  []*WeChatConfig


var WeChatURL = "https://api.weixin.qq.com"
var WeChatFileURL = "http://file.api.weixin.qq.com"
var requestUri []string = []string{
                           "/cgi-bin/token",
                           "/sns/oauth2/access_token",
                           "/sns/userinfo",
                           "/cgi-bin/ticket/getticket",
                           "/cgi-bin/media/get",
                           } 
var WeChatUserAuthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?"


const (
    WE_REQ_URI_SERVER_TOKEN = iota
    WE_REQ_URI_USER_TOKEN
    WE_REQ_URI_GET_USER_INFO
    WE_REQ_URI_GET_JS_AUTH
    WE_REQ_URI_GET_MEDIA_FILE
)

type WeChatConfig struct {

    AppId        string

    Secret       string

    GrantType    string

    Token        string

    Expeires     int

    Timestamp    time.Time

    ResponseURL  string

    requestUri   []string

    Js           *WeChatJS
}


type WeChatJS struct {
    AppId       string
    Ticket      string 
    JsApiList   []string
}


func InitWeChatConfig(path string) []*WeChatConfig {
     if data, err := ioutil.ReadFile(path); err == nil {
           var confs []*WeChatConfig
           if err := json.Unmarshal(data, &confs); err == nil {
                   for _, wcc := range confs {
                         wcc.requestUri = requestUri
                         wcc.Js = &WeChatJS{AppId : wcc.AppId}
                   }
                   return confs
           } else {
                   gotom.LE("read wechat config faild %s   path:%s\n", err, path)
           }
     } else {
          gotom.LE("read wechat config faild %s   path:%s\n", err, path)
     }
     return nil
}


func GetDefaultWeChatConfig() * WeChatConfig {
     if WeChatConfs != nil && len(WeChatConfs) > 0 {
           return WeChatConfs[0]
     }

     gotom.LP("Doesn initalize wechat config yet!\n")
     return nil
}


func DC() * WeChatConfig {
    return GetDefaultWeChatConfig()
}


func (wcc * WeChatConfig) IsValid() bool {
      return len(wcc.AppId) > 0 && len(wcc.Secret) > 0
}

func (wcc * WeChatConfig) IsExperied() int {
      return -1
}


func (wcc * WeChatConfig) AuthServer() bool {

      var ret bool = false
      url := wcc.getRequestUrl(WE_REQ_URI_SERVER_TOKEN)
      gotom.LD(" auth server url %s\n", url)
      ar := &AuthResponse{}
      err := readDataFromServer(url, ar)
      gotom.LD("===>%s\n", ar)
      if err == nil {
           wcc.Token    = ar.Access_token
           wcc.Expeires = ar.Expires_in
           ret = true
      }
      return ret
}


func (wcc * WeChatConfig) AuthJS() {
      url := wcc.getRequestUrl(WE_REQ_URI_GET_JS_AUTH) +"access_token=" + wcc.Token +"&type=jsapi"
      ar := &JSAuthResponse{}
      err := readDataFromServer(url, ar)
      gotom.LD("===get js auth response : %s\n", ar)
      if err == nil {
           wcc.Js.Ticket = ar.Ticket           
      }
}


func (wcc * WeChatConfig) DownloadMediaFile(mid string, fp string) (bool) {
      if wcc.Token == "" {
           gotom.LE(" doesn't get token yet\n")
           return false
      }
      url := wcc.getRequestUrl(WE_REQ_URI_GET_MEDIA_FILE) +"&media_id=" + mid
      resp, err := http.Get(url)
      if err != nil {
           gotom.LE("Create response failed %s\n", err)
           return false
      }
      content, err := ioutil.ReadAll(resp.Body)
      defer resp.Body.Close()
      if err != nil {
          gotom.LE("Get Media file failed %s\n", err)
          return false
      }

      err = ioutil.WriteFile(fp, content, os.ModePerm)
      if err != nil {
          gotom.LE("Write file failed %s\n", err)
          return false
      }
      return true
   
}




func (wcc * WeChatConfig) getRequestUrl(ty int) string {
 
     switch {
          case ty == WE_REQ_URI_SERVER_TOKEN:
               return WeChatURL + wcc.requestUri[WE_REQ_URI_SERVER_TOKEN] +"?appid=" + wcc.AppId+"&secret=" + wcc.Secret+"&grant_type=client_credential"
          case ty == WE_REQ_URI_USER_TOKEN:
               return WeChatURL + wcc.requestUri[WE_REQ_URI_USER_TOKEN] +"?appid=" + wcc.AppId+"&secret=" + wcc.Secret+"&grant_type=authorization_code"
          case ty == WE_REQ_URI_GET_USER_INFO:
               return WeChatURL + wcc.requestUri[WE_REQ_URI_GET_USER_INFO] +"?appid=" + wcc.AppId+"&lang=zh_CN"
          case ty == WE_REQ_URI_GET_JS_AUTH:
               return WeChatURL + wcc.requestUri[WE_REQ_URI_GET_JS_AUTH] +"?"
          case ty == WE_REQ_URI_GET_MEDIA_FILE:
               return WeChatFileURL + wcc.requestUri[WE_REQ_URI_GET_MEDIA_FILE] +"?access_token="+wcc.Token

     }
     return ""
}


func (wjs *WeChatJS) GetSign(nonce string, timestamp time.Time, url string) (string) {
     ts := timestamp.Unix()
     str := "jsapi_ticket="+wjs.Ticket+"&noncestr="+nonce+"&timestamp="+strconv.Itoa(int(ts))+"&url=" + url
     ret := sha1.Sum([]byte(str))
     sign := fmt.Sprintf("%x", ret)
     return sign
}

func (wjs *WeChatJS) S(url string) (string, time.Time, string) {
     letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
     rand.Seed(time.Now().UnixNano())
     b := make([]rune, 12)
     for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
     }

     nonce := string(b)
     tr    := time.Now()
     sign  := wjs.GetSign(nonce, tr, url)
     return nonce, tr, sign
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



func (wcc * WeChatConfig) InitWeChatUser() *WeChatUser {
      if wcc == nil || wcc.IsValid() == false {
             gotom.ErrorMsg("WeChat Config doesn't initalize yet ==>%s\n", wcc)
             return nil
      }

      return &WeChatUser{State : WC_USER_INIT, Conf : wcc}
}


func (wcu * WeChatUser) BuildAuthUrl(redirectUrl string) {
      if wcu == nil {
            gotom.LE(" we chat user is nil\n")
            return
      }
      if wcu.Conf == nil {
            gotom.LE(" we chat config is nil\n")
            return
      }

      var redirect = redirectUrl
      if len(redirect) <=0 {
            redirect = wcu.Conf.ResponseURL
            gotom.LI("Use default redirect url %s\n", redirect)
      }
      url :=  WeChatUserAuthURL +  
              "appid=" + wcu.Conf.AppId +
              "&redirect_uri=" + url.QueryEscape(redirect) +
              "&response_type=code" +
              "&scope=snsapi_userinfo" +
              "&state=wechatuserauth" +
              "#wechat_redirect"

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
}


func (wc * WeChatUser) AuthToken(handler WeChatRespHandler) {

    if wc.State != WC_USER_CODE {
          gotom.LI("Can not get user auth token not in code state\n")
          handler.OnResponse(RESPONSE_TYPE_USER_AUTH, wc, false, nil)
          return
    }
    authTokenUrl := wc.Conf.getRequestUrl(WE_REQ_URI_USER_TOKEN)+"&code=" + wc.Code
    gotom.LD("get token url %s\n", authTokenUrl)

    ar := &AuthResponse{}
    err := readDataFromServer(authTokenUrl, ar)
    var ret bool = false
    if err != nil {
         ret = false
    } else {
         ret = true
    }
    gotom.LD("  parse %s   %s \n", err, ar)

    wc.State  =  WC_USER_TOKEN
    wc.Token  =  ar.Access_token
    wc.OpenId =  ar.Openid
 
    handler.OnResponse(RESPONSE_TYPE_USER_AUTH, wc, ret, ar)
}

func (wc * WeChatUser) GetUserInfoFromServer(handler WeChatRespHandler) error {
     var ret bool = true
     if wc.State != WC_USER_TOKEN {
           return gotom.ErrorMsg(" User does not get token yet")
     }

     gotom.LD("===== start to get user info\n")
     userinfourl := wc.Conf.getRequestUrl(WE_REQ_URI_GET_USER_INFO) + "&access_token=" + wc.Token+"&openid=" + wc.OpenId
     ur := &UserInfoResp{}
     err := readDataFromServer(userinfourl, ur)
     if err == nil {
           ret = true
     } else {
           ret = false
     }

     handler.OnResponse(RESPONSE_TYPE_GET_USER_INFO, wc, ret, ur)
     return nil
}


const (
     RESPONSE_TYPE_SERVER_AUTH  = iota
     RESPONSE_TYPE_USER_AUTH 
     RESPONSE_TYPE_GET_USER_INFO
)

type  WeChatRespHandler interface {
      OnResponse(t int, user * WeChatUser, ret bool, data interface{})
}



type AuthResponse struct {
      Errcode         int     `json:"errcode"`
      Errmsg          string  `json:"errmsg"`
      Access_token    string  `json:"access_token"`
      Expires_in      int     `json:"expires_in"`
      Refresh_token   string
      Openid          string
}

type JSAuthResponse struct {
      Errcode         int     `json:"errcode"`
      Errmsg          string  `json:"errmsg"`
      Ticket          string  `json:"ticket"`
      Expires_in      int     `json:"expires_in"`
}


type UserInfoResp struct {
   
      Errcode         int
      Errmsg          string
      Openid          string
      NickName        string
      City            string
      UnionId         string
      Headimgurl      string
      Sex             int
      Country         string
}




func readDataFromServer(url string, ar interface{}) error {
      
      if len(url) <= 0 {
           return gotom.ErrorMsg(" URL not correct :%s\n", url)
      }

      gotom.LD("[reader] read data from %s\n", url)
      resp, err := http.Get(url);
      if  err != nil {
          return err;
      }

      defer resp.Body.Close() 
      if body, err := ioutil.ReadAll(resp.Body); err == nil {
              gotom.LD("[wechat] <===  %s\n", body)
              if err := json.Unmarshal(body, &ar); err == nil {
                     gotom.LD("[wechat][unmarshal]   %s\n", ar)
                     return nil
              } else {
                     gotom.LE(" json unmarshal failed  %s\n", err) 
                     return err
              }
      } else {
             gotom.LE(" read data from server failed  %s\n", err) 
             return gotom.ErrorMsg(" Read data from server failed ")
      }

}
  
