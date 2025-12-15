package scheduler

func Run(_ interface{}, job func() error) {
	_ = job()
}
