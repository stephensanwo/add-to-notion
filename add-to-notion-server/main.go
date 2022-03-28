package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"add-to-notion-server/routes"
	"github.com/joho/godotenv"
)

type IndexPage struct {
	Title       string
	Description string
	Long        string
	AuthUrl     string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	HTML_AUTH_URL := os.Getenv("HTML_AUTH_URL")
	fmt.Println(HTML_AUTH_URL)
	page := IndexPage{
		Title:       "Add to Notion",
		Description: "Add to Notion",
		Long:        "Leverage the power of Artificial Intelligence with Open AI's GPT3 to write blog posts, compose emails, write marketing briefs, social media copies, etc. all within your notion workspace, without writing a single line of code!. Try Add to Notion for free, by authenticating with Notion and selecting the pages you want Notion AI to have access to.",
		AuthUrl:     HTML_AUTH_URL,
	}

	t, _ := template.ParseFiles("./static/index.html")
	fmt.Println(t.Execute(w, page))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/auth", routes.AuthHandler)
	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	log.Println("Server running on port 8100...")
	// http.ListenAndServe(":8100", nil)
	http.ListenAndServeTLS(":8100", "./ssl/go-server.crt", "./ssl/go-server.key", nil)

}
