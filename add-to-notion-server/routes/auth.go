package routes

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"net/http"
	"os"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello")

	query := r.URL.Query()
	code := query.Get("code")
	w.WriteHeader(200)
	w.Write([]byte(code))

	body := struct {
		grant_type   string
		code         string
		redirect_uri string
	}{
		grant_type:   "authorization_code",
		code:         code,
		redirect_uri: "https://192.168.0.149:4452/auth",
	}

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	NOTION_CLIENT_ID := os.Getenv("NOTION_CLIENT_ID")
	NOTION_CLIENT_SECRET := os.Getenv("NOTION_CLIENT_SECRET")

	// env := struct {
	// 	NotionAuthUrl     string
	// 	NotionRedirectUrl string
	// }{
	// 	NotionAuthUrl:     os.Getenv("NOTION_AUTH_URL"),
	// 	NotionRedirectUrl: os.Getenv("NOTION_REDIRECT_URL"),
	// }
	// body := `{
	// 	"grant_type": "authorization_code",
	// 	"code": "",
	// 	"redirect_uri": "https://192.168.0.149:8100/auth"
	// }`

	// t := ttemp.Must(ttemp.New("").Parse(body))
	// fmt.Println(body)
	// fmt.Println(t)
	jsonData := []byte(`{
		"grant_type": "authorization_code",
		"code": "` + body.code + `",
		"redirect_uri": "https://192.168.0.149:8100/auth"
	}`)

	// headers := struct {
	// 	NotionVersion string
	// }{
	// 	NotionVersion: "2021-08-16",
	// }

	// fmt.Println(headers)

	request, error := http.NewRequest("POST", "https://api.notion.com/v1/oauth/token", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Notion-Version", "2021-08-16")
	request.Header.Set("Authorization", "Bearer "+code)
	request.SetBasicAuth(NOTION_CLIENT_ID, NOTION_CLIENT_SECRET)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}

	fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(data))

	if error != nil {
		http.RedirectHandler("https://192.168.0.149:8100", http.StatusUnauthorized)
	} else {
		http.RedirectHandler("https://192.168.0.149:8100/profile.html", http.StatusTemporaryRedirect)
	}
}
