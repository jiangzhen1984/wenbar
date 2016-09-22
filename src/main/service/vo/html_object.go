

package vo


type ItemPage struct {
    
     CurPage     uint32
     PageSize    uint32
     ItemOffset  uint32
}

type HotListHtml struct {

    Title      string
    TopicList  []TopicHtml
}


type TopicHtml struct {
     Tid           Wid
     Title         string
     AnsUserName   string
     AnsUserTitle  string
     UserCount     uint32
     ErrMsg        string
}


type LoginHtml struct {

    PhoneNumber string
    ErrMsg      string
    From        string
}



type PersonalHtml struct {


    Name         string
    Title        string
    Incoming     float32 
    Revenue      float32
    Fee          float32
}


type MyInquiryHtml struct {

    TopicList   []TopicHtml

    Page        ItemPage
}

type MyAnswerHtml struct {

    TopicList   []TopicHtml

    Page        ItemPage
}

