package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/vaxx99/raxx/xmp"
)

type stat struct {
	Nm string
	Sn string
	Ln string
}

func raxx() {
	var st []stat
	doc, _ := goquery.NewDocument("https://www.internet-radio.com")
	doc.Find(".text-danger").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		_, e := a.Attr("href")
		if e != false {
			c := a.Text()
			st = append(st, stat{c, "", ""})
		}
	})
	doc.Find("small.hidden-xs").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		b, e := a.Attr("onclick")
		if e != false {
			b = b[strings.Index(b, "http://")+len("http://"):]
			b = b[:strings.Index(b, "/")]
			st[i].Ln = b
		}
	})
	doc.Find("td.text-center").Each(func(i int, s *goquery.Selection) {
		st[i].Sn = s.Next().Find("b").Text()
	})
	t, _ := template.New("raxx").Parse(xmp.Tmpl)
	f, e := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if e != nil {
		log.Fatalln(e)
	}
	defer f.Close()
	t.Execute(f, st)
}

func main() {
	raxx()
}
