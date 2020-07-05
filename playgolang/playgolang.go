package playgolang

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var code = `
package main
import "fmt"

func main() {
    fmt.Println("Hello WebDriver!\n")
}
`

func Spawn() {
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

	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	elem.Clear()
	elem.SendKeys(code)

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	btn.Click()

	// Get the result
	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")

	output := ""
	// Wait for run to finish
	for {
		output, _ = div.Text()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Got: %s\n", output)
}
