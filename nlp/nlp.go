//nlp 相关处理过程
package nlp

import (
	"github.com/yanyiwu/gojieba"
	"regexp"
	"sync"
	"word_concurrence/utils"
)

func RemovePunctuation(text string) string {
	re, err := regexp.Compile("[^a-zA-Z0-9\u4e00-\u9fa5]+")
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")
}

//去除 标点
func RemovePunctuationAll(texts [] string) [] string {
	re, err := regexp.Compile("[^a-zA-Z0-9\u4e00-\u9fa5]+")
	if err != nil {
		panic(err)
	}
	var postTexts [] string
	for _, text := range texts {
		postTexts = append(postTexts, re.ReplaceAllString(text, ""))
	}
	return postTexts
}

func RemoveCharacter(text string) string {
	re, err := regexp.Compile("hellip|quot|ldquo")
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")
}

func SearchWord(word string, candidates [] string) bool {
	for _, w := range candidates {
		if word == w {
			return true
		}
	}
	return false
}

func RemoveStopwords(words, stopwords [] string, ) [] string {
	var clean [] string
	for _, w := range words {
		if SearchWord(w, stopwords) {
			continue
		} else {
			clean = append(clean, w)
		}
	}
	return clean
}

func SplitSents(texts [] string) [] string {
	var ret [] string
	re := regexp.MustCompile("[;,.，。；!?？！]")
	for _, text := range texts {
		ret = append(ret, re.Split(text, -1)...)
	}
	return ret
}

//预处理过程: 1.分句(Optional)，2.去标点，3.去特殊字符，4. 分词
func Preprocess(texts [] string, splitSent bool,stopwordsFilePath string) [] [] string {
	if splitSent {
		texts = SplitSents(texts)
	}
	jieba := gojieba.NewJieba()
	var ret [][]string
	ch := make(chan [] string)
	var wg sync.WaitGroup

	stopwords := utils.ReadFile(stopwordsFilePath)

	for _, t := range texts {
		text := t
		wg.Add(1)
		go func() {
			defer wg.Done()
			text = RemovePunctuation(text)
			text = RemoveCharacter(text)
			words := jieba.Cut(text, true)
			words = RemoveStopwords(words, stopwords)
			ch <- words
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for words := range ch {
		ret = append(ret, words)
	}

	return ret
}

func WordCount(wordSlice [][]string) map[string]int {
	wc := make(map[string]int)
	for _, text := range wordSlice {
		for _, word := range text {
			wc[word] += 1
		}
	}
	return wc
}
