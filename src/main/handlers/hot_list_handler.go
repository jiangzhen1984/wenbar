

package handlers

import (
    "gotom"
    "main/vo"
)


func HotListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LD("===hotlist   tpls %s\n", tpls)

     if tpls == nil {
          return nil, nil, nil
     }

   
     data := vo.HotListHtml{Title :"sss" , TopicList : []vo.Topic{{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"}}}
     return &gotom.GTTemplate{tpls.Tpls["hot_list"]}, data, nil
}



