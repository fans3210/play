package di

import (
	"balling/balling/app"
	"balling/balling/usecase"
	"balling/domain"
)

type prodContainer struct {
	ballingScoreAPI domain.BallingScoreAPI
}

func NewProductionContainer() *prodContainer {
	return &prodContainer{
		ballingScoreAPI: app.NewCoreCalculator(),
	}
}

func (c *prodContainer) MakeCalculateBallingScoreUseCase(input [10][]uint) domain.CalBallingScoreUseCase {
	uc := usecase.NewCalBallingScoreuseCase(input, c.ballingScoreAPI)
	return uc
}
