/************************************************************
** @Description: tria
** @Author: george hao
** @Date:   2018-02-22 10:21
** @Last Modified by:  george hao
** @Last Modified time: 2018-02-22 10:21
*************************************************************/
package double_trie

import (
	"fmt"
	"testing"
)

var (
	cd    *Dat
	words = []string{
		"a", "aa", "ab", "ac", "abc", "abd",
		"abcd", "abde", "abdf", "abcdef", "abcde",
		"abcdefghijklmn", "bcd", "b", "xyz",
		"中国", "中国北京", "中国上海", "中国广州",
		"中华", "中华文明", "中华民族", "中华人民共和国",
		"this", "this is", "this is a sentence.", "郝大全",
	}
)

func TestBasic(t *testing.T) {
	loadTestData()
	// check the consistency
	checkConsistency(cd)
}

func checkConsistency(cd *Dat) {
	for i, word := range words {
		id, err := cd.Jump([]byte(word), 0)
		if i%4 == 0 {
			if err == ErrNoPath {
				continue
			}
			_, err := cd.Value(id)
			if err == ErrNoValue {
				continue
			}
			panic("not deleted")
		}
		key, err := cd.Key(id)
		if err != nil {
			panic(err)
		}
		if string(key) != word {
			panic("key error")
		}
		value, err := cd.Value(id)
		if err != nil || value != i {
			fmt.Println(word, i, value, err)
			panic("value error")
		}
	}
}

func loadTestData() {
	if cd != nil {
		return
	}
	cd = New()
	// cd.Ordered = false

	// add the keys
	for i, word := range words {
		//fmt.Println(word)
		if err := cd.Insert([]byte(word), i); err != nil {
			panic(err)
		}
	}

	PrintDt(*cd)

	for _, word := range words {
		if err := cd.Delete([]byte(word)); err != nil {
			panic(err)
		}
	}

	for i, word := range words {
		if err := cd.Update([]byte(word), i); err != nil {
			panic(err)
		}
	}

	// delete some keys
	for i := 0; i < len(words); i += 4 {
		if err := cd.Delete([]byte(words[i])); err != nil {
			panic(err)
		}
	}
	return
}

func PrintDt(cd Dat) {

	for i, v := range cd.Array {
		fmt.Println(i, v.Check, v.Value)
	}

	for i, v := range cd.Ninfos {
		fmt.Println(i, v.Sibling, v.Child)
	}

}
