package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	totalDuration := time.Duration(5)
	allowedDuration := totalDuration * time.Second

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(allowedDuration))
	defer cancel()

	name, err := getNameContext(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, name)
}

func getNameContext(ctx context.Context) (string, error) {
	var err error
	name := "DefaultName"

	c := make(chan error, 1)

	go func() {
		name, err = getName(os.Stdin, os.Stdout)
		c <- err
	}()

	select {
	case <-ctx.Done():
		return name, ctx.Err()
	case err := <-c:
		return name, err
	}
}

func getName(r io.Reader, w io.Writer) (string, error) {

	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when done"

	fmt.Fprintln(w, msg)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you entered an empty name")
	}

	return name, nil
}
