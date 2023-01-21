package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/sync/errgroup"
)

func MultiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		processSmth(ctx)

		nums := r.URL.Query()["ns"]

		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			span.LogKV("message", "got nums", "nums", nums)
		}

		resText := strings.Builder{}
		var lock sync.Mutex

		if len(nums) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "no numbers in request")
		}

		eg, egCtx := errgroup.WithContext(ctx)

		for _, num := range nums {
			num := num
			eg.Go(func() error {
				res, err := requestAnotherService(egCtx, num)
				if err != nil {
					return err
				}

				lock.Lock()
				fmt.Fprintf(&resText, "fib(%s) = %s\n", num, res)
				lock.Unlock()

				return nil
			})
		}

		err := eg.Wait()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}

		w.Write([]byte(resText.String()))
	})
}

func processSmth(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "processSmth")
	defer span.Finish()

	time.Sleep(time.Millisecond * 1)
}

var errStatusCode = errors.New("wrong status code")

func requestAnotherService(ctx context.Context, n string) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "requestAnotherService")
	span.SetTag("n", n)
	ext.SpanKindRPCClient.Set(span)
	defer span.Finish()

	query := url.Values{}
	query.Add("n", n)
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://127.0.0.1:8080/fibonacci?"+query.Encode(),
		nil,
	)
	if err != nil {
		return "", err
	}

	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errStatusCode
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
