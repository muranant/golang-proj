// concurrent implementation
package main

import "fmt"
import "io/ioutil"
import "strings"
import "os"
import "bufio"
import "sync"
import "strconv"

//struct to save details of the search
type MapResults struct {
	filename string
	mr       map[string]int
}

func main() {
	fmt.Println("Start main")
	dirname := "/go/src/15-text"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	mapresults := make(chan MapResults, 135)

	cindex := 0
	var wg sync.WaitGroup
	for _, rfile := range files {
		wg.Add(1)
		go func(filename string, count int) {
			defer wg.Done()
			fullname := dirname + "/" + filename
			FreqCount(fullname, mapresults)
		}(rfile.Name(), cindex)
		cindex++
	}
	wg.Wait()

	acc := make(map[string]int)
	for {
		select {
		case result := <-mapresults:
			for k, v := range result.mr {
				acc[k] += v
			}
		default:
			{
				// Dump out final accumulated result:
				fmt.Println("Final map: /tmp/results.txt")
				WriteToFile(acc, "/tmp/results.txt")
				return
			}
		}
	}

}

func FreqCount(filename string, mapresults chan MapResults) {

	//Allocate object
	lmapr := new(MapResults)
	lmapr.mr = make(map[string]int)
	lmapr.filename = filename
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s\n", filename)
		os.Exit(1)
	}
	//fmt.Printf("Reading %s:\n", filename)
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		// replace the punctuations with space so that we dont lose the word
		rr := prunechars(line, "0123456789")
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
		r := pReplacer.Replace(rr)
		s := strings.TrimSpace(r)
		for _, word := range strings.Split(s, " ") {
			if word != "" {
				lmapr.mr[word]++
			}
		}
	}
	mapresults <- *lmapr
	//close(mapresults)
}

func prunechars(str, prune string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(prune, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func WriteToFile(result map[string]int, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()
	for word, n := range result {
		str := word + ":" + strconv.Itoa(n) + "\n"
		f.WriteString(str)
	}
}
