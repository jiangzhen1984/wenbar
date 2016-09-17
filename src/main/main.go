

package main

import (
    "fmt"
    "net/http"
    "regexp"
    "gotom"
)


type Session struct {
    id string
}


func handler(w http.ResponseWriter, r * http.Request) {
   
   coi := http.Cookie{Name:"sss1", Value:"eee1"}
   r.ParseForm()
   fmt.Printf("%s\n",r.Form["s1"])
   fmt.Printf("form value==>   s=%s\n", r.Form)
   fmt.Printf("form value==>   s=%s\n", r.FormValue("s"))
   fmt.Printf("%s", http.ServerContextKey)
   http.SetCookie(w, &coi)
}




func static_handler(rw http.ResponseWriter, r * http.Request) {
  rex := regexp.MustCompilePOSIX("(.jpg|.png|.css)$")
  fmt.Printf("====%s\n", r.URL.Path)
  if rex.MatchString(r.URL.Path) {
    http.ServeFile(rw, r, r.URL.Path)
    return
  } else {
    handler(rw, r) 
  }
}



func main() {
    gotom.InitServer(nil)
}
