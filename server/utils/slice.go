package utils

func Filter[T any](slice []T, fn func(T) bool) []T {
	filtered := []T{}

	for _, v := range slice {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
