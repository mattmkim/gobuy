package service

import (
	"fmt"
	"gobuy/uniqlo"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type T struct {
	uniqlo *uniqlo.T
}

func Start() {
	//var s T
	// f := config.NewINIFile("config.ini")
	// config, err := f.Load()

	go func() {
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	var opts []selenium.ServiceOption
	service, err := selenium.NewChromeDriverService(os.Getenv("CHROMEDRIVER_PATH"), port, opts...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{Path: os.Getenv("GOOGLE_CHROME_BIN")})

	path, err := exec.LookPath(os.Getenv("CHROMEDRIVER_PATH"))
	if err != nil {
		fmt.Println("Browser binary path not found")
	}

	fmt.Println(path)
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Error(err.Error())
	}

	defer wd.Quit()

	time.Sleep(time.Minute * 10)

	// if s.uniqlo, err = uniqlo.TestSpawn(wd, config); err != nil {
	// 	log.Error("couldn't start uniqlo service")
	// }

	// // block until OS signal received
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	defer wg.Done()
	// 	log.Info("waiting for signal")
	// 	receivedSignal := <-c
	// 	log.Info(fmt.Sprintf("received signal: %s", receivedSignal))
	// 	s.Stop()
	// }()

	// wg.Wait()
}

func (s *T) Stop() {
	s.uniqlo.Stop()
}
