
package gotom


import (
    "net/http"
    "sync"
    "math/rand"
    "strconv"
)


const (
     NO_ERR              = iota
     ERR_KEY_DUPLICATE
     ERR_PARMETER_NIL
)

type EnumRet uint


var (
    SerCtx * GTServerContext
)


type GTServerContext struct {

    sess      map[uint64]*GTSession
    mu        sync.Mutex

    mapping   map[string]*Mapping
    mpmu      sync.Mutex
}


type GTSession struct {

    id         uint64

    attrs      map[string]string
    mu         sync.Mutex

    ctx        *GTServerContext

    valid    bool
}


type GTRequest struct {

    req        *http.Request

    attrs      map[string]string
    mu         sync.Mutex

    sess       *GTSession
    ctx        *GTServerContext
}


type GTResponse struct {

    resp       *http.ResponseWriter
}



type Mapping struct {

    uri        string

    hld        GoTomHandler
}


type MappingHandler interface {
    OnHandle(GTResponse, *GTRequest)
}



func (ctx * GTServerContext) GetSession(id uint64) * GTSession {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    if ctx.sess == nil {
        return nil
    }

    return ctx.sess[id]
}


func (ctx * GTServerContext) RemoveSession(id uint64) {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    if ctx.sess == nil {
        return
    } 

    ctx.mu.Lock()
    defer ctx.mu.Unlock()
    delete(ctx.sess, id)
}



func (ctx * GTServerContext) CreateSession() *GTSession {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }


    ctx.mu.Lock()
    defer ctx.mu.Unlock()
    if ctx.sess == nil {
        ctx.sess = make(map[uint64]*GTSession)
    }

    sid  := uint64(rand.Int63())
    session := &GTSession{id:sid, attrs:nil} 
    session.ctx =  SerCtx
    session.valid = true
    
    ctx.sess[sid] = session 
    return session
}



func (ctx * GTServerContext) AddMappingUri(uri string, hdr GoTomHandler) EnumRet {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    if hdr == nil {
        return ERR_PARMETER_NIL
    }

    ctx.mpmu.Lock()
    defer ctx.mpmu.Unlock()
 
    if ctx.mapping == nil {
         ctx.mapping = make(map[string]*Mapping)
    } else {
         mh := ctx.mapping[uri]
         if mh != nil {
                return ERR_KEY_DUPLICATE
         }
    }

    ctx.mapping[uri] = & Mapping{uri : uri, hld : hdr}
    
    return NO_ERR
}

func (ctx * GTServerContext) AddMapping(mapping * Mapping) EnumRet {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    if mapping == nil {
        return ERR_PARMETER_NIL
    }

    ctx.mpmu.Lock()
    defer ctx.mpmu.Unlock()
    if ctx.mapping == nil {
         ctx.mapping = make(map[string]*Mapping)
    } else {
         mh :=  ctx.mapping[mapping.uri]
         if mh != nil {
                return ERR_KEY_DUPLICATE
         }
    }

    ctx.mapping[mapping.uri] = mapping
    return NO_ERR
}

func (ctx * GTServerContext) RemoveMapping(uri string) {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    ctx.mpmu.Lock()
    defer ctx.mpmu.Unlock()
    delete(ctx.mapping, uri)
}

func (ctx * GTServerContext) UpdateMapping(uri string, hdr  GoTomHandler) {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }

    if hdr == nil {
        return
    }

    ctx.mpmu.Lock()
    defer ctx.mpmu.Unlock()

    mp := ctx.mapping[uri]
    if mp != nil {
         mp.hld = hdr
    } else {
         ctx.mapping[uri] = & Mapping {uri : uri, hld : hdr}
    }
}


func (ctx * GTServerContext) GetMapping(uri string) * Mapping {

    if ctx == nil {
        LP("Conext doesn't initalize\n")
    }
 
    if ctx.mapping == nil {
         return nil
    }

    return ctx.mapping[uri]
}



func (sess * GTSession) GetAttribute(key string) string {

    if sess == nil {
        LP("Session is null\n")
    }

    if sess.attrs == nil {
         return ""
    }

    return sess.attrs[key]
}

func (sess * GTSession) SetAttribute(key string, value string) {

    if sess == nil {
        LP("Session is null\n")
    }

    sess.mu.Lock()
    defer sess.mu.Unlock()
    if sess.attrs == nil {
         sess.attrs = make(map[string]string)
    }

    sess.attrs[key] = value
}

func (sess * GTSession) RemoveAttribute(key string) {

    if sess == nil {
        LP("Session is null\n")
    }

    sess.mu.Lock()
    defer sess.mu.Unlock()
    if sess.attrs == nil {
         return
    }

    delete(sess.attrs, key)
}


func (sess * GTSession) Invalidate() {
    if sess == nil {
        LP("Session is null\n")
    }

    sess.ctx.RemoveSession(sess.id) 
    sess.valid = false
}


func (sess GTSession) String() string {
     return "(" + strconv.FormatUint(sess.id, 10) +", "+strconv.FormatBool(sess.valid)+")"
}

func (req * GTRequest) GetAttribute(key string) string {
    if req.attrs == nil {
         return ""
    }

    return req.attrs[key]
}

func (req * GTRequest) SetAttribute(key string, value string) {

    req.mu.Lock()
    defer req.mu.Unlock()
    if req.attrs == nil {
         req.attrs = make(map[string]string)
    }

    req.attrs[key] = value
}

func (req * GTRequest) RemoveAttribute(key string) {
    req.mu.Lock()
    defer req.mu.Unlock()
    if req.attrs == nil {
         return
    }

    delete(req.attrs, key)
}


