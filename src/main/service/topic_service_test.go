
package ws

import (
     "gotom"
     "main/service/vo"
     //"time"
     //"gopkg.in/mgo.v2/bson"
     //"strconv"
     "testing"
)



func TestAddTopicAnswer(t * testing.T) {
     InitDB(DBConfiguration{DBUrl : "localhost"})
     sess := CreateDBSession() 
     if sess == nil {
          t.Fatal(" create db session failed \n")
     }
     topic := &vo.Topic{}
     _, err := AddTopicAnswer(sess, topic)
     if err == nil || err.Error() != "Parameter not enough" {
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed for parameter not enough \n")
     }
     ans := &vo.Answer{}
     _, err = AddTopicAnswer(sess, ans)
     if err == nil || err.Error() != "Parameter not enough" {
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed for parameter not enough \n")
     }

     
     _, err = AddTopicAnswer(sess, topic, *ans)
     if err == nil || err.Error() != "Parameter type not *vo.Topic and *vo.Answer" {
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed for type mismatch\n")
     }

     _, err = AddTopicAnswer(sess, topic, ans)
     if err == nil || err.Error() != "No user" {
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed for user mismatch\n")
     }

     ans.AnsUser = &vo.User{Uid : "ssssss"}
     _, err = AddTopicAnswer(sess, topic, ans)
     if err == nil || err.Error() != "not found"{
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed for mismatch\n")
     }

     _, err = CreateTopic(sess, topic)
     if err != nil {
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Prepare test data failed\n")
     }
     ans.Content =" test content"
     ans.AudioPath =" path"
     _, err = AddTopicAnswer(sess, topic, ans)
     if err != nil{
          gotom.LE("=== > msg %s\n", err)
          t.Fatal(" Assert failed \n")
     }

     if ans.Id == "" || len(ans.Id) <= 0 {
     }
     
}
