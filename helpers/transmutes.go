package helpers

// []*T => map[T.K]*T
func SliceToMap[T any, K comparable](list []*T, getkey func(item *T) K) map[K]*T {
	ret := make(map[K]*T)
	if len(list) == 0 {
		return ret
	}
	for _, item := range list {
		key := getkey(item)
		ret[key] = item
	}
	return ret
}
