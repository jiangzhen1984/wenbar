

package gotom


import (
     "html/template"
     "strconv"
     "fmt"
)

type GTConfig struct {

     DebugMode         bool

     Port              string
     
     Tpldir            string
   
     TplSufix          string

     Mapping           []*Mapping 

     TplMapping        map[string]*GTTemplateMapping

     SessionExpires    int
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
         s := fmt.Sprintf("uri: %s  func: %p", val.Uri, val.Hld)
         buf += "          " + s + "\n"
     }

     buf += "      Templates:\n"
     for _, val := range gt.TplMapping {
         s := fmt.Sprintf("uri: %s  tpls: %s", val.Uri, val.Tpls)
         buf += "          " + s + "\n"
     }
     return buf
}


var GConf * GTConfig
