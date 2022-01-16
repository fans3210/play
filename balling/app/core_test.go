package app_test

import (
	"balling/balling/app"
	dm "balling/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cal = app.NewCoreCalculator()

func TestFrameValidation(t *testing.T) {
	input := dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}
	assert := assert.New(t)
	var ipt dm.Game
	var err error

	ipt = input
	_, err = cal.Calculate(ipt)
	assert.NoError(err)

	// frame 1-10, num throws = 0
	ipt = input
	ipt[2] = make([]uint, 0)
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 1-9, num throws >=3,
	ipt = input
	ipt[1] = []uint{1, 2, 3}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 10 num throws >= 4
	ipt = input
	ipt[9] = []uint{1, 2, 3, 4}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 1-9, num throws = 1 but 1st throw is not strike
	ipt = input
	ipt[2] = []uint{5}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 10, num throws = 3 but 1st throw is not strike
	ipt = input
	ipt[9] = []uint{1, 2, 3}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 10, num throws < 3 but first throw is strike
	// The last frame has three throws only if a bowler makes a strike on the first throw. means 10, 10, 10 is valid
	ipt = input
	ipt[9] = []uint{10, 9}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 1-9, score sum > 10(maximum)
	ipt = input
	ipt[1] = []uint{6, 9}
	_, err = cal.Calculate(ipt)
	assert.Error(err)

	// frame 10, score sum > 30(maximum)
	ipt = input
	ipt[9] = []uint{10, 10, 11}
	_, err = cal.Calculate(ipt)
	assert.Error(err)
}

func TestCal(t *testing.T) {
	// spare samples
	ipt := dm.Game{{5, 2}, {8, 1}, {6, 4}, {0, 0}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {2, 6}}

	output, err := cal.Calculate(ipt)
	assert.NoError(t, err)

	expected := dm.Scores{7, 16, 26, 26, 31, 39, 48, 56, 63, 71}
	assert.Equal(t, output, expected)

	// provided samples
	ipt = dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {0, 5}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}

	output, err = cal.Calculate(ipt)
	assert.NoError(t, err)

	expected = dm.Scores{7, 16, 26, 41, 46, 54, 63, 71, 78, 96}
	assert.Equal(t, output, expected)

	// consecutive strike case
	ipt = dm.Game{{5, 2}, {8, 1}, {6, 4}, {10}, {10}, {2, 6}, {8, 1}, {5, 3}, {6, 1}, {10, 2, 6}}
	output, err = cal.Calculate(ipt)
	assert.NoError(t, err)

	expected = dm.Scores{7, 16, 26, 54, 72, 80, 89, 97, 104, 122}
	assert.Equal(t, output, expected)

}
