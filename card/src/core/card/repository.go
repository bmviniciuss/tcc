package card

type CardRepository interface {
	Generate(generateCardDTO *GenerateCardRepoInput) (*Card, error)
	GetByToken(token string) (*Card, error)
}
