
package vo

import "gopkg.in/mgo.v2"
import "fmt"

import "testing"


func TestSession(t * testing.T) {
   sess, err := mgo.Dial("localhost")
   fmt.Printf("===%s   %s\n", sess, err)
   sess1 := sess.Copy()

   fmt.Printf("===sess1 %s\n", sess1)
   err1 := sess1.Ping()
   fmt.Printf("=sess1===%s\n", err1)

   err = sess.Ping()
   fmt.Printf("=sess===%s\n", err)


   sess1.Close()
   err = sess.Ping()
   fmt.Printf("=sess1 closed sess ping===%s\n", err)
   


   sess2 :=sess.Clone()
   err = sess2.Ping()
   fmt.Printf("=sess2 ping===%s\n", err)

   sess2.Close()
   err = sess.Ping()
   fmt.Printf("=sess2 closed sess ping===%s\n", err)

   sess3 := sess.New()
   err = sess3.Ping()
   fmt.Printf("=sess3 ping===%s\n", err)
   sess3.Close()
   err = sess.Ping()
   fmt.Printf("=sess ping===%s\n", err)


   sess.Close()
   sess3.Close()
   
}
