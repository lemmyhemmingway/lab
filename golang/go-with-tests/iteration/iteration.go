package iteration

const repeatCount = 5

func Repeat(s string) string  {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		repeated += s

	}
	return repeated
}
