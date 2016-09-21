

package ws

import (
     "gotom"
     "main/service/vo"
     "time"
)



type TopicList struct {
     TL  []*vo.Topic
}

func GetHotList(dbs * DBSession, p ...*gotom.Object) *gotom.Object {
     var topicList []*vo.Topic
     var ptime time.Time
  
     if p == nil || len(p) == 0 {
         ptime = time.Now()
     } else {
         ti, ok := (*p[0]).(time.Time)
         if ok ==  false {
              ti = time.Now()
         }
         ptime = ti
     }
     gotom.LD("===%s\n", ptime)
     
     topicList = make([]*vo.Topic, 0, 10)
     for i :=0; i < 20; i++ {
          topicList = append(topicList, &vo.Topic{Id :"1", Title : "sss"})
     }
     
     gobj := gotom.Object(topicList)
     return &gobj
}


func SearchTopic(dbs * DBSession, p ...*gotom.Object) *gotom.Object{
     var topicList []*vo.Topic

     gobj := gotom.Object(topicList)
     return &gobj
}
