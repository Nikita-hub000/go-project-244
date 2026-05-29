package cliapp

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestRunRequiresTwoArguments(t *testing.T) {
	var out bytes.Buffer
	app := New(&out)
	app.ExitErrHandler = func(_ context.Context, _ *cli.Command, _ error) {}

	err := app.Run(context.Background(), []string{"gendiff", "only-one-arg"})
	if err == nil {
		t.Fatal("expected error for invalid argument count")
	}

	if !strings.Contains(err.Error(), pathArgumentError) {
		t.Fatalf("unexpected error: %v", err)
	}
}
