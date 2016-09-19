

package gotom


import (
     "html/template"
)

type GTConfig struct {

     DebugMode   bool

     Port        string
     
     Tpldir      string
   
     TplSufix    string

     Mapping     []*Mapping 

     TplMapping  map[string]*GTTemplateMapping
}



type GTTemplateMapping struct {

     Uri         string

     Tpls        map[string]*GTTemplate
}



type GTTemplate struct {
     
     Name       string
 
     Path       string

     NativeTpl  * template.Template
}


var GConf * GTConfig
