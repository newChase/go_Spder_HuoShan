package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}
type HuoShanRequest struct {
	Text     string `json:"text"`
	language string `json:"language"`
}

type HuoShanResponse struct {
	Words []struct {
		Source  int    `json:"source"`
		Text    string `json:"text"`
		PosList []struct {
			Type      int `json:"type"`
			Phonetics []struct {
				Type int    `json:"type"`
				Text string `json:"text"`
			} `json:"phonetics"`
			Explanations []struct {
				Text     string `json:"text"`
				Examples []struct {
					Type      int `json:"type"`
					Sentences []struct {
						Text      string `json:"text"`
						TransText string `json:"trans_text"`
					} `json:"sentences"`
				} `json:"examples"`
				Synonyms []interface{} `json:"synonyms"`
			} `json:"explanations"`
			Relevancys []interface{} `json:"relevancys"`
		} `json:"pos_list"`
	} `json:"words"`
	Phrases  []interface{} `json:"phrases"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

// func query(word string) {
// 	client := &http.Client{}
// 	request := DictRequest{TransType: "en2zh", Source: word}
// 	buf, err := json.Marshal(request)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var data = bytes.NewReader(buf)
// 	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Set("Connection", "keep-alive")
// 	req.Header.Set("DNT", "1")
// 	req.Header.Set("os-version", "")
// 	req.Header.Set("sec-ch-ua-mobile", "?0")
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
// 	req.Header.Set("app-name", "xy")
// 	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
// 	req.Header.Set("Accept", "application/json, text/plain, */*")
// 	req.Header.Set("device-id", "")
// 	req.Header.Set("os-type", "web")
// 	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
// 	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
// 	req.Header.Set("Sec-Fetch-Site", "cross-site")
// 	req.Header.Set("Sec-Fetch-Mode", "cors")
// 	req.Header.Set("Sec-Fetch-Dest", "empty")
// 	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
// 	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
// 	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	bodyText, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if resp.StatusCode != 200 {
// 		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
// 	}
// 	var dictResponse DictResponse
// 	err = json.Unmarshal(bodyText, &dictResponse)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
// 	for _, item := range dictResponse.Dictionary.Explanations {
// 		fmt.Println(item)
// 	}
// }
func query2(word string) {

	// client := &http.Client{}
	// //var data = strings.NewReader(`{"text":"Hello","language":"en"}`)
	//
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDVlo83KSWsEQPt/pLwt&_signature=_02B4Z6wo000010zg.jAAAIDCx6tl0ziaf2NM4PqAALFZ5yVJZRswPmgYBtpLRjtmWoXaTdDHF.CIEOVMu5qysqx810cmorVhD7.JkthtixtPL9e6.E2ioPN95A5EGKfrU8KJEvxXdm4OsdO30d", data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf("{\"text\":\"%s\",\"language\":\"en\"}", word))

	// request := HuoShanRequest{Text: "good\n\n", language: "en"}
	// buf, err := json.Marshal(request)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var data = bytes.NewReader(buf)
	//fmt.Println(data)
	// return
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzswVLQDGX3W3KSWMLVPt/pLvF&_signature=_02B4Z6wo00001BbzKPwAAIDBnbizHB2us1gW8yxAAGfS5yVJZRswPmgYBtpLRjtmWoXaTdDHF.CIEOVMu5qysqx810cmorVhD7.JkthtixtPL9e6.E2ioPN95A5EGKfrU8KJEvxXdm4OsdO37b", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16519424297964386; i18next=zh-CN; s_v_web_id=verify_10325c7cf807f12cc4fe16d7e6690ed9; _tea_utm_cache_2018=undefined; ttcid=7fb2c419bf9c475492edd61aad1429a385; tt_scid=L72lF3.oAMQuiYmMV6cc2WhAdqrUsekJCNQlnesZRPqRvojU6KLdPX.eFR0K.FT5d8e2")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=good%0A%0A")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	resp, err := client.Do(req)
	//fmt.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	// if resp.StatusCode != 200 {
	// 	log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	// }

	var HuoResponse HuoShanResponse
	err = json.Unmarshal(bodyText, &HuoResponse)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(word, "UK:", HuoResponse.Words.Prons.En, "US:", HuoResponse.Dictionary.Prons.EnUs)
	for _, item := range HuoResponse.Words[0].PosList {
		fmt.Println(item.Explanations)
	}
	// //fmt.Println(HuoResponse.Words.PosList.Explanations.Text)
	// //fmt.Printf("%s\n", bodyText)
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	//query(word)
	query2(word)
}
