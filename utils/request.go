/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/19 16:14
  */

package utils

import (
	"io/ioutil"
	"net/http"
	"log"
	"WeiboCrawler/g"
)

func GetRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest Error:", err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("cookie", g.CNCookie)
	req.Header.Set("DNT", "1")
	req.Header.Set("Host", "weibo.cn")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", g.HEADERS[RandHeaders()])

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("DefaultClient Error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println("body:", string(body))
	defer resp.Body.Close()

	return body
}
