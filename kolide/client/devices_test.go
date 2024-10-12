package kolide_client

import (
	"encoding/json"
	"net/http"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Devices", Label("endpoint:devices"), func() {
	var target string
	var fixture *Device
	var result = Device{}

	JustBeforeEach(func() {
		f, _ := json.Marshal(fixture)

		httpmock.RegisterResponder("GET", target, func(request *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusOK, string(f))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		res, err := kolide.GetDeviceById(fixture.Id)
		Ω(err).ShouldNot(HaveOccurred())

		b, _ := json.Marshal(res)

		err = json.Unmarshal(b, &result)
		Ω(err).ShouldNot(HaveOccurred())

	})

	Context("GetDeviceById", func() {
		BeforeEach(func() {
			fixture = &Device{
				Id:   "1",
				Name: "ikebana",
			}
			target = baseUrl + "/devices/" + fixture.Id

		})

		It("retrieves the specified Device", func() {
			Ω(result).To(MatchFields(IgnoreExtras, Fields{
				"Id":   Equal(fixture.Id),
				"Name": Equal(fixture.Name),
			}))
		})
	})
})
