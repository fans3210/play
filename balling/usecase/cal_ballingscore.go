package usecase

import "balling/domain"

// currently usecase only have api component(dependency), in the future if necessary can add more compontents:
// eg: dataStore(repo)
type calBallingScoreUseCase struct {
	input [10][]uint
	api   domain.BallingScoreAPI
}

func NewCalBallingScoreuseCase(ipt [10][]uint, api domain.BallingScoreAPI) domain.CalBallingScoreUseCase {
	return &calBallingScoreUseCase{
		input: ipt,
		api:   api,
	}
}

func (uc *calBallingScoreUseCase) Run() ([10]uint, error) {
	return uc.api.Calculate(uc.input)
}
