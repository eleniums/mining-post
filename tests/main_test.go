package tests

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/eleniums/mining-post/client"
)

const (
	testPlayer = "tstark"
)

var (
	baseURI    string
	gameClient *client.GameClient
)

var (
	httpHost           string
	httpPort           string
	enableTLS          bool
	insecureSkipVerify bool
)

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	// parse configuration - flags have highest priority, then env vars, and then defaults
	flag.StringVar(&httpHost, "http-host", getEnvStr("HTTP_HOST", "127.0.0.1"), "HTTP_HOST: host to serve endpoint on")
	flag.StringVar(&httpPort, "http-port", getEnvStr("HTTP_PORT", "9090"), "HTTP_PORT: port to serve endpoint on")
	flag.BoolVar(&enableTLS, "enable-tls", getEnvBool("ENABLE_TLS", false), "ENABLE_TLS: true to enable a secure TLS connection")
	flag.BoolVar(&insecureSkipVerify, "insecure-skip-verify", getEnvBool("INSECURE_SKIP_VERIFY", false), "INSECURE_SKIP_VERIFY: true to skip verifying the certificate chain and host name")
	flag.Parse()

	// create http client
	var config client.Config
	if enableTLS {
		baseURI = fmt.Sprintf("https://%s:%s", httpHost, httpPort)
		config.TLSConfig = &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		}
	} else {
		baseURI = fmt.Sprintf("http://%s:%s", httpHost, httpPort)
	}
	gameClient = client.NewGameClient(baseURI, config)

	// run tests
	result := m.Run()

	return result
}

// getEnvStr will return the value for the requested environment variable or a default value.
func getEnvStr(envVar string, defaultValue string) string {
	v := os.Getenv(envVar)
	if v == "" {
		return defaultValue
	}
	return v
}

// getEnvBool will return the value for the requested environment variable or a default value.
func getEnvBool(envVar string, defaultValue bool) bool {
	v := os.Getenv(envVar)
	if v == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}
