package tools


type MySet map[any]struct{}

// 切片去重 map方式
func DuplicateRemovingMap[T any](s []T) []T {
	res := make([]T, 0, len(s))
	mySet := make(MySet)
	for _, t := range s {
	  if _, ok := mySet[t]; !ok {
		res = append(res, t)
		mySet[t] = struct{}{}
	  }
	}
	return res
}