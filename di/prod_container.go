package di

import (
	"balling/balling/app"
	"balling/balling/usecase"
	dm "balling/domain"
)

type prodContainer struct {
	ballingScoreAPI dm.BallingScoreAPI
}

func NewProductionContainer() *prodContainer {
	return &prodContainer{
		ballingScoreAPI: app.NewCoreCalculator(),
	}
}

func (c *prodContainer) MakeCalculateBallingScoreUseCase(input dm.Game) dm.CalBallingScoreUseCase {
	uc := usecase.NewCalBallingScoreuseCase(input, c.ballingScoreAPI)
	return uc
}
