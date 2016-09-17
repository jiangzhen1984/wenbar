

package gotom

import (
    "net/http"
    "strconv"
)


func (gth GoTomHandler) OnHandler(resp http.ResponseWriter, req * http.Request) {
    var sess * GTSession
    if cki, err := req.Cookie("gotom_session_id"); err != nil && cki != nil {
         if idv, err := strconv.ParseUint(cki.Value, 10, 64); err == nil {
              sess = SerCtx.GetSession(idv)
         }
    }
    LI(" request session :%s\n", sess)
    gth(GTResponse{resp : &resp}, &GTRequest{req : req, sess : sess, ctx : SerCtx})
}







type GoTomHandler func(resp GTResponse, req * GTRequest)


