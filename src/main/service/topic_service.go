

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
     qr := sess.DB("test1").C("topic").Find(bson.M{"date": bson.M{"$lte" : time.Now()}}).Sort("-date").Limit(20).All(&topicList)

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
          ti.Id = vo.Wid(bson.NewObjectId().Hex())
          err := sess.DB("test1").C("topic").Insert(&ti)
          gotom.LD("===create result:%s\n", err)
          gotom.LD("===create result:%s\n", ti)
          return p[0],nil
     }
     return nil, gotom.ErrorMsg(" paramter is no vo.Topic")
}



const (
   QUESTION_QUERY  = iota
   ANSWER_QUEY  
)

func GetPersonalTopicList(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
     var topicList []*vo.Topic
     var err error

     if p == nil || len(p) < 3 {
          return nil, gotom.ErrorMsg("Parameter failed")
     }

     ty, ok := (*p[0]).(int)
     if ok == false {
          gotom.LD("query type error %s\n", p[1])
          return nil, gotom.ErrorMsg("Type not support")
     }

     ti, ok := (*p[1]).(time.Time)
     if ok ==  false {
          ti = time.Now()
     }
  
     tid, ok := (*p[2]).(uint64)
     if ok ==  false {
          gotom.LD("query type error %s\n", p[2])
          return nil, gotom.ErrorMsg("Type not support")
     }
      
     sess := dbs.GetMongoSession()
     switch ty {
          case QUESTION_QUERY:
          gotom.LD("own question query ==>%s ==%d\n", ti, 1)
          err = sess.DB("test1").C("topic").Find(bson.M{"date": bson.M{"$lte" : ti}, "creator.nativeid" : tid}).Sort("-date").Limit(20).All(&topicList)
     }

     gotom.LD("=== topic len :%d   %s\n", len(topicList), err)
     gobj := gotom.Object(topicList)
     return &gobj, nil
}




func GetTopicById(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
     
     var topic vo.Topic

     tid, ok :=  (*p[0]).(string) 
     if ok == false {
     }

  
     sess := dbs.GetMongoSession()
     sess.DB("test1").C("topic").Find(bson.M{"id" : tid}).One(&topic)
     gobject := gotom.Object(topic)     
     return &gobject, nil
}



func  RecordTopicViewUser(dbs * DBSession, p ...*gotom.Object) (*gotom.Object, error) {
     if p == nil || len(p) <= 0{
     }
     vt, ok := (*p[0]).(vo.ViewTopic)
     if ok == false {
     }
   
     vt.Date = time.Now()
     sess := dbs.GetMongoSession()
     sess.DB("test1").C("view_topic").Insert(&vt)
    
     return nil,nil
}
