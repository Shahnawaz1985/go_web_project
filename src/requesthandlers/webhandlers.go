package requesthandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"typedef"
)

func ViewHandler(writer http.ResponseWriter, request *http.Request){
	pwd, _ := os.Getwd()
	fmt.Println("pwd :", pwd)
	signatures := typedef.GetStrings(pwd+"/src/templates/signatures.txt")
	html, err := template.ParseFiles(pwd+"/src/templates/view.html")
	typedef.Check(err)
	guestbook := typedef.Guestbook{
		SignatureCount : len(signatures),
		Signatures: signatures,
	}

	err = html.Execute(writer, guestbook)
	typedef.Check(err)
}

func NewHandler(writer http.ResponseWriter, request *http.Request){
	pwd, _ := os.Getwd()
	fmt.Println("pwd :", pwd)
	html, err := template.ParseFiles(pwd+"/src/templates/new.html")
	typedef.Check(err)
	err = html.Execute(writer, nil)
	typedef.Check(err)
}

func CreateHandler(writer http.ResponseWriter, request *http.Request) {
	pwd, _ := os.Getwd()
	fmt.Println("pwd :", pwd)
	signature := request.FormValue("signature") +"\n"
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(pwd+"/src/templates/signatures.txt", options, os.FileMode(0600))
	typedef.Check(err)
	_, err = fmt.Fprintln(file, signature)
	typedef.Check(err)
	err = file.Close()
	typedef.Check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
