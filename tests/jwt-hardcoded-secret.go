package rules

import "os"

type signer interface{ SignedString(key interface{}) (string, error) }

func sign(t signer) {
	// ruleid: jwt-hardcoded-secret-go
	t.SignedString([]byte("my-super-secret"))

	key := []byte("another-secret")
	// ruleid: jwt-hardcoded-secret-go
	t.SignedString(key)

	// ok: jwt-hardcoded-secret-go
	t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
