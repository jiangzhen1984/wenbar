
package tpl

import (
   "html/template"
)


type Test1 struct {

    Title string
    Name  string
    Val int
}


var Tpls  * template.Template = template.Must(template.ParseFiles("view/login.html"))
