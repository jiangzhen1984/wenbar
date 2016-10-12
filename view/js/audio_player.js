

var AudioPlayer = function(ae) {
     this.state =  0;
     if (ae == undefined) {
         return;
     }
     if (ae != undefined){
         if (typeof ae == "string") {
             this.ae = new Audio(ae);
         } else {
             this.ae = ae;
         }
     }
};


AudioPlayer.prototype.play = function() {
    if (this.state == 1) {
        this.stop();
    }
    if (this.ae == undefined) {
        console.log("Not initalized yet");
        return;
    }
    this.state = 1;
    console.log(this.ae);
    this.ae.play();
}

AudioPlayer.prototype.stop = function() {
    if (this.ae != undefined) {
       this.ae.pause();
       this.ae.currentTime = 0;
       this.state = 0;
    }
}

AudioPlayer.prototype.playOrStopUrl = function(url) {
    if (this.cache_url == url) {
        if (this.state == 1) {
            console.log("===>stop");
            this.stop();
        } else {
            console.log("===>start");
            this.ae = new Audio(url);
            this.cache_url = url;
            this.play();
        }
    } else {
        console.log("mismatch===>start");
        this.stop();
        this.ae = new Audio(url);
        this.cache_url = url;
        this.play();
    }
}

