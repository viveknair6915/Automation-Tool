package pages

import (
	"fmt"
	"math/rand"
	"time"

	"automation/utils"

	"github.com/tebeka/selenium"
)

type SignupPage struct {
	wd selenium.WebDriver
}

func NewSignupPage(wd selenium.WebDriver) *SignupPage {
	return &SignupPage{wd: wd}
}

func (p *SignupPage) Open() error {
	err := p.wd.Get("https://magento.softwaretestingboard.com/customer/account/create/")
	if err != nil {
		return err
	}
	if blocked, _ := utils.IsErrorOrCaptchaPage(p.wd); blocked {
		return fmt.Errorf("Blocked or captcha page detected on signup")
	}
	return nil
}

func (p *SignupPage) Register(firstName, lastName, email, password string) error {
	el, err := utils.WaitForElement(p.wd, selenium.ByID, "firstname", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(firstName); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "lastname", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(lastName); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "email_address", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(email); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "password", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(password); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "password-confirmation", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(password); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByXPATH, "//button[@title='Create an Account']", 5)
	if err != nil {
		return err
	}
	if err := el.Click(); err != nil {
		return err
	}
	return nil
}

func GenerateRandomEmail() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("testuser%d@example.com", rand.Intn(1000000))
}
