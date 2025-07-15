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
	if err := utils.TakeScreenshot(wd, "e2e_signup.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}

	// 2. Sign out after signup
	home := pages.NewHomePage(wd)
	if err := home.SignOut(); err != nil {
		t.Fatalf("Sign out after signup failed: %v", err)
	}
	if err := utils.TakeScreenshot(wd, "e2e_signout1.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}

	// 3. Login with same account
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
	if err := utils.TakeScreenshot(wd, "e2e_login.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}

	// 4. Sign out again
	if err := home.SignOut(); err != nil {
		t.Fatalf("Sign out after login failed: %v", err)
	}
	if err := utils.TakeScreenshot(wd, "e2e_signout2.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}

	// 5. Login again to change password
	if err := login.Open(); err != nil {
		t.Fatalf("Failed to open login page for password change: %v", err)
	}
	if err := login.Login(email, password); err != nil {
		t.Fatalf("Login before password change failed: %v", err)
	}
	loggedIn, _ = home.IsLoggedIn()
	if !loggedIn {
		t.Fatalf("Login before password change did not succeed")
	}

	// 6. Change password
	change := pages.NewChangePasswordPage(wd)
	if err := change.Open(); err != nil {
		t.Fatalf("Failed to open change password page: %v", err)
	}
	newPassword := "Test@54321"
	if err := change.ChangePassword(password, newPassword); err != nil {
		t.Fatalf("Change password failed: %v", err)
	}
	if err := utils.TakeScreenshot(wd, "e2e_changepassword.png"); err != nil {
		t.Logf("Failed to take screenshot: %v", err)
	}
}
