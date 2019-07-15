package main

import (
	"fmt"
	"word_concurrence/core"
	"word_concurrence/nlp"
	"word_concurrence/utils"
)

func main() {
	texts := utils.ReadFile("data/comments100.txt")
	wordBase := nlp.Preprocess(texts, true,"data/stopwords.txt")
	wordCount := nlp.WordCount(wordBase)
	support,confidence:=0.002,0.1
	supportCount:=float64(len(wordBase))*support
	headElems, headAddr := core.BuildHeadElems(wordCount, int(supportCount))
	filteredWordBase := core.FilterWordBase(headAddr, wordBase)

	root := core.FPNode{}
	core.BuildFPTree(&root, filteredWordBase, headAddr)

	core.ConditionalPattern(&root, headElems, 1, headAddr)

	coWords := core.WordConcurrence("物流", headAddr, confidence)
	fmt.Println(coWords)
}
