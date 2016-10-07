
package sim

import "gotom"
import "net/http"
import "encoding/json"
import "main/service/wechat"



var server * http.Server = &http.Server{
    Addr    :  ":8082",
}

type MockWechatServer struct {
}

func (wcc * MockWechatServer) ServeHTTP(resp http.ResponseWriter, r * http.Request) {
     gotom.LD("========request :%s\n", r.URL.Path[:])
     uri := r.URL.Path[:]
     if uri == "/cgi-bin/token" || uri == "/sns/oauth2/access_token"{
           resp.Write([]byte("{\"access_token\":\"test\",\"expires_in\":7200, \"openid\": \"abcd\", \"scope\":\"web\"}"))
     } else if uri == "/connect/oauth2/authorize" {
           http.Redirect(resp, r, r.FormValue("redirect_uri") +"?code=1", 302)
     } else if uri == "/redirect" {
           gotom.LD("====> get redirect %s\n", r.FormValue("code"))
     } else if uri == "/sns/userinfo" {
 
           st := wechat.UserInfoResp {
                   Openid       : "abcd",
                   NickName     : "1235", 
                   City         :" beijing", 
                   UnionId      :"aaaaa", 
                   Headimgurl   :"http://localhost:8081/images/p8.jpg",
                   } 
           content, _ := json.Marshal(st)
           resp.Write(content)
     }
}


func InitSIMWeChatServer() {
    wechat.WeChatURL = "http://localhost:8082"
    wechat.WeChatUserAuthURL = "http://localhost:8082/connect/oauth2/authorize?"
    server.Handler =  &MockWechatServer{}
    gotom.LD("=== try to start server\n")
    server.ListenAndServe()
}
