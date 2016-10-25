



package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
    "main/service/wechat"
    "encoding/json"
    "strconv"
    "time"
)


func QuestionDetailHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if req.Req.Method == "GET" {
         if req.P("type") == "wcjs" {
               outputJsConfig(&resp, req)
               return nil, nil, nil
         }
         qid := req.Req.FormValue("qid") 
         if qid == "" || len(qid) <= 0 {
              //TODO redirect to error page
              return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
         }
         if gobject, err := ws.DoService(ws.GetTopicById, qid); err == nil {
             topic, ok:= gobject.(*vo.Topic)
             //FIXME should use back-end event to do udpate
             ws.DoService(ws.UpdateTopicViewCount, topic.Id)
             if user := GetLoggedUser(req); user != nil {
                  vt := new(vo.ViewTopic)
                  vt.ViewUser = user
                  vt.Topic = topic
                  ws.DoService(ws.RecordTopicViewUser, vt)
             }
             topicHtml := new(vo.TopicHtml)
             if topic != nil {
                  topicHtml.PopulateTopicReHtml(topic) 
             } else {
                  gotom.LE("===query topic failed %s   %s\n", topic, ok)
             }
             relatedList := make([]vo.TopicHtml, 0, 4)
             for i := 0; i < 4; i ++  {
                   rt := vo.TopicHtml{}
                   rt.Title ="内分泌失调怎么办?"
                   rt.AskToName ="刘博士"
                   rt.AskToTitle = "北京中医药大学博士"
                   relatedList = append(relatedList, rt)
             }
 
             topicHtml.RelatedList = relatedList
             return tpls.Tpls["questiondetail"], &topicHtml, nil
         }
         gotom.LD("========not found question\n")
         return tpls.Tpls["questiondetail"], nil, nil
     } else if req.Req.Method == "POST" {
         cuser :=  GetLoggedUser(req)
         if cuser == nil {
             //TODO need to login first
         }
         qid := req.P("qid")
         vid := req.P("vid")
         updateAns(qid, vid, cuser)
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}


type JsConfig struct {
       Appid        string `json:"appid"`
       Timestamp    string `json:"timestamp"`
       Nonce        string `json:"nonce"`
       Sign         string `json:"sign"`
}

func outputJsConfig(resp * gotom.GTResponse, req * gotom.GTRequest) {
     non, tm, sign := wechat.DC().Js.S("http://www.wenbar.cn/question?qid="+req.P("qid"))
     jc := JsConfig {Appid : wechat.DC().AppId, Timestamp : strconv.Itoa(int(tm.Unix())),  Nonce : non, Sign : sign}
     gotom.LD("=== config :%s\n", jc)

     content, err := json.Marshal(jc)
     if err != nil {
           gotom.LE(" json marshall failed %s\n", err)
     } else {
           (*resp.Resp).Write(content)
     }
}


func updateAns(qid string, vid string, user * vo.User) (int) {
     gotom.LD("==== qid %s   vid :%s\n", qid, vid)
     uripath := qid +"_" +strconv.Itoa(int(time.Now().Unix())) + ".amr"
     //TODO distribute to different folder
     mfile := HostConf.AudioRoot + uripath
     ts := time.Now().Unix()    
     ret:= wechat.DC().DownloadMediaFile(vid, mfile)
     te := time.Now().Unix()
     gotom.LI("Get wechat media ret:%b  cost :%d\n", ret, (te - ts))
     if gobject, err := ws.DoService(ws.GetTopicById, qid); err == nil {
           ans := &vo.Answer{AudioPath : uripath, AnsUser : user}
           ws.DoService(ws.AddTopicAnswer, gobject, ans)
     }
     return  0
}


