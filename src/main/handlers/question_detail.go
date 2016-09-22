



package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
)


func QuestionDetailHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         qid := req.Req.FormValue("qid") 
         if qid == "" || len(qid) <= 0 {
              //TODO return error
         }
         gobj := gotom.Object(qid)
         if gobject, err := ws.DoService(ws.GetTopicById, &gobj); err == nil {
             topic, _:= (*gobject).(vo.Topic)
             if user := GetLoggedUser(req); user != nil {
                  vt := new(vo.ViewTopic)
                  vt.ViewUser = user
                  vt.Topic = &topic
                  gobj = gotom.Object(vt)
                  ws.DoService(ws.RecordTopicViewUser, &gobj)
             }
             topicHtml := vo.TopicHtml{Title : topic.Title, Content : topic.Content, CreatorName : topic.Creator.Name, CreatorTitle : topic.Creator.Title}
             return tpls.Tpls["questiondetail"], &topicHtml, nil
         }
         gotom.LD("========not found question\n")
         return tpls.Tpls["questiondetail"], nil, nil
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}


