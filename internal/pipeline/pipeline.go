package pipeline

import (
	"context"
	"log"
	"sync"

	"github.com/inflame-ue/pastiche/internal/formatter"
	"golang.design/x/clipboard"
)

type Pipeline struct {
	rawSource chan []byte
	wg        sync.WaitGroup
	cancel    context.CancelFunc
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		rawSource: make(chan []byte),
		wg:        sync.WaitGroup{},
	}
}

func (p *Pipeline) Submit(src []byte) {
	p.rawSource <- src
}

func (p *Pipeline) Run(ctx context.Context, registry *formatter.FormatterRegistry) {
	ctx, p.cancel = context.WithCancel(ctx)
	for {
		select {
		case src := <-p.rawSource:
			p.wg.Add(1)
			out, err := registry.Format(src)
			if err != nil {
				log.Print(err)
				continue
			}
			log.Println("writing formatted source back to the clipboard...")
			clipboard.Write(clipboard.FmtText, out)
			p.wg.Done()
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pipeline) Stop() {
	p.cancel()
	p.wg.Wait()
}
