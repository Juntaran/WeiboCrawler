/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/18 18:14
  */

package utils

import (
	"os"
	"log"
	"bufio"
	"strings"
	"io"
)

func ReadLineTxt(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	var nameList []string
	if err != nil {
		log.Println("Open File Error:", err)
		return nil, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			nameList = append(nameList, line)
		}
		if err != nil {
			if err == io.EOF {
				log.Println("Read File Finish")
				return nameList, nil
			}
			log.Println("Read File Error:", err)
			return nil, err
		}
	}
	return nil, err
}

