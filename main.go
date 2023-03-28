// this is a job application project.
//the applicant fills the form and press the "apply now" button to submit.
// the result will then be sent to an email address
// response is then sent to the applicant

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

// function that sends the form data to the email address

func sendEmail(Full_Name string, Email string, D_O_B string,
	Job_Role string, Address string, City string) error {

	from := "sender@gmail.com"
	password := ""
	to := "recipient@gmail.com"
	subject := "Job Application Submission"

	// authentication

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nFull Name: %s\r\nEmail: %s\r\nDate of Birth: %s\r\nJob Role: %s\r\nAddress: %s\r\nCity: %s\r\n", from, to, subject, Full_Name, Email, D_O_B, Job_Role, Address, City)

	// send email

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func main() {

	// serving the html and css files in the root directory

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")

		// receive the form data

		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Parseform() err: %v", err)
			return
		}

		// save form data in a variable

		Full_Name := r.FormValue("full_name")
		Email := r.FormValue("email")
		D_O_B := r.FormValue("date")
		Job_Role := r.FormValue("job-role")
		Address := r.FormValue("address")
		City := r.FormValue("city")
		//CV := r.FormValue("cv")

		// function call to send form data

		err := sendEmail(Full_Name, Email, D_O_B, Job_Role, Address, City)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// print response to the applicant

		if _, err := fmt.Fprintf(w, "Application Submitted Successfully!"); err != nil {
			fmt.Fprintf(w, "sunmission error: %v", err)
		}

	})

	// start server

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
