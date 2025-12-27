package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60 // количество метров в километре.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста
	walkingCaloriesCoefficient = 0.5 // коэффициент для расчета калорий при ходьбе
)

func Distance(steps int, height float64) float64 {
// TODO: реализовать функцию
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
// TODO: реализовать функцию
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	hours := duration.Hours()
	return dist / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
// TODO: реализовать функцию	
if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input data")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	return (weight * speed * minutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
// TODO: реализовать функцию	
if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input data")
	}
	speed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	cal := (weight * speed * minutes) / minInH
	return cal * walkingCaloriesCoefficient, nil
}
