/************************************************************
** @Description: trie
** @Author: george hao
** @Date:   2018-02-07 14:04
** @Last Modified by:  george hao
** @Last Modified time: 2018-02-07 14:04
*************************************************************/
package tria

import (
	"bytes"
	"fmt"
)

type Node struct {
	Letter []byte  //存储字母
	Child  []*Node //子节点
	IsWord bool	   //是否是一个单词
}

type Tria struct {
	Root Node
}

//判断是否是一个单词
func (T *Tria)Search(word string) bool {
	current := &T.Root
	letterChars := []byte(word)
	for i := 0; i < len(letterChars); i++ {
		letterChar := letterChars[i : i+1]
		nodes := current.Child
		index,isword := T.BinarySearch(nodes,letterChar)
		if !isword {
			return false
		}
		current = nodes[index]
	}
	return current.IsWord
}

//删除一个单词，
func (T *Tria)Remove(word string) int {
	current := &T.Root
	letterChars := []byte(word)
	len := len(letterChars)
	for i := 0; i < len; i++ {
		letterChar := letterChars[i : i+1]
		nodes := current.Child
		index , isword := T.BinarySearch(nodes,letterChar)
		if isword && i == (len - 1) {
			current.IsWord = false
			fmt.Println("找到我了",i,string(current.Letter),current.IsWord,isword,index)
			return index
		}
		current = nodes[index]
		fmt.Println(current.Letter)
	}
	return -1
}

//插入一个单词
func (T *Tria)Insert(word string) {
	//获取节点
	currentNode := &T.Root
	letterChars := []byte(word)
	//fmt.Println(  "->", word ,len(letterChars))

	for i := 0; i < len(letterChars); i++ {
		letterChar := letterChars[ i : i+1 ]

		nodes := &currentNode.Child
		index, isWord := T.BinarySearch(*nodes , letterChar)

		//fmt.Println(string(letterChar[0]),i,index,isWord)
		//PrintNode(*nodes)

		if !isWord {
			//fmt.Println(index ,string(letterChar), isWord)
			*nodes = append(*nodes,nil)
			//fmt.Println(*nodes)
			copy((*nodes)[index+1:],(*nodes)[index:])
			//fmt.Println(*nodes)
			(*nodes)[index] = &Node{Letter:letterChar,IsWord:false}
			//fmt.Println(string((*nodes)[0].Letter),index)
		}
		currentNode.Child = *nodes
		currentNode = (*nodes)[index]
		//fmt.Println("-------------")
	}

	currentNode.IsWord = true
}

//二分法
func (T *Tria)BinarySearch(nodes []*Node, letter []byte) (int , bool) {
	start := 0
	end := len(nodes) - 1
	//fmt.Println("end",end)
	if end == -1 {
		return 0 , false
	}

	compareFirst := bytes.Compare(letter , nodes[0].Letter)
	//fmt.Println(string(letter),string(nodes[0].Letter),compareFirst)
	if compareFirst < 0 {
		return 0 , false
	}else if compareFirst == 0 {
		return 0 , true
	}

	compareLast := bytes.Compare(letter , nodes[end].Letter)
	//fmt.Println(string(letter),string(nodes[end].Letter),compareLast)

	if compareLast < 0 {
		return end + 1 , false
	}else if compareLast == 0 {
		return end , true
	}
	current := end / 2
	//fmt.Println("end:",end,start,current)
	if end - start > 1 {
		compareCurrent := bytes.Compare(letter , nodes[current].Letter)

		if compareCurrent < 0 {
			//start = current
			//current = (end + start) / 2
		}else if compareCurrent == 0 {
			end = current
			//current = (end + start) / 2
		}
	}
	return end , true
}



//打印node
func PrintNode(Ns []*Node)  {

	//fmt.Println(Ns)
	if len(Ns) == 0 || Ns == nil {
		fmt.Println("over")
	}
	for index, node := range Ns {
		fmt.Println(index,"--",string(node.Letter),"==",node.IsWord)
		if len(node.Child) != 0 {
			PrintNode(node.Child)
		}
	}
}
