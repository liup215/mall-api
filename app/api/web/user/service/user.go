package service

import (
	"fmt"
	"mall/lib/strings"
)

func (s *Service) userHash(password, salt string) string {
	return strings.Sha1(fmt.Sprintf("%s-%s-%s", password, salt, s.authKey))
}
