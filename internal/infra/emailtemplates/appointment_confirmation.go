package emailtemplates

import (
	"bytes"
	"text/template"
)

const AppointmentConfirmationSubject = "Appointment Confirmation"

const appointmentConfirmationTemplate = `Hello {{.Name}}, 
                                         Your appointment has been successfully scheduled. 
                                         start time: {{.StartTime}}
                                         end time: {{.EndTime}}
                                        If you need to reschedule or cancel, please reply to this email or contact our support team.
                                        Thank you.
`

type AppointmentConfirmationData struct {
	Name      string
	StartTime string
	EndTime   string
}

func RenderAppointmentConfirmation(data AppointmentConfirmationData) (string, error) {
	tmpl, err := template.New("appointment_confirmation").Parse(appointmentConfirmationTemplate)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
