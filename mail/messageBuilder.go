package mail

import (
	"strings"

	"github.com/open-function-computers-llc/secret-share/config"
)

var message = `
<p>A new secret is ready for you to view</p>
<p>View your secret here: <a href="%%URL%%">Secret</a></p>
`

func notificiationEmail(c config.Config, key string) string {
	return strings.ReplaceAll(message, "%%URL%%", c.BaseURL+"/show/"+key)
}
