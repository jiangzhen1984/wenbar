
package vo


type User struct {

     Uid       uint32
     
     Name      string 

     Title     string

     Avatar1   string
  
     Avatar2   string

   
     Personal  *UserPersonal
}



type UserPersonal struct {

     Incoming     float32

     Revenue      float32

     QuesFee      float32 
}
