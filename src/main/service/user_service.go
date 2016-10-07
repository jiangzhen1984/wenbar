

package ws

import (
     "gotom"
     "main/service/vo"
     "gopkg.in/mgo.v2/bson"
)

func GetUserWS(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {
     return nil,nil
}



func GetUserById(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {

     var user vo.User

     if o == nil || len(o) <= 0 {
           return nil, gotom.ErrorMsg("Parameter not statist\n")
     }
     if uid, ok := (o[0]).(vo.Wid); ok == true {
           sess := dbs.GetMongoSession()
           gotom.LD("==== query uid :%s\n", uid)
           err := sess.DB("test1").C("user").Find(bson.M{"_id" : uid}).One(&user) 
           if err != nil {
                return nil, gotom.E(" query failed %s\n", err)
           } else {
                return &user, nil
           }
     } else {
           return nil, gotom.E(" type not support for uid %s\n", o[0])

     }
     
}


func GetUserByOpenId(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {
    
    if o == nil || len(o) <= 0 {
         return nil, gotom.ErrorMsg("Parameter not enough")
    }

    var user vo.User
    if openid, ok := o[0].(string); ok == true {
         sess := dbs.GetMongoSession()
         err := sess.DB("test1").C("user").Find(bson.M{"wechat.openid" : openid}).One(&user)
         return &user, err
    } else {
         return nil, gotom.ErrorMsg("UID not string type")
    }
}


func CreateUser(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {

     if user, ok := o[0].(*vo.User); ok == true {
          sess := dbs.GetMongoSession()
          user.Uid = vo.Wid(bson.NewObjectId().Hex())
          err := sess.DB("test1").C("user").Insert(*user)
          return nil, err
     }

     return nil, gotom.ErrorMsg("Parameter incorrect")
     
}



func UpdateUserPersonal(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {
     if user, ok := (o[0]).(*vo.User); ok == true {
          sess := dbs.GetMongoSession()
          gotom.LD("===%s\n", user.WeChat.Unionid)
          query := bson.M{"_id" : string(user.Uid)} 
          updater := bson.M{
                            "$set" : bson.M{
                                      "personal.incoming" : user.Personal.Incoming,
                                      "personal.revenue"  : user.Personal.Revenue,
                                      "personal.beviewed" : user.Personal.BeViewed,
                                      "personal.unans"    : user.Personal.UnAns,
                                      "personal.quesfee"  : user.Personal.QuesFee,
                                      "personal.ans"      : user.Personal.Ans,
                                     },
                           }
                              
          err := sess.DB("test1").C("user").Update(query, updater)
          gotom.LI("Update personal informaton :%s \n", err)
          return nil, err
     }
     return nil, gotom.ErrorMsg("Parameter incorrect")
}

func UpdateUserWeChat(dbs * DBSession, o... gotom.Object) (gotom.Object, error) {
     if user, ok := (o[0]).(*vo.User); ok == true {
          if user.WeChat == nil || user.Personal == nil {
               return nil, gotom.ErrorMsg("Parameter incorrect")
          }
          sess := dbs.GetMongoSession()
          
          if len(user.Uid) <= 0 {
               dbuserobj, err := GetUserByOpenId(dbs, user.WeChat.OpenId)
               gotom.LD("===>>>> %s  %s\n", dbuserobj, err)
               if err != nil {
                     gotom.LE(" query user by open id failed :%s\n", err)
                     user.Uid = vo.Wid(bson.NewObjectId().Hex())
                     err := sess.DB("test1").C("user").Insert(*user)
                     gotom.LI("Create new user for wechat auth=> %s\n", user)
                     return nil, err
               } else {
                     cu, ok := dbuserobj.(*vo.User)
                     if ok == false {
                             gotom.LP(" return type not matched %s\n", dbuserobj)
                     }
                     user.Uid      = cu.Uid
                     user.Name     = cu.Name 
                     user.Title    = cu.Title 
                     user.Avatar1  = cu.Avatar1 
                     user.Avatar2  = cu.Avatar2 
               }
          }
          query := bson.M{"_id" : string(user.Uid) , "wechat.openid" : user.WeChat.OpenId} 
          updater := bson.M{
                      "$set":
                          bson.M{
                             "wechat.unionid"    : user.WeChat.Unionid,
                             "wechat.openid"     : user.WeChat.OpenId,
                             "wechat.city"       : user.WeChat.City,
                             "wechat.country"    : user.WeChat.Country,
                             "wechat.nickname"   : user.WeChat.NickName,
                             "wechat.sex"        : user.WeChat.Sex,
                             "wechat.token"      : user.WeChat.Token,
                             "wechat.tokentime"  : user.WeChat.TokenTime,
                             "wechat.avatar"     : user.WeChat.Avatar,
                          },
                    }
          err := sess.DB("test1").C("user").Update(query, updater)
          gotom.LI("Update personal informaton :%s \n", err)
          return nil, err
     }
     return nil, gotom.ErrorMsg("Parameter incorrect")
}

