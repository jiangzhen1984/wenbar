



package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
)


func InquiryHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if UserLoginCheck(req) != true {
          Redirect(resp, req, "/login?from=/inquiry") 
          return nil, nil, nil
     }

     if req.Req.Method == "GET" {
          return tpls.Tpls["inquiry"], nil, nil
     } else if req.Req.Method == "POST" {
          con := req.Req.FormValue("content")
          if len(con) <= 0 {
              return tpls.Tpls["inquiry"], vo.TopicHtml{ErrMsg : "请输入问题内容"}, nil
          }
              
          gtopic  := gotom.Object(vo.Topic{Title:"sss", Content:con, Creator : GetLoggedUser(req)})
          _, err := ws.DoService(ws.CreateTopic, &gtopic) 
          if err != nil {
                return tpls.Tpls["inquiry"], vo.TopicHtml{ErrMsg : "提交问题失败"}, nil
          }
          Redirect(resp, req, "/my_inquiry") 

     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}



