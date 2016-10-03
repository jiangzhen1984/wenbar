

package ws

import (
     "gotom"
     "main/service/vo"
     "time"
     "gopkg.in/mgo.v2/bson"
     "strconv"
)



type TopicList struct {
     TL  []*vo.Topic
}

func GetHotList(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     var topicList []*vo.Topic
     var ptime = time.Now().Unix()
     if p != nil && len(p) > 0 {
          if ti, ok := (p[0]).(string); ok == true {
              gotom.LD("  time %s  \n", ti)
              ptime, _ = strconv.ParseInt(ti, 10, 64)
          } else {
              gotom.LD(" Parse type error use default time %d  \n", ptime)
          }
     }

     gotom.LD(" query time :  %d   %s\n", ptime, time.Unix(ptime, 0))
     
     sess := dbs.GetMongoSession()
     qr := sess.DB("test1").C("topic").Find(bson.M{"timestamp": bson.M{"$lt" : ptime}, "ispublic" : true}).Sort("-count", "-timestamp").Limit(10).All(&topicList)

     gotom.LD("=== topic len :%d   %s\n", len(topicList), qr)
     return topicList, nil
}

func GetNewestList(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     var topicList []*vo.Topic
     var ptime = time.Now().Unix()
     if p != nil && len(p) > 0 {
          if ti, ok := (p[0]).(string); ok == true {
              gotom.LD("  time %s  \n", ti)
              ptime, _ = strconv.ParseInt(ti, 10, 64)
          } else {
              gotom.LD(" Parse type error use default time %d  \n", ptime)
          }
     }

     sess := dbs.GetMongoSession()
     qr := sess.DB("test1").C("topic").Find(bson.M{"timestamp": bson.M{"$lt" : ptime}, "ispublic" : true}).Sort("-timestamp").Limit(10).All(&topicList)

     gotom.LD("=== topic len :%d   %s\n", len(topicList), qr)
     return topicList, nil
}


func SearchTopic(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     var topicList []*vo.Topic

     return topicList, nil
}



func CreateTopic(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     sess := dbs.GetMongoSession()      
     
     ti, ok := (p[0]).(vo.Topic)
     if ok == true {
          ti.Date = time.Now()
          ti.Id = vo.Wid(bson.NewObjectId().Hex())
          ti.TimeStamp = ti.Date.Unix()
          err := sess.DB("test1").C("topic").Insert(&ti)
          gotom.LD("===create result:%s\n", err)
          gotom.LD("===create result:%s\n", ti)
          return p[0],nil
     }
     return nil, gotom.ErrorMsg(" paramter is no vo.Topic")
}



const (
   QUESTION_QUERY  = iota
   ANSWER_QUERY  
   VIEWED_QUERY
)

func GetPersonalTopicList(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     var topicList []*vo.Topic
     var err error
     var result []struct{ TopicId string `bson:"topicid"` }

     if p == nil || len(p) < 3 {
          return nil, gotom.ErrorMsg("Parameter failed")
     }

     ty, ok := (p[0]).(int)
     if ok == false {
          gotom.LD("query type error %s\n", p[1])
          return nil, gotom.ErrorMsg("Type not support")
     }

     ti, ok := (p[1]).(time.Time)
     if ok ==  false {
          ti = time.Now()
          gotom.LD("use default time to query personal topic %s\n", ti)
     }
     gotom.LD("use time to query personal topic %s\n", ti)
  
     tid, ok := (p[2]).(vo.Wid)
     if ok ==  false {
          gotom.LD("query type error %s\n", p[2])
          return nil, gotom.ErrorMsg("Type not support")
     }
      
     sess := dbs.GetMongoSession()
     switch ty {
          case QUESTION_QUERY:
              gotom.LD("own question query ==>%s\n", ti)
              err = sess.DB("test1").C("topic").Find(bson.M{"date": bson.M{"$lte" : ti}, "creator.uid" : tid}).Sort("-date").Limit(10).All(&topicList)
          case ANSWER_QUERY:
              gotom.LD("Ask to me query ==>%s\n", ti)
              err = sess.DB("test1").C("topic").Find(bson.M{"date": bson.M{"$lte" : ti}, "askto" : tid}).Sort("-date").Limit(10).All(&topicList)
          case VIEWED_QUERY:
              gotom.LD("my viewed query ==>%s\n", tid)
              query := bson.M{"date": bson.M{"$lte" : ti}, "viewuserid" : tid}
              project := bson.M{"topicid" : 1}
              err := sess.DB("test1").C("view_topic").Find(query).Sort("-date").Limit(10).Select(project).All(&result)
              idlist := []string{}
              if err != nil {
              }
              for _, val := range result {
                   idlist = append(idlist, val.TopicId)
              }

              err = sess.DB("test1").C("topic").Find(bson.M{"_id" : bson.M{"$in" : idlist}}).All(&topicList)
              
     }

     gotom.LD("=== topic len :%d   %s\n", len(topicList), err)
     return topicList, nil
}




func GetTopicById(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     
     var topic vo.Topic

     tid, ok :=  (p[0]).(string) 
     if ok == false {
          return nil, gotom.ErrorMsg("NO such ID")
     }
  
     sess := dbs.GetMongoSession()
     err := sess.DB("test1").C("topic").FindId(tid).One(&topic)
     gotom.LD(" query topic by Id %s  %s\n", tid, err) 
     if err == nil {
         return &topic, nil
     } else {
         return nil, gotom.ErrorMsg(" Query failed  !")
     }
}



func UpdateTopicViewCount(dbs *DBSession, p...gotom.Object) (gotom.Object, error) {

     var tid string
     tid, ok :=  (p[0]).(string) 
     if ok == false {
          return nil, gotom.ErrorMsg("NO such ID")
     }

     gotom.LD(" Update topic[%s] count\n", tid)

     sess := dbs.GetMongoSession()
     err := sess.DB("test1").C("topic").UpdateId(tid, bson.M{"$inc" : bson.M{"count" : 1}})
     if err != nil {
           gotom.LE(" Update view count error  %s", err)
     }
     return nil, nil
}



func  RecordTopicViewUser(dbs * DBSession, p ...gotom.Object) (gotom.Object, error) {
     if p == nil || len(p) <= 0{
     }
     vt, ok := (p[0]).(*vo.ViewTopic)
     if ok == false {
           gotom.LP("==== convert failed %s  \n", p[0])
     }
   
     vt.Date       = time.Now()
     vt.TopicId    = vt.Topic.Id
     vt.ViewUserId = vt.ViewUser.Uid
     sess := dbs.GetMongoSession()
     sess.DB("test1").C("view_topic").Insert(vt)
    
     return nil,nil
}
