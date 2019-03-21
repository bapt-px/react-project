package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)


type Result struct {
	average float64
	nbComment int
}

func makeNumber(chaine string) int {
	s := RegSplit(chaine, "[^0-9]")
	textNumber := ""
	for i := range s {
		textNumber += s[i]
	}
	return parseInt(textNumber)
}

func contains(s []string, e string) bool {	// return true if "e" exist into "s"
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RegSplit(text string, delimeter string) []string {	// function that split a strings depending on regular expression
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes) + 1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

func parseInt(nb string) int {	// convert a string to an int
	flag.Parse()
	i, err := strconv.Atoi(nb)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func scrapLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	r.ParseForm()

	url := r.PostForm.Get("url")

	c := colly.NewCollector( )

	result := Result{average: 0, nbComment: 0}

	c.OnHTML(".hotels-hotel-review-community-content-review-list-parts-ReviewFilters__filters_wrap--3q5X7 div div ul li span:nth-child(4)", func(e *colly.HTMLElement) {
		fmt.Println("el: ", e.Text[1])

		fmt.Println("stop", e.Index)

		nbVote := makeNumber(e.Text)
		result.average += float64(nbVote * (5 - e.Index))
		result.nbComment += nbVote

	})



	//// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)
	if result.nbComment > 0 {
		result.average /= float64(result.nbComment)
	}
	str := fmt.Sprintf("%s%f%s%d%s", "{\"average\" : ", result.average, ", \"nbComment\" : ", result.nbComment, "}")
	fmt.Fprintf(w, str)

	fmt.Println(str)

}



func main() {

	router := httprouter.New()
	router.POST("/", scrapLink)

	log.Fatal(http.ListenAndServe(":8083", router))


}