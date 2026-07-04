package hasher

import "simple-blog.com/impl/hasher/internal"

type Hasher interface {
	Hash(plaintext string) (string, error)
	Compare(hash, plaintext string) error
}

func New() Hasher {
	return internal.NewBcrypt()
}
