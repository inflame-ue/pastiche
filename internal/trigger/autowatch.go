package trigger

import (
	"bytes"
	"context"

	"github.com/inflame-ue/pastiche/internal/pipeline"
	"golang.design/x/clipboard"
)

func heuristic(src []byte, heuristicThreshold int) bool {
	var hereusticHits int

	if bytes.Count(src, []byte("\n")) >= 3 {
		hereusticHits++
	}

	if bytes.ContainsAny(src, "{(=:;#") {
		hereusticHits++
	}

	if bytes.Contains(src, []byte("==")) || bytes.Contains(src, []byte("!=")) {
		hereusticHits++
	}

	if bytes.Contains(src, []byte("=>")) || bytes.Contains(src, []byte("->")) {
		hereusticHits++
	}

	if bytes.Contains(src, []byte("//")) || bytes.Contains(src, []byte("/*")) {
		hereusticHits++
	}

	return hereusticHits >= heuristicThreshold
}

func FormatAutowatch(ctx context.Context, p *pipeline.Pipeline, heuristicThreshold int) {
	var lastSeen []byte
	for data := range clipboard.Watch(ctx, clipboard.FmtText) {
		src := data.Bytes

		// this is at most a two cycle, because we write back to the clipboard
		if bytes.Equal(src, lastSeen) {
			continue
		}

		lastSeen = src
		if heuristic(src, heuristicThreshold) {
			p.Submit(src)
		}
	}
}
