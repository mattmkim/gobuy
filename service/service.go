package service

import (
	"fmt"
	"gobuy/uniqlo"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	"github.com/zpatrick/go-config"
)

type T struct {
	uniqlo *uniqlo.T
}

func Start() {
	var s T
	f := config.NewINIFile("config.ini")
	config, err := f.Load()

	caps := selenium.Capabilities{"browserName": "chrome"}
	time.Sleep(time.Second * 5)
	wd, err := selenium.NewRemote(caps, "http://host.docker.internal:4444/wd/hub")
	if err != nil {
		log.Error(err.Error())
	}

	defer wd.Quit()

	if s.uniqlo, err = uniqlo.Spawn(wd, config); err != nil {
		log.Error("couldn't start uniqlo service")
	}

	// block until OS signal received
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Info("waiting for signal")
		receivedSignal := <-c
		log.Info(fmt.Sprintf("received signal: %s", receivedSignal))
		s.Stop()
	}()

	wg.Wait()
}

func (s *T) Stop() {
	s.uniqlo.Stop()
}
