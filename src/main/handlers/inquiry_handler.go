



package handlers

import (
    "gotom"
    "main/service"
    "main/service/vo"
    "strconv"
)


func InquiryHandler(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping)  (*gotom.GTTemplate, gotom.Object, error) {

     if UserLoginCheck(req) != true {
          Redirect(resp, req, "/login?from=/inquiry&anu=" + req.Req.FormValue("anu")) 
          return nil, nil, nil
     }

     if req.Req.Method == "GET" {
          inq := vo.InquiryHtml{}
          uid := req.P("anu")
          if len(uid) <= 0  {
               return tpls.Tpls["inquiry"], inq, nil
          }

          gotom.LD(" query user id according to %s\n", uid)
          pouser := gotom.Object(vo.Wid(uid))
          if gobject, err := ws.DoService(ws.GetUserById, &pouser); err == nil {
               user := (*gobject).(*vo.User)
               //shuold always  okay
               
               inq.InqName    = user.Name
               inq.InqTitle   = user.Title 
               inq.InqUid     = user.Uid         
               if user.Personal != nil {
                    inq.InqUserAnsweredQues  = user.Personal.Ans
                    inq.InqUserRevenue       = user.Personal.Revenue
               }
          } else {
               gotom.LE(" get user by id failed %s\n", err)
          }
          return tpls.Tpls["inquiry"], inq, nil
          
     } else if req.Req.Method == "POST" {
          asq, _ :=  strconv.ParseInt(req.P("iuq"), 10, 32)
          iq, _  :=  strconv.ParseFloat(req.P("iur"), 32)
          inq := vo.InquiryHtml{}
          inq.InqName             = req.P("inqname")
          inq.InqTitle            = req.P("inqtitle")
          inq.InqUserAnsweredQues = int(asq) 
          inq.InqUserRevenue      = float32(iq)
          inq.InqUid              = vo.Wid(req.P("anu"))
          con := req.P("content")
          if len(con) <= 0 {
              inq.ErrMsg = "请输入问题内容"
              return tpls.Tpls["inquiry"], inq, nil
          }
              
          gtopic  := gotom.Object(vo.Topic{Title:"sss", Content:con, Creator : GetLoggedUser(req)})
          _, err := ws.DoService(ws.CreateTopic, &gtopic) 
          if err != nil {
                inq.ErrMsg = "提交问题失败"
                return tpls.Tpls["inquiry"], inq, nil
          }
          Redirect(resp, req, "/my_inquiry") 
     }
    
     return nil, nil, gotom.ErrorMsg("Not support method %s", req.Req.Method)
}



