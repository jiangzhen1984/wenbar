

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
}







type GoTomHandler func(resp GTResponse, req * GTRequest)


