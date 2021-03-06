package card

type CardRepository interface {
	Generate(card *Card) error
	GetByToken(token string) (*Card, error)
}
