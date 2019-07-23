package core

import (
	"fmt"
	"log"
	"testing"
	"word_concurrence/nlp"
)

func TestFPNode_Insert(t *testing.T) {
	//tree := &FPTree{}
	root := &FPRoot{}
	p := (*FPNode)(root)
	node1 := FPNode{word: "apple", count: 3, parent: nil, next: nil}
	p.Insert(&node1)
	node2 := FPNode{word: "banana", count: 4, parent: nil, next: nil}
	p.Insert(&node2)
	log.Println(root)
}

func TestFPTree_BuildFPTree(t *testing.T) {
	wordBase := nlp.Preprocess("../test/validate.txt", false, "../data/stop_words.txt", 10)
	wordCount := nlp.WordCount(wordBase)
	_, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)
	log.Println(filteredWordBase)
	root := &FPRoot{}
	root.BuildFPTree(filteredWordBase, headAddr)
	log.Println(fmt.Sprint(root))
}

func TestFPTree_ConditionalPattern(t *testing.T) {
	wordBase := nlp.Preprocess("../test/validate.txt", false, "../data/stop_words.txt", 10)
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)
	root := &FPRoot{}
	root.BuildFPTree(filteredWordBase, headAddr)
	root.ConditionalPattern(headElems, 2, headAddr, 10)
	for word, elem := range headAddr {
		log.Println(word, elem.pattern)
	}
}
