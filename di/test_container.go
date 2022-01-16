package di

import (
	"balling/balling/app"
	"balling/balling/usecase"
	"balling/domain"
)

type testContainer struct {
	ballingScoreAPI domain.BallingScoreAPI
}

func NewTestContainer() *testContainer {
	return &testContainer{
		ballingScoreAPI: app.NewCoreCalculator(),
	}
}

func (c *testContainer) MakeCalculateBallingScoreUseCase(input [10][]uint) domain.CalBallingScoreUseCase {
	uc := usecase.NewCalBallingScoreuseCase(input, c.ballingScoreAPI)
	return uc
}
