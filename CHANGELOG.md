## v0.4.0 [2024-08-05]

_What's new?_

- New tables, not fully able to test these on Core and Max plans

  - kolide_auth_log
  - kolide_exemption_request
  - kolide_registration_request

- Bug fix for unprocessable entity error introduced in v0.3.1, #129
- Support for fixes in the Kolide API, where documentation and observed values were mismatched
- Support for the revised JSON output format used in Steampipe v0.23
- Documentation to support local development and testing
- Dependency updates


## v0.3.1 [2024-05-08]

_What's new?_

- End-to-end test harness covering basic info queries on all implemented tables
- Bug fixes on several tables, potentially not fully resolving the issues with `kolide_person*` tables
- Documentation clarifications on the coverage matrix, reflecting these
- Dependency updates

## v0.3.0 [2024-05-03]

_What's new?_

- New tables, not fully able to test these on Core and Max plans

  - kolide_device_group
  - kolide_device_group_device
  - kolide_person
  - kolide_person_open_issue
  - kolide_person_registered_device
  - kolide_person_group

- Documentation coverage matrix, showing support and behaviour for Core, Max and K2 plans
- Small fixes on missing columns and filters, better synthetic 'title' columns
- Dependency updates

## v0.2.1 [2024-04-11]

_What's new?_

- Renamed tables to omit "k2", i.e. it is now simply `kolide_device`

## v0.2.0 [2024-04-05]

_What's new?_

- New tables

  - kolide_k2_check
  - kolide_k2_device_open_issue
  - kolide_k2_issue
  - kolide_k2_package

- Refactored underlying client to support alternate search field names and keyed collections
- Pre-commit hooks for local development
- Hardened all GitHub Actions and introduced CodeQL and OSSF Scorecard scanning
- Update Go to v1.21.9

- Documentation clarifications and updates

  - Contributing guidelines, issue and pull request templates
  - Discussions, support and security documentation
  - Examples

## v0.1.0 [2024-03-31]

_What's new?_

- Initial release with tables:

  - kolide_k2_admin_user
  - kolide_k2_audit_logs
  - kolide_k2_deprovisioned_person
  - kolide_k2_devices
