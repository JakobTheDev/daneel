package tools

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"

	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func RunSubfinder(domain string) ([]string, error) {
	subfinderOpts := &runner.Options{
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 10,
		ProviderConfig:     "config/subfinder-provider-config.yaml",
		Silent:             true,
	}

	// disable timestamps in logs / configure logger
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		log.Fatalf("failed to create subfinder runner: %v", err)
		return nil, err
	}

	output := &bytes.Buffer{}

	if err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		log.Fatalf("failed to enumerate single domain: %v", err)
		return nil, err
	}

	return strings.Split(output.String(), "\n"), err
}
