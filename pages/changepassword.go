package pages

import (
	"automation/utils"
	"fmt"

	"github.com/tebeka/selenium"
)

type ChangePasswordPage struct {
	wd selenium.WebDriver
}

func NewChangePasswordPage(wd selenium.WebDriver) *ChangePasswordPage {
	return &ChangePasswordPage{wd: wd}
}

func (p *ChangePasswordPage) Open() error {
	err := p.wd.Get("https://magento.softwaretestingboard.com/customer/account/edit/")
	if err != nil {
		return err
	}
	if blocked, _ := utils.IsErrorOrCaptchaPage(p.wd); blocked {
		return fmt.Errorf("Blocked or captcha page detected on change password")
	}
	return nil
}

func (p *ChangePasswordPage) ChangePassword(current, newpass string) error {
	el, err := utils.WaitForElement(p.wd, selenium.ByID, "change-password", 5)
	if err != nil {
		return err
	}
	if err := el.Click(); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "current-password", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(current); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "password", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(newpass); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByID, "password-confirmation", 5)
	if err != nil {
		return err
	}
	if err := el.SendKeys(newpass); err != nil {
		return err
	}
	el, err = utils.WaitForElement(p.wd, selenium.ByXPATH, "//button[@title='Save']", 5)
	if err != nil {
		return err
	}
	if err := el.Click(); err != nil {
		return err
	}
	return nil
}
