

package gotom

import (
    "net/http"
    "strconv"
)


func (gth GoTomHandler) OnHandler(resp http.ResponseWriter, req * http.Request) {
    var sess * GTSession
    cki, err := req.Cookie(GOTOM_SESSION_ID)
    if err == nil && cki != nil {
         if idv, err := strconv.ParseUint(cki.Value, 10, 64); err == nil {
              sess = SerCtx.GetSession(idv)
         }
    }
    LI(" request session :%s\n", sess)
    greq := &GTRequest{Req : req, sess : sess, Ctx : SerCtx}
    gth(GTResponse{Resp : &resp}, greq)
    sess = greq.GetSession(false)
    if sess != nil {
        http.SetCookie(resp, &http.Cookie{Name : GOTOM_SESSION_ID, Value : strconv.FormatUint(sess.Id, 10)})
    }
}







type GoTomHandler func(resp GTResponse, req * GTRequest)


