package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	classroom "google.golang.org/api/classroom/v1"
	"github.com/gen2brain/beeep"
)

const credentialFile = "credentials.json"
const tokenFile = "token.json"

func getClient(config *oauth2.Config) *http.Client {
	tok, err := tokenFromFile(tokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, classroom.ClassroomCoursesReadonlyScope, classroom.ClassroomAnnouncementsReadonlyScope, classroom.ClassroomCourseworkMeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := classroom.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Classroom client: %v", err)
	}

	courses, err := srv.Courses.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses: %v", err)
	}

	fmt.Println("Courses:")
	for _, course := range courses.Courses {
		fmt.Printf("%s (%s)\n", course.Name, course.Id)
	}

	lastAnnouncementTime := time.Now()

	for {
		for _, course := range courses.Courses {
			announcements, err := srv.Courses.Announcements.List(course.Id).OrderBy("updateTime desc").Do()
			if err != nil {
				log.Fatalf("Unable to retrieve announcements: %v", err)
			}

			for _, announcement := range announcements.Announcements {
				announcementTime, err := time.Parse(time.RFC3339, announcement.UpdateTime)
				if err != nil {
					log.Fatalf("Unable to parse time: %v", err)
				}

				if announcementTime.After(lastAnnouncementTime) {
					beeep.Notify("New Announcement", fmt.Sprintf("Course: %s\nAnnouncement: %s", course.Name, announcement.Text), "")
				}
			}

			assignments, err := srv.Courses.CourseWork.List(course.Id).OrderBy("updateTime desc").Do()
			if err != nil {
				log.Fatalf("Unable to retrieve coursework: %v", err)
			}

			for _, assignment := range assignments.CourseWork {
				assignmentTime, err := time.Parse(time.RFC3339, assignment.UpdateTime)
				if err != nil {
					log.Fatalf("Unable to parse time: %v", err)
				}

				if assignmentTime.After(lastAnnouncementTime) {
					beeep.Notify("New Assignment", fmt.Sprintf("Course: %s\nAssignment: %s", course.Name, assignment.Title), "")
				}
			}
		}

		lastAnnouncementTime = time.Now()
		time.Sleep(60 * time.Second)
	}
}
