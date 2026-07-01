package pipeline

type Pipeline struct {
	rawSource       chan []byte
	formattedSource chan []byte
	done            chan any
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		rawSource:       make(chan []byte),
		formattedSource: make(chan []byte),
		done:            make(chan any),
	}
}

func (p *Pipeline) Submit(src []byte) {
	p.rawSource <- src
}


