package card

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func MaskPANNumber(pan string) string {
	return pan[:4] + strings.Repeat("*", len(pan)-8) + pan[len(pan)-4:]
}

func getCardExpiration() (int, int) {
	EXPIRATION_STEP_YEARS := 5 // in years
	now := time.Now()
	year := now.Year() + EXPIRATION_STEP_YEARS
	month := int(now.Month())
	return year, month
}

func generateToken() string {
	token := uuid.NewString() + uuid.NewString()
	return strings.ReplaceAll(token, "-", "")
}
