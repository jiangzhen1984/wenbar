
var AudioPlayer = function(ae) {
	Object.defineProperty(this, "AP_STATE_INIT", {
	     value : 0,
	     writable : false,
	});
	Object.defineProperty(this, "AP_STATE_START", {
	     value : 1,
	     writable : false,
	});

	Object.defineProperty(this, "AP_STATE_DOWNLOADING", {
	     value : 2,
	     writable : false,
	});

	Object.defineProperty(this, "AP_STATE_PLAYING", {
	     value : 3,
	     writable : false,
	});

	Object.defineProperty(this, "AP_STATE_PAUSED", {
	     value : 4,
	     writable : false,
	});

	Object.defineProperty(this, "AP_STATE_ENDED", {
	     value : 5,
	     writable : false,
	});

	Object.defineProperty(this, "AP_STATE_ERROR", {
	     value : 6,
	     writable : false,
	});

     this.state =  this.AP_STATE_INIT;
     if (ae == undefined) {
         return;
     }
     if (ae != undefined){
         if (typeof ae == "string") {
             this.ae = new Audio(ae);
             this.registerAudioCB(this.ae);
         } else {
             this.ae = ae;
         }
     }
};



AudioPlayer.prototype.play = function() {
    if (this.state == this.AP_STATE_PLAYING) {
        this.stop();
    }
    if (this.ae == undefined) {
        console.log("Not initalized yet");
        return;
    }
    this.state = this.AP_STATE_INIT;
    console.log(this.ae);
    this.ae.play();
}

AudioPlayer.prototype.stop = function() {
    if (this.ae != undefined) {
       this.ae.pause();
       this.ae.currentTime = 0;
       this.state = this.AP_STATE_PAUSED;
    }
}

AudioPlayer.prototype.playOrStopUrl = function(url) {
    if (this.cache_url == url) {
        if (this.state == this.AP_STATE_PLAYING || this.state == this.AP_STATE_DOWNLOADING ) {
            this.stop();
        } else {
            this.ae = new Audio(url);
            this.registerAudioCB(this.ae);
            this.cache_url = url;
            this.play();
        }
    } else {
        this.stop();
        this.ae = new Audio(url);
        this.registerAudioCB(this.ae);
        this.cache_url = url;
        this.play();
    }
}



AudioPlayer.prototype.registerAudioCB = function(ad) {
      if (ad == undefined) {
            return;
      }
      var ap = this;
      ad.addEventListener("abort", function() {
            console.log("abort");
      });

      ad.addEventListener("canplay", function() {
            if (!ad.paused) {
                 ap.state = ap.AP_STATE_PLAYING;
                 if (ap.evt_cb != undefined) {
                      ap.evt_cb(ap, ap.AP_STATE_PLAYING);
                 }
            }
      });
      ad.addEventListener("ended", function() {
            ap.state = ap.AP_STATE_ENDED;
                 if (ap.evt_cb != undefined) {
                      ap.evt_cb(ap, ap.AP_STATE_ENDED);
                 }
            console.log("end");
      });
      ad.addEventListener("loadstart", function() {
            console.log("start load");
      });
      ad.addEventListener("play", function() {
            console.log("started");
      });
      ad.addEventListener("pause", function() {
            if (ap.evt_cb != undefined) {
                 ap.evt_cb(ap, ap.AP_STATE_PAUSED);
            }
      });
      ad.addEventListener("progress", function() {
            if (ap.state == ap.AP_STATE_START) {
                 ap.state = ap.AP_STATE_DOWNLOADING;
                 if (ap.evt_cb != undefined) {
                      ap.evt_cb(ap, ap.AP_STATE_DOWNLOADING);
                 }
            }
      });
      ad.addEventListener("error", function() {
            console.log("error");
            ap.state = ap.AP_STATE_ERROR;
                 if (ap.evt_cb != undefined) {
                      ap.evt_cb(ap, ap.AP_STATE_ERROR);
                 }
      });
};


AudioPlayer.prototype.addStateListener = function(cb) {
     console.log(typeof cb);
     this.evt_cb = cb;
};



