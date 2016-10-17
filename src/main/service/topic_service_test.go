
package ws

import (
     "gotom"
     "main/service/vo"
     "time"
     //"strconv"
     "testing"
     "fmt"
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




func TestSearchTopic(t *testing.T) {
     InitDB(DBConfiguration{DBUrl : "localhost"})
     sess := CreateDBSession() 
     if sess == nil {
          t.Fatal(" create db session failed \n")
     }

     var err error
     _, err = SearchTopic(sess)
     if err.Error() !=  "Parameter not enough" {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: Parameter not enough\n")
     }

     _, err = SearchTopic(sess, nil)
     if err.Error() !=  "Parameter type mismatch  string" {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: Parameter type mismatch string")
     }

     _, err = SearchTopic(sess, "ss", nil)
     if err.Error() !=  "Parameter type mismatch  time.Time" {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: Parameter type mismatch  time.Time")
     }


     var ti time.Time = time.Now()
     _, err = SearchTopic(sess, ti, "ss", nil)
     if err.Error() !=  "Parameter type mismatch  time.Time" {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: Parameter type mismatch  time.Time")
     }


     for i:= 0; i < 20; i++ {
         tic := new(vo.Topic)
         tic.Title = fmt.Sprint(" title ",i)
         tic.Content = fmt.Sprint("content test ",i)
         CreateTopic(sess, tic)
     } 


     var gobj gotom.Object
     gobj, err = SearchTopic(sess, "content", time.Now()) 
     if err != nil {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: search failed")
     }

     tl := gobj.([]*vo.Topic)
     if len(tl) != 10 {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: result not match")
     }

     gobj, err = SearchTopic(sess, "test", time.Now()) 
     if err != nil {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: search failed")
     }

     tl = gobj.([]*vo.Topic)
     if len(tl) == 0 {
          gotom.LE("msg :%s", err)
          t.Fatal(" assert failed: result not match")
     }

}
