package kolide_client // test suit needs to be inside the kolide_client package to intercept HTTP communications

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var kolide *Client
var baseUrl string = os.Getenv("KOLIDE_API_URL")

func init() {
	if baseUrl == "" {
		baseUrl = "https://api.kolide.com"
	}
}

var _ = BeforeSuite(func() {
	kolide = New()
	// Block all HTTP requests made by the Kolide client
	httpmock.ActivateNonDefault(kolide.c.GetClient())
})

var _ = BeforeEach(func() {
	// Remove any mock responders so each test has a clean slate
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	// Unblock our HTTP requests
	httpmock.DeactivateAndReset()
})

func setupHTTPMock(target string, fixture interface{}) {
	f, _ := json.Marshal(fixture)
	httpmock.RegisterResponder("GET", target, func(request *http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(http.StatusOK, string(f))
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	})
}
