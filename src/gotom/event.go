
package gotom

import (
     "sync"
)




type GTEventType interface {

     GetEType()  int
}


type GTShutDownEventType struct {
}

func (sdt GTShutDownEventType) GetEType() (int) {
     return -1
}


type GTEventHandler interface {

    HandleEvent(evt * GTEvent)
}


type GTEvent struct {

    Id       uint64

    EType    GTEventType

    Data     interface{}
}


type GTControlEvent struct {
    gt       int
}


type GTLooper struct {
    
     Name             string

     handlerMap       map[int]GTEventHandler

     isReady          bool
     lock             sync.Mutex 

     queue           chan * GTEvent
     shutdownFlag     bool
}


var DefaultLooper * GTLooper = &GTLooper{
                                Name        : "default",
                                queue       : make(chan * GTEvent),
                                handlerMap  : map[int]GTEventHandler{},
                                }


func (looper * GTLooper) PostEvent(et GTEventType, data interface{}) (*GTEvent) {
     st := looper.IsReady() 
     if st == false {
          LP("Looper not ready yet, call WaitingForEvent fisrt !\n")
     }
     ev := new(GTEvent)
     ev.EType = et
     ev.Data  = data
     postEventToLooper(ev, looper)
     return ev
}


func (looper * GTLooper) AddEventHandler(et GTEventType, handler GTEventHandler) {
     LI("Add Event(%s) Handler(%p)", et, handler)
     looper.handlerMap[et.GetEType()] = handler
}

func (looper * GTLooper) RemoveEventHandler(et GTEventType) {
     delete(looper.handlerMap, et.GetEType())
}


func (looper * GTLooper) WaitingForEvent() {
     looper.lock.Lock()
     defer looper.lock.Unlock()
     if looper.shutdownFlag == true {
         LW(" event looper is already shutdown\n")
         return
     }
     if looper.isReady == false {
         looper.isReady = true
     } else {
         LE(" event looper is already ready\n")
         return
     }
     looper.lock.Unlock()
     LI(" start event loop ")
     looper.handlerEvent()
     LI(" event loop quit")
}


func (looper * GTLooper) IsReady() (bool) {
     return looper.isReady
}


func (looper * GTLooper) shutdown() {
     looper.lock.Lock()
     looper.shutdownFlag = true
     looper.PostEvent(GTShutDownEventType{}, nil)
     looper.lock.Unlock()
}


func (looper * GTLooper) handlerEvent() {
     for looper.shutdownFlag != true {
         ev := <-looper.queue
         handler := looper.handlerMap[ev.EType.GetEType()]
         LD("===forward evt(%s)  to handler(%p)\n", ev, handler)
         if handler == nil {
              LW("No Such Handler :%s %s\n", ev.EType)
              continue
         }
         handler.HandleEvent(ev)
     }
}



func InitEventLooper(name string, size int) * GTLooper {
     var looper * GTLooper = new(GTLooper)
     looper.queue = make(chan * GTEvent, size)
     looper.Name  = name
     looper.handlerMap = make(map[int]GTEventHandler, 50)

     return looper 
}



func AddEventHandler(et GTEventType, handler GTEventHandler) {
     LI("Add To Default Looper Event(%s) Handler(%p)", et, handler)
     DefaultLooper.AddEventHandler(et, handler)
}

func RemoveEventHandler(et GTEventType) {
     DefaultLooper.RemoveEventHandler(et)
}



func PostEvent(et GTEventType, data interface{}) (*GTEvent) {
     ev := new(GTEvent)
     ev.EType = et
     ev.Data  = data
     if DefaultLooper.IsReady() == false {
           go DefaultLooper.WaitingForEvent()
           //FIXME should wait few seconds to make sure subroutine alread run
     }
     postEventToLooper(ev, DefaultLooper)
     return ev
}


func postEventToLooper(ev * GTEvent, looper * GTLooper) (bool) {
     looper.queue<- ev
     return true
}
