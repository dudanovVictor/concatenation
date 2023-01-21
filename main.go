package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

func findAllConcatenatedWordsInADict1(words []string) []string {
	dict := make(map[string][]string)
	for i, v1 := range words {
		for j, v2 := range words {
			if i != j && strings.Contains(v2, v1) {
				dict[v2] = append(dict[v2], v1)
			}
		}
	}
	var result []string
	for k, v := range dict {
		if len(v) > 1 {
			sort.Slice(v, func(i, j int) bool {
				return len(v[i]) > len(v[j])
			})
			kk := k
			for _, i := range v {
				kk = strings.ReplaceAll(kk, i, "")
			}
			if len(kk) == 0 {
				result = append(result, k)
			}
		}
	}
	return result
}
func findAllConcatenatedWordsInADict2(words []string) []string {
	dict := make([][]string, len(words))
	for i, v1 := range words {
		for j, v2 := range words {
			if i != j && strings.Contains(v2, v1) {
				if dict[j] == nil {
					dict[j] = make([]string, 0, 2)
				}
				dict[j] = append(dict[j], v1)
			}
		}
	}
	var result []string
	for i, v := range dict {
		if len(v) > 1 {
			/*sort.Slice(v, func(i, j int) bool {
				return len(v[i]) > len(v[j])
			})
			/**/
			w := words[i]
			for _, j := range v {
				w = strings.ReplaceAll(w, j, "")
			}
			if len(w) == 0 {
				result = append(result, words[i])
			}
		}
	}
	return result
}
func findAllConcatenatedWordsInADict3(words []string) []string {
	dict := make([][]string, len(words))
	for i, v1 := range words {
		for j, v2 := range words {
			if i != j && strings.Contains(v2, v1) {
				if dict[j] == nil {
					dict[j] = make([]string, 0, 2)
				}
				dict[j] = append(dict[j], v1)
			}
		}
	}
	var result []string
	for i, v := range dict {
		if len(v) > 0 {
			sort.Slice(v, func(i, j int) bool {
				return len(v[i]) > len(v[j])
			})
			w := words[i]
			for len(v) > 0 {
				for _, j := range v {
					w = strings.ReplaceAll(w, j, "")
				}
				if len(w) == 0 {
					result = append(result, words[i])
					break
				}
				v = v[1:]
			}
		}
	}
	return result
}

func findAllConcatenatedWordsInADict4(words []string) []string {
	dict := make([][]string, len(words))
	for i, v1 := range words {
		for j, v2 := range words {
			if i != j && strings.Contains(v2, v1) {
				if dict[j] == nil {
					dict[j] = make([]string, 0, 2)
				}
				dict[j] = append(dict[j], v1)
			}
		}
	}
	var result []string
	for i, v := range dict {
		if len(v) > 0 {
			//var prm [][]string
			//generateIntPermutations(v, len(v), &prm)
			prm := permutations(v)
			for _, p := range prm {
				w := words[i]
				for _, j := range p {
					w = strings.ReplaceAll(w, j, "")
				}
				if len(w) == 0 {
					result = append(result, words[i])
					break
				}
			}
		}
	}
	return result
}
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig []string, p []int) []string {
	result := append([]string{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

type Point struct {
	links map[byte]*Point
	last  bool
}

var Points []*Point            //make([]Point, 10000)
var PointMap []map[byte]*Point // make([]map[byte]*Point, 10000)
var lastPoint int

const MaxNodes = 10000

var Pool [MaxNodes]Point

func newPoint() *Point {
	if Points == nil {
		Points = make([]*Point, MaxNodes)
		PointMap = make([]map[byte]*Point, MaxNodes)
		for i := 0; i < MaxNodes; i++ {
			PointMap[i] = make(map[byte]*Point)
			Points[i] = &Pool[i]
		}
		lastPoint = 0
	}
	p := Points[lastPoint]
	p.links = PointMap[lastPoint]
	lastPoint++
	return p
	//return &Point{last: false, links: make(map[byte]*Point)}
}

type Trie struct {
	root *Point
}

func (t Trie) Add(word string) {
	node := t.root
	for _, letter := range word {
		p, ok := node.links[byte(letter)]
		if !ok {
			node.links[byte(letter)] = newPoint()
			node = node.links[byte(letter)]
		} else {
			node = p
		}

	}
	node.last = true
}
func (t Trie) IsConcate(word string, i int) bool {
	node := t.root
	for i < len(word) {
		v, ok := node.links[byte(word[i])]
		if !ok {
			return false
		}
		node = v
		if node.last && (i == len(word)-1 || t.IsConcate(word, i+1)) {
			return true
		}
		i++
	}
	return false
}
func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	result := make([]string, 0, len(words))
	trie := Trie{root: newPoint()}
	for _, word := range words {
		if trie.IsConcate(word, 0) {
			result = append(result, word)
		} else {
			trie.Add(word)
		}
	}
	return result
}

func main() {
	validate := validator.New()
	str := "2022-02-01T08:29:12Z"
	fmt.Println(str)
	errTo := validate.Var(str, "datetime=2006-01-02T15:04:05Z07:00")
	//errTo := validate.Var(str, "datetime=time.RFC3339")
	if errTo != nil {
		fmt.Println(errTo)
	}

	orig := []string{"11", "22", "33"}
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		fmt.Println(getPerm(orig, p))
	}
}
