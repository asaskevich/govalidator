package govalidator

import (
	"context"
	"math"
)

// Abs returns absolute value of number
func Abs(value float64) float64 {
	return math.Abs(value)
}

// Sign returns signum of number: 1 in case of value > 0, -1 in case of value < 0, 0 otherwise
func Sign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// IsNegative returns true if value < 0
func IsNegative(value float64) bool {
	return value < 0
}

// IsPositive returns true if value > 0
func IsPositive(value float64) bool {
	return value > 0
}

// IsNonNegative returns true if value >= 0
func IsNonNegative(value float64) bool {
	return value >= 0
}

// IsNonPositive returns true if value <= 0
func IsNonPositive(value float64) bool {
	return value <= 0
}

// InRange returns true if value lies between left and right border
func InRangeInt(ctx context.Context, value, left, right interface{}) bool {
	value64, _ := ToInt(ctx, value)
	left64, _ := ToInt(ctx, left)
	right64, _ := ToInt(ctx, right)
	if left64 > right64 {
		left64, right64 = right64, left64
	}
	return value64 >= left64 && value64 <= right64
}

// InRangeFloat32 returns true if value lies between left and right border
func InRangeFloat32(value, left, right float32) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRangeFloat64 returns true if value lies between left and right border
func InRangeFloat64(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRange returns true if value lies between left and right border, generic type to handle int, float32 or float64, all types must the same type
func InRange(ctx context.Context, value interface{}, left interface{}, right interface{}) bool {
	switch value.(type) {
	case int:
		intValue, _ := ToInt(ctx, value)
		intLeft, _ := ToInt(ctx, left)
		intRight, _ := ToInt(ctx, right)
		return InRangeInt(ctx, intValue, intLeft, intRight)
	case float32, float64:
		intValue, _ := ToFloat(value)
		intLeft, _ := ToFloat(left)
		intRight, _ := ToFloat(right)
		return InRangeFloat64(intValue, intLeft, intRight)
	case string:
		return value.(string) >= left.(string) && value.(string) <= right.(string)
	default:
		return false
	}
}

// IsWhole returns true if value is whole number
func IsWhole(value float64) bool {
	return math.Remainder(value, 1) == 0
}

// IsNatural returns true if value is natural number (positive and whole)
func IsNatural(value float64) bool {
	return IsWhole(value) && IsPositive(value)
}
