/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 20:28
  */

package crawler

import (
	"log"
	"strconv"
	"weiboCrawler/g"
	"weiboCrawler/utils"
)

func getFollow(username, realID, pageNum string)  {
	// 获取关注list 	https://weibo.cn/{realID}/follow
	body := utils.GetRequest("https://weibo.cn/" + realID + "/follow?page=" + pageNum)
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
		g.FOLLOWS.Lock()
		g.FOLLOWS.Map[username] = append(g.FOLLOWS.Map[username], ret)
		g.FOLLOWS.Unlock()

		g.Tasks <- ret
		pos = pos2 + 1
	}
}

func CrawUserFollow(username, realID string, followNum int) {
	for i := 1; i <= followNum/10+1; i++ {
		if i > 20 {
			break
		}
		pageNum := strconv.Itoa(i)
		getFollow(username, realID, pageNum)
	}
	g.FOLLOWS.RLock()
	log.Println(username, "'s Follow", g.FOLLOWS.Map[username])
	g.FOLLOWS.RUnlock()
	//log.Println(len(g.FOLLOWS[username]))
	return
}

func getFans(username, realID, pageNum string)  {
	// 获取关注list 	https://weibo.cn/{realID}/follow
	body := utils.GetRequest("https://weibo.cn/" + realID + "/fans?page=" + pageNum)
	//log.Println(string(body))
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

		g.FANS.Lock()
		g.FANS.Map[username] = append(g.FANS.Map[username], ret)
		g.FANS.Unlock()
		g.Tasks <- ret

		pos = pos2 + 1
	}
}

func CrawUserFans(username, realID string, fansNum int) {
	for i := 1; i <= fansNum/10+1; i++ {
		if i > 20 {
			break
		}
		pageNum := strconv.Itoa(i)
		getFans(username, realID, pageNum)
	}
	g.FANS.RLock()
	log.Println(username, "'s Follow", g.FANS.Map[username])
	g.FANS.RUnlock()
	//log.Println(len(g.FANS[username]))
	return
}

// 通过 weibo.cn/u/{realId}
// 每个用户只能抓取 200 个粉丝
// 每个用户只能抓取 200 个关注
func CrawUserName(username string)  {
	realID, err := GetRealID(username)
	if err != nil {
		return
	}

	// 根据 bloomFilter 判断是否已经处理
	if g.BloomFilter.Test([]byte(username)) == true {
		log.Println(username, "已经处理")
		return
	}
	g.BloomFilter.Add([]byte(username))

	followBody := utils.GetRequest("https://weibo.cn/" + realID + "/follow")
	pos1 := utils.KMP(followBody, 0, len(followBody)-1, []byte("关注["))
	pos2 := utils.KMP(followBody, pos1, len(followBody)-1, []byte("]<"))
	pos3 := utils.KMP(followBody, pos2, len(followBody)-1, []byte("粉丝["))
	pos4 := utils.KMP(followBody, pos3, len(followBody)-1, []byte("]<"))

	followNum, _ := strconv.Atoi(string(followBody)[pos1+7:pos2])
	fansNum,   _ := strconv.Atoi(string(followBody)[pos3+7:pos4])

	log.Println(username, "关注:", followNum)
	log.Println(username, "粉丝:", fansNum)

	go CrawUserFollow(username, realID, followNum)
	go CrawUserFans(username, realID, followNum)
	go CrawPicture(username)
}

