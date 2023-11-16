package utils

// usage
// isEmpty := IF(str=="", "true", "false")
// isEmpty = "true"

// IF 函数实现三元运算符 (cond ? true : false)
func IF[T any](cond bool, trueResult, falseResult T) T {
	if cond {
		return trueResult
	}

	return falseResult
}
