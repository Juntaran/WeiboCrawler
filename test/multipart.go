/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 20:53
  */

package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"weiboCrawler/g"
)

// Creates a new file upload http request with optional extra params
func newMultipartRequest(url string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()
	return http.NewRequest("POST", url, body)
}

func main() {

	extraParams := map[string]string{
		"keyword":      "硬卧客",
		"suser":      	"2",
	}
	request, err := newMultipartRequest("https://weibo.cn/find/user", extraParams)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cookie", "_T_WM=d5096b743fbcd884471ef1c1a4fcb047; SUB=_2A253M-vSDeRhGeVM7VUW8yzPzj2IHXVU3_WarDV6PUJbkdANLUzCkW1NHetkTxY_OE71CO77aMUSfNuTIHoMFovD; SUHB=0G0etrtEr1ZHeO; SCF=ApEhSu57Cv7Mepw7kAceuu3GtWgU4RH7AuxnARuEGdSBTR5xEcCZPRazKJjD-rBxYlw4mfheDAFtFtKdxO8KS7U.; SSOLoginState=1513593730; WEIBOCN_FROM=1110006030")
	request.Header.Set("DNT", "1")
	request.Header.Set("Origin", "https://weibo.cn")
	request.Header.Set("Referer", "https://weibo.cn/find/user")
	request.Header.Set("Host", "weibo.cn")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("User-Agent", g.HEADERS[0])
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}
}