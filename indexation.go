package allcaps

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

var db = map[string][]string{
	"aia": []string{"Allcaps is awesome!!"},
}

func Index(phrase string) {
	caps := capsOf(phrase)
	// fmt.Println(caps)
	//db[caps] = append(db[caps], phrase)
	indexWithSubCaps(caps, phrase)
}

func indexWithSubCaps(caps, phrase string) {
	// TODO in a Trie, index all subcaps
	// https://en.wikipedia.org/wiki/Trie

	// Naive impl (poor runtime perf, high memory usage)
	for i := 0; i+5 <= len(caps); i++ {
		for j := i + 5; j <= len(caps); j++ {
			sub := caps[i:j]
			db[sub] = append(db[sub], phrase)
		}
	}

	// TODO introduce error-tolerance when length>=9 ?
}

func capsOf(phrase string) string {
	var caps bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(phrase))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		caps.WriteString(scanner.Text()[:1])
		// TODO consider 1st rune, not 1st byte
	}
	letters := string(caps.Bytes())
	letters = nonAscii.ReplaceAllLiteralString(letters, "")
	letters = digits.ReplaceAllLiteralString(letters, "")
	// TODO turn accented letters into ascii equivalent: é -> e
	lower := strings.ToLower(letters)
	return lower
}

var nonAscii = regexp.MustCompile("[[:^ascii:]]")
var digits = regexp.MustCompile("[0-9]")
