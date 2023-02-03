package main

import (
	"context"
	"log"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

func WaitGroup(files []string) {
	var wg sync.WaitGroup

	for _, file := range files {
		path := file

		wg.Add(1)

		go func() {
			defer wg.Done()

			data, err := os.ReadFile(path)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(string(data))
			}
		}()
	}

	wg.Wait()
}

func ErrorGroup(files []string) {
	var g errgroup.Group

	for _, file := range files {
		path := file

		g.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			log.Print(string(data))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Print(err)
	}
}

func ErrorGroupContext(ctx context.Context, files []string) {
	g, ctx := errgroup.WithContext(ctx)

	for _, file := range files {
		path := file

		g.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				log.Print(ctx.Err())
				return ctx.Err()
			default:
				log.Print(string(data))
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Print(err)
	}
}
