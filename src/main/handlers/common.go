

package handlers

import (
     "gotom"
     "net/http"
     "main/service/vo"
)


func UserLoginCheck(req *gotom.GTRequest) bool {

     sess := req.GetSession()
     if sess == nil {
          return false
     }

     user := sess.GetAttribute("user")

     if user == nil {
           return false
     }
  
     return true 
}




func Redirect(resp gotom.GTResponse,  req * gotom.GTRequest, url string) {
         http.Redirect(*resp.Resp, req.Req, url, 302) 
}



func GetLoggedUser(req *gotom.GTRequest) * vo.User {

     if req == nil {
           return nil
     }

     sess := req.GetSession()
     if sess == nil {
           return nil
     }

     gobject := sess.GetAttribute("user")
     user, _ := gobject.(*vo.User)
     return user
}
