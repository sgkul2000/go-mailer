package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/jordan-wright/email"

	"github.com/labstack/echo/v4"
)

// Mailer handles sending of mails
func Mailer(c echo.Context) error {
	response := struct {
		Success bool `json:"success"`
	}{
		true,
	}

	mail, err := c.FormFile("content")
	if err != nil {
		return err
	}

	// Source
	src, err := mail.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var buff bytes.Buffer
	io.Copy(&buff, src)

	emailUser := &struct {
		Username    string
		Password    string
		EmailServer string
		Port        int
	}{
		os.Getenv("MAILID"),
		os.Getenv("PASSWORD"),
		"smtp.gmail.com",
		587,
	}

	auth := smtp.PlainAuth("",
		emailUser.Username,
		emailUser.Password,
		emailUser.EmailServer,
	)

	// Get the address to send the email to
	var to []string
	err = json.Unmarshal([]byte(c.FormValue("to")), &to)
	if err != nil {
		return err
	}
	fmt.Println(to[0])
	e := &email.Email{
		To:      to,
		From:    fmt.Sprintf("%v <%v>", os.Getenv("FROM"), os.Getenv("MAILID")),
		Subject: c.FormValue("subject"),
		HTML:    buff.Bytes(),
	}

	err = e.Send(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port), auth)
	if err != nil {
		fmt.Println("Error sending mail: ", err)
		return err
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
