package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CaptchaService string `yaml:"captchaService"`
	CaptchaKey     string `yaml:"captchaKey"`
}

func ReadConfig(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func WriteConfig(filename string, config Config) error {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ChangeConfig(config *Config) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Would you like to change the Captcha Key? (y/n): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))

	if answer == "y" || answer == "yes" {
		fmt.Print("Enter the new Captcha Key: ")
		captchaKey, _ := reader.ReadString('\n')
		captchaKey = strings.TrimSpace(captchaKey)

		config.CaptchaKey = captchaKey
	}
}

func WriteDataToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", data)
	if err != nil {
		return err
	}

	return nil
}
