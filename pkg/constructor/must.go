package constructor

func MustBeValid[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}

	return result
}
