package captcha

import (
	"fmt"

	"github.com/ZennoLab/capmonstercloud-client-go/pkg/client"
	"github.com/ZennoLab/capmonstercloud-client-go/pkg/tasks"
)

func SolveCapMonsterHcaptcha(apiKey string, siteURL string, siteKey string) (*tasks.HCaptchaTaskSolution, error) {
	client := client.New(apiKey)

	// Get balance
	if balance, err := client.GetBalance(); err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	} else {
		fmt.Println(balance)
	}

	// Solve hCaptcha (without proxy)
	hcaptchaTask := tasks.NewHCaptchaTaskProxyless(siteURL, siteKey)
	noCache := false
	result, err := client.SolveHCaptchaProxyless(hcaptchaTask, noCache, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to solve hCaptcha: %v", err)
	}

	return result, nil
}

func SolveCapMonsterRecaptchaV2(apiKey, siteURL, siteKey string) (*tasks.RecaptchaV2TaskSolution, error) {
	client := client.New(apiKey)

	// Get balance
	if balance, err := client.GetBalance(); err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	} else {
		fmt.Println(balance)
	}

	// Solve reCAPTCHA V2 (without proxy)
	recaptchaV2Task := tasks.NewRecaptchaV2TaskProxyless(siteURL, siteKey)
	noCache := false
	result, err := client.SolveRecaptchaV2Proxyless(recaptchaV2Task, noCache, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to solve reCAPTCHA V2: %v", err)
	}

	return result, nil
}

func SolveCapMonsterRecaptchaV3(apiKey, siteURL, siteKey string) (*tasks.RecaptchaV3TaskTaskSolution, error) {
	client := client.New(apiKey)

	// Get balance
	if balance, err := client.GetBalance(); err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	} else {
		fmt.Println(balance)
	}

	// Solve reCAPTCHA V3 (without proxy)
	recaptchaV3Task := tasks.NewRecaptchaV3TaskProxyless(siteURL, siteKey) // Adjust the minimum score as needed
	noCache := false
	result, err := client.SolveRecaptchaV3Proxyless(recaptchaV3Task, noCache, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to solve reCAPTCHA V3: %v", err)
	}

	return result, nil
}
