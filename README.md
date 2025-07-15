# Magento User Flow Automation (Go + Selenium)

## Overview
This project automates the critical user account flows on the Magento e-commerce demo site (https://magento.softwaretestingboard.com/) using Go and Selenium WebDriver. It is designed for QA assessment and covers:

- Sign up flow
- Login with the same account
- Sign out
- Change password

All flows are tested in a single, robust end-to-end scenario.

## Project Structure
```
Automation/
├── README.md
├── .gitignore
├── go.mod
├── go.sum
├── testcases.xlsx
├── main.go
├── pages/
│   ├── signup.go
│   ├── login.go
│   ├── home.go
│   └── changepassword.go
├── utils/
│   └── webdriver.go
├── tests/
│   ├── e2e_flow_test.go
│   └── screenshots/
│        ├── e2e_signup.png
│        ├── e2e_login.png
│        ├── e2e_signout1.png
│        ├── e2e_signout2.png
│        └── e2e_changepassword.png
```

## Setup Instructions

### Prerequisites
- Go 1.18 or higher
- Chrome browser
- ChromeDriver (matching your Chrome version)
- Git

### Steps
1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd Automation
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Download ChromeDriver:**
   - Download from: https://chromedriver.chromium.org/downloads
   - Place `chromedriver.exe` in your project root or ensure it’s in your system PATH.

## Running the End-to-End Test

Run the following command to execute the full user flow:
```sh
go test -v ./tests/e2e_flow_test.go
```

- This will run the sign up, login, sign out, and change password flows in a single browser session.
- Screenshots for each step will be saved in `tests/screenshots/` for proof of execution.

## Proof of Execution
Below are example screenshots captured during automated test runs:

### Sign Up
![Sign Up](tests/screenshots/e2e_signup.png)

### Login
![Login](tests/screenshots/e2e_login.png)

### Sign Out (after signup)
![Sign Out 1](tests/screenshots/e2e_signout1.png)

### Sign Out (after login)
![Sign Out 2](tests/screenshots/e2e_signout2.png)

### Change Password
![Change Password](tests/screenshots/e2e_changepassword.png)

## Test Cases
- All test cases are documented in `testcases.xlsx`.

## Notes
- Uses Page Object Model (POM) for maintainability.
- All code is robust, clear, and focused on the assessment objectives.
- Only essential files are included for submission.

## License
This project is for educational and assessment purposes.