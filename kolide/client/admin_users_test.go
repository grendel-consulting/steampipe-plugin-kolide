package kolide_client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("AdminUser", Label("endpoint:admin_users"), func() {
	var target string
	var fixture *AdminUser
	var result = AdminUser{}

	JustBeforeEach(func() {
		f, _ := json.Marshal(fixture)

		httpmock.RegisterResponder("GET", target, func(request *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(http.StatusOK, string(f))
			resp.Header.Set("Content-Type", "application/json")
			return resp, nil
		})

		res, err := kolide.GetAdminUserById(fixture.Id)
		Ω(err).ShouldNot(HaveOccurred())

		b, _ := json.Marshal(res)

		err = json.Unmarshal(b, &result)
		Ω(err).ShouldNot(HaveOccurred())

	})

	Context("GetAdminUserById", func() {
		BeforeEach(func() {
			fixture = &AdminUser{
				Id:        "1",
				FirstName: "Bob",
				LastName:  "Bobson",
				Email:     "b@b.com",
				CreatedAt: time.Now(),
				Access:    "full",
			}
			target = baseUrl + "/admin_users/" + fixture.Id

		})

		It("retrieves the specified Admin User", func() {
			Ω(result).To(MatchFields(IgnoreExtras, Fields{
				"Id":        Equal(fixture.Id),
				"FirstName": Equal(fixture.FirstName),
				"LastName":  Equal(fixture.LastName),
				"Email":     Equal(fixture.Email),
				"CreatedAt": BeTemporally("<=", fixture.CreatedAt),
				"Access":    Equal(fixture.Access),
			}))
		})
	})
})
