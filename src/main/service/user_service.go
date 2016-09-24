

package ws

import (
     "gotom"
     "main/service/vo"
     "gopkg.in/mgo.v2/bson"
)

func GetUserWS(dbs * DBSession, o... *gotom.Object) (*gotom.Object, error) {
     return nil,nil
}



func GetUserById(dbs * DBSession, o... *gotom.Object) (*gotom.Object, error) {

     var user vo.User

     if o == nil || len(o) <= 0 {
           return nil, gotom.ErrorMsg("Parameter not statist \n")
     }
     if uid, ok := (*o[0]).(vo.Wid); ok == true {
           sess := dbs.GetMongoSession()
           err := sess.DB("test1").C("user").Find(bson.M{"uid" : uid}).One(&user) 
           if err != nil {
                return nil, gotom.E(" query failed %s\n", err)
           } else {
                gobject := gotom.Object(&user)
                return &gobject, nil
           }
     } else {
           return nil, gotom.E(" type not support for uid %s\n", o[0])

     }
     
}
