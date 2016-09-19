

package gotom

import (
    "net/http"
    "strconv"
)


func (gth GoTomTplHandler) OnHandle(resp http.ResponseWriter, req * http.Request) {
    var sess * GTSession
    cki, err := req.Cookie(GOTOM_SESSION_ID)
    if err == nil && cki != nil {
         if idv, err := strconv.ParseUint(cki.Value, 10, 64); err == nil {
              sess = SerCtx.GetSession(idv)
         }
    }

    if sess != nil {
         LI(" request session :%s ==> %s(%p)\n", sess, gth, gth)
    }
    greq := &GTRequest{Req : req, sess : sess, Ctx : SerCtx}
    tplMapping := GConf.TplMapping[req.URL.Path]
    tpl, da, err := gth(GTResponse{Resp : &resp}, greq, tplMapping)
    if err == nil && tpl != nil {
        tpl.NativeTpl.Execute(resp, da)
    }
}






type GoTomTplHandler func(resp GTResponse, req * GTRequest, tpls * GTTemplateMapping) (*GTTemplate, Object, error)



