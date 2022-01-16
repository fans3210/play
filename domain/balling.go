package domain

type Game = [10][]uint
type Scores = [10]uint

type BallingScoreAPI interface {
	Calculate(ipt Game) (Scores, error)
}

type CalBallingScoreUseCase interface {
	Run() (Scores, error)
}

type CalBallingScoreUseCaseFactory interface {
	MakeCalculateBallingScoreUseCase(input Game) CalBallingScoreUseCase
}
