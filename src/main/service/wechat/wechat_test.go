

package wechat

import (
     "testing"
     "net/http"
     "gotom"
     "net/url"
     "io/ioutil"
     "encoding/json"
)




func TestInitWeChatConfig(t * testing.T) {
     if confs := InitWeChatConfig(""); confs != nil {
          t.Fatal(" should not initalize ")
     }

     if confs := InitWeChatConfig("test_config.json"); confs != nil && len(confs) > 0 {
          t.Fatal(" should not json unmarshal ")
     }

     if confs := InitWeChatConfig("test_config1.json"); confs != nil {
          if confs[0].AppId != "1234" && confs[0].Secret != "5678" {
              t.Fatal(" not match %s ", confs[0])
          }
     } else {
          t.Fatal(" parse failed ")
     }
     
}


var server * http.Server = &http.Server{
    Addr    :  ":8080",
}



func (wcc * WeChatConfig) ServeHTTP(resp http.ResponseWriter, r * http.Request) {
     gotom.LD("========request :%s\n", r.URL.Path[:])
     uri := r.URL.Path[:]
     if uri == "/cgi-bin/token" { if r.FormValue("appid") == "1234" {
               resp.Write([]byte("{\"access_token\":\"test\",\"expires_in\":7200, \"openid\": \"abcd\", \"scope\":\"web\"}"))
           } else {
               resp.Write([]byte("{\"errocode\": 1234,\"errmsg\" : \"11111\"}"))
           }
     } else if uri == "/connect/oauth2/authorize" {
           http.Redirect(resp, r, r.FormValue("redirect_uri") +"?code=1", 302)
     } else if uri == "/redirect" {
           gotom.LD("====> get redirect %s\n", r.FormValue("code"))
     } else if uri == "/sns/userinfo" {
 
           st := UserInfoResp {
                   Openid  : "abcd",
                   NickName     : "1235", 
                   City         :" beijing", 
                   UnionId      :"aaaaa", 
                   Headimgurl   :"http://localhost:8081/test",
                   } 
           content, _ := json.Marshal(st)
           resp.Write(content)

     }
}

func morkAuthServer(wcc * WeChatConfig) {
    server.Handler =  wcc
    gotom.LD("=== try to start server\n")
    server.ListenAndServe()
    
}

func TestAuthServer(t * testing.T) {
    weChatURL = "http://localhost:8080"
    requestUri = []string{"/cgi-bin/token", "/sns/oauth2/access_token", "/sns/userinfo"} 
    WeChatUserAuthURL = "http://localhost:8080/connect/oauth2/authorize?"

   go  morkAuthServer(&WeChatConfig{})
   if confs := initWeChatConfig("ctest_config1.json"); confs != nil {
       
        for _, c := range confs {
              if c.AuthServer() == false {
                    t.Fatal(" auth server failed ")
              } else if c.Token != "test" {
                    t.Fatal(" token parse failed ")
              }
        }
   } else {
          t.Fatal(" Init wechatconfig failed ")
   } 


    weChatURL = ""
    requestUri = []string{"", "", ""} 
    WeChatUserAuthURL = ""

   if confs := initWeChatConfig("ctest_config1.json"); confs != nil {
       
        for _, c := range confs {
              if c.AuthServer() == true {
                    t.Fatal(" test failed due to auth server return true")
              }
        }
   } else {
          t.Fatal(" Init wechatconfig failed ")
   } 


}




func TestBuildUserAuthUrl(t * testing.T) {
   weChatURL = "http://localhost:8080"
   requestUri = []string{"/cgi-bin/token", "/sns/oauth2/access_token", "/sns/userinfo"} 
   WeChatUserAuthURL = "http://localhost:8080/connect/oauth2/authorize?"

   go  morkAuthServer(&WeChatConfig{})

   if confs := initWeChatConfig("ctest_config1.json"); confs != nil && len(confs) > 0{
         u := confs[0].InitWeChatUser()
         if u == nil {
                t.Fatal(" Init user failed ")
         }
         reurl, _ := url.Parse("http://localhost:8080/redirect")
         u.BuildAuthUrl(url.QueryEscape(reurl.String()))
         gotom.LD(" build url %s\n", u.WebAuthUrl)
         resp, err := http.Get(u.WebAuthUrl)
         if err != nil {
             t.Fatal(" read error %s\n", err)
         }
         ioutil.ReadAll(resp.Body)
         defer resp.Body.Close()
   } else {
         t.Fatal(" Init wechatconfig failed ")
   } 
}



func TestUpdateAuthCode(t * testing.T) {
   if confs := initWeChatConfig("ctest_config1.json"); confs != nil && len(confs) > 0{
         u := confs[0].InitWeChatUser()        
         if u == nil {
                t.Fatal(" Init user failed ")
         }
         u.UpdateAuthCode("abc")
         if u.Code != "abc" || u.State != WC_USER_CODE {
             t.Fatal(" update code failed ")
         }
   } else {
         t.Fatal(" Init wechatconfig failed ")
   } 
}



func (u * WeChatUser)  OnResponse(t int, user * WeChatUser, ret bool, data interface{}) {
  switch data.(type) {
        case *AuthResponse:
            ar := data.(*AuthResponse) 
            if ar.Openid != "abcd" {
                     gotom.LE("===== failed \n")
            } else {
                     gotom.LI("===== success \n")
            }
       case *UserInfoResp:
            ur := data.(*UserInfoResp) 
             gotom.LD("===%s\n", ur)
  }
    gotom.LD("==== get user token response :%b   %s\n", ret,   data)
}

func TestUserAuthToken(t * testing.T) {
    weChatURL = "http://localhost:8080"
    requestUri = []string{"/cgi-bin/token", "/sns/oauth2/access_token", "/sns/userinfo"} 
    WeChatUserAuthURL = "http://localhost:8080/connect/oauth2/authorize?"

   go  morkAuthServer(&WeChatConfig{})
   if confs := initWeChatConfig("ctest_config1.json"); confs != nil {
         u := confs[0].InitWeChatUser()        
         if u == nil {
                t.Fatal(" Init user failed ")
         }
         u.UpdateAuthCode("w")
         u.AuthToken(u)
   } else {
          t.Fatal(" Init wechatconfig failed ")
   } 
}


func TestGetUserInfo(t * testing.T) {
    weChatURL = "http://localhost:8080"
    requestUri = []string{"/cgi-bin/token", "/sns/oauth2/access_token", "/sns/userinfo"} 
    WeChatUserAuthURL = "http://localhost:8080/connect/oauth2/authorize?"

   go  morkAuthServer(&WeChatConfig{})
   if confs := initWeChatConfig("ctest_config1.json"); confs != nil {
         u := confs[0].InitWeChatUser()        
         if u == nil {
                t.Fatal(" Init user failed ")
         }
         err :=  u.GetUserInfoFromServer(u)
         if err == nil {
               t.Fatal(" should not pass due to incorrect user state %s ", err)
         }
         u.UpdateAuthCode("ssdf")
         u.AuthToken(u)

        
         err =  u.GetUserInfoFromServer(u)
         if err != nil {
               t.Fatal(" get failed %s ", err)
         }
   } else {
          t.Fatal(" Init wechatconfig failed ")
   }
}


func initWeChatConfig(p string) [] * WeChatConfig {
     if confs := InitWeChatConfig("test_config1.json"); confs != nil  && len(confs) > 0{
           return confs
     } else {
           return nil
     }
}
