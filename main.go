package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
)

type flags struct {
	match  string
	sortOn string
	out    string
	file   string
}

type element struct {
	matched string
	sortOn  string
	output  string
}

// Implement sort.Interface.
type elementSlice []element

func (es elementSlice) Len() int           { return len(es) }
func (es elementSlice) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es elementSlice) Less(i, j int) bool { return es[i].sortOn < es[j].sortOn }

func main() {
	f := flags{}
	flag.StringVar(&f.match, "match", "", "regex string to match")
	flag.StringVar(&f.sortOn, "sort-on", "", "The string to sort on, should include variables extracted by the match regex")
	flag.StringVar(&f.out, "out", "", "The string to output, should include variables extracted by the match regex")
	flag.StringVar(&f.file, "file", "", "The file to run against")
	flag.Parse()

	b, err := ioutil.ReadFile(f.file)
	if err != nil {
		panic(err)
	}
	matches := regexp.MustCompile(f.match).FindAllStringSubmatch(string(b), -1)
	elements := elementSlice{}
	for _, match := range matches {
		e := element{
			matched: match[0],
			sortOn:  interpolate(f.sortOn, match),
			output:  interpolate(f.out, match),
		}
		elements = append(elements, e)
	}
	sort.Sort(elements)
	for _, e := range elements {
		fmt.Print(e.output)
	}
}

func interpolate(format string, vars []string) string {
	out := format
	for i, v := range vars {
		out = regexp.MustCompile("\\$"+strconv.Itoa(i)).ReplaceAllString(out, v)
		out = regexp.MustCompile("\\\\n").ReplaceAllString(out, "\n")
	}
	return out
}
