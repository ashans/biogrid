package algo

import "biogrid/internal/entities"

func matchValue(a string, b string, config entities.AlignConfig) int {
	if a == b {
		return config.Match
	}
	return config.Mismatch
}

func max(a ...int) (r int) {
	r = a[0]
	for _, v := range a {
		if r < v {
			r = v
		}
	}
	return r
}
