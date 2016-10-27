
package gotom


import "fmt"
import "testing"
import "runtime"


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


func TestMemory(t * testing.T) {
    var size = 3;
    VLogLevel = VWarn
    InitContext()
    var rm *runtime.MemStats = new(runtime.MemStats)
    var sessArr []*GTSession = make([]*GTSession, 0, size)
    runtime.GC()
            runtime.ReadMemStats(rm)
            LW("  Test size: %d\n", size) 
            LW("  Memory: \n") 
            LW("     Alloc %d\n", rm.Alloc) 
            LW("     TotalAlloc %d\n", rm.TotalAlloc) 
            LW("     Sys %d\n", rm.Sys) 
            LW("     Mallocs %d\n", rm.Mallocs) 
            LW("     Frees %d\n", rm.Frees) 

    for i := 0; i < size; i++ {
        sess :=SerCtx.CreateSession()
        sessArr = append(sessArr, sess)
    }
    LW(" =============Preparing release session=============== \n\n")
        runtime.ReadMemStats(rm)
            LW("  Memory: \n") 
            LW("     Alloc %d\n", rm.Alloc) 
            LW("     TotalAlloc %d\n", rm.TotalAlloc) 
            LW("     Sys %d\n", rm.Sys) 
            LW("     Mallocs %d\n", rm.Mallocs) 
            LW("     Frees %d\n", rm.Frees) 
    for idx, sess := range(sessArr) {
        sess.Invalidate() 
        sessArr[idx] = nil
    }

    runtime.GC()
    runtime.ReadMemStats(rm)
    LW("  Memory: \n") 
    LW("     Alloc %d\n", rm.Alloc) 
    LW("     TotalAlloc %d\n", rm.TotalAlloc) 
    LW("     Sys %d\n", rm.Sys) 
    LW("     Mallocs %d\n", rm.Mallocs) 
    LW("     Frees %d\n", rm.Frees) 
}

