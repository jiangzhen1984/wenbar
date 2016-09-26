


package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
    "time"
    "strconv"
    "encoding/json"
)


func MyViewedHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, gotom.ErrorMsg("No template Mapping")
     }

     user := GetLoggedUser(req)
     if user == nil {
          Redirect(resp, req, "/login?from=/my_viewed")
          return nil, nil, nil 
     }

     ts := req.P("ts")
     var timestamp time.Time

     if len(ts) <= 0 {
           timestamp = time.Now()
     } else {
           tsint, _ := strconv.ParseInt(ts, 10, 64)
           timestamp = time.Unix(tsint, 0)
     }

     gotom.LD(" My Inquriy timestamp :%d  %s\n", ts, timestamp)
     gotype := gotom.Object(ws.VIEWED_QUERY)
     gotime := gotom.Object(timestamp)
     gonativeId := gotom.Object(user.Uid)
     gdata, err := ws.DoService(ws.GetPersonalTopicList, &gotype, &gotime, &gonativeId)

     topiclist := (*gdata).([]*vo.Topic)
     gotom.LD("====>%d  \n", len(topiclist))

     if req.P("rfrom") == "ajax" {
          ret, _ := json.Marshal(topiclist) 
          (*resp.Resp).Write(ret)
          return nil, nil, nil
     } else {
          if err != nil {
               ///TODO  check error for query
               gotom.LE("===>%s\n", err)
          }
     
          data := new(vo.HotListHtml)   
          data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
          for _, val := range topiclist {
               vt := vo.TopicHtml{}
               vt.PopulateTopic(val)
               data.TopicList = append(data.TopicList, vt)
          }
     
          return tpls.Tpls["myviewed"], data, nil
     }
}



