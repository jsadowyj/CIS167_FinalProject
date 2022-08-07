package ansi

func Red(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}

func Blue(s string) string {
	return "\x1b[34m" + s + "\x1b[0m"
}

func Green(s string) string {
	return "\x1b[32m" + s + "\x1b[0m"
}

func Yellow(s string) string {
	return "\x1b[33m" + s + "\x1b[0m"
}

func Bold(s string) string {
	return "\x1b[1m" + s + "\x1b[0m"
}

func Faint(s string) string {
	return "\x1b[2m" + s + "\x1b[0m"
}
