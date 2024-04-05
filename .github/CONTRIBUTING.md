# Contributing Guidelines

First off, thanks for taking the time to contribute!

We're conducting an experiment here by
[working in the open](https://visitmy.website/2020/01/25/blogging-working-open/). We're finding
out what works, and for that other perspectives matter.

## Our Code of Conduct

Our project and everyone participating in it are governed by our
[Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to
uphold this code. Please report unacceptable behavior to the project team at
[abuse@grendel-consulting.com][contact] or through the options to report an abusive
[issue](https://docs.github.com/en/github/building-a-strong-community/reporting-abuse-or-spam#reporting-an-issue-or-pull-request)
or
[comment](https://docs.github.com/en/github/building-a-strong-community/reporting-abuse-or-spam#reporting-a-comment)

## Getting Started

Please [start a conversation](https://github.com/grendel-consulting/steampipe-plugin-kolide/discussions/new/choose) or
[raise an issue](https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/new/choose)
about the feature or issue you've found; that provides us an opportunity
to understand what you've spotted, where it challenges our approach and where
it augments it.

## Your Commits

We request that prospective contributors include themselves in our [Contributors](../CONTRIBUTORS.md)
within their first pull request, to indicate they have read these guidelines and
agree to uphold our [Code of Conduct](CODE_OF_CONDUCT.md).

We require that contributors:

- [Sign off their commits](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository#about-commit-signoffs)
- [Sign their commits](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits)

You can read about the [difference between signing-off and signing](https://medium.com/@MarkEmeis/git-commit-signoff-vs-signing-9f37ee272b14)

## Our Conventions and Styleguides

We practise [scaled trunk-based developmet](https://trunkbaseddevelopment.com/) with
[short-lived feature branches](https://trunkbaseddevelopment.com/short-lived-feature-branches/)
and [continuous integration](https://trunkbaseddevelopment.com/continuous-integration/)
for everything being worked on by humans. Bots handle the heavy lifting in the
subsequent pull requests.

We use Linters and Golang's native formatters to maintain a consistent opinionated style.

We use Code Scanners to help spot bugs, issues and vulnerabilities.

We pin dependencies and then keep them evergreen automagically.

[contact]: mailto:abuse@grendel-consulting.com
