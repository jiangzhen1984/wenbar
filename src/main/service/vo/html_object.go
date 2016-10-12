

package vo

import (
    "gotom"
    "html/template"
    "strings"
    "time"
    "strconv"
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
     Tid           Wid                 `json:"id" `
     Title         string              `json:"title"` 
     Content       template.HTML       `json:"content"`
     UserCount     uint32              `json:"count"`
     ErrMsg        string
     CreatorName   string
     CreatorTitle  string
     CreatorId     string
     AskToName     string
     AskToId       string              `json:"askto"`
     AskToTitle    string
     AnsList       []TopicAnswerList
     Date          string              `json:"date,string"`
     RecCount      uint32
     Price         float32
     TimeStamp     time.Time           `json:"timestamp,string"`
     RelatedList   []TopicHtml
     AudioUrl      string
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
     //TODO use normal audio url
     th.AudioUrl     = "http://kolber.github.io/audiojs/demos/mp3/juicy.mp3"
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
     //TODO use normal audio url
     th.AudioUrl     = "http://kolber.github.io/audiojs/demos/mp3/juicy.mp3?" + strconv.Itoa(int(th.TimeStamp.Unix())
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

