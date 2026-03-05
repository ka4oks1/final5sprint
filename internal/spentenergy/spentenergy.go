package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	switch {
	case weight <= 0:
		return 0, errors.New("incorrect weight")

	case steps <= 0:
		return 0, errors.New("incorrect steps")

	case height <= 0:
		return 0, errors.New("incorrect height")

	case duration <= 0:
		return 0, errors.New("incorrect duration")

	}

	meanSpeed := MeanSpeed(steps, height, duration)

	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH

	walkingSpentCalories := spentCalories * walkingCaloriesCoefficient

	return walkingSpentCalories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	switch {
	case weight <= 0:
		return 0, errors.New("incorrect weight")

	case steps <= 0:
		return 0, errors.New("incorrect steps")

	case height <= 0:
		return 0, errors.New("incorrect height")

	case duration <= 0:
		return 0, errors.New("incorrect duration")

	}
	meanSpeed := MeanSpeed(steps, height, duration)

	spentCalories := (weight * meanSpeed * duration.Minutes()) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	meanSpeed := Distance(steps, height) / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {

	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength / mInKm
	return distance

}
