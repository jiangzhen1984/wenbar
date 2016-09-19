

package handlers

import (
    "gotom"
    "main/vo"
    "fmt"
)


func HotListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, fmt.Errorf("No template Mapping")
     }

   
     data := vo.HotListHtml{Title :"sss" , TopicList : []vo.TopicHtml{{Title:"s", AnsUserName : " aa", AnsUserTitle :" sss" },{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"}}}
     return tpls.Tpls["hot_list"], data, nil
}



