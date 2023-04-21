package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	ShowHTML "news.com/events/src/ShowHTML"
	SignUp "news.com/events/src/SignUp"
)

type User struct {
	Username string
	Password string
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// url := fmt.Sprintf("/?userID=1111")
	// Định nghĩa hàm xử lý request HTTP

	ShowHTML.RequestHTTP()
	SignUp.SignUpHandler()

	http.ListenAndServe(":8080", nil)
}

func getHtmlContent(htmlName string) []byte {
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
