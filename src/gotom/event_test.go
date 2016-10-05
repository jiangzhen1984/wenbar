

package gotom

import "testing"
import "time"


type TestEvent struct {
     id      int
}


func (te * TestEvent) GetEType()  (int) {
     return 1
}

func (te * TestEvent) HandleEvent(evt * GTEvent) {
    LD("===> handle event %s\n", evt)
}


func TestInitLooper(t * testing.T) {
    looper := InitEventLooper("aaa", 10)
    if looper.Name != "aaa" || looper.IsReady() == true {
           t.Fatal("Name not correct or is ready")
    }
}


func TestWaitingForEvent(t * testing.T) {
    looper := InitEventLooper("aaa", 10)
    if looper.Name != "aaa" || looper.IsReady() == true {
           t.Fatal("Name not correct or is ready")
    }
    te := &TestEvent{}
    looper.AddEventHandler(GTEventType(te),GTEventHandler(te))
    go looper.WaitingForEvent()
    time.Sleep(time.Duration(3)*time.Second)
    e := looper.PostEvent( &TestEvent{}, nil)
    LD("==== event create %s\n", e)

    for i :=1; i< 10; i++ {
          go sendEvent(i * 100, looper)
    }
    time.Sleep(time.Duration(10)*time.Second)
}


func sendEvent(idx int, looper * GTLooper) {
    for i :=1; i< 10; i++ {
        te := &TestEvent{id : idx + i}
        looper.PostEvent(te, te)
        time.Sleep(30 * time.Millisecond )
    }
}



func TestDefaultEventLooper(t * testing.T) {
    te := &TestEvent{}
    AddEventHandler(GTEventType(te),GTEventHandler(te))
    for i :=1; i< 10; i++ {
          go sendEventToDefault(i * 100)
    }
    time.Sleep(time.Duration(10)*time.Second)
}


func sendEventToDefault(idx int) {
    for i :=1; i< 10; i++ {
        te := &TestEvent{id : idx + i}
        PostEvent(te, te)
        time.Sleep(30 * time.Millisecond )
    }
}


