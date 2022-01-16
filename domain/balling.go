package domain

type BallingScoreAPI interface {
	Calculate(ipt [10][]uint) ([10]uint, error)
}

type CalBallingScoreUseCase interface {
	Run() ([10]uint, error)
}

type CalBallingScoreUseCaseFactory interface {
	MakeCalculateBallingScoreUseCase(input [10][]uint) CalBallingScoreUseCase
}
