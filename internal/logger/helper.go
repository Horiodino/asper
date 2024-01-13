package logger

type Asperstring *string

func String(s string) Asperstring {
	return Asperstring(&s)
}
