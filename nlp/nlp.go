//nlp 相关处理过程
package nlp

import (
	"bufio"
	"github.com/yanyiwu/gojieba"
	"os"
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

func RemoveStopwords(words, stopwords [] string) [] string {
	var clean [] string
	for _, w := range words {
		if SearchWord(w, stopwords) || w == " " {
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

func SplitSent(text string) [] string {
	re := regexp.MustCompile("[;,.，。；!?？！]")
	return re.Split(text, -1)
}

//预处理过程: 1.分词 2.去停用词 3. 去重
func Preprocess(filePath string, splitSent bool, stopwordsFilePath string, paraNum int) [] [] string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	textCh := make(chan string)
	go func() {
		sc := bufio.NewScanner(file)
		defer close(textCh)
		for sc.Scan() {
			text := sc.Text()
			if splitSent {
				texts := SplitSent(text)
				for _, t := range texts {
					textCh <- t
				}
			} else {
				textCh <- text
			}
		}
	}()

	jieba := gojieba.NewJieba()
	var ret [][]string
	ch := make(chan [] string)
	var wg sync.WaitGroup
	stopwords := utils.ReadFile(stopwordsFilePath)

	for i := 0; i < paraNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for text := range textCh {
				words := jieba.Cut(text, true)
				words = RemoveStopwords(words, stopwords)
				// 去重
				words = DropDuplicates(words)
				ch <- words
			}
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

// 去重
func DropDuplicates(candidates [] string) [] string {
	elemDict := make(map[string]bool)
	for _, can := range candidates {
		elemDict[can] = true
	}
	var dCandidates [] string
	for elem := range elemDict {
		dCandidates = append(dCandidates, elem)
	}
	return dCandidates
}

func WordCount(wordBase [][]string) map[string]int {
	wc := make(map[string]int)
	for _, words := range wordBase {
		for _, word := range words {
			wc[word] += 1
		}
	}
	return wc
}
