package core

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"word_concurrence/nlp"
	"word_concurrence/utils"
)

func TestBuildHeadElems(t *testing.T) {
	wordCount := map[string]int{"apple": 2, "banana": 3, "watermelon": 1}
	headSlice, _ := BuildHeadElems(wordCount, 2)
	if !reflect.DeepEqual(headSlice, [] HeadElem{{word: "banana", count: 3, treeNode: nil, pattern: nil}}) {
		t.Error()
	}
}

func TestFilterWordBase(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := nlp.Preprocess(texts, false,"../data/stopwords.txt")
	wordCount := nlp.WordCount(wordBase)
	_, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)
	fmt.Println(filteredWordBase)
}

func TestBuildFPTree(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := nlp.Preprocess(texts, false,"../data/stopwords.txt")
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)

	root := FPNode{}
	BuildFPTree(&root, filteredWordBase, headAddr)
	var treeNodes [] *FPNode
	for _, elem := range headElems {
		if headAddr[elem.word].treeNode != nil {
			treeNodes = append(treeNodes, headAddr[elem.word].treeNode)
		} else {
			fmt.Println(elem)
		}
	}
	if len(treeNodes) != len(headElems) {
		t.Error()
	}
}

func TestConditionalPattern(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := nlp.Preprocess(texts, true,"../data/stopwords.txt")
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)

	root := FPNode{}
	BuildFPTree(&root, filteredWordBase, headAddr)

	ConditionalPattern(&root, headElems, 1, headAddr)

	fmt.Println(headAddr["物流"])
}

func TestWordConcurrence(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := nlp.Preprocess(texts, true,"../data/stopwords.txt")
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)

	root := FPNode{}
	BuildFPTree(&root, filteredWordBase, headAddr)

	ConditionalPattern(&root, headElems, 1, headAddr)

	coWords := WordConcurrence("物流", headAddr, 0.1)
	sort.Strings(coWords)
	if !reflect.DeepEqual(coWords, [] string{"公司", "很快", "服务态度"}) {
		t.Error()
	}
}
