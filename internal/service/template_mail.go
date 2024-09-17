package service

import (
	"bytes"
	"text/template"
)

type EmailTemplateData struct {
	Username string
	Message  string
}

// Load and parse the HTML email template
func GenerateEmailBody(username, message string) (string, error) {
	emailTemplate := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Your OTP Code</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                line-height: 1.6;
                color: #333;
                max-width: 600px;
                margin: 0 auto;
                padding: 20px;
            }
            .container {
                background-color: #f9f9f9;
                border-radius: 5px;
                padding: 20px;
                box-shadow: 0 2px 5px rgba(0,0,0,0.1);
            }
            h1 {
                color: #2c3e50;
                text-align: center;
            }
            .otp-code {
                font-size: 32px;
                font-weight: bold;
                text-align: center;
                letter-spacing: 5px;
                margin: 20px 0;
                color: #3498db;
            }
            .qr-code {
                text-align: center;
                margin: 20px 0;
            }
            .qr-code img {
                max-width: 200px;
                height: auto;
            }
            .instructions {
                background-color: #ecf0f1;
                border-left: 4px solid #3498db;
                padding: 10px;
                margin-top: 20px;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Your One-Time Password (OTP)</h1>
            <p>Hello {{.Username}},</p>
            <p>You have requested a one-time password for authentication. Please use the following OTP:</p>
            <div class="otp-code">[{{.Message}}]</div>
            <div class="instructions">
                <p><strong>Instructions:</strong></p>
                <ol>
                    <li>Enter the OTP code shown above in the authentication field.</li>
                    <li>Complete the authentication process within 1 minute.</li>
                </ol>
            </div>
            <p>If you didn't request this OTP, please ignore this email or contact our support team immediately.</p>
        </div>
    </body>
    </html>`

	// Parse the template
	t, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return "", err
	}

	// Data to pass to the template
	data := EmailTemplateData{
		Username: username,
		Message:  message,
	}

	// Execute the template with the data
	var body bytes.Buffer
	err = t.Execute(&body, data)
	if err != nil {
		return "", err
	}

	return body.String(), nil
}
