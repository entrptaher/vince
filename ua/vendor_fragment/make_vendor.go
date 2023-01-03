package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"

	"github.com/gernest/vince/ua"
)

func main() {
	var b bytes.Buffer

	fmt.Fprintln(&b, "// DO NOT EDIT Code generated by ua/os/make_os.go")
	fmt.Fprintln(&b, " package vince")
	fmt.Fprintln(&b, `

type vendorRe struct{
	re *ReMatch
	name        string  
}
type vendorResult struct {
	name        string  
}

`)
	err := generate(&b, "vendorfragments.yml")
	if err != nil {
		log.Fatal(err)
	}
	r, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("ua_vendor.go", r, 0600)
}

type Vendor struct {
	name string
	re   []string
}

type VSLice []*Vendor

func (x VSLice) Len() int           { return len(x) }
func (x VSLice) Less(i, j int) bool { return x[i].name < x[j].name }
func (x VSLice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func generate(b *bytes.Buffer, path string) error {
	var m map[string][]string

	err := ua.Read(path, &m)
	if err != nil {
		return err
	}
	var items []*Vendor
	for k, v := range m {
		items = append(items, &Vendor{
			name: k, re: v,
		})
	}
	sort.Sort(VSLice(items))

	var s bytes.Buffer
	var started bool
	for i, d := range items {
		if started {
			s.WriteByte('|')
		} else {
			started = true
		}
		if i != 0 {
			s.WriteByte('|')
		}
		for _, r := range d.re {
			s.WriteString(r)
		}
	}
	if ua.IsStdRe(s.String()) {
		fmt.Fprintf(b, " var vendorAllRe= MatchRe(`%s`)\n", ua.Clean(s.String()))
	} else {
		fmt.Fprintf(b, " var vendorAllRe= MatchRe2(`%s`)\n", ua.Clean(s.String()))
	}
	fmt.Fprintf(b, "var vendorAll=[]*vendorRe{\n")
	var buf bytes.Buffer
	for _, d := range items {
		buf.Reset()
		s.Reset()
		for k, v := range d.re {
			if k != 0 {
				s.WriteByte('|')
			}
			s.WriteString(v)
		}
		r := ua.Clean(s.String())
		if ua.IsStdRe(s.String()) {
			fmt.Fprintf(&buf, "re:MatchRe(`%s`)", r)
		} else {
			fmt.Fprintf(&buf, "re: MatchRe2(`%s`)", r)
		}
		fmt.Fprintf(b, "{%s,name:%q", &buf, d.name)
		fmt.Fprintf(b, "},\n")
	}
	fmt.Fprintln(b, "}")
	return nil
}
