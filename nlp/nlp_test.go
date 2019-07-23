package nlp

import (
	"log"
	"reflect"
	"sort"
	"testing"
	"word_concurrence/utils"
)

func TestRemovePunctuation(t *testing.T) {
	text := "！@；'天气不错，。/：'，。/，,./"
	text = RemovePunctuation(text)
	log.Println(text)
	if text != "天气不错" {
		t.Error()
	}
}

func TestRemovePunctuationAll(t *testing.T) {
	texts := utils.ReadFile("../test/comments100.txt")
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

func TestDropDuplicates(t *testing.T) {
	words := [] string{"好的", "不错", "好的"}
	words = DropDuplicates(words)
	sort.Strings(words)
	if reflect.DeepEqual(words, [] string{"好的", "不错"}) {
		t.Error()
	}
}

func TestSplitSents(t *testing.T) {
	texts := [] string{"今天天气不错，心情也不错。纠结ing"}
	texts = SplitSents(texts)
	log.Println(texts)
	if len(texts) != 3 {
		t.Error()
	}
}

func TestPreprocess(t *testing.T) {
	wordBase := Preprocess("../test/validate.txt", false, "../data/stop_words.txt", 10)
	log.Println(wordBase)
	if wordBase == nil {
		t.Error()
	}
}

func TestWordCount(t *testing.T) {
	wordBase := Preprocess("../test/comments100.txt", false, "../data/stop_words.txt", 10)
	wc := WordCount(wordBase)
	log.Println(wc)
}

func TestRemoveStopwords(t *testing.T) {
	stopwords := utils.ReadFile("../data/stop_words.txt")
	words := RemoveStopwords([] string{"今天", "天气", "不错", "的", "了", "呢", ","}, stopwords)
	log.Println(words)
	if !reflect.DeepEqual(words, [] string{"今天", "天气", "不错"}) {
		t.Error()
	}
}
