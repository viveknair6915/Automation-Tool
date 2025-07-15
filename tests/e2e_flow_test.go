package tests

import (
	"automation/pages"
	"automation/utils"
	"testing"
)

func TestEndToEndUserFlow(t *testing.T) {
	wd, err := utils.StartWebDriver()
	if err != nil {
		t.Fatalf("Failed to start WebDriver: %v", err)
	}
	defer wd.Quit()

	// 1. Sign up flow
	signup := pages.NewSignupPage(wd)
	if err := signup.Open(); err != nil {
		t.Fatalf("Failed to open signup page: %v", err)
	}
	firstName := "Test"
	lastName := "User"
	email := pages.GenerateRandomEmail()
	password := "Test@12345"
	if err := signup.Register(firstName, lastName, email, password); err != nil {
		t.Fatalf("Signup failed: %v", err)
	}
	if err := utils.TakeScreenshot(wd, "signup.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}

	// 2. Login with same account
	home := pages.NewHomePage(wd)
	if err := home.SignOut(); err != nil {
		t.Fatalf("Sign out after signup failed: %v", err)
	}

	login := pages.NewLoginPage(wd)
	if err := login.Open(); err != nil {
		t.Fatalf("Failed to open login page: %v", err)
	}
	if err := login.Login(email, password); err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	loggedIn, _ := home.IsLoggedIn()
	if !loggedIn {
		t.Fatalf("Login did not succeed")
	}

	// 3. Change password
	change := pages.NewChangePasswordPage(wd)
	if err := change.Open(); err != nil {
		t.Fatalf("Failed to open change password page: %v", err)
	}
	newPassword := "Test@54321"
	if err := change.ChangePassword(password, newPassword); err != nil {
		t.Fatalf("Change password failed: %v", err)
	}
	if err := utils.TakeScreenshot(wd, "changepassword.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}
}
