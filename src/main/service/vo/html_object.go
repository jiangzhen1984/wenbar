

package vo

import (
    "gotom"
)


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
     Content       string
     UserCount     uint32
     ErrMsg        string
     CreatorName   string
     CreatorTitle  string
     CreatorId     string
     AnsList       []TopicAnswerList
     Date          string
     RecCount      uint32
     Price        float32
}


func (th * TopicHtml) PopulateTopic(val *Topic) {
     th.Tid          = val.Id
     th.Title        = val.Title
     th.Content      = val.Content
     th.CreatorName  = val.Creator.Name
     th.CreatorTitle = val.Creator.Title
     th.CreatorId    = val.Creator.GetNativeID()
     th.UserCount    = val.Count
     th.Date         = val.GetElapsedTime()
     gotom.LD("=====%s\n", th.Date)
}


type TopicAnswerList struct {
     AnsUserName  string
     AnsUserTitle string
     Ans          string 
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


type InquiryHtml struct {
    InqUid                Wid
    InqName               string
    InqTitle              string
    InqUserAnsweredQues   int
    InqUserRevenue        float32
    IsPublic              bool
    ErrMsg                string
}

