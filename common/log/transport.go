package log

type Transport interface {
	Write(string, string)
}
