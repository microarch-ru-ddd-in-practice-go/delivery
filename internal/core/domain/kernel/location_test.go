package kernel

import (
	"delivery/pkg/constructor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocation(t *testing.T) {
	tests := []struct {
		name       string
		vertical   int
		horizontal int
		want       *Location
		wantErrs   []error
	}{
		{
			name:       "Малая вертикаль",
			vertical:   -1,
			horizontal: 5,
			want:       nil,
			wantErrs:   []error{ErrLocationInvalidVertical},
		},
		{
			name:       "Большая вертикаль",
			vertical:   11,
			horizontal: 5,
			want:       nil,
			wantErrs:   []error{ErrLocationInvalidVertical},
		},
		{
			name:       "Малая горизонталь",
			vertical:   5,
			horizontal: -1,
			want:       nil,
			wantErrs:   []error{ErrLocationInvalidHorizontal},
		},
		{
			name:       "Большая горизонталь",
			vertical:   5,
			horizontal: 11,
			want:       nil,
			wantErrs:   []error{ErrLocationInvalidHorizontal},
		},
		{
			name:       "Не валидная вертикаль и не валидная горизонталь",
			vertical:   -1,
			horizontal: -1,
			want:       nil,
			wantErrs:   []error{ErrLocationInvalidVertical, ErrLocationInvalidHorizontal},
		},
		{
			name:       "Валидный вызов",
			vertical:   5,
			horizontal: 5,
			want: &Location{
				vertical:   5,
				horizontal: 5,
			},
			wantErrs: nil,
		},
	}

	for _, tt := range tests {
		result, err := NewLocation(tt.vertical, tt.horizontal)

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, result, tt.want)

			for _, wantErr := range tt.wantErrs {
				assert.ErrorIs(t, err, wantErr)
			}
		})
	}
}

func TestNewRandomLocation(t *testing.T) {
	t.Run("Генерация рандомного валидного Location (проверяем отсутствие паники)", func(t *testing.T) {
		for i := 0; i <= 50; i++ {
			var result *Location

			assert.NotPanics(t, func() { result = NewRandomLocation() })
			assert.NotNil(t, result)
		}
	})
}

func TestLocation_Distance(t *testing.T) {
	tests := []struct {
		name         string
		loc1         *Location
		loc2         *Location
		wantDistance int
	}{
		{
			name:         "Расчет дистанции для nil и валидного Location (проверяем отсутствие паники)",
			loc1:         nil,
			loc2:         constructor.MustBeValid(NewLocation(10, 1)),
			wantDistance: 11,
		},
		{
			name:         "Расчет дистанции для валидного и nil Location (проверяем отсутствие паники)",
			loc1:         constructor.MustBeValid(NewLocation(1, 10)),
			loc2:         nil,
			wantDistance: 11,
		},
		{
			name:         "Расчет дистанции для валидных Location",
			loc1:         constructor.MustBeValid(NewLocation(1, 10)),
			loc2:         constructor.MustBeValid(NewLocation(10, 1)),
			wantDistance: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantDistance, tt.loc1.Distance(tt.loc2))
			assert.Equal(t, tt.wantDistance, tt.loc2.Distance(tt.loc1))
		})
	}
}

func TestLocation_Equals(t *testing.T) {
	tests := []struct {
		name       string
		loc1       *Location
		loc2       *Location
		wantEquals bool
	}{
		{
			name:       "Расчет равенства для валидного и nil Location (проверяем отсутствие паники)",
			loc1:       nil,
			loc2:       constructor.MustBeValid(NewLocation(10, 1)),
			wantEquals: false,
		},
		{
			name:       "Расчет равенства для nil и валидного Location (проверяем отсутствие паники)",
			loc1:       constructor.MustBeValid(NewLocation(10, 1)),
			loc2:       nil,
			wantEquals: false,
		},
		{
			name:       "Расчет равенства для разных валидных Location",
			loc1:       constructor.MustBeValid(NewLocation(1, 10)),
			loc2:       constructor.MustBeValid(NewLocation(10, 1)),
			wantEquals: false,
		},
		{
			name:       "Расчет равенства для одинаковых валидных Location",
			loc1:       constructor.MustBeValid(NewLocation(1, 10)),
			loc2:       constructor.MustBeValid(NewLocation(1, 10)),
			wantEquals: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantEquals, tt.loc1.Equals(tt.loc2))
			assert.Equal(t, tt.wantEquals, tt.loc2.Equals(tt.loc1))
		})
	}
}
