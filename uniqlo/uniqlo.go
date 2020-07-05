package uniqlo

import (
	"gobuy/notifier"
	"log"

	"github.com/tebeka/selenium"
)

func Spawn(wd selenium.WebDriver) error {
	service := "Uniqlo"
	message := "Gudetama shirt not available."

	// Get interface
	err := wd.Get("https://www.uniqlo.com/us/en/ut-graphic-tees/gudetama")
	if err != nil {
		log.Printf("couldn't get website: %s", err.Error())
		return err
	}

	_, err = wd.FindElement(selenium.ByCSSSelector, "#categoryspecific")
	if err != nil {
		log.Printf("No products available: %s", err.Error())
		notifier.Notify(service, message)
		return nil
	}

	return nil
}
