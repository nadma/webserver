package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	jsonStr := []byte(`{
  "mock_data": "true",
  "ip_address": "92.188.61.181",
  "email": "user@example.com",
  "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_4) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.100 Safari/534.30",
  "url": "http://example.com/"
}`)
	responsy := make([]byte, 9000)

	req, err := http.NewRequest("POST", "https://enlhyhwrs0m0qfe.m.pipedream.net", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.Body.Read(responsy))

}
