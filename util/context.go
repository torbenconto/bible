package util

import (
	"context"
	"github.com/torbenconto/bible"
	"log"
)

func GetFromContext(ctx context.Context) *bible.Bible {
	// Get the Bible from the context
	ctxBible, ok := ctx.Value("bible").(*bible.Bible)
	if !ok {
		log.Fatal("Could not get Bible from context")
	}
	return ctxBible
}
