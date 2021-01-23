# Go-Mailer

Go-Mailer is a simple go microservice that can be used to send emails to multiple users using Gmail account (potentially any email address as it creates a smtp server).

## Project setup

This project requires three environment variables namely:

- `MAILID` : The email id that you want to send **from**
- `PASSWORD` : The password of the above email address
- `FROM` : Name of the sender (This makes the email prettier)

once this is done you can simply run `go run main.go`

## Endpoints

- [`GET /`](#get)
- [`POST /mail`](#post-mail)

This microservice has only two endpoints (actually only one, but the other one s just to test if the app is working).

### `GET - /`

> An endpoint to test if the app is running and potentially wake it up if its running on a serverless.

### `POST /mail`

> An endpoint to actually send mail.

**Parameters** (Remember all the data passed should be in formdata and is required):

- `content` (_File_) : The actual content of the email (either an html or a txt file)
- `to` (_[]string_) : The list of addresses to which the email is to be sent to (A list of strings).
- `subject`(_string_) : The subject of the email (string).


_Feel free to contribute. Enjoy!_