package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	transmitter := func(in Bi, done In, out Out) {
		for {
			select {
			case <-done:
				close(in)
				return
			default:
				select {
				case <-done:
					close(in)
					return
				case tmp, ok := <-out:
					if ok {
						in <- tmp
					} else {
						close(in)
						return
					}
				}
			}
		}
	}
	for _, stage := range stages {
		input := make(Bi)
		go transmitter(input, done, in)
		in = stage(input)
	}
	return in
}
