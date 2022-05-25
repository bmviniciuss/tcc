package carddetails

type CardDetails struct {
	Number string
	Cvv    string
}

type GeneratorService interface {
	Generate() (*CardDetails, error)
}

type CardDetailGenerator interface {
	Generate() (*CardDetails, error)
}

type generator struct {
	cardDetailsGenerator CardDetailGenerator
}

func NewCardDetailsGeneratorService(cardDetailsGenerator CardDetailGenerator) *generator {
	return &generator{
		cardDetailsGenerator: cardDetailsGenerator,
	}
}

func (g *generator) Generate() (*CardDetails, error) {
	cardDetails, err := g.cardDetailsGenerator.Generate()

	if err != nil {
		return nil, err
	}

	return cardDetails, nil
}
