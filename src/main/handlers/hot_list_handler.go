

package handlers

import (
    "time"
    "gotom"
    "main/service/vo"
    "main/service"
    "encoding/json"
)




func HotListHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {
     gotom.LF()

     if tpls == nil {
          gotom.LE("No template mapping \n")
          return nil, nil, gotom.ErrorMsg("No template Mapping")
     }

     ti  := gotom.Object(time.Now())
     fs  := gotom.Object(DEFAULT_FETCH_SIZE)
     gdata, err := ws.DoService(ws.GetHotList, &ti, &fs)
     if err != nil {
     }
     topiclist := (*gdata).([]*vo.Topic)
     
     gotom.LD("====>%d  \n", len(topiclist))

    
     if "ajax" == req.P("rfrom") {
          ret, _ := json.Marshal(topiclist) 
          (*resp.Resp).Write(ret)
          return nil, nil, nil
     } else {
          data := new(vo.HotListHtml)   
          data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
          for _, val := range topiclist {
               vth := vo.TopicHtml{}
               vth.PopulateTopic(val)
               data.TopicList = append(data.TopicList, vth)
          }
          return tpls.Tpls["hot_list"], data, nil
     }
}



