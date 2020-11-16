package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath     = "vendor/selenium-server-standalone-3.141.59.jar"
    chromeDriverPath = "vendor/chromedriver"
	port             = 8080
	username		 = "atikayuwdn"
	password		 = "atkywdn22897"
)

type saleList struct{
	element selenium.WebElement
	name string
	price int
}

func main() {
	opts := []selenium.ServiceOption{
		// selenium.StartFrameBuffer(),
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	if err!= nil {
		panic(err)
	}
	// defer wd.Quit()

	if err := wd.Get("https://www.sociolla.com/flash-sale"); err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(60 * time.Second)

	// login 
	// find username input form
	if !login(wd) {
		fmt.Printf("Cant Login. Something is wrong")
	}

	flashsaleList, err := wd.FindElements(selenium.ByCSSSelector, ".loaded-item");
	if err!= nil{
		panic(err)
	}
	fmt.Printf("HADUH KONTOL %d\n",len(flashsaleList))

	// fmt.Print(flashsaleList)
}

func login(wd selenium.WebDriver) bool{
	loginForm, err := wd.FindElement(selenium.ByCSSSelector, ".mybag_wrapper")
	if err!= nil{
		panic(err)
	}
	loginForm.Click()
	userNameInput, err := wd.FindElement(selenium.ByCSSSelector, "input#username");
	if err!= nil{
		panic(err)
	}
	err = userNameInput.SendKeys(username)
	if err!= nil{
		panic(err)
	}

	passwordInput, err := wd.FindElement(selenium.ByCSSSelector, "input[name=password]");
	if err!= nil{
		panic(err)
	}
	err = passwordInput.SendKeys(password)
	if err!= nil{
		panic(err)
	}
	submitButton, err  := wd.FindElement(selenium.ByCSSSelector, ".formaction")
	submitButton.Click()
	return true
}
