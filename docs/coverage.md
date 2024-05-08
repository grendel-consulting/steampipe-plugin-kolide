# Coverage and API Behaviour

Documenting the observed behaviour of the Kolide API utilising credentials representing
different billing plans, we have seen the following. All tests were manual and performed
using the interactive API documentation at https://kolideapi.readme.io/reference

| Action | Endpoint                                                   | Covered?           | Core | Max  | K2   |
|--------|------------------------------------------------------------|--------------------|------|------|------|
| GET    | /audit_logs                                                | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /audit_logs/{id}                                           | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /devices/{deviceId}/open_issues                            | :question:         |    ? |    ? | [^2] |
| PATCH  | /devices/{deviceId}/authentication_mode                    | :no_entry_sign:    |      |      |      |
| PUT    | /devices/{deviceId}/authentication_mode                    | :no_entry_sign:    |      |      |      |
| DELETE | /devices/{deviceId}/registration                           | :no_entry_sign:    |      |      |      |
| POST   | /devices/{deviceId}/check_refreshes                        | :no_entry_sign:    |      |      |      |
| GET    | /devices                                                   | :question:         |    ? |    ? | [^2] |
| GET    | /devices/{id}                                              | :question:         |    ? |    ? | [^2] |
| DELETE | /devices/{id}                                              | :no_entry_sign:    |      |      |      |
| GET    | /issues                                                    | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /issues/{id}                                               | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /deprovisioned_people                                      | :question:         |    ? |    ? | [^3] |
| GET    | /packages                                                  | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /packages/{id}                                             | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /people/{personId}/registered_devices                      | :question:         |    ? |    ? | [^4] |
| GET    | /people/{personId}/open_issues                             | :exclamation:      |    ? |    ? | [^4] |
| GET    | /people                                                    | :question:         |    ? |    ? | [^4] |
| GET    | /people/{id}                                               | :question:         |    ? |    ? |   Ok |
| GET    | /person_groups                                             | :white_check_mark: |    ? |    ? | [^1] |
| GET    | /person_groups/{id}                                        | :white_check_mark: |    ? |    ? | [^1] |
| GET    | /auth_logs                                                 | #27                |    ? |    ? |      |
| GET    | /auth_logs/{id}                                            | #28                |    ? |    ? |      |
| GET    | /device_groups/{deviceGroupId}/devices                     | :white_check_mark: |    ? |    ? | [^1] |
| POST   | /device_groups/{deviceGroupId}/memberships                 | :no_entry_sign:    |      |      |      |
| DELETE | /device_groups/{deviceGroupId}/memberships/{id}            | :no_entry_sign:    |      |      |      |
| GET    | /device_groups                                             | :white_check_mark: |    ? |    ? | [^1] |
| GET    | /device_groups/{id}                                        | :white_check_mark: |    ? |    ? | [^1] |
| GET    | /live_query_campaigns/{liveQueryCampaignId}/query_results  | #32                |    ? |    ? | [^1] |
| GET    | /live_query_campaigns                                      | #33                |    ? |    ? | [^1] |
| GET    | /live_query_campaigns/{id}                                 | #34                |    ? |    ? | [^1] |
| GET    | /checks                                                    | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /checks/{id}                                               | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /exemption_requests                                        | #36                |    ? |    ? |   Ok |
| GET    | /exemption_requests/{id}                                   | #37                |    ? |    ? |   Ok |
| PATCH  | /exemption_requests/{id}                                   | :no_entry_sign:    |      |      |      |
| PUT    | /exemption_requests/{id}                                   | :no_entry_sign:    |      |      |      |
| GET    | /registration_requests                                     | #38                |    ? |    ? |   Ok |
| GET    | /registration_requests/{id}                                | #39                |    ? |    ? |   Ok |
| PATCH  | /registration_requests/{id}                                | :no_entry_sign:    |      |      |      |
| PUT    | /registration_requests/{id}                                | :no_entry_sign:    |      |      |      |
| GET    | /reporting/tables/{tableName}/table_records                | #42                |    ? |    ? | [^1] |
| GET    | /reporting/tables                                          | #40                |    ? |    ? | [^1] |
| GET    | /reporting/tables/{name}                                   | #41                |    ? |    ? | [^1] |
| GET    | /reporting/queries                                         | #43                |    ? |    ? | [^1] |
| GET    | /reporting/queries/{id}                                    | #44                |    ? |    ? | [^1] |
| GET    | /admin_users                                               | :white_check_mark: |    ? |    ? |   Ok |
| GET    | /admin_users/{id}                                          | :white_check_mark: |    ? |    ? |   Ok |

[^1]: Returns a 403, not included in billing plan
[^2]: Returns ok, but registration information is missing, i.e. date and owner
[^3]: Returns existing active users
[^4]: Returns empty list, despite existing active user
