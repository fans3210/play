package app

import (
	"balling/domain/errs"
)

type ScoreCalculator struct {
}

func NewCoreCalculator() ScoreCalculator {
	return ScoreCalculator{}
}

func (s ScoreCalculator) Calculate(input [10][]uint) ([10]uint, error) {
	var output [10]uint
	isValid := s.isValidInput(input)

	if !isValid {
		return output, errs.ErrInvalidInput
	}

	score := [10]uint{}
	for i := len(input) - 1; i >= 0; i-- {
		frame := input[i]
		score[i] = sum(frame)

		if isFrameStrike(frame) {
			if i+1 < len(input)-1 {
				nextFrameOriginalScore := sum(input[i+1])
				score[i] += nextFrameOriginalScore
			}
		}
	}

	return [10]uint{1}, nil
}

// each frame must contain 1-2 elements for 1th to 9th frame and up to 3 for 10th frame
func (ScoreCalculator) isValidInput(ipt [10][]uint) bool {

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
