
var current_timestamp = null;
    function load() {
        $.ajax({
             url: '/newest_list?rfrom=ajax&ts='+encodeURIComponent(Math.floor(current_timestamp.getTime() / 1000)),
             dataType: 'json',
             success: function(json) {
                 if (json == null || json == "undefined") {
                     console.log("NO more data.");
                     return;
                  }
                  json = json.list;
                  var el = $('#content');
                  for (var i = 0; i < json.length; i++) {
                      var str ="";
                      str += '<div class="topic"> '; 
                      str += '    <div class="topic_tit">';
                      str += '       <a href="/question?qid='+json[i].id+'" target="_self" >'+ json[i].title +'</a>';
                      str += '    </div>';
                      str += '    <div class="topic_detail">';
                      str += '       <div> <span class="topic_detail_author">'+json[i].creator.Name+'|'+json[i].creator.Title+'</div>';
                      str += '       <div class="topic_detail_left">';
             	      str += '       <a href="/inquiry?anu='+json[i].creator.Uid+'" target="_self" ><img class="topic_detail_photo" src="images/3_10.jpg" align="middle"></a>';
                      str += '       <a href="/question?qid='+json[i].id+'" ><span class="topic_detail_words"><p><img class="topic_detail_bg_voice" align="top" src="images/icon_voice_13.png">1元偷偷看</p></span></a>';
                      str += '    </div>';
       		      str += '    <div class="topic_detail_right"> 看过 '+json[i].count+' </div>  ';
                      str += '    <div class="clear"></div>';
                      str += '   </div> ';
	              str += '</div>'
                      el.append(str);
                  }
                  current_timestamp = new Date(json[json.length -1].date);
               }
           });

    }
