/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/19 17:38
  */

package g

import "github.com/willf/bloom"

const MaxGoroutines = 2			// 最多同时处理用户数，每个用户处理的时候会额外开启两个 goroutine 用来同时获取 follow fans
const TaskLoad      = 10000		// 缓冲 channel 的容量

const BFN = uint(10000)								// BloomFilter 容量
var BloomFilter = bloom.New(20 * BFN, 5)			// 20*BFN 为总位数、5 为使用的 hash 函数个数

var Tasks = make(chan string, TaskLoad)
var Done  = make(chan struct{}, MaxGoroutines)