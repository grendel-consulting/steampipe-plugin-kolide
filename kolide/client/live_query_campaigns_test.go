package kolide_client

import (
	"testing"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLiveQueryCampaigns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LiveQueryCampaigns Suite")
}

var _ = Describe("LiveQueryCampaigns", func() {
	var (
		client *Client
	)

	BeforeEach(func() {
		client = New()
		httpmock.ActivateNonDefault(client.c.GetClient())
	})

	AfterEach(func() {
		httpmock.DeactivateAndReset()
	})

	Describe("GetLiveQueryCampaigns", func() {
		It("should fetch the list of live query campaigns", func() {
			mockResponse := `{
				"data": [
					{
						"id": "1",
						"query": "SELECT * FROM users;",
						"status": "completed",
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T00:00:00Z"
					}
				],
				"pagination": {
					"next_cursor": ""
				}
			}`

			httpmock.RegisterResponder("GET", "https://api.kolide.com/live_query_campaigns/",
				httpmock.NewStringResponder(200, mockResponse))

			res, err := client.GetLiveQueryCampaigns("", 10)
			Expect(err).To(BeNil())
			Expect(res).To(HaveLen(1))

			campaigns := res.(*LiveQueryCampaignListResponse)
			Expect(campaigns.LiveQueryCampaigns[0].Id).To(Equal("1"))
			Expect(campaigns.LiveQueryCampaigns[0].Query).To(Equal("SELECT * FROM users;"))
			Expect(campaigns.LiveQueryCampaigns[0].Status).To(Equal("completed"))
		})
	})

	Describe("GetLiveQueryCampaignById", func() {
		It("should fetch a specific live query campaign by its ID", func() {
			mockResponse := `{
				"id": "1",
				"query": "SELECT * FROM users;",
				"status": "completed",
				"created_at": "2023-01-01T00:00:00Z",
				"updated_at": "2023-01-01T00:00:00Z"
			}`

			httpmock.RegisterResponder("GET", "https://api.kolide.com/live_query_campaigns/1",
				httpmock.NewStringResponder(200, mockResponse))

			res, err := client.GetLiveQueryCampaignById("1")
			Expect(err).To(BeNil())

			campaign := res.(*LiveQueryCampaign)
			Expect(campaign.Id).To(Equal("1"))
			Expect(campaign.Query).To(Equal("SELECT * FROM users;"))
			Expect(campaign.Status).To(Equal("completed"))
		})
	})
})
