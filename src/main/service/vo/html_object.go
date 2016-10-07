

package vo

import (
    "gotom"
    "html/template"
    "strings"
    "time"
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
     Content       template.HTML
     UserCount     uint32
     ErrMsg        string
     CreatorName   string
     CreatorTitle  string
     CreatorId     string
     AskToName     string
     AskToId       string
     AskToTitle    string
     AnsList       []TopicAnswerList
     Date          string
     RecCount      uint32
     Price         float32
     TimeStamp     time.Time 
     RelatedList   []TopicHtml
}


func (th * TopicHtml) PopulateTopicReHtml(val *Topic) {
     if val == nil {
         gotom.LP(" topic is nil")
     }
     th.Tid          = val.Id
     th.Title        = val.Title
     th.Content      = template.HTML(strings.Replace(val.Content, "\n", "<br />", -1))
     if val.Creator != nil {
         th.CreatorName  = val.Creator.Name
         th.CreatorTitle = val.Creator.Title
         th.CreatorId    = val.Creator.GetNativeID()
     } else {
         gotom.LE("Ilegal user information for topic %s\n", val)
     }
     th.UserCount    = val.Count
     th.Date         = val.GetElapsedTime()
     th.TimeStamp    = val.Date
}

func (th * TopicHtml) PopulateTopic(val *Topic) {
     if val == nil {
         gotom.LP(" topic is nil")
     }
     th.Tid          = val.Id
     th.Title        = val.Title
     th.Content      = template.HTML(val.Content)
     if val.Creator != nil {
         th.CreatorName  = val.Creator.Name
         th.CreatorTitle = val.Creator.Title
         th.CreatorId    = val.Creator.GetNativeID()
     }
     th.UserCount    = val.Count
     th.Date         = val.GetElapsedTime()
     th.TimeStamp    = val.Date
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
    BeViewed     int
    Avatar       string
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

