package core

import (
	"log"
	"reflect"
	"testing"
	"word_concurrence/nlp"
)

func TestBuildHeadElems(t *testing.T) {
	wordCount := map[string]int{"apple": 2, "banana": 3, "watermelon": 1}
	headElems, _ := BuildHeadElems(wordCount, 0)
	words_test := [] string{"banana", "apple", "watermelon"}
	var words [] string
	for _, elem := range headElems {
		words = append(words, elem.word)
	}
	if !reflect.DeepEqual(words, words_test) {
		t.Error()
	}
}

func TestFilterWordBase(t *testing.T) {
	wordBase := nlp.Preprocess("../test/validate.txt", true, "../data/stop_words.txt", 10)
	wordCount := nlp.WordCount(wordBase)
	headElem, headAddr := BuildHeadElems(wordCount, 2)
	log.Println(headElem)
	filteredWordBase := FilterWordBase(headAddr, wordBase)
	log.Println(filteredWordBase)
}

func TestWordConcurrence(t *testing.T) {
	wordBase := nlp.Preprocess("../test/validate.txt", true, "../data/stop_words.txt", 10)
	wordCount := nlp.WordCount(wordBase)
	headElems, headAddr := BuildHeadElems(wordCount, 2)
	filteredWordBase := FilterWordBase(headAddr, wordBase)

	root := &FPRoot{}
	root.BuildFPTree(filteredWordBase, headAddr)
	root.ConditionalPattern(headElems, 2, headAddr, 10)

	freqItems := WordConcurrence(headAddr, 0)
	log.Println(freqItems)
}

func BenchmarkWordConcurrence(b *testing.B) {
	paraNum := 10
	wordBase := nlp.Preprocess("../test/comments100w.txt", true, "../data/stop_words.txt", paraNum)
	wordCount := nlp.WordCount(wordBase)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		headElems, headAddr := BuildHeadElems(wordCount, 3)
		filteredWordBase := FilterWordBase(headAddr, wordBase)

		root :=&FPRoot{}
		root.BuildFPTree(filteredWordBase, headAddr)
		root.ConditionalPattern(headElems, 3, headAddr, 10)
		//
		freqItems := WordConcurrence(headAddr, 0)
		log.Println(freqItems)
	}
}
