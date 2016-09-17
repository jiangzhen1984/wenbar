

package main

import (
    "fmt"
    "net/http"
    "tpl"
    "regexp"
)


type Session struct {
    id string
}


func handler(w http.ResponseWriter, r * http.Request) {
   t1 := tpl.Test1{"aa", "000000000000000", 123}
   
   coi := http.Cookie{Name:"sss1", Value:"eee1"}
   r.ParseForm()
   fmt.Printf("%s\n",r.Form["s1"])
   fmt.Printf("form value==>   s=%s\n", r.Form)
   fmt.Printf("form value==>   s=%s\n", r.FormValue("s"))
   fmt.Printf("%s", http.ServerContextKey)
   http.SetCookie(w, &coi)
   tpl.Tpls.ExecuteTemplate(w, "login.html", t1)
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
   http.Handle("/", http.FileServer(http.Dir("./view/")))
   http.HandleFunc("/test", handler)
   http.ListenAndServe(":8080", nil)
}
