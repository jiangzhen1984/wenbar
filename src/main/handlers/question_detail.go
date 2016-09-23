



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
              //TODO redirect to error page
              return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
         }
         gobj := gotom.Object(qid)
         if gobject, err := ws.DoService(ws.GetTopicById, &gobj); err == nil {
             topic, ok:= (*gobject).(*vo.Topic)
             //TODO update topic view count
             
             ws.DoService(ws.UpdateTopicViewCount, &gobj)
             if user := GetLoggedUser(req); user != nil {
                  vt := new(vo.ViewTopic)
                  vt.ViewUser = user
                  vt.Topic = topic
                  gobj = gotom.Object(vt)
                  ws.DoService(ws.RecordTopicViewUser, &gobj)
             }
             topicHtml := new(vo.TopicHtml)
             if topic != nil {
                  topicHtml.PopulateTopic(topic) 
             } else {
                  gotom.LE("===query topic failed %s   %s\n", topic, ok)
             }
             return tpls.Tpls["questiondetail"], &topicHtml, nil
         }
         gotom.LD("========not found question\n")
         return tpls.Tpls["questiondetail"], nil, nil
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}


