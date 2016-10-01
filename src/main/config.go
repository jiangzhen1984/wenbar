

package main

import (
    "gotom"
    "main/handlers"
    "main/service/wechat"
)


var  WechatConf * wechat.WeChatConfig = &wechat.WeChatConfig {
   AppId       : "wxc95bca05a3e20730",
   Secret      : "9b6ad04958e230c395257fed5784731a",
   ResponseURL : "http://www.wenbar.com/wechat", 
}

var conf *gotom.GTConfig = &gotom.GTConfig { 
     DebugMode      : true,
     Port           : ":8080", 
     Tpldir         : "./view/",
     Mapping        : []*gotom.Mapping {
                           {"/hot_list",      handlers.HotListHandler},
                           {"/newest_list",   handlers.NewestListHandler},
                           {"/login",         handlers.LoginHandler},
                           {"/personal",      handlers.PersonalHandler},
                           {"/logout",        handlers.LogoutHandler},
                           {"/my_inquiry",    handlers.MyInquiryHandler},
                           {"/my_viewed",     handlers.MyViewedHandler},
                           {"/my_answer",     handlers.MyAnsHandler},
                           {"/question",      handlers.QuestionDetailHandler},
                           {"/inquiry",       handlers.InquiryHandler},
                           {"/wechat",        handlers.WeChatHandler},
                      },
     TplMapping     : map[string]*gotom.GTTemplateMapping {
                           "/hot_list"    : {
                                             Uri  : "/hot_list",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "hot_list" : {Name : "hot_list" , Path : "view/hot_list.html"},
                                                    },
                                            }  ,
                           "/newest_list" : {
                                             Uri  : "/newest_list",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "newest_list" : {Name : "newest_list" , Path : "view/newest_list.html"},
                                                    },
                                            }  ,
                           "/login"       : {
                                             Uri  : "/login",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "login" : {Name : "login"          , Path : "view/login.html"},
                                                    },
                                            }  ,
                           "/personal"    : {
                                             Uri  : "/personal",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "personal" : {Name : "login"       , Path : "view/personal.html"},
                                                    },
                                            }  ,
                           "/my_inquiry"  : {
                                             Uri  : "/my_inquiry",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "myinquiry" : {Name : "myinquiry"  , Path : "view/my_inquiry.html"},
                                                    },
                                            }  ,
                           "/my_answer"  : {
                                             Uri  : "/my_answer",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "myanswer" : {Name : "myanswer"  , Path : "view/my_ans.html"},
                                                    },
                                            }  ,
                           "/my_viewed"  : {
                                             Uri  : "/my_viewed",
                                             Tpls : map[string]*gotom.GTTemplate{
                                                        "myviewed" : {Name : "myviewed"  , Path : "view/my_viewed.html"},
                                                    },
                                            }  ,
                           "/question"    : {
                                                Uri  : "/question",
                                                Tpls : map[string]*gotom.GTTemplate{
                                                        "questiondetail" : {Name : "questiondetail"  , Path : "view/question_detail.html"},
                                                        },
                                            }  ,
                           "/inquiry"    : {
                                                Uri  : "/inquiry",
                                                Tpls : map[string]*gotom.GTTemplate{
                                                        "inquiry" : {Name : "inquiry"  , Path : "view/inquire.html"},
                                                        },
                                            }  ,
                      },
}

