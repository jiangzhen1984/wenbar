

package handlers

import (
    "time"
    "gotom"
    "main/service/vo"
    "main/service"
)


const DEFAULT_FETCH_SIZE = 20


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
     data := new(vo.HotListHtml)   
     data.TopicList = make([]vo.TopicHtml, 0, len(topiclist))
     for _, val := range topiclist {
          gotom.LD("======>%s\n", val.Id)
          data.TopicList = append(data.TopicList, vo.TopicHtml{Tid : val.Id, Title : val.Title})
     }
    
     return tpls.Tpls["hot_list"], data, nil
}



