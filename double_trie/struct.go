/************************************************************
** @Description: double_trie
** @Author: george hao
** @Date:   2018-03-21 15:12
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-21 15:12
*************************************************************/
package double_trie

// struct
type Dat struct {
	*dat
}

type dat struct {
	Array    []node
	Ninfos   []ninfo
	Blocks   []block
	Reject   [257]int
	BheadF   int
	BheadC   int
	BheadO   int
	Capacity int
	Size     int
	Ordered  bool
	MaxTrial int
}

type node struct {
	Value int
	Check int
}

type ninfo struct {
	Sibling, Child byte
}

type block struct {
	Prev, Next, Num, Reject, Trial, Ehead int
}

func (b *block) init() {
	b.Num = 256
	b.Reject = 257
}
