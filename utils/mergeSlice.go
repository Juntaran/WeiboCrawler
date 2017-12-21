/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/12/21 16:32
  */

package utils

func MergeSlice(s ...[]string) (slice []string)  {
	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := MergeSlice(s[1:]...)		//...将数组元素打散
		slice = make([]string, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}
	return
}