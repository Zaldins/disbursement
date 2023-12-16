package helpers

func ErrorWrapper(errInfo []string, errorDescription string) []string {
	return append(errInfo, errorDescription)
}