# weiboCrawler

*weibo crawler implemented in Go*

一个微博爬虫，出于无聊帮师兄写的  

双站点登陆：`m.weibo.cn`    `weibo.cn`  

## 微博用户爬取

已有用户 -> 遍历关注/粉丝 list -> Bloom Filter -> 入队  

## channel

第一版采用 channel 作为消息队列  

优点： 爬虫方面，channel 是一个天生的并发安全、可以传递任何类型、支持阻塞等待的队列  
缺点： 如果程序挂了，channel 中缓冲的数据会丢失  


## 进度

| 日期  | 进度  |
|---|---|
| 2017.12.19  | 基本框架搭建、完成用户关注者爬取  |

## 使用方法

修改 `g/g_crawler.go` 中的 `MCNCookie` 和 `CNCookie`  
分别对应 `m.weibo.cn` 的 Cookie 和 `weibo.cn`  的 Cookie  

在 `nicknames.txt` 中，填入你要爬取的用户名，一行一个