

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
     IsPublic  bool
     AskTo     Wid
     AnsList   []*Answer       
}


func (t * Topic) GetCount() string {
     return strconv.FormatUint(uint64(t.Count), 10)
}

func (t * Topic) GetElapsedTime() string {
    var el []string = []string{"秒", "分钟", "小时", "天", "周", "月" , "年"}
    duration := time.Now().Sub(t.Date)
    if duration.Hours() >= 1 {
       hr := int(duration.Hours()) + 1
       if m := hr / (24 * 30); m >= 1 {
           return strconv.Itoa(m) + el[5]
       } else if w := hr/ (24 * 7); w >= 1 {
           return strconv.Itoa(w) + el[4]
       } else if d := hr/ 24; d >= 1 {
           return strconv.Itoa(d) + el[3]
       } else {
           return strconv.Itoa(hr) + el[2]
       }
    } else if duration.Minutes() > 1 {
       return strconv.Itoa(int(duration.Minutes())) + el[1]
    } else {
       return strconv.Itoa(int(duration.Seconds())) + el[0]
    }
    return el[1]
}



type Answer struct {

    Id        Wid
    Date      time.Time
    AnsUser   *User
    Content   string
    AudioPath string
}


type ViewTopic struct {

    Topic       *Topic `bson:"-"`
    TopicId     Wid
    ViewUser    *User  `bson:"-"`	
    ViewUserId  Wid
    Date        time.Time

}

