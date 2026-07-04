package internal

// Sanitize strips anything not on the allowlist (script tags, on* attrs,
// javascript: URLs, ...) so HTML authored via a WYSIWYG client can be stored
// and later output as-is without risking stored XSS.
func (s *Sanitizer) Sanitize(html string) string {
	return s.policy.Sanitize(html)
}
