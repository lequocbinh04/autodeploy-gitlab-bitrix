package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func SendMessage(text, chatId string) {
	token := os.Getenv("BitrixToken")
	clientId := os.Getenv("BotClientId")

	url := "https://crondata.bitrix24.vn/rest/8/" + token + "/imbot.message.add.json?BOT_ID=26&CLIENT_ID=" + clientId + "&DIALOG_ID=" + chatId
	method := "POST"
	text = strings.ReplaceAll(text, "\"", "\\\"")
	text = strings.ReplaceAll(text, "\n", "\\n")
	payload := strings.NewReader(`{"MESSAGE": "` + text + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "BITRIX_SM_SALE_UID=0; qmb=.")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
}
