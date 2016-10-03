

package handlers

import (
    "time"
    "gotom"
    "main/service/vo"
    "main/service"
    "encoding/json"
    "strconv"
)


func NewestListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()
     var timestamp string

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, gotom.ErrorMsg("No template Mapping")
     }

     timestamp  = req.P("ts")
     if len(timestamp) <= 0 {
          timestamp = strconv.FormatInt(time.Now().Unix(), 10)
     }
     gotom.LI(" query time >>> %s\n", timestamp)
     gdata, err := ws.DoService(ws.GetNewestList, timestamp, DEFAULT_FETCH_SIZE)
     if err != nil {
     }
     topiclist := gdata.([]*vo.Topic)
     
     gotom.LD("====>%d  \n", len(topiclist))
    
     if "ajax" == req.P("rfrom") {
          ret, _ := json.Marshal(topiclist) 
          (*resp.Resp).Write(ret)
          return nil, nil, nil
     } else {
          data := new(vo.HotListHtml)   
          data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
          for _, val := range topiclist {
               vth := &vo.TopicHtml{}
               vth.PopulateTopic(val)
               gotom.LD("=======>%s\n", val)
               data.TopicList = append(data.TopicList, *vth)
          }
    
          return tpls.Tpls["newest_list"], data, nil
     }
}



