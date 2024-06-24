package functional

func Fmap[A any, B any](fn (func (A) B), s []A) []B {
	retVal := []B{}
	for i := 0; i < len(s); i++ {
		retVal = append(retVal, fn(s[i]))
	}
	return retVal
}
