/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/19 00:46
  */

package utils

import (
	"math/rand"
	"time"
	"WeiboCrawler/g"
)

func RandHeaders() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(len(g.HEADERS))
}