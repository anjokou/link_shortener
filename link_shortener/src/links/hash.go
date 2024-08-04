package links

import (
	"crypto/sha256"
	"encoding/base64"
)

type Hasher interface {
	Hash(target string) string
}

type Sha256Hasher struct {
}

func (hasher *Sha256Hasher) Hash(target string) string {
	hash := sha256.Sum256([]byte(target))
	sha := base64.URLEncoding.EncodeToString(hash[:8])
	return sha
}

func CreateSha256Hasher() *Sha256Hasher {
	return new(Sha256Hasher)
}
