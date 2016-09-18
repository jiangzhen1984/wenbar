


package handlers

import (
    "gotom"
    "html/template"
    "main/vo"
)


func LoginHandler(resp gotom.GTResponse, req * gotom.GTRequest) {

     t, err := template.ParseFiles("./view/hot_list.html")
     gotom.LE(" parse hot list file err:%s\n", err)
     if err != nil {
            gotom.LE(" parse hot list file err:%s\n", err)
            return
     }

     data := vo.HotListHtml{Title :"sss" , TopicList : []vo.Topic{{Title:"s"},{Title:"s"},{Title:"s"}}}
     t.Execute(*resp.Resp, data)
     gotom.LD("=====output  \n")
}



