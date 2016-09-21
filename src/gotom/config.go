

package gotom


import (
     "html/template"
     "strconv"
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


func (gt GTConfig) String() string {
     var buf string
     buf += "[GTConfig]\n"
     buf += "      DebugMode: " + strconv.FormatBool(gt.DebugMode) +"\n"
     buf += "      Port: " + gt.Port+"\n"
     buf += "      Mappings:\n"
     for _, val := range gt.Mapping {
         buf += "          uri:" + val.Uri + "   func:" + val.String() +"\n"
     }
     return buf
}


var GConf * GTConfig
