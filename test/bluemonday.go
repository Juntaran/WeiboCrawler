/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 22:51
  */

package main

import (
	"github.com/microcosm-cc/bluemonday"
	"fmt"
)

func main() {
	p := bluemonday.UGCPolicy()
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
}