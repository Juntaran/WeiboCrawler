/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/21 11:41
  */

package crawler

import (
	"WeiboCrawler/utils"
	"log"
	"github.com/tidwall/gjson"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"errors"
	"strconv"
)

func getPictureUrl(cid, page string) ([]string, error) {
	// url = "https://m.weibo.cn/api/container/getIndex?count=25&containerid={}&page={}".format(containerid, page)
	// url = "https://m.weibo.cn/api/container/getIndex?type=uid&value={}&containerid={}&page={}".format(cid, "107603" + cid, page)
	url := "https://m.weibo.cn/api/container/getIndex?type=uid&value=" + cid + "&containerid=107603" + cid + "&page=" + page
	body := utils.GetRequest(url)

	if len(gjson.Get(string(body), "data.cards").Array()) == 0 {
		return nil, errors.New("Error: Page Done")
	}
	value := gjson.Get(string(body), "data.cards.#.mblog.retweeted_status.pics.#.large.url")
	var photos []string
	for _, v := range value.Array() {
		//log.Println(v.Type)
		//log.Println(v.String())
		for _, i := range v.Array() {
			//log.Println(i.String())
			photos = append(photos, i.String())
		}
	}
	return photos, nil
}

func saveImage(nickname, url string)  {
	_, err := os.Stat("WeiboAlbum/" + nickname)
	if err != nil {
		os.MkdirAll("WeiboAlbum/" + nickname, 0777)
	}
	imageSavePath := "WeiboAlbum/" + nickname + "/" + strings.Split(url, "/")[len(strings.Split(url, "/"))-1]
	_, err = os.Stat(imageSavePath)
	if err == nil {
		log.Println("File Already Exists:", imageSavePath)
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read data failed:", url, err)
		return
	}
	image, err := os.Create(imageSavePath)
	if err != nil {
		log.Println("create file failed:", imageSavePath, err)
		return
	}

	defer image.Close()
	image.Write(data)

}

func CrawPicture(nickname string)  {
	cid, err := GetRealID(nickname)
	if err != nil {
		return
	}
	var allImages []string
	for page := 1; ; page++ {
		urls, err := getPictureUrl(cid, strconv.Itoa(page))
		if err != nil {
			break
		}
		allImages = utils.MergeSlice(allImages, urls)
	}
	count := len(allImages)
	log.Println("Get", nickname, "Pictures:", count)

	for k, v := range allImages {
		log.Printf("%v %v/%v", nickname, k+1, count)
		saveImage(nickname, v)
	}
	return
}