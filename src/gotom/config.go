

package gotom


import (
     "html/template"
)

type GTConfig struct {

     Port        string
     
     Tpldir      string
   
     TplSufix    string

     Mapping     []*Mapping 

     TplMapping  map[string]*GTTemplateMapping
}



type GTTemplateMapping struct {

     Uri         string

     Tpls        map[string]*template.Template
}



type GTTemplate struct {

     NativeTpl  * template.Template
}


type Object interface {}

var GConf * GTConfig
