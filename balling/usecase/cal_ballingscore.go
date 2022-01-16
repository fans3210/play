package usecase

import (
	dm "balling/domain"
)

// currently usecase only have api component(dependency), in the future if necessary can add more compontents:
// eg: dataStore(repo)

type calBallingScoreUseCase struct {
	input dm.Game
	api   dm.BallingScoreAPI
}

func NewCalBallingScoreuseCase(ipt dm.Game, api dm.BallingScoreAPI) dm.CalBallingScoreUseCase {
	return &calBallingScoreUseCase{
		input: ipt,
		api:   api,
	}
}

func (uc *calBallingScoreUseCase) Run() (dm.Scores, error) {
	return uc.api.Calculate(uc.input)
}
