

package ws

import (
     "gotom"
     "main/service/vo"
     "time"
     "gopkg.in/mgo.v2/bson"
)



type TopicList struct {
     TL  []*vo.Topic
}

func GetHotList(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
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
     
     sess := dbs.GetMongoSession()
     qr := sess.DB("test1").C("topic").Find(bson.M{"date": bson.M{"$lte" : time.Now()}}).Limit(20).All(&topicList)

     gotom.LD("=== topic len :%d   %s\n", len(topicList), qr)
     gobj := gotom.Object(topicList)
     return &gobj, nil
}


func SearchTopic(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
     var topicList []*vo.Topic

     gobj := gotom.Object(topicList)
     return &gobj, nil
}



func CreateTopic(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
     sess := dbs.GetMongoSession()      
     
     ti, ok := (*p[0]).(vo.Topic)
     if ok == true {
          ti.Date = time.Now()
          err := sess.DB("test1").C("topic").Insert(&ti)
          gotom.LD("===create result:%s\n", err)
          gotom.LD("===create result:%s\n", ti)
          return p[0],nil
     }
     return nil, gotom.ErrorMsg(" paramter is no vo.Topic")
}
