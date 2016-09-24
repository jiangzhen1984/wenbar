
package vo

import (
    "strconv"
)


type User struct {

     Uid       Wid
     
     Name      string 

     Title     string

     Avatar1   string
  
     Avatar2   string

     NativeId  uint64

     OutId     string

     Personal  *UserPersonal
}


func (u User) GetNativeID() string {
     return strconv.FormatUint(u.NativeId, 10)
}



type UserPersonal struct {

     Incoming     float32

     Revenue      float32

     QuesFee      float32 

     Ans          int
}



