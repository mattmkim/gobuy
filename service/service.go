package service

import (
	"fmt"
	"gobuy/uniqlo"

	"github.com/pkg/errors"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func Start() error {
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless", // <<<
			"--no-sandbox",
		},
	}
	caps.AddChrome(chromeCaps)
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	if err != nil {
		fmt.Println(err)
	}

	defer wd.Quit()

	if err = uniqlo.Spawn(wd); err != nil {
		return errors.Wrap(err, "couldn't start uniqlo service")
	}

	return nil

}
