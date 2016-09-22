
package vo


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



type UserPersonal struct {

     Incoming     float32

     Revenue      float32

     QuesFee      float32 
}



