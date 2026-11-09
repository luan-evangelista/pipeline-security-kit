package rules

import "crypto/tls"

func clients() (*tls.Config, *tls.Config) {
	// ruleid: tls-verification-disabled-go
	bad := &tls.Config{MinVersion: tls.VersionTLS12, InsecureSkipVerify: true}

	// ok: tls-verification-disabled-go
	good := &tls.Config{MinVersion: tls.VersionTLS12}
	return bad, good
}
