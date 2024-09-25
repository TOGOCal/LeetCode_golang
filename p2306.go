package main

import (
	"fmt"
)

/*
给你一个字符串数组 ideas 表示在公司命名过程中使用的名字列表。公司命名流程如下：

从 ideas 中选择 2 个 不同 名字，称为 ideaA 和 ideaB 。
交换 ideaA 和 ideaB 的首字母。
如果得到的两个新名字 都 不在 ideas 中，那么 ideaA ideaB（串联 ideaA 和 ideaB ，中间用一个空格分隔）是一个有效的公司名字。
否则，不是一个有效的名字。
返回 不同 且有效的公司名字的数目。
*/

func main() {

	strs := []string{"coffee", "donuts", "time", "toffee"}

	fmt.Println(distinctNames(strs))
}

func distinctNames(ideas []string) (ans int64) {
	satisfy := [26][26]int{} // satisfy[i][j] 以i开头能接受j的后缀

	// 以后缀String为key，查看每个String有哪些首字母
	group := make(map[string][]int) // 统计各后缀中以x开头的数目
	for _, s := range ideas {
		//检查这个后缀String是否存在
		a, ok := group[s[1:]]
		if !ok { //如果不存在
			a = make([]int, 26) // 创建一个数组，长度为26
			group[s[1:]] = a    //放入map中
		}
		a[s[0]-'a']++ //有多少个这样的首字母
	}

	for _, valueIntArray := range group { // 遍历后缀String map

		for i, count := range valueIntArray { //遍历后缀String对应的所有首字母
			if count == 0 {
				continue
			} // 如果不存在以当前字母value_int_arr'+i开头的后缀

			for possibleInitialCharacter := range valueIntArray {
				if valueIntArray[possibleInitialCharacter] == 0 { // 以value_int_arr'+i开头，但是可以接受value_int_arr'+j的数量
					satisfy[i][possibleInitialCharacter] += count
				}
			}
		}
	}

	for i := range satisfy {
		for j := range satisfy[i] {
			ans += int64(satisfy[i][j]) * int64(satisfy[j][i])
		}
	}
	return ans

}

//func distinctNames(ideas []string) int64 {
//
//	var ans int64
//
//	charMap := make(map[uint8]map[string]bool)
//
//	for _, idea := range ideas {
//
//		_, ok := charMap[idea[0]]
//
//		//没有创建过这个集合
//		if !ok {
//
//			charMap[idea[0]] = make(map[string]bool)
//		}
//
//		m := charMap[idea[0]]
//
//		//subString(1, len(idea)) 获取字符串的子串
//		newString := idea[1:]
//		m[newString] = true
//	} //完成加入
//
//	//开始遍历
//	for i := range charMap {
//
//		for j := range charMap {
//
//			sameElement := same(charMap[i], charMap[j])
//
//			a := len(charMap[i]) - sameElement
//			b := len(charMap[j]) - sameElement
//
//			ans += int64(a * b)
//		}
//	}
//
//	return ans
//}

func same(m1, m2 map[string]bool) int {

	a := 0

	for key := range m1 {

		if _, ok := m2[key]; ok {
			a++
		}
	}

	return a
}

func difference(m1, m2 map[string]bool) int {

	a := 0

	for key := range m1 {

		if _, ok := m2[key]; ok {
			continue
		} else {

			a++
		}
	}

	return a
}
