
package ws

import (
    "gotom"
    "gopkg.in/mgo.v2"
)


var (
     localDBConf * DBConfiguration
     localGlobalSession * mgo.Session
)


type DBSession struct {
    
     db   interface{}
}


type DBConfiguration struct {

     DBUrl       string
     DBUname     string
     DBUpwd      string
     DBName      string
}


func (dbs * DBSession) Close() {
     CloseDBSession(dbs)
}


func (dbs * DBSession) GetMongoSession() *mgo.Session {
     msess, ok := dbs.db.(*mgo.Session)
     if !ok {
          panic("==== dbs.db cast failed\n")
     }
     return msess
}


func InitDB(dbc DBConfiguration) {
      localDBConf = & DBConfiguration {
                         DBUrl     : dbc.DBUrl  ,
                         DBUname   : dbc.DBUname  ,
                         DBUpwd    : dbc.DBUpwd ,
                         DBName    : dbc.DBName ,
                    }

     sess, err :=  mgo.Dial(localDBConf.DBUrl)
     if err != nil {
           panic(err)
     } 

     localGlobalSession = sess
     gotom.LD(" Initial DB Session successfully ==>%p \n", localGlobalSession)
}


func CreateDBSession() (* DBSession) {
   return & DBSession { db : localGlobalSession.New() }
}


func CloseDBSession(dbs * DBSession) {
     msess, ok := dbs.db.(*mgo.Session)
     if !ok {
          panic("==== dbs.db cast failed\n")
     }
     msess.Close()
}


     
type WService func(o ...gotom.Object)  (gotom.Object)




type WSServiceFunc func(ds * DBSession, o ...gotom.Object) (gotom.Object, error)


func DoService(wf WSServiceFunc, o ...gotom.Object) ( gotom.Object, error) {
      gotom.LI("====> %s\n", wf)
      gotom.LI("====> %s\n", o)
      sess := CreateDBSession() 
      gobject, err := wf(sess, o...)
      CloseDBSession(sess)
      gotom.LI("<==== %s err :%s \n", wf, err)
      
      return gobject, err
}



