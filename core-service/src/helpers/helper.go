package helpers

import (
	"errors"
	"net/mail"
	"reflect"
	"regexp"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

func IsInvitationTokenValid(regInvitation domain.RegistrationInvitation, token string) error {
	if regInvitation.Token != token {
		return errors.New("invalid token")
	}
	return nil
}

func IsValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func IsAdminOPD(au *domain.JwtCustomClaims) bool {
	return au.Role.ID == domain.RoleAdministrator || au.Role.ID == domain.RoleGroupAdmin
}

func IsSuperAdmin(au *domain.JwtCustomClaims) bool {
	return au.Role.ID == domain.RoleSuperAdmin
}

func InArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func RegexReplaceString(c echo.Context, str string, repl string) string {
	re := regexp.MustCompile(`[^ a-zA-Z0-9_]`)
	return re.ReplaceAllString(str, repl)
}

func RegexKeywordReplaceString(c echo.Context, str string, repl string) string {
	regexPattern := `[^ a-zA-Z0-9_,\.\(\)/'"\-&]` // custom regex by requirements
	re := regexp.MustCompile(regexPattern)
	return re.ReplaceAllString(str, repl)
}

func RegexCustomReplaceString(c echo.Context, str string, repl string) string {
	regexPattern := `[^ a-zA-Z0-9_,\.\'"\/\/-]`
	re := regexp.MustCompile(regexPattern)
	return re.ReplaceAllString(str, repl)
}

// Sanity slug text with standard regex rules
func RegexReplaceSlug(str string) string {
	reg := regexp.MustCompile("[^a-z0-9]+")
	str = reg.ReplaceAllString(strings.ToLower(str), "-")

	return strings.Trim(str, "-") // Remove leading and trailing
}

func IsCompletedDocumentArchive(body *domain.DocumentArchiveRequest) bool {
	return (body.Title != "" && body.Description != "" && body.Category != "")
}

func IsDisallowed(input string, disallowedKeywords []string) bool {
	for _, keyword := range disallowedKeywords {
		if strings.Contains(strings.ToLower(input), strings.ToLower(keyword)) {
			return true
		}
	}
	return false
}
