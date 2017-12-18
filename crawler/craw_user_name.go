/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 20:28
  */

package crawler

import (
	"io/ioutil"
	"net/http"
	"log"
	"weiboCrawler/g"
	"weiboCrawler/utils"
	"strconv"
)

func getFollow(username, realID, pageNum string) []string {
	// 获取关注list 	https://weibo.cn/{realID}/follow
	req, err := http.NewRequest("GET", "https://weibo.cn/" + realID + "/follow?page=" + pageNum, nil)
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
	//log.Println("body:", string(body))
	defer resp.Body.Close()

	pos := 0

	for {
		pos1 := utils.KMP(body, pos, len(body)-1, []byte("</a></td><td valign=\"top\"><a href="))
		if pos1 == -1 {
			break
		}
		pos2 := utils.KMP(body, pos1+35, len(body)-1, []byte("\">"))
		//log.Println(pos1, pos2)
		pos3 := utils.KMP(body, pos2+2, len(body)-1, []byte("</a"))
		ret := string(body)[pos2+2:pos3]
		//log.Println(ret)
		g.FOLLOWS[username] = append(g.FOLLOWS[username], ret)
		pos = pos2 + 1
	}

	return nil
}

// 通过 weibo.cn/u/{realId}
// 每个用户只能抓取 200 个粉丝
// 每个用户只能抓取 200 个关注
func CrawUserName(username string)  {
	realID, err := GetRealID(username)
	if err != nil {
		return
	}
	req, err := http.NewRequest("GET", "https://weibo.cn/" + realID + "/follow", nil)
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
	//log.Println("body:", string(body))
	defer resp.Body.Close()

	ret1 := utils.KMP(body, 0, len(body)-1, []byte("关注["))
	ret2 := utils.KMP(body, ret1, len(body)-1, []byte("]<"))
	log.Println(username, "关注了:", string(body)[ret1+7:ret2], "个人")

	followNum, _ := strconv.Atoi(string(body)[ret1+7:ret2])
	for i := 1; i <= followNum/10+1; i++ {
		if i > 20 {
			break
		}
		pageNum := strconv.Itoa(i)
		getFollow(username, realID, pageNum)
	}

	log.Println(g.FOLLOWS)
	log.Println(len(g.FOLLOWS[username]))
}

