

package handlers

import (
    "time"
    "gotom"
    "main/service/vo"
    "main/service"
    "encoding/json"
    "strconv"
)



func HotListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LD("==== Cookies %s \n", req.Req.Cookies())
     var timestamp string

     if "ts" == req.P("type") {
          searchTopic(resp, req)
          return nil,nil,nil
     }

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, gotom.ErrorMsg("No template Mapping")
     }

     timestamp  = req.P("ts")
     if len(timestamp) <= 0 {
          timestamp = strconv.FormatInt(time.Now().Unix(), 10)
     }
     gotom.LI(" query time >>> %s\n", timestamp)
     gdata, err := ws.DoService(ws.GetHotList, timestamp, DEFAULT_FETCH_SIZE)
     if err != nil {
     }
     topiclist := gdata.([]*vo.Topic)
     
     gotom.LD("====>%d  \n", len(topiclist))

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
    
     if "ajax" == req.P("rfrom") {
          ret, _ := json.Marshal(data)
          (*resp.Resp).Write(ret)
          return nil, nil, nil
     } else {
          return tpls.Tpls["hot_list"], data, nil
     }
}



func searchTopic(resp gotom.GTResponse, req * gotom.GTRequest) {
     var st string
     var ts int
     var ti time.Time 
     var err error
     st = req.P("ts")
     ts, err = strconv.Atoi(req.P("ts"))
     if err != nil {
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


