package pages

import (
	"automation/utils"
	"fmt"

	"github.com/tebeka/selenium"
)

type LoginPage struct {
	wd selenium.WebDriver
}

func NewLoginPage(wd selenium.WebDriver) *LoginPage {
	return &LoginPage{wd: wd}
}

func (p *LoginPage) Open() error {
	err := p.wd.Get("https://magento.softwaretestingboard.com/customer/account/login/")
	if err != nil {
		return err
	}
	if blocked, _ := utils.IsErrorOrCaptchaPage(p.wd); blocked {
		return fmt.Errorf("Blocked or captcha page detected on login")
	}
	return nil
}

func (p *LoginPage) Login(email, password string) error {
	el, err := utils.WaitForElement(p.wd, selenium.ByID, "email", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(email); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "pass", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(password); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "send2", 5)
	if err != nil {
		return err
	}
	if err := el.Click(); err != nil {
		return err
	}
	return nil
}
