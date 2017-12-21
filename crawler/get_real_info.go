/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 21:28
  */

package crawler

import (
	"io/ioutil"
	"strings"
	"net/http"
	"log"
	"errors"
	"WeiboCrawler/g"
	"WeiboCrawler/utils"
)

// 根据用户名返回真实的 userID
func GetRealID(username string) (string, error) {
	// 访问 m.weibo.com/n/{username}
	req, err := http.NewRequest("GET", "https://m.weibo.com/n/" + username, nil)
	if err != nil {
		log.Println("NewRequest Error:", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("DefaultClient Error:", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	
	//log.Println("body:", string(body))
	defer resp.Body.Close()

	if utils.KMP(body, 0, len(body)-1, []byte("用户不存在")) != -1 {
		log.Println("Error: 用户不存在")
		err := errors.New("Error: 用户不存在")
		return "", err
	}

	// 1005053267732351 -> 100505 + real_id
	realID := strings.Split(resp.Request.URL.Path, "/")[2][6:]
	//log.Println("RealID:", realID)

	return realID, nil
}


// 根据链接返回用户名
func GetUserName(userLink string) string {
	// 访问 weibo.cn/xxx
	req, err := http.NewRequest("GET", userLink, nil)
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

	sub := "</a></td><td valign=\"top\"><div class=\"ut\"><span class=\"ctt\">"
	pos1 := utils.KMP(body, 0, len(body)-1, []byte(sub))
	pos2 := utils.KMP(body, pos1+len(sub)+1, len(body)-1, []byte("<img"))
	pos3 := utils.KMP(body, pos1+len(sub)+1, len(body)-1, []byte("<a"))
	pos4 := utils.KMP(body, pos1+len(sub)+1, len(body)-1, []byte("&nbsp"))
	if pos4 < pos3 && pos4 != -1 {
		pos3 = pos4
	}
	if pos3 < pos2 && pos3 != -1 {
		pos2 = pos3
	}

	log.Println(string(body)[pos1+len(sub):pos2])

	userName := string(body)[pos1+len(sub):pos2]
	log.Println("userName:", userName)

	return userName
}
