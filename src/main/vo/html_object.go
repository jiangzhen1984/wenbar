

package vo



type HotListHtml struct {

    Title      string
    TopicList  []TopicHtml
}


type TopicHtml struct {
     Tid           uint32
     Title         string
     AnsUserName   string
     AnsUserTitle  string
     UserCount     uint32
}


type LoginHtml struct {

    PhoneNumber string
    ErrMsg      string
}



type PersonalHtml struct {


    Name         string
    Title        string
    Incoming     float32 
    Revenue      float32
    Fee          float32
}
