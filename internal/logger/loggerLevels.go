package logger

type status int

const (
	INFO    = "Info"
	ERROR   = "Error"
	WARNING = "Warning"
)

func (s status) String() string {
	return [...]string{"INFO", "ERROR", "WARNING"}[s]
}

func validateLevel(level string) bool {
	return level == INFO || level == ERROR || level == WARNING
}
