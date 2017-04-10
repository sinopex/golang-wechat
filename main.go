package main

import (
	"fmt"
	"github.com/sidbusy/weixinmp"
	"github.com/sinopex/golang/wechat/toutiao"
	"log"
	"net/http"
	"os"
)

var APP_ID = os.Getenv("WECHAT_APP_KEY")
var APP_SECRET = os.Getenv("WECHAT_APP_SECRET")
var TOKEN = os.Getenv("WECHAT_TOKEN")

var NewsCount = 5

func main() {
	// 注册处理函数
	// 注册处理函数
	http.HandleFunc("/", receiver)
	log.Fatal(http.ListenAndServe(":8002", nil))
}

func receiver(w http.ResponseWriter, r *http.Request) {
	// 仅被动响应消息时可不填写appid、secret
	// 仅主动发送消息时可不填写token
	mp := weixinmp.New(TOKEN, APP_ID, APP_SECRET)
	// 检查请求是否有效
	// 仅主动发送消息时不用检查
	if !mp.Request.IsValid(w, r) {
		return
	}
	// 判断消息类型
	if mp.Request.MsgType == weixinmp.MsgTypeText {
		founds := toutiao.GetArticles(mp.Request.Content)
		if len(founds) > 0 {
			articles := make([]weixinmp.Article, 0)
			for i, found := range founds {
				if i > NewsCount {
					break
				}

				articles = append(articles, weixinmp.Article{
					Title:       found.Title,
					Description: found.Summary,
					PicUrl:      "http://ww3.sinaimg.cn/large/7cc829d3gw1ezft1ollblj20p00dwtab.jpg",
					Url:         found.Link,
				})
			}

			err := mp.ReplyNewsMsg(w, &articles)
			if err != nil {
				fmt.Println("error", err)
			}
		} else {
			// 回复消息
			mp.ReplyTextMsg(w, "暂无["+mp.Request.Content+"]对应的文章")
		}
	}
}
