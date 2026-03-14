package emailtemplates

import (
	"bytes"
	"text/template"
)

const AccountCreatedSubject = "Account created"

const accountCreatedTemplate = `Hello {{.Email}},
                                     Thank you for creating your account.
                                     If you did not create this account, you can safely ignore this email.
                                     Best regards.
`

type AccountCreatedData struct {
	Email string
}

func RenderAccountCreated(data AccountCreatedData) (string, error) {
	tmpl, err := template.New("account_created").Parse(accountCreatedTemplate)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
