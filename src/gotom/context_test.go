
package gotom


import "fmt"
import "testing"


type DefaultMapping struct {

}



func (dm DefaultMapping) OnHandle(resp GTResponse, req * GTRequest) {
   fmt.Printf("===handler test\n") 
   
}



func stest(resp GTResponse, req * GTRequest) {
}





func TestAddMappingUri(t *testing.T) {
   SerCtx.AddMappingUri("/tst", stest)

   mp := SerCtx.GetMapping("/tst")

   t.Logf("====>%s", mp)
   if mp == nil  || mp.uri != "/tst" || mp.hld == nil {
      t.FailNow()
   }

   t.Logf("===== finish test TestAddMapping")
}




func TestCreateSession(t * testing.T) {
  InitContext()

  sess := SerCtx.CreateSession()

  if sess == nil || sess.id <= 0 {
      t.Fatal("create session failed")
  }

  sess = SerCtx.GetSession(sess.id)

  if sess == nil {
      t.Fatal("get session")
  }
}

