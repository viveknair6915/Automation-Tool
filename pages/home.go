package pages

import (
	"automation/utils"
	"fmt"

	"github.com/tebeka/selenium"
)

type HomePage struct {
	wd selenium.WebDriver
}

func NewHomePage(wd selenium.WebDriver) *HomePage {
	return &HomePage{wd: wd}
}

func (p *HomePage) IsLoggedIn() (bool, error) {
	_, err := utils.WaitForElement(p.wd, selenium.ByXPATH, "//a[text()='Sign Out']", 5)
	if err == nil {
		return true, nil
	}
	return false, nil
}

func (p *HomePage) SignOut() error {
	el, err := utils.WaitForElement(p.wd, selenium.ByXPATH, "//a[text()='Sign Out']", 5)
	if err != nil {
		return err
	}
	if err := el.Click(); err != nil {
		return err
	}
	if blocked, _ := utils.IsErrorOrCaptchaPage(p.wd); blocked {
		return fmt.Errorf("Blocked or captcha page detected after signout")
	}
	return nil
}
