/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/19 00:45
  */

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<10; i++ {
		fmt.Println(r.Intn(100))
	}
}