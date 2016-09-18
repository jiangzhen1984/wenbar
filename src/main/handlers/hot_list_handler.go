

package handlers

import (
    "gotom"
    "gotom/tpl"
    "html/template"
    "main/vo"
)


func HotListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * tpl.GTTemplateMapping)  *gotom.GTTemplate, gotom.data, error {

     if tpls == nil {
          return nil, nil, nil
     }
     data := vo.HotListHtml{Title :"sss" , TopicList : []vo.Topic{{Title:"s"},{Title:"s"},{Title:"s"}}}

     return nil, data, nil
}



