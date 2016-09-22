

package vo


import (
    "time"
)



type Topic struct {

     Id        Wid
     Title     string 
     Content   string
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
    Content   string
    AudioPath string
}


type ViewTopic struct {

    Topic       *Topic
    ViewUser    *User
    Date        time.Time

}

