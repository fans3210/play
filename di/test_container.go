package di

import (
	"balling/balling/app"
	"balling/balling/usecase"
	dm "balling/domain"
)

type testContainer struct {
	ballingScoreAPI dm.BallingScoreAPI
}

func NewTestContainer() *testContainer {
	return &testContainer{
		ballingScoreAPI: app.NewCoreCalculator(),
	}
}

func (c *testContainer) MakeCalculateBallingScoreUseCase(input dm.Game) dm.CalBallingScoreUseCase {
	uc := usecase.NewCalBallingScoreuseCase(input, c.ballingScoreAPI)
	return uc
}
