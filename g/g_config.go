/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/19 17:38
  */

package g


const MaxGoroutines = 2			// 最多同时处理用户数，每个用户处理的时候会额外开启两个 goroutine 用来同时获取 follow fans
const TaskLoad      = 16

var Tasks = make(chan string, TaskLoad)
var Done  = make(chan struct{}, MaxGoroutines)