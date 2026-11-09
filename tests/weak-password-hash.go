package rules

import (
	"crypto/md5"
	"crypto/sha256"
)

func hashes(p []byte) {
	// ruleid: weak-password-hash-go
	_ = md5.Sum(p)

	// ok: weak-password-hash-go
	_ = sha256.Sum256(p)
}
