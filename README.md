# Google Classroom Notifier

Google Classroom Notifier is a desktop application built with Go and the Fyne GUI framework that notifies you of new announcements and assignments posted in your Google Classroom courses.

## Features

- Real-time notifications for new announcements and assignments
- Desktop notifications using `beeep`
- Simple and user-friendly GUI using the Fyne framework

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.16 or later)
- Fyne (version 2.0 or later) (*only for desktop GUI app*)

## Setup

### Step 1: Create a Google Cloud Project

1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Click on the project drop-down and select "New Project".
3. Give your project a name and click "Create".

### Step 2: Enable Google Classroom API

1. In the Google Cloud Console, navigate to `APIs & Services > Library`.
2. Search for "Google Classroom API" and click on it.
3. Click "Enable" to enable the API for your project.

### Step 3: Create OAuth 2.0 Credentials

1. Navigate to `APIs & Services > Credentials`.
2. Click on "Create Credentials" and select "OAuth client ID".
3. Configure the consent screen by providing necessary information like application name and email.
4. After configuring the consent screen, select "Desktop app" as the application type.
5. Download the `credentials.json` file and save it in your project directory.

### Step 4: Clone the repository

```bash
git clone https://github.com/sanjibdahal/google-classroom-notifier.git
```

### Step 5: Install Required Go Packages

```bash
go get fyne.io/fyne/v2
go install fyne.io/fyne/v2/cmd/fyne@latest
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/classroom/v1
go get github.com/gen2brain/beeep
```
Place your credentials.json file in the same directory as main.go or in desktop-app-gui folder for desktop app.

### Step 6: Run the application

```bash
go run main.go
```

### Step 7: Authorize the Application

1. When you run the application for the first time, it will print a URL in the terminal. Open this URL in your web browser.
2. Log in with your Google account and authorize the application to access your Google Classroom data.
3. You will be given an authorization code. Copy this code and paste it back into your terminal where your application is running.
4. The application will exchange the authorization code for tokens and save them in a file named 'token.json'.

**Now, you will get notifications in the desktop whenever new announcements are posted.**