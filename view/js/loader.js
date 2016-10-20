

var PageObject = function(uri, contentId, ts, par) {
      this.uri = uri;
      this.contentId = contentId;
      this.timestamp = ts;
      this.parameters = par;
};


PageObject.prototype.registerScroll = function() {
    $(window).scroll(function() {
        var wd = $(window);
        if(wd.scrollTop() + wd.height() >= $(document).height()) {
            this.load();
        }
    });
};

PageObject.prototype.load = function() {
  var po = this;
  $.ajax({
     url: this.uri + '?rfrom=ajax&ts='+encodeURIComponent(Math.floor(this.timestamp.getTime() / 1000)),
     dataType: 'json',
     success: function(json) {
         if (json == null || json == "undefined") {
             console.log("NO more data.");
             return;
         }
         json = json.list;
         var el = $('#' + po.contentId);
         for (var i = 0; i < json.length; i++) {
             el.append(po.getElement(json[i]));
         }

         if (json.length <= 0) {
             po.timestamp = new Date();
         } else {
             po.timestamp = new Date(json[json.length -1].date);
         }
       }
   });

}

PageObject.prototype.getElement=function(topic) {
    var str  ="";
        str += '<div class="topic"> '; 
        str += '    <div class="topic_tit">';
        str += '       <a href="/question?qid='+topic.id+'" target="_self" >'+ topic.title +'</a>';
        str += '    </div>';
        str += '    <div class="topic_detail">';
        str += '       <div> <span class="topic_detail_author">'+(topic.creator == undefined ? "" : topic.creator.Name+'|'+topic.creator.Title)+'</div>';
        str += '       <div class="topic_detail_left">';
        str += '       <a href="/inquiry?anu='+(topic.creator == undefined ? "" :topic.creator.Uid)+'" target="_self" ><img class="topic_detail_photo" src="images/3_10.jpg" align="middle"></a>';
        str += '       <span class="topic_detail_words" onclick="onAnsClicked(\''+topic.AudioUrl+'\')" ><p><img class="topic_detail_bg_voice" align="top" src="images/icon_voice_13.png">1元偷偷看</p></span>';
        str += '    </div>';
        str += '    <div class="topic_detail_right"> 看过 '+topic.count+' </div>  ';
        str += '    <div class="clear"></div>';
        str += '   </div> ';
        str += '</div>'
    return str;
};


PageObject.prototype.searchTopic = function (text) {
    var po = this;
    $.ajax({
       url: this.uri +'?rfrom=ajax&type=ts&text='+text +'&' + this.par,
       dataType: 'json',
       success: function(json) {
             json = json.list;
             var el = $('#'+po.contentId);
             el.empty();
             for (var i = 0; i < json.length; i++) {
                 el.append(po.getElement(json[i]));
             }
       }
    
    });
};
