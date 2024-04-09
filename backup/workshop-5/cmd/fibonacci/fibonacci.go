package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func getNumber(ctx context.Context, n int) int {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getNumber")
	span.SetTag("n", n)
	defer span.Finish()

	switch n {
	case 0, 1:
		return n
	default:
		return getNumber(ctx, n-1) + getNumber(ctx, n-2)
	}
}
