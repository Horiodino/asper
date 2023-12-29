package logger

type Asperstring string

func String(s string) Asperstring {
	return Asperstring(s)
}

type asperint int

func Int(i int) asperint {
	return asperint(i)
}
