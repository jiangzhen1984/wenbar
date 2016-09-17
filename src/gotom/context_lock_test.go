
package gotom


import "testing"
import "time"
import "fmt"


func RuntimeSession(id int) {
   fmt.Printf("[%d] ===> request\n", id)
   for i := 0; i < 500; i++ {
       sess := SerCtx.CreateSession()
       fmt.Printf("[%d] ======> %s  finish \n", id, sess)
   }
   
}


func TestSessionLock(t *testing.T) {
   InitContext()
   idx := 0
   for i := 0; i < 500; i++ {
      go RuntimeSession(idx)
      idx ++
   }

   time.Sleep(200000)
}
