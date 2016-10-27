
package gotom



const (
     evt_ID_NEW_SESS = 1
)

type nativeNewSessEvt struct {
}

func (sdt nativeNewSessEvt) GetEType() (int) {
     return evt_ID_NEW_SESS
}




type nativeEventHandler struct {
}

func (neh * nativeEventHandler) HandleEvent(evt * GTEvent) {
     if evt == nil {
          LE(" nil event coming ")
          return
     }

     if evt.EType.GetEType() == evt_ID_NEW_SESS {
         LW("Event received ")
         //TODO add to list
     }
}

//TODO add native go routine to clean sesion if sess expired





