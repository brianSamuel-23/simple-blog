package sanitizer

import "simple-blog.com/impl/sanitizer/internal"

type Sanitizer interface {
	Sanitize(html string) string
}

func New() Sanitizer {
	return internal.NewSanitizer()
}
