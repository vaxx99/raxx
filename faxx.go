package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type stat struct {
	Nm string
	Sn string
	Ln string
	Un string
}

var ss []stat

func raxx() []stat {
	var st []stat
	doc, _ := goquery.NewDocument("https://www.internet-radio.com")
	doc.Find(".text-danger").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		_, e := a.Attr("href")
		if e != false {
			c := a.Text()
			st = append(st, stat{c, "", "", ""})
		}
	})
	doc.Find("small.hidden-xs").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		b, e := a.Attr("onclick")
		if e != false {
			c := b[strings.Index(b, "http://") : strings.Index(b, ")")-1]
			resp, _ := http.Get(c)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			u := string(body)
			u = strings.Replace(u, "\n", "", -1)
			if strings.Index(u, "File1=") > 0 {
				u = u[strings.Index(u, "File1=")+len("File1="):]
			}
			if strings.Index(u, "Title") > 0 {
				u = u[:strings.Index(u, "Title")]
			}
			lf := ""
			if strings.Index(c, ".pls") > 0 {
				lf = ";"
			}
			st[i].Un = u + lf
			b = b[strings.Index(b, "http://")+len("http://"):]
			b = b[:strings.Index(b, "/")]
			st[i].Ln = b
		}
	})

	doc.Find("td.text-center").Each(func(i int, s *goquery.Selection) {
		st[i].Sn = s.Next().Find("b").Text()
	})
	return st
}

func Stat(w http.ResponseWriter, r *http.Request) {
	i := 0
	s := ""
	fn := r.URL.RawQuery
	fn = fn[:2]
	id := r.URL.Query().Get(fn)
	i, _ = strconv.Atoi(id)
	switch fn {
	case "ln":
		s = "http://" + ss[i].Ln
	case "nm":
		s = ss[i].Nm
	case "un":
		s = ss[i].Un
	}
	fmt.Fprintf(w, "%s", s)
}

func main() {
	ss = raxx()
	//for i, j := range ss {
	//	fmt.Println(i, j.Nm, j.Sn, j.Ln, j.Un)
	//}
	http.HandleFunc("/stat", Stat)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
