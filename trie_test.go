/************************************************************
** @Description: tria
** @Author: george hao
** @Date:   2018-02-08 09:32
** @Last Modified by:  george hao
** @Last Modified time: 2018-02-08 09:32
*************************************************************/
package tria

import (
	"testing"
	"fmt"
)

func TestTria_Insert(t *testing.T) {
	tree := new(Tria)

	//新增
	tree.Insert("h")
	tree.Insert("http")
	tree.Insert("https")
	tree.Insert("httds")
	tree.Insert("htta")
	tree.Insert("httb")
	//tree.Insert("https://www.haodaquan.com")
	//tree.Insert("a")
	//tree.Insert("an")
	//tree.Insert("a computer")
	//tree.Insert("georgehao")

	//查找
	isword := tree.Search("httpsan")

	fmt.Println("1================")
	fmt.Println(isword)
	PrintNode(tree.Root.Child)

	//删除
	index := tree.Remove("httb")
	fmt.Println("2================")
	fmt.Println(index)
	PrintNode(tree.Root.Child)
}