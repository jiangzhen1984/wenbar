

package vo


import (
    "time"
    "strconv"
)



type Topic struct {

     Id        Wid      `json:"id" bson:"_id,omitempty"`
     Title     string 
     Content   string
     Price     float32
     Count     uint32
     Date      time.Time
     Creator   *User
     AnsList   []*Answer       
}


func (t Topic) GetCount() string {
     return strconv.FormatUint(uint64(t.Count), 10)
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

