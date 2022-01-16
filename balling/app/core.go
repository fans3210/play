package app

import (
	dm "balling/domain"
	"balling/domain/errs"
)

type ScoreCalculator struct {
}

func NewCoreCalculator() ScoreCalculator {
	return ScoreCalculator{}
}

func (s ScoreCalculator) Calculate(ipt dm.Game) (dm.Scores, error) {
	var output dm.Scores
	isValid := s.isValidInput(ipt)

	if !isValid {
		return output, errs.ErrInvalidInput
	}

	// calculate pure scores of each frame
	for i := len(ipt) - 1; i >= 0; i-- {
		output[i] = sum(ipt[i])

		if isFrameStrike(ipt[i]) {
			if i+1 < len(ipt)-1 {
				nextFrameScore := sum(ipt[i+1])
				output[i] += nextFrameScore
			}

			if i+2 < len(ipt)-1 && isFrameStrike(ipt[i+1]) {
				frameAfterNextScore := sum(ipt[i+2])
				output[i] += frameAfterNextScore
			}
		}
	}

	// calculate accumulated scores of each frame
	for i := 1; i < len(output); i++ {
		output[i] += output[i-1]
	}

	return output, nil
}

// each frame must contain 1-2 elements for 1th to 9th frame and up to 3 for 10th frame
func (ScoreCalculator) isValidInput(ipt dm.Game) bool {

	for idx, frame := range ipt {

		if len(frame) <= 0 {
			return false
		}

		if idx < len(ipt)-1 && len(frame) >= 3 {
			return false
		}

		if idx == len(ipt)-1 && len(frame) >= 4 {
			return false
		}

		if idx < len(ipt)-1 && len(frame) == 1 && !isFrameStrike(frame) {
			return false
		}

		if idx == len(ipt)-1 && len(frame) == 3 && !isFrameStrike(frame) {
			return false
		}

		if idx == len(ipt)-1 && len(frame) < 3 && isFrameStrike(frame) {
			return false
		}

		// frame 1-9, sum of throws should be 0-10
		if idx < len(ipt)-1 && sum(frame) > 10 {
			return false
		}

		// frame 10, sum of throws should be 0-30
		if idx == len(ipt)-1 && sum(frame) > 30 {
			return false
		}
	}
	return true
}

// without considering whether frame is last frame
func isFrameStrike(f []uint) bool {
	if len(f) <= 0 || len(f) > 3 {
		panic("invalid frame")
	}

	return f[0] == 10
}

func sum(nums []uint) (ret uint) {
	for _, v := range nums {
		ret += v
	}
	return
}
