# weiboCrawler

一个微博爬虫，出于无聊帮师兄写的  
  
双站点登陆：`m.weibo.cn`    `weibo.cn`

## 进度

| 日期  | 进度  |
|---|---|
| 2017.12.18  | 基本框架搭建、完成用户关注者爬取  |
| 2017.12.19  | 并发模式框架、用户粉丝爬取  |

## 使用方法

修改 `g/g_crawler.go` 中的 `MCNCookie` 和 `CNCookie`   
分别对应 `m.weibo.cn` 的 Cookie 和 `weibo.cn`  的 Cookie   

所有可以改动的配置都在 `g` 目录下，

在 `nicknames.txt` 中，填入你要爬取的用户名，一行一个  
然后会以这些用户作为起点，爬取 Fans & Follows  
再把这些 Fans & Follows 入队，无限循环  

## 存在的问题

1. 当前版本未关闭爬取 goroutine，如果想试用请自己加个 `return`   
2. 当前版本未加入限速，请自己加个 `Sleep()`

## 吐槽

### 为什么使用 go

喜欢  

### 微博用户爬取

已有用户 -> 遍历关注/粉丝 list -> Bloom Filter -> 入队  

### channel

第一版采用 channel 作为消息队列  

优点： 爬虫方面，channel 是一个天生的并发安全、可以传递任何类型、支持阻塞等待的队列  
缺点： 如果程序挂了，channel 中缓冲的数据会丢失  

### html 解析

鉴于 go 没有较好的 html 解析工具，`goqueriy` 也不尽如人意  
采用了最粗暴的方法：KMP 扫描  
希望未来 go 会有像 python 的 lxml beautifulsoup 这些东西  


