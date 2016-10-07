

function initjs() {
    $.ajax({
            url: '/question?qid='+qid+'&rfrom=ajax&type=wcjs',
            dataType: 'json',
            success: function(json) {
                 if (json == null || json == "undefined") {
                      console.log("get js auth error");
                      return
                 }
                 console.log("===>appid:" + json.appid);
                 console.log("===>ts:" + json.timestamp);
                 console.log("===>nonce:" + json.nonce);
                 console.log("===>sign" + json.sign);
                 initwx(json);
            },
            error: function (data) {
                 console.log("=== failed");
            }
          });
}


function initwx(json) {
    wx.config({
        debug     : false,
        appId     : json.appid,
        timestamp : json.timestamp,
        nonceStr  : json.nonce,
        signature : json.sign,
        jsApiList : ["startRecord", "stopRecord", "onVoiceRecordEnd","uploadVoice"]
     });
}


var gstat = false;

function startRecord() {
   gstat = true;
   wx.startRecord();
}


function stopRecord() {
    gstat = false;
    wx.stopRecord({
        success: function (res) {
            uploadAns(res.localId);
        }
    });
}

function uploadAns(id) {
    wx.uploadVoice({
        localId: id, // 需要上传的音频的本地ID，由stopRecord接口获得
        isShowProgressTips: 1, // 默认为1，显示进度提示
        success: function (res) {
             var serverId = res.serverId; // 返回音频的服务器端ID
             //send to server 
             updateAnsVid(serverId);
        }
    });
}


function updateAnsVid(vid) {
    $.ajax({
            url: '/question?qid='+qid+'&rfrom=ajax&vid='+vid,
            type : "POST", 
            dataType: 'json',
            success: function(json) {
                 console.log("get js auth error");
            },
    });
}
