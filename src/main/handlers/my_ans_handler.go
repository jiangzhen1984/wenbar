


package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
    "time"
)


func MyAnsHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, gotom.ErrorMsg("No template Mapping")
     }

     user := GetLoggedUser(req)
     if user == nil {
          Redirect(resp, req, "/login")
          return nil, nil, nil 
     }
     gotype := gotom.Object(ws.ANSWER_QUERY)
     gotime := gotom.Object(time.Now())
     gonativeId := gotom.Object(user.Uid)
     gdata, err := ws.DoService(ws.GetPersonalTopicList, &gotype, &gotime, &gonativeId)

     if err != nil {
           ///TODO  check error for query
           gotom.LE("===>%s\n", err)
     }
     topiclist := (*gdata).([]*vo.Topic)
     
     gotom.LD("====>%d  \n", len(topiclist))
     data := new(vo.HotListHtml)   
     data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
     for _, val := range topiclist {
          vt := vo.TopicHtml{}
          vt.PopulateTopic(val)
          data.TopicList = append(data.TopicList, vt)
     }
     
     return tpls.Tpls["myanswer"], data, nil
}



