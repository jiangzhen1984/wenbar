
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
  
     WeChat    *PersonalWeChat
}


func (u User) GetNativeID() string {
     return strconv.FormatUint(u.NativeId, 10)
}



type UserPersonal struct {

     Incoming     float32

     Revenue      float32

     QuesFee      float32 

     Ans          int
 
     BeViewed     int

     UnAns        int
}


type PersonalWeChat struct {

     OpenId        string

     NickName      string
  
     Sex           int

     Province      string

     City          string

     Country       string

     Unionid       string

     Avatar        string

     Token         string

     TokenTime     int64

     TokenExprisIn int
}



