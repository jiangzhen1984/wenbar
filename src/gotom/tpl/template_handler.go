

package tpl

import (
    "net/http"
    "strconv"
    "gotom"
)


func (gtth GoTomTplHandler) OnHandler(resp http.ResponseWriter, req * http.Request) {
    var sess * gotom.GTSession
    cki, err := req.Cookie(gotom.GOTOM_SESSION_ID)
    if err == nil && cki != nil {
         if idv, err := strconv.ParseUint(cki.Value, 10, 64); err == nil {
              sess = gotom.SerCtx.GetSession(idv)
         }
    }

    gotom.LI(" request session :%s\n", sess)
    greq := &gotom.GTRequest{Req : req, sess : sess, Ctx : gotom.SerCtx}
    //TODO get template according to uri
    tplMapping := gotom.GConf.TplMapping[req.URL.Path()]
    tpl, da, err := gtth(gotom.GTResponse{Resp : &resp}, greq, tplMapping)
    if err != nil {
        tpl.NativeTpl.Execute(resp, da)
    }
}







type GoTomTplHandler func(resp gotom.GTResponse, req * gotom.GTRequest, tpls * gotom.GTTemplateMapping) (*gotom.GTTemplate, gotom.Object, error)
