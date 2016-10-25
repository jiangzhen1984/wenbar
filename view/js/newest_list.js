


var player = new AudioPlayer();

player.addStateListener(function(ap, evt) {
     switch (evt) {
         case ap.AP_STATE_START:
             showNotification("开始准备数据");
             break;
         case ap.AP_STATE_DOWNLOADING:
             showNotification("正在下载音频数据");
             break;
         case ap.AP_STATE_PLAYING:
             showNotification("正在正在播放音频");
             break;
         case ap.AP_STATE_PAUSED:
             console.log("===get event>>> " + evt+"  p ");
             break;
     }
});


function onAnsClicked(ul) {
    if (ul == undefined || ul == "") {
        showNoAns();
    } else {
        player.playOrStopUrl(ul);
    }
}


var g_toast;
function showNotification(str) {
    if (g_toast != undefined) {
         g_toast.reset();
    }
    g_toast = $.toast({text : str, position: 'mid-center', hideAfter : 0,}); 
}

function showNoAns() {
    $.toast({text : '该问题还在悬赏。。', position: 'mid-center', hideAfter : 500,}); 
}
