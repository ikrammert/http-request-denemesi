package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func main() {
	// HTTP isteği yapmak istediğiniz URL'i belirtin
	var url string
	fmt.Print("URL'yi girin: ")
	fmt.Scanln(&url)
	if url != "" {
		fmt.Println("Url Boş Olamaz")
	}

	// İstek oluşturma
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("İstek oluşturulurken bir hata oluştu:", err)
		return
	}
	//veya
	//checkError(err)

	// İstek başlıklarını ayarlama (isteğe bağlı)
	req.Header.Set("User-Agent", "My-Go-Program")

	// İstek gönderme
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("İstek gönderilirken bir hata oluştu:", err)
		return
	}
	defer resp.Body.Close()

	// Yanıtı okuma
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Yanıt okunurken bir hata oluştu:", err)
		return
	}

	// Yanıtı ekrana yazdırma
	fmt.Println("Yanıt:", string(body))
}
