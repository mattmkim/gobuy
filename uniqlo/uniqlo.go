package uniqlo

import (
	"fmt"
	"uniqgo/notifier"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

type T struct {
	wd     selenium.WebDriver
	config map[string]string
	stopCh chan struct{}
}

func (s *T) Check() {
	url := "https://www.uniqlo.com/us/en/ut-graphic-tees/gudetama"
	err := s.wd.Get(url)
	if err != nil {
		log.Error(fmt.Sprintf("couldn't get %s: %s", url, err.Error()))
		notifier.Notify(s.config, fmt.Sprintf("Web driver couldn't get %s", url))
	}

	_, err = s.wd.FindElement(selenium.ByCSSSelector, "#categoryspecific")
	if err != nil {
		log.Printf("Gudetama shirts not available: %s", err.Error())
	} else {
		log.Print("Gudetama shirts available")
		notifier.Notify(s.config, "Gudetama shirts available!")
	}
}

func Spawn(wd selenium.WebDriver, config map[string]string) (*T, error) {
	s := &T{
		wd:     wd,
		config: config,
		stopCh: make(chan struct{}),
	}

	cron := cron.New()
	_, err := cron.AddFunc("0 * * * *", s.Check)
	if err != nil {
		log.Error("couldn't add function to cron job")
		return s, err
	}

	cron.Start()

	go func() {
		select {
		case <-s.stopCh:
			return
		}
	}()

	return s, nil
}

func (s *T) Stop() {
	close(s.stopCh)
}
