

package ws

import (
    "testing"
    "main/service/vo"
    "gotom"
    "time"
)


func TestCreateUser(t * testing.T) {

     InitDB(DBConfiguration{DBUrl : "localhost"})
     sess := CreateDBSession() 
     if sess == nil {
          t.Fatal(" create db session failed \n")
     }

     u := &vo.User{Name : "aaa" , Personal : &vo.UserPersonal{Incoming : 12.5}, WeChat : &vo.PersonalWeChat{OpenId: "1111"}}
     _, err := CreateUser(sess, u)
     if err != nil {
           t.Fatal(" create user failed  %s", err)
     }

     if len(u.Uid) <=0 {
            t.Fatal("===   uid not return % s\n ", u)
     }

     sess.Close()
}



func TestUpdateUser(t * testing.T) {
     InitDB(DBConfiguration{DBUrl : "localhost"})
     sess := CreateDBSession() 
     if sess == nil {
          t.Fatal(" create db session failed \n")
     }

     u := &vo.User{Name : "aaa" , Personal : &vo.UserPersonal{Incoming : 12.5}, WeChat : &vo.PersonalWeChat{OpenId: "1111"}}
     _, err := CreateUser(sess, u)
     if err != nil {
           t.Fatal(" create user failed  %s", err)
     }

     if len(u.Uid) <=0 {
            t.Fatal("===   uid not return % s\n ", u)
     }


     gotom.LD("===>%s\n", u.Uid)
     gret, err := GetUserById(sess, vo.Wid(u.Uid))
     if err != nil {
          t.Fatal("=== find prepared data failed %s\n", err)
     }

     u = (gret).(*vo.User) 
     gotom.LD("==>%s \n", u)
     u.Name = "bbbb"
     u.WeChat.Unionid ="eeeeeeaaaaa"
     u.Personal.UnAns = 100
     u.Personal.QuesFee = 10.3
     u.Personal.BeViewed = 200
     u.Personal.Ans = 20
     u.Personal.Revenue = 20.5
     _, err = UpdateUserPersonal(sess, u)
     if err != nil {
            t.Fatal(" failed ====> %s\n", err)
     }

     gret, err = GetUserById(sess, vo.Wid(u.Uid))
     if err != nil {
          t.Fatal("=== assert failed not found %s\n", err)
     }

     u = (gret).(*vo.User) 

     if u.Personal.Ans != 20 || u.Personal.UnAns != 100 || u.Personal.BeViewed != 200 {
          t.Fatal("=== assert failed not match %s\n", u)
     }
     

     sess.Close()
}



func TestUpdateUserWeChat(t * testing.T) {
     InitDB(DBConfiguration{DBUrl : "localhost"})
     sess := CreateDBSession() 
     if sess == nil {
          t.Fatal(" create db session failed \n")
     }

     u := &vo.User{Name : "aaa" , Personal : &vo.UserPersonal{Incoming : 12.5}, WeChat : &vo.PersonalWeChat{OpenId: "1111"}}
     _, err := CreateUser(sess, u)
     if err != nil {
           t.Fatal(" create user failed  %s", err)
     }

     if len(u.Uid) <=0 {
            t.Fatal("===   uid not return % s\n ", u)
     }


     gotom.LD("===>%s\n", u.Uid)
     gret, err := GetUserById(sess, vo.Wid(u.Uid))
     if err != nil {
          t.Fatal("=== find prepared data failed %s\n", err)
     }

     u = (gret).(*vo.User) 
     gotom.LD("==>%s \n", u)
     u.Name = "bbbb"
     u.WeChat.Unionid = "testunion"
     tr := time.Now().Unix()
     u.WeChat.TokenTime = tr
     u.WeChat.Token = "abcd"
     u.WeChat.City = "a"
     u.WeChat.Country = "b"
     u.WeChat.NickName = "test-nick"
     u.WeChat.Sex = "1"
     _, err = UpdateUserWeChat(sess, u)
     if err != nil {
            t.Fatal(" failed ====> %s\n", err)
     }

     gret, err = GetUserById(sess, vo.Wid(u.Uid))
     if err != nil {
          t.Fatal("=== assert failed not found %s\n", err)
     }

     u = (gret).(*vo.User) 

     gotom.LD("==>%s\n", u)
     if u.WeChat.Sex != "1" || u.WeChat.City != "a" || u.WeChat.Token != "abcd"  ||u.WeChat.TokenTime != tr || u.WeChat.NickName != "test-nick"  || u.WeChat.Unionid != "testunion" {
          t.Fatal("=== assert failed not match %s\n", u)
     }

}
