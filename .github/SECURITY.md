# Security Policy and Procedures

Our security policies and procedures for this project are outlined below. Broadly,
we wish to avoid leaving the ecosystem worse than we found it.

## Supported Versions

Any issues with minor versions will be investigated whilst this project is in early 
development, however we may recommend updated to the latest pre-v1.0 version.

| Version | Supported          |
| ------- | ------------------ |
| < 1.0   | :white_check_mark: |

## Reporting a Bug or Vulnerability

We take all security bugs in our projects seriously. Thank you for improving the
security of them. We appreciate your efforts and responsible disclosure, and will
make every effort to acknowledge your contributions. At this time we do not run
a formal bug bounty programme.

Report security bugs by emailing the lead maintainer at
[security@grendel-consulting.com][security].

They will acknowledge your email within 72 hours, and will send a more detailed
response within a further 72 hours indicating the next steps in handling your
report. After the initial reply to your report, we will endeavor to keep you
informed of the progress towards a fix and full announcement, and may ask for
additional information or guidance.

Report security bugs in third-party modules should be to the person or team
maintaining said module.

## Disclosure Policy

We are advocates of [responsible vulnerability disclosure][disclosure]. If you’ve
found a vulnerability, we would like to know so we can fix it.

Disclosures should be sent to [security@grendel-consulting.com][security], including:

- Your name and affiliation
- Sufficient details of the vulnerability to allow it to be understood and
  reproduced; this would include the website, page or repository where the
  vulnerability can be observed
- Optionally, the type of vulnerability and any related [OWASP category][category]
- Relevant HTTP requests and responses, HTML snippets, screenshots or any other
  supporting evidence. Redact any personal data before reporting
- Proof of concept code (if available), or non-destructive exploitation details
- The impact of the vulnerability
- Any references or further reading that may be appropriate

Our investigation process is straight-forward. We will work to:

- Confirm the problem and determine the affected versions.
- Audit code to find any potential similar problems.
- Prepare fixes for all releases still under maintenance

## Security Checklist and Recommendations

We have baked some baseline security checks into our toolchain, to be reflected
in this section together with things to watch out for.

### Our Security Toolchain

- GitHub [Advisories](https://github.com/grendel-consulting/risen-world/security/advisories)
- [Renovate](https://renovate.whitesourcesoftware.com/)

### Our Security Checklist

- [ ] You MUST encode, escape and validate any inputs
- [ ] You MUST NOT commit secrets, passwords or keys
- [ ] You SHOULD pin any new dependencies

### Recommendations

Prospective contributors are encouraged to familiarise themselves, if not already,
with existing techniques and good practise.

## Providing Feedback

If you have suggestions on how this process could be improved please submit a
pull request.

## Versions

All notable changes to this policy should be noted below. We use
[SemVer](https://semver.org) for versioning, with the following intents:

- We will increment the MAJOR version when we change contact information,
  encryption keys, or a field in security.txt in a backwards-incompatible manner
- We will increment the MINOR version when we otherwise change this file or the
  security.txt in a backwards-compatible manner
- We will increment the PATCH version for minor typos or similar

### Version History

- 1.0.0 (2024-04-03) - Initial policy and procedures

## Attribution

Thanks to [@trewaters](https://github.com/trewaters) for their thoughts on
[structuring a SECURITY readme](https://github.com/Trewaters/security-README)
together with the team behind [security.txt](https://securitytxt.org/)

Based in part on the excellent material in the [standardjs security readme](https://github.com/standard/.github/blob/master/SECURITY.md)

Licensed under [CC BY-SA 4.0 Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/)

[security]: mailto:security@grendel-consulting.com
[disclosure]: https://cheatsheetseries.owasp.org/cheatsheets/Vulnerability_Disclosure_Cheat_Sheet.html#responsible-or-coordinated-disclosure
[category]: https://owasp.org/www-project-top-ten/