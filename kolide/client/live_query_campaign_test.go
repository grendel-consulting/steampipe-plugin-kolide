package kolide_client

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("LiveQueryCampaign", Label("endpoint:live_query_campaign"), func() {
	var target string
	var fixture *LiveQueryCampaign
	var result = LiveQueryCampaign{}

	JustBeforeEach(func() {
		setupHTTPMock(target, fixture)

		res, err := kolide.GetLiveQueryCampaignById(fixture.Id)
		Expect(err).ShouldNot(HaveOccurred())

		// Jumping through hoops to cast from `interface{}` to `LiveQueryCampaign`
		b, _ := json.Marshal(res)
		err = json.Unmarshal(b, &result)
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("GetLiveQueryCampaignById", func() {
		BeforeEach(func() {
			fixture = &LiveQueryCampaign{
				Id:                     "1",
				Name:                   "Dump Environment",
				OsquerySql:             "SELECT * FROM default_environment",
				CreatedAt:              time.Now(),
				Published:              true,
				Revision:               0,
				TablesUsed:             []string{"default_environment"},
				SuccessfulDevicesCount: 0,
				ErroredDevicesCount:    0,
				WaitingDevicesCount:    1,
			}
			target = baseUrl + "/live_query_campaigns/" + fixture.Id
		})

		It("retrieves the specified Admin User", Label("plan:core"), func() {
			Expect(result).To(MatchFields(IgnoreExtras, Fields{
				"Id":                     Equal(fixture.Id),
				"Name":                   Equal(fixture.Name),
				"OsquerySql":             Equal(fixture.OsquerySql),
				"CreatedAt":              BeTemporally("<=", fixture.CreatedAt),
				"Published":              Equal(fixture.Published),
				"Revision":               Equal(fixture.Revision),
				"TablesUsed":             Equal(fixture.TablesUsed),
				"SuccessfulDevicesCount": Equal(fixture.SuccessfulDevicesCount),
				"ErroredDevicesCount":    Equal(fixture.ErroredDevicesCount),
				"WaitingDevicesCount":    Equal(fixture.WaitingDevicesCount),
			}))
		})
	})
})
