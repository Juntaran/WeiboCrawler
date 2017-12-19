/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/12/18 18:24
 */

package g

const NICKNAMES_FILE = "nicknames.txt"

var HEADERS = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 5.1; rv:7.0.1) Gecko/20100101 Firefox/7.0.1",
	"Mozilla/5.0 (X11; ; Linux x86_64; rv:1.8.1.6) Gecko/20070802 Firefox; ; Linux x86_64; rv:1.8.1.6) Gecko/20070802 Firefox",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12.0; rv:48.0.2) Gecko/20100101 Firefox/to 48.0.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.109 Safari/537.36",
	"Mozilla/5.0 (Unknown; Linux) AppleWebKit/538.1 (KHTML, like Gecko) Chrome/v1.0.0 Safari/538.1",
}

var FOLLOWS = make(map[string][]string)
var FANS    = make(map[string][]string)

type User struct {
	USERNAME string
	FANS     map[string][]string
	FOLLOWS  map[string][]string
}

var MCNCookie = ""
var CNCookie  = ""
