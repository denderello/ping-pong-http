package log

type WriterBridge struct {
	Logger Logger
}

func (lb *WriterBridge) Write(p []byte) (n int, err error) {
	lb.Logger.Info(string(p[:]))
	return len(p), nil
}
