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
	"sync"
)


// worker 处理传入 tasks 的工作
func worker(tasks chan string, f func(string), wg sync.WaitGroup, worker int, done chan struct{})  {
	defer wg.Done()
	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 通道已经关闭
			log.Printf("Worker: %d : Shutting Down\n", worker)
			done <- struct{}{}
			return
		}
		log.Printf("Work %d Start: %v\n", worker, task)
		f(task)
	}
}


func main() {
	// 读文件
	nameList, err := utils.ReadLineTxt(g.NICKNAMES_FILE)
	if err != nil {
		return
	}
	log.Println("nameList:", nameList)

	//// 对每个写在配置中的微博开启一个 goroutine
	//for _, v := range nameList {
	//	go crawler.CrawUserName(v)
	//}

	// 启动 goroutine 开始工作
	var wg sync.WaitGroup
	wg.Add(g.MaxGoroutines)

	for i := 0; i < g.MaxGoroutines; i++ {
		go worker(g.Tasks, crawler.CrawUserName, wg, i, g.Done)
	}
	wg.Wait()
	log.Println("处理完毕")

	//var end chan struct{}
	//end <- struct{}{}
}
