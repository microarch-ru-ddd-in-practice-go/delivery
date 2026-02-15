package must

func NotNil(v any, name string) {
	if v == nil {
		panic(name + " must not be nil")
	}
}
