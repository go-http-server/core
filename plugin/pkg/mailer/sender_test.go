package mailer

import (
	"testing"

	"github.com/go-http-server/core/utils"
	"github.com/stretchr/testify/require"
)

func TestSendMailWithTemplate(t *testing.T) {
	env, err := utils.LoadEnviromentVariables("./../../../")
	require.NoError(t, err)
	require.NotEmpty(t, env)

	sender := NewGmailSender(env.EMAIL_USERNAME_SENDER, env.EMAIL_ADDRESS_SENDER, env.EMAIL_PASSWORD_SENDER)
	require.NotNil(t, sender)
	to := UserReceive{
		Username:     utils.RandomString(6),
		EmailAddress: "21A100100257@students.hou.edu.vn",
		Code:         utils.RandomCode(),
		Fullname:     "Pham Hai Nam",
	}
	err = sender.SendWithTemplate("[Go core] Kích hoạt tài khoản", "../../../templates/verify_email.html", to, []string{"../../../docker-compose.yml"})
	require.NoError(t, err)
}
