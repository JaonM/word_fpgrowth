// FP-Growth 频繁相集挖掘算法
package core

import (
	"math"
	"sort"
	"sync"
	"word_concurrence/nlp"
)

//项头表建立
//返回项头表及词地址哈希表(快速取值)
func BuildHeadElems(wordCount map[string]int, supportCount int) (HeadElems, map[string]*HeadElem) {
	var headElems HeadElems
	var headAddr = make(map[string]*HeadElem)
	for word, count := range wordCount {
		if count > supportCount {
			elem := HeadElem{word: word, count: count, treeNode: nil, pattern: nil}
			headAddr[word] = &elem
			headElems = append(headElems, elem)
		}
	}
	return headElems, headAddr
}

func searchHeadElem(word string, headAddr map[string]*HeadElem) *HeadElem {
	return headAddr[word]
}

// 根据项头表过滤WordBase
func FilterWordBase(headAddr map[string]*HeadElem, wordBase [][]string) []map[string]int {
	var filteredWordBase []map[string]int
	for _, words := range wordBase {
		var filteredWords = make(map[string]int)
		for _, word := range words {
			if searchHeadElem(word, headAddr) != nil {
				//filteredWords = append(filteredWords, word)
				filteredWords[word] += 1
			}
		}
		filteredWordBase = append(filteredWordBase, filteredWords)
	}
	return filteredWordBase
}

//建立FPNode
func BuildFPTree(root *FPNode, wordBase []map[string]int, headAddr map[string]*HeadElem) {
	for _, wordCount := range wordBase {
		var nodes []*FPNode
		//Sort map by key
		var keys [] string
		for k := range wordCount {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			nodes = append(nodes, &FPNode{key, wordCount[key], nil, nil, nil, nil})
		}
		insertNodeToTree(nodes, root, headAddr)
	}
}

func insertNodeToTree(nodes []*FPNode, root *FPNode, headAddr map[string]*HeadElem) {
	p := root
	for _, node := range nodes {
		if p.child == nil {
			p.Insert(node)
			insertHeadElemTreeNode(node, headAddr)
			p = p.child
		} else {
			currentNode := p.child
			for currentNode != nil {
				if currentNode.word == node.word {
					currentNode.count += node.count
					p = currentNode
					break
				}
				currentNode = currentNode.neighbor
			}
			if currentNode == nil {
				p.Insert(node)
				insertHeadElemTreeNode(node, headAddr)
			}
		}
	}
}

func insertHeadElemTreeNode(node *FPNode, headAddr map[string]*HeadElem) {
	tmp := headAddr[node.word].treeNode
	if tmp == nil {
		headAddr[node.word].treeNode = node
	} else {
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = node
	}
}

//获取条件模式基
func ConditionalPattern(root *FPNode, headElems HeadElems, supportCount int, headAddr map[string]*HeadElem) {
	//升序排序
	sort.Sort(headElems)
	var wg sync.WaitGroup
	for _, headElem := range headElems {
		elemCopy := headAddr[headElem.word]
		wg.Add(1)
		go func(headElem *HeadElem) {
			defer wg.Done()
			node := headElem.treeNode
			pattern := make(map[string]int)
			for node != nil {
				p := node
				count := p.count
				for p.parent != root {
					tempCount := pattern[p.parent.word] + count
					if tempCount > supportCount {
						pattern[p.parent.word] = tempCount
					}
					p = p.parent
				}
				node = node.next
			}
			headElem.pattern = pattern
		}(elemCopy)
	}
	wg.Wait()
}

func WordConcurrence(word string, headAddr map[string]*HeadElem, confidence float64) [] string {
	var candidates [] string
	for w := range headAddr {
		candidates = append(candidates, w)
	}
	if !nlp.SearchWord(word, candidates) {
		return [] string{}
	}
	pattern := headAddr[word].pattern

	if len(pattern) == 0 {
		return [] string{}
	}
	var coWords [] string
	for word, count := range pattern {
		c1 := math.Min(float64(count), float64(headAddr[word].count))
		confi := c1 / float64(headAddr[word].count)
		if confi > confidence {
			coWords = append(coWords, word)
		}
	}
	return coWords
}
