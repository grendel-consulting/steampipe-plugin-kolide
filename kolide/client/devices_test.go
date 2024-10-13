package kolide_client

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Devices", Label("endpoint:devices"), func() {
	var target string
	var fixture *Device
	var result = Device{}

	JustBeforeEach(func() {
		setupHTTPMock(target, fixture)

		res, err := kolide.GetDeviceById(fixture.Id)
		Expect(err).ShouldNot(HaveOccurred())

		b, _ := json.Marshal(res)

		err = json.Unmarshal(b, &result)
		Expect(err).ShouldNot(HaveOccurred())

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
			Expect(result).To(MatchFields(IgnoreExtras, Fields{
				"Id":   Equal(fixture.Id),
				"Name": Equal(fixture.Name),
			}))
		})
	})
})
