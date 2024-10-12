package kolide_client // needs to be inside kolide_client package to intercept HTTP communications

import (
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
var baseUrl string = "https://api.kolide.com"

var _ = BeforeSuite(func() {
	kolide = New()
	// Block all HTTP requests made by the Kolide client
	httpmock.ActivateNonDefault(kolide.c.GetClient())
})

var _ = BeforeEach(func() {
	// Remove any mock responders
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	// Unblock our HTTP requests
	httpmock.DeactivateAndReset()
})
