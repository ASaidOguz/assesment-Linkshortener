package entity

import "strings"

type URL struct {
	Original string
}

func (u *URL) ValidateURL() bool {
	return strings.HasPrefix(u.Original, "http://") || strings.HasPrefix(u.Original, "https://")
}
