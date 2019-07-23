// FP-Growth 频繁相集挖掘算法
package core

import (
	"sort"
	"strconv"
)

//项头表建立
//返回项头表及词地址哈希表(快速取值)
func BuildHeadElems(wordCount map[string]int, supportCount int) (HeadElems, map[string]*HeadElem) {
	var headElems HeadElems
	var headAddr = make(map[string]*HeadElem)
	for word, count := range wordCount {
		if count >= supportCount {
			elem := HeadElem{word: word, count: count, treeNode: nil, pattern: nil}
			headAddr[word] = &elem
			headElems = append(headElems, elem)
		}
	}
	// 项头表支持度降序排序
	sort.Sort(sort.Reverse(headElems))
	return headElems, headAddr
}

// 根据项头表过滤WordBase
func FilterWordBase(headAddr map[string]*HeadElem, wordBase [][]string) [][] string {
	var filteredWordBase [][] string
	for _, words := range wordBase {
		var filteredWords [] string
		var wordSupport Pairs
		for _, word := range words {
			if headAddr[word] != nil {
				wordSupport = append(wordSupport, Pair{word, headAddr[word].count})
			}
		}
		// 按支持度降序排序,此时与项头表采用相同的排序方法，保证二者相对顺序一致
		sort.Sort(sort.Reverse(wordSupport))
		for _, pair := range wordSupport {
			filteredWords = append(filteredWords, pair.key)
		}
		if len(filteredWords) > 0 {
			filteredWordBase = append(filteredWordBase, filteredWords)
		}
	}
	return filteredWordBase
}

// 二频繁项集合
type TwoFreqItem struct {
	BaseWord     string
	Word         string
	SupportCount int
	Confidence   float64
}

// 输出所有二频繁项
func WordConcurrence(headAddr map[string]*HeadElem, confidence float64) [] TwoFreqItem {
	var freqItems [] TwoFreqItem
	for baseWord, headElem := range headAddr {
		for coWord, supportCount := range headElem.pattern {
			con := float64(supportCount) / float64(headElem.count)
			if con >= confidence {
				freqItems = append(freqItems, TwoFreqItem{baseWord, coWord, supportCount, con})
			}
			con = float64(supportCount)/float64(headAddr[coWord].count)
			if con >= confidence {
				freqItems=append(freqItems,TwoFreqItem{coWord,baseWord,supportCount,con})
			}
		}
	}

	return freqItems
}


func FreqItemsToStrings(items [] TwoFreqItem) [][] string {
	var ret [] [] string
	for _, item := range items {
		var value [] string
		value = append(value, item.BaseWord)
		value = append(value, item.Word)
		value = append(value, strconv.Itoa(item.SupportCount))
		value = append(value, strconv.FormatFloat(item.Confidence, 'f', -1, 64))
		ret = append(ret, value)
	}
	return ret
}