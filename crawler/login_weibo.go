/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 18:47
  */

package crawler

import (
	"io/ioutil"
	"net/http"
	"log"
	"WeiboCrawler/g"
	"WeiboCrawler/utils"
)

func LoginMWeiboCN()  {
	req, err := http.NewRequest("GET", "https://m.weibo.cn", nil)
	if err != nil {
		log.Println("NewRequest Error:", err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("cookie", g.MCNCookie)
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "m.weibo.cn")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", g.HEADERS[utils.RandHeaders()])

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("DefaultClient Error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("body:", string(body))
	defer resp.Body.Close()
}

func LoginWeiboCN()  {
	req, err := http.NewRequest("GET", "https://weibo.cn", nil)
	if err != nil {
		log.Println("NewRequest Error:", err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("cookie", g.CNCookie)
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "weibo.cn")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", g.HEADERS[utils.RandHeaders()])

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("DefaultClient Error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("body:", string(body))
	defer resp.Body.Close()
}