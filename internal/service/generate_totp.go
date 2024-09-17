package service

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"inverntory_management/internal/exception"

	"github.com/pquerna/otp/totp"
)

// // GenerateTOTPSecret implements AuthServiceImpl.
func GenerateTOTPSecret(email string) (string, string, error) {
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "InventoryManagement",
		AccountName: email,
	})

	if err != nil {
		return "", "", exception.ErrInternal
	}

	qrCode, err := secret.Image(200, 200)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	png.Encode(&buf, qrCode)

	qrCodeBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	qrCodeImage := "data:image/png;base64," + qrCodeBase64

	return secret.Secret(), qrCodeImage, nil
}
