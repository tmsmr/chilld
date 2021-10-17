package fancurve

import "math"

const (
	coolingStart = 35
)

// LinearFanSpeedFor calculates the fan speed in % for a given CPU temperature.
// For temperatures < coolingStart, 0% is returned. This only makes sense for a fan that is able to stop.
// The target speed is calculated using a simple linear function y = mx + b (Capped to 100%).
// To avoid frequent RPM changes, the result is rounded to the nearest 5%.
func LinearFanSpeedFor(temp float64) int {
	if temp < coolingStart {
		return 0
	}
	// 0% @ 20°C
	// 20% @ 35°C
	// 100% @ 75°C
	speed := 2*temp - 50
	if speed > 100 {
		return 100
	}
	speed = math.Round(speed/5) * 5
	return int(speed)
}
