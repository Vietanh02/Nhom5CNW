package ShowHTML

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func RequestHTTP() {
	ShowHTMLpage("/", "homepage.html")
	ShowHTMLpage("/html/mainpage/news-page", "html/mainpage/news-page.html")
	ShowHTMLpage("/html/mainpage/courses-page", "html/mainpage/courses-page.html")
	ShowHTMLpage("/html/sign-in", "html/sign-in/sign-in.html")
	// ShowHTMLpage("/html/sign-up", "html/sign-up/sign-up.html")
}

func ShowHTMLpage(pagePath string, pageName string) {
	http.HandleFunc(pagePath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(GetHTMLContent(pageName))
	})
}

func GetHTMLContent(htmlName string) []byte {
	htmlFile, err := os.Open(htmlName)
	if err != nil {
		fmt.Println(err)

	}
	defer htmlFile.Close()

	htmlContent, err := ioutil.ReadAll(htmlFile)
	if err != nil {
		fmt.Println(err)

	}
	return htmlContent
}
