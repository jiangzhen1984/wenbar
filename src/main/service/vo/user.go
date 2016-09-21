
package vo


type User struct {

     Uid       Wid
     
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
