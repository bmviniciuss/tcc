package postgrescardrepository

import (
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"gorm.io/gorm"
)

type PostgresCard struct {
	gorm.Model
	Id             string `gorm:"primaryKey"`
	Number         string `gorm:"type:varchar(16);not null"`
	Cvv            string `gorm:"type:varchar(3);not null"`
	CardholderName string `gorm:"type:varchar(256);not null"`
	Token          string `gorm:"type:varchar(256);not null"`
	MaskedNumber   string `gorm:"type:varchar(16);not null"`
	Active         bool   `gorm:"type:boolean;default:true"`
	IsCredit       bool   `gorm:"type:boolean;default:true"`
	IsDebit        bool   `gorm:"type:boolean;default:true"`
}

func (PostgresCard) TableName() string {
	return "cards"
}

type postgresCardRepository struct {
	db gorm.DB
}

func NewPostgresCardRepository(db *gorm.DB) *postgresCardRepository {
	return &postgresCardRepository{
		db: *db,
	}
}

func (r *postgresCardRepository) Generate(generateCardDTO *card.GenerateCardRepoInput) (*card.Card, error) {
	pgCard := &PostgresCard{
		Number:         generateCardDTO.Number,
		Cvv:            generateCardDTO.Cvv,
		CardholderName: generateCardDTO.CardholderName,
		Token:          generateCardDTO.Token,
		MaskedNumber:   generateCardDTO.MaskedNumber,
		Active:         generateCardDTO.Active,
		IsCredit:       generateCardDTO.IsCredit,
		IsDebit:        generateCardDTO.IsDebit,
	}

	result := r.db.Create(pgCard)

	if result.Error != nil {
		return nil, result.Error
	}

	return &card.Card{
		Id:             pgCard.Id,
		Number:         pgCard.Number,
		Cvv:            pgCard.Cvv,
		CardholderName: pgCard.CardholderName,
		Token:          pgCard.Token,
		MaskedNumber:   pgCard.MaskedNumber,
		Active:         pgCard.Active,
		IsCredit:       pgCard.IsCredit,
		IsDebit:        pgCard.IsDebit,
	}, nil
}
