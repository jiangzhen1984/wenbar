
package gotom


import "fmt"
import "testing"


type DefaultMapping struct {

}



func (dm DefaultMapping) OnHandle(resp GTResponse, req * GTRequest) {
   fmt.Printf("===handler test\n") 
   
}








func TestAddMappingUri(t *testing.T) {
   SerCtx.AddMappingUri("/tst", nil)

   mp := SerCtx.GetMapping("/tst")

   t.Logf("====>%s", mp)
   if mp == nil  || mp.Uri != "/tst" || mp.Hld == nil {
      t.FailNow()
   }

   t.Logf("===== finish test TestAddMapping")
}




func TestCreateSession(t * testing.T) {
  InitContext()

  sess := SerCtx.CreateSession()

  if sess == nil || sess.Id <= 0 {
      t.Fatal("create session failed")
  }

  sess = SerCtx.GetSession(sess.Id)

  if sess == nil {
      t.Fatal("get session")
  }
}

