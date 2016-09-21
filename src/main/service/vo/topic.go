

package vo


import (
    "time"
)



type Topic struct {

     Id        Wid
     Title     string 
     Price     float32
     Count     uint32
     Date      time.Time
     Creator   *User
     AnsList   []*Answer       
}


type Answer struct {

    Id        Wid
    Date      time.Time
    AnsUser   *User
}
