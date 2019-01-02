package gex

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/transcom/mymove/pkg/server"
)

// gexRequestTimeout is how long to wait on Gex request before timing out (30 seconds).
const gexRequestTimeout = time.Duration(30) * time.Second

// SendToGex is an interface for sending and receiving a request
type SendToGex interface {
	Call(edi string, transactionName string) (resp *http.Response, err error)
}

// SendToGexHTTP represents a struct to contain an actual gex request function
type SendToGexHTTP struct{ URL string }

// Call sends an edi file string as a POST to the gex api
func (s SendToGexHTTP) Call(edi string, transactionName string) (resp *http.Response, err error) {
	// Ensure that the transaction body ends with a newline, otherwise the GEX EDI parser will fail silently
	edi = strings.TrimSpace(edi) + "\n"

	URL := s.URL
	hasSlash, _ := regexp.MatchString("(.*/$)", URL)
	if !hasSlash {
		URL = URL + "/"
	}
	request, err := http.NewRequest(
		"POST",
		URL+transactionName,
		strings.NewReader(edi),
	)
	if err != nil {
		return resp, errors.Wrap(err, "Creating GEX POST request")
	}

	// We need to provide basic auth credentials for the GEX server, as well as
	// our client certificate for the proxy in front of the GEX server.
	request.SetBasicAuth(os.Getenv("GEX_BASIC_AUTH_USERNAME"), os.Getenv("GEX_BASIC_AUTH_PASSWORD"))

	config, err := getTLSConfig()
	if err != nil {
		return resp, errors.Wrap(err, "Creating TLS config")
	}
	tr := &http.Transport{TLSClientConfig: config}

	client := &http.Client{Transport: tr, Timeout: gexRequestTimeout}
	resp, err = client.Do(request)
	if err != nil {
		return resp, errors.Wrap(err, "Sending GEX POST request")
	}

	return resp, err
}

// getTLSConfig gets the configuration certs for the GEX connection
func getTLSConfig() (*tls.Config, error) {
	clientCA := os.Getenv("MOVE_MIL_DOD_CA_CERT")
	clientCert := os.Getenv("MOVE_MIL_DOD_TLS_CERT")
	clientKey := os.Getenv("MOVE_MIL_DOD_TLS_KEY")
	// At this time, GEX does not already trust the intermediate CA that signed our certs; so include it with our cert
	clientCertPlusCA := strings.Join([]string{clientCert, clientCA}, "\n")

	certificate, err := tls.X509KeyPair([]byte(clientCertPlusCA), []byte(clientKey))
	if err != nil {
		return nil, errors.Wrap(err, "error creating key pair")
	}

	// Load DOD CA certs so that we can validate GEX's server cert
	pkcs7Package, err := ioutil.ReadFile(os.Getenv("DOD_CA_PACKAGE"))
	if err != nil {
		return nil, err
	}
	rootCAs, err := server.LoadCertPoolFromPkcs7Package(pkcs7Package)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      rootCAs,
	}, nil
}
