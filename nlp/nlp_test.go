package nlp

import (
	"fmt"
	"reflect"
	"testing"
	"word_concurrence/utils"
)

func TestRemovePunctuation(t *testing.T) {
	text := "！@；'天气不错，。/：'，。/，,./"
	text = RemovePunctuation(text)
	fmt.Println(text)
	if text != "天气不错" {
		t.Error()
	}
}

func TestRemovePunctuationAll(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	texts = RemovePunctuationAll(texts)
	if len(texts) != 100 {
		t.Error()
	}
}

func TestRemoveCharacter(t *testing.T) {
	text := RemoveCharacter("hellipldquoquot")
	if text != "" {
		t.Error()
	}
}

func TestSplitSents(t *testing.T) {
	texts := [] string{"今天天气不错；心情也不错。纠结ing"}
	texts = SplitSents(texts)
	fmt.Println(texts)
	if len(texts) != 3 {
		t.Error()
	}
}

func TestPreprocess(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := Preprocess(texts, false)
	fmt.Println(wordBase)
	if wordBase == nil {
		t.Error()
	}
}

func TestWordCount(t *testing.T) {
	texts := utils.ReadFile("../data/comments100.txt")
	wordBase := Preprocess(texts, false)
	wc := WordCount(wordBase)
	fmt.Println(wc)
}

func TestRemoveStopwords(t *testing.T) {
	stopwords := utils.ReadFile("../data/stopwords.txt")
	words := RemoveStopwords([] string{"今天", "天气", "不错", "的", "了", "呢"}, stopwords)
	fmt.Println(words)
	if !reflect.DeepEqual(words, [] string{"今天", "天气", "不错"}) {
		t.Error()
	}
}
