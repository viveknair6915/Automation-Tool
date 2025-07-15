package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const (
	SeleniumPath = "chromedriver" // Assumes chromedriver is in PATH or project root
	Port         = 9515
)

func StartWebDriver() (selenium.WebDriver, error) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	service, err := selenium.NewChromeDriverService(SeleniumPath, Port)
	if err != nil {
		return nil, fmt.Errorf("failed to start ChromeDriver service: %w", err)
	}
	// Wait for service to start
	time.Sleep(1 * time.Second)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", Port))
	if err != nil {
		service.Stop()
		return nil, fmt.Errorf("failed to start WebDriver: %w", err)
	}
	return &wrappedDriver{WebDriver: wd, service: service}, nil
}

type wrappedDriver struct {
	selenium.WebDriver
	service *selenium.Service
}

func (w *wrappedDriver) Quit() error {
	err := w.WebDriver.Quit()
	w.service.Stop()
	return err
}

func TakeScreenshot(wd selenium.WebDriver, filename string) error {
	img, err := wd.Screenshot()
	if err != nil {
		return err
	}
	if _, err := os.Stat("screenshots"); os.IsNotExist(err) {
		os.Mkdir("screenshots", 0755)
	}
	return os.WriteFile(fmt.Sprintf("screenshots/%s", filename), img, 0644)
}

// WaitForElement waits for an element to appear up to timeoutSeconds, polling every 100ms
func WaitForElement(wd selenium.WebDriver, by, value string, timeoutSeconds int) (selenium.WebElement, error) {
	var el selenium.WebElement
	var err error
	for i := 0; i < timeoutSeconds*10; i++ {
		el, err = wd.FindElement(by, value)
		if err == nil {
			return el, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil, fmt.Errorf("element not found: %s=%s after %d seconds", by, value, timeoutSeconds)
}

// IsErrorOrCaptchaPage checks for common error/captcha indicators in the page source
func IsErrorOrCaptchaPage(wd selenium.WebDriver) (bool, error) {
	source, err := wd.PageSource()
	if err != nil {
		return false, err
	}
	if containsAny(source, []string{"captcha", "Access Denied", "error-page", "blocked", "verify you are human"}) {
		return true, nil
	}
	return false, nil
}

func containsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if len(sub) > 0 && (len(s) > 0 && (contains(s, sub))) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(substr) > 0 && len(s) > 0 && (len(s) >= len(substr)) && (s == substr || (len(s) > len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}
