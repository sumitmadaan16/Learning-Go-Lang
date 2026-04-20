package main

func main() {
	GoRoutineDemo()
	WaitGroupDemo()
	BasicMutexDemo()
	ReadWriteMutexDemo()
	UnbufferedChannelDemo()
	BufferedChannelDemo()
	ChannelDirectionDemo()
	// Final application: worker pool combining goroutines + channels
	ConcurrencyDemo()
	TimeTickerDemo()
}
