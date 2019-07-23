package main

import (
	"flag"
	"word_concurrence/core"
	"word_concurrence/nlp"
	"word_concurrence/utils"
)

func main() {
	supportCount := flag.Int("s", 1, "support Count")
	confidence := flag.Float64("c", 0.001, "confidence")
	paraNum := flag.Int("np", 10, "number parallel")
	filePath := flag.String("fp", "test/comments100.txt", "data file path")
	isSplit := flag.Bool("split", true, "whether split sentence")
	stopwordsPath := flag.String("stopwords", "data/stop_words.txt", "stop words file path")
	outputPath := flag.String("out", "test/output.csv", "output csv path")
	flag.Parse()

	wordBase := nlp.Preprocess(*filePath, *isSplit, *stopwordsPath, *paraNum)
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := core.BuildHeadElems(wordCount, *supportCount)
	filteredWordBase := core.FilterWordBase(headAddr, wordBase)

	root := &core.FPRoot{}
	root.BuildFPTree(filteredWordBase, headAddr)

	root.ConditionalPattern(headElems, *supportCount, headAddr, *paraNum)

	coWords := core.WordConcurrence(headAddr, *confidence)

	data := core.FreqItemsToStrings(coWords)
	headers := [] string{"word", "co_word", "support_count", "confidence"}

	utils.WriteCSV(*outputPath, headers, data)
}
