package bible

import (
	"context"
	"log"
)

func GetFromContext(ctx context.Context) *Bible {
	// Get the Bible from the context
	bible, ok := ctx.Value("bible").(*Bible)
	if !ok {
		log.Fatal("Could not get Bible from context")
	}
	return bible
}
