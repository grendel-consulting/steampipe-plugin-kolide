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
