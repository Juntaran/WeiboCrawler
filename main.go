/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 18:23
  */

package main

import (
	"weiboCrawler/utils"
	"weiboCrawler/g"
	"weiboCrawler/crawler"
	"log"
)


func main() {
	// 读文件
	nameList, err := utils.ReadLineTxt(g.NICKNAMES_FILE)
	if err != nil {
		return
	}
	log.Println("nameList:", nameList)

	var end chan struct{}

	// 对每个写在配置中的微博开启一个 goroutine
	for _, v := range nameList {
		go crawler.CrawUserName(v)
	}

	end <- struct{}{}
}
