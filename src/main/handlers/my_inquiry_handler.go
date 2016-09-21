


package handlers

import (
    "gotom"
    "main/service/vo"
    "fmt"
)


func MyInquiryHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, fmt.Errorf("No template Mapping")
     }

   
     data := vo.MyInquiryHtml{TopicList : []vo.TopicHtml{{Title:"s", AnsUserName : " aa", AnsUserTitle :" 院北京基因组研究所,遗传学博士 asdfsdfsp偷偷看" },{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"},{Title:"s"}}}
     return tpls.Tpls["myinquiry"], data, nil
}



