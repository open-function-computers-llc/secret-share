package mail

import (
	"strings"
)

var message = `
<p>A new secret is ready for you to view</p>
<p>View your secret here: <a href="%%URL%%">Secret</a></p>
`

func notificiationEmail(url string) string {
	return strings.ReplaceAll(message, "%%URL%%", url)
}
