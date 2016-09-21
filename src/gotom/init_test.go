
package gotom


import (
    "testing"
    "strconv"
)


func TestInitContext(t *testing.T) {
    if SerCtx != nil {
         t.Fatal(" SerCtx initalized ")
    }

    InitContext()
    if SerCtx == nil {
         t.Errorf("==== Context doesn't inital \n")
    }
    t.Logf("===== finish test TestInitContext")
}



func TestInitMapping(t *testing.T) {
LP("ssss")
    InitContext()

    mappings := make([]*Mapping, 10, 10)

    for idx := 0 ; idx < cap(mappings); idx++ {
         mappings[idx] = &Mapping{Uri:"/test" + strconv.FormatInt(int64(idx), 10)}
    } 
    
    InitMappings(mappings)

    for _, mp := range mappings {
        tmp := SerCtx.GetMapping(mp.Uri)
        if tmp == nil {
              t.Fatal(" failed at %s\n", mp)
        }
    }
    
}


func testOnHandler(resp GTResponse, req * GTRequest) {
     LD("==== on Hanlder \n")
}


func TestInitServer(t * testing.T) {
    InitContext()

    mappings := make([]*Mapping, 10, 10)

    for idx := 0 ; idx < cap(mappings); idx++ {
         mappings[idx] = &Mapping{Uri:"/test" + strconv.FormatInt(int64(idx), 10), Hld : nil}
    } 
    
    InitMappings(mappings)

    InitServer(nil)
}
