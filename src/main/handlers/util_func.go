
package handlers

import (
    "time"
    "gotom"
    "main/service/vo"
    "main/service"
    "encoding/json"
    "strconv"
)

func searchTopic(resp gotom.GTResponse, req * gotom.GTRequest) {
     var st string
     var ts int
     var ti time.Time 
     var err error
     st = req.P("text")
     ts, err = strconv.Atoi(req.P("ts"))
     if err != nil && ts > 0 {
         ti = time.Unix(int64(ts), 10) 
     } else {
         ti = time.Now()
     }
     gdata, err := ws.DoService(ws.SearchTopic, st, ti)
     topiclist := gdata.([]*vo.Topic)

     data := new(vo.HotListHtml)   
     data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
     for _, val := range topiclist {
          vth := vo.TopicHtml{}
          vth.PopulateTopic(val)
          if val.AnsList != nil && len(val.AnsList) > 0 {
              vth.AudioUrl = HostConf.AudioHost + val.AnsList[0].AudioPath
          }
          data.TopicList = append(data.TopicList, vth)
     }

     ret, _ := json.Marshal(data)
     (*resp.Resp).Write(ret)
}


