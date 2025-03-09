package utils

type FilterError interface {
	ErrEmptyList() string
	ErrNotFound(params ...any) string
	// ErrNotFound(T any) string
}

func Filter[T FilterError](ipList []T, cond func(t T) bool) ([]T, string) {
	opList := []T{}
	if (len(ipList) > 0) && (cap(ipList) > 0) {
		for _, v := range ipList {
			if cond(v) {
				opList = append(opList, v)
			}
		}
		if (len(opList) > 0) && (cap(opList) > 0) {
			return opList, ""
		} else {
			var tErr T
			return []T{}, tErr.ErrNotFound()
		}
	} else {
		var tErr T
		return []T{}, tErr.ErrEmptyList()
	}
}

func First[T FilterError](ipList []T) (T, string) {
	if (len(ipList) > 0) && (cap(ipList) > 0) {
		return ipList[0], ""
	} else {
		var none T
		return none, none.ErrEmptyList()
	}
}
