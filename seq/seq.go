package main

import "fmt"
import "net/http"
import "io/ioutil"
import "strings"
import "os"
import "strconv"
import "path"

func main() {
	fmt.Println("Start main")
	filename := "http://www.gutenberg.org/files/15/text/moby-000.txt"
	baseName := path.Base(filename)
	bytes := ReadFile(filename)
	var wordCount map[string]int = FreqCount(bytes)
	outFileName := "/tmp/" + baseName + ".out"
	WriteToFile(wordCount, outFileName)
	fmt.Printf("Check %s for output\n", outFileName)
}

func FreqCount(text []byte) map[string]int {
	fmap := make(map[string]int)
	r := strings.ToLower(strings.TrimSpace(string(text)))
	// replace the punctuations with space so that we dont lose the word
	pReplacer := strings.NewReplacer(
		",", " ",
		"%", " ",
		"{", " ",
		"}", " ",
		"?", " ",
		"\r", " ",
		"\n", " ",
		".", " ",
		";", " ",
		":", " ",
		"]", " ",
		"[", " ",
		"'", " ",
		"+", " ",
		"*", " ",
		"/", " ",
		"<", " ",
		">", " ",
		"\"", " ",
		"(", " ",
		")", " ",
		"-", " ",
		"!", "",
		"^[0-9]+", "",
	)
	s := pReplacer.Replace(r)
	for _, word := range strings.Split(s, " ") {
		if word != "" {
			fmap[word]++
		}
	}
	return fmap
}

func ReadFile(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func WriteToFile(result map[string]int, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()
	for word, n := range result {
		str := word + ":" + strconv.Itoa(n) + "\n"
		f.WriteString(str)
	}
}
