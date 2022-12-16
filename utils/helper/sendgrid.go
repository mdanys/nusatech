package helper

import (
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func helloEmail(address string, name string) []byte {
	sender := "mochammaddany@gmail.com"
	nama := "Dany"
	from := mail.NewEmail(nama, sender)
	subject := "Email Verification"
	to := mail.NewEmail(name, address)
	content := mail.NewContent("text/plain", "please verify your account")
	m := mail.NewV3MailInit(from, subject, to, content)
	email := mail.NewEmail(name, address)
	m.Personalizations[0].AddTos(email)
	return mail.GetRequestBody(m)
}

func SendHelloEmail(address string, name string) *rest.Response {
	request := sendgrid.GetRequest(os.Getenv("YOUR_SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = helloEmail(address, name)
	request.Body = Body
	response, _ := sendgrid.API(request)

	return response
}
