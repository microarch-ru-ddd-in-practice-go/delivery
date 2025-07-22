package kernel

import (
	"delivery/pkg/constructor"
	"errors"
	"math/rand"
)

const (
	LocationMinHorizontal = 1
	LocationMaxHorizontal = 10

	LocationMinVertical = 1
	LocationMaxVertical = 10
)

type Location struct {
	vertical   int
	horizontal int
}

var (
	ErrLocationInvalidVertical   = errors.New("kernel.location.invalid_vertical")
	ErrLocationInvalidHorizontal = errors.New("kernel.location.invalid_horizontal")
)

func NewLocation(vertical, horizontal int) (_ *Location, err error) {
	if LocationMinVertical > vertical || vertical > LocationMaxVertical {
		err = errors.Join(err, ErrLocationInvalidVertical)
	}

	if LocationMinHorizontal > horizontal || horizontal > LocationMaxHorizontal {
		err = errors.Join(err, ErrLocationInvalidHorizontal)
	}

	if err != nil {
		return nil, err
	}

	return &Location{
		vertical:   vertical,
		horizontal: horizontal,
	}, nil
}

func NewRandomLocation() *Location {
	var (
		vertical   = rand.Intn(LocationMaxVertical-LocationMinVertical) + LocationMinVertical
		horizontal = rand.Intn(LocationMaxHorizontal-LocationMinHorizontal) + LocationMinHorizontal
	)

	return constructor.MustBeValid(NewLocation(vertical, horizontal))
}

func (loc *Location) Vertical() int {
	if loc == nil {
		return 0
	}

	return loc.vertical
}

func (loc *Location) Horizontal() int {
	if loc == nil {
		return 0
	}

	return loc.horizontal
}

func (loc *Location) Equals(loc2 *Location) bool {
	return loc.Vertical() == loc2.Vertical() &&
		loc.Horizontal() == loc2.Horizontal()
}

func (loc *Location) Distance(loc2 *Location) int {
	horizontal := loc.Horizontal() - loc2.Horizontal()
	if horizontal < 0 {
		horizontal *= -1
	}

	vertical := loc.Vertical() - loc2.Vertical()
	if vertical < 0 {
		vertical *= -1
	}

	return horizontal + vertical
}
