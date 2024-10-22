# Developing

We recommend reviewing our contributing guidelines and the [writing plugins](https://steampipe.io/docs/develop/writing-plugins) tutorial from Turbot

## Prerequisites

Ensure the following are installed in your workspace at the [supported version](./github/SUPPORT.md):

* [Steampipe](https://steampipe.io/downloads)
* [Golang](https://golang.org/doc/install)

## Testing

## Unit Tests

### Prerequisties

Ensure the following are installedd in your workspace to run tests

* [Ginkgo](https://onsi.github.io/ginkgo/)

Dependencies on other testing tools, such as Gomega, are present in `go.mod` whilst tests need to exist in the `kolide_client` package

### Approach and Structure

Found in `kolide/client/*_test.go` files, and are written using Ginkgo and Gomega, popular Go testing frameworks.

We separate tests by endpoint, which is how the client SDK is structured, with hard-coded fixtures that we use with HttpMock to provide a mock for the Kolide API. Given isolation from the real Kolide API, they are structured to allow us to mimic various fleets, pricing plans and so forth.

Each test set should ensure we can run query the endpoint, it contains the expected number of results, and that each of the fields returned are in line with expectations for both happy path and unhappy path scenarios.

### Running Tests

To run the whole test suite, from the repo root. You may omit the `-v` flag for less verbose output.

```bash
ginkgo -v kolide/client
```

To run tests for a specific endpoint:

```bash
ginkgo --label-filter="endpoint:devices" kolide/client
```

To run tests for a specific pricing plan:

```bash
ginkgo --label-filter="plan:k2" kolide/client
```

### Writing Tests

We encourage tests to be written using the [Arrange-Act-Assert](https://softwarepatternslexicon.com/patterns-go/testing/arrange-act-assert/) pattern, using `BeforeEach` where it simplifies or makes test code more readable

## End-to-End Tests

### Prerequisites

Ensure the following are installed in your workspace to run tests

* [BATS](https://bats-core.readthedocs.io/en/stable/)
* [JQ](https://jqlang.github.io/jq/)

### Approach and Structure

Found in `test/end-to-end`, and are written using BATS, a bash testing framework.

We separate tests by query, with a test harness, a query file and an expected results file. We rely on common naming to reduce the heavy-lifting in setting up tests (though we miss some of the ease of table-driven and aspect-driven testing found in other frameworks that could reduce this further).

For example, running the basic info query for the `kolide_{table}` table is accomplished with:

* `test/end-to-end/kolide_{table}.bats`
* `test/end-to-end/_query/kolide_{table}.sql`
* `test/end-to-end/_results/kolide_{table}.bash`

Each test typically ensures we can run the query, it contains the expected number of results, and that each of the fields returned are in line with expectations. Given these are live queries against a dynamic fleet of workstations, tests should minimise fragility and be written with that volatility in mind. By their nature, some tables will be more volatile than others.

### Running Tests

Running these tests locally will require your own Kolide API key and hence require changes to the expected results found in `results/*.bash`. We have structured our end-to-end tests deliberately to facilitate these changes.

To run the whole test suite, from the repo root

```bash
bats test/end-to-end/kolide_*.bats
```

To run a subset of smoke tests, which will hit every API endpoint available under your billing plan:

```bash
bats test/end-to-end/kolide_*.bats --filter-tags scope:smoke
```

To run tests for a specific table:

```bash
bats test/end-to-end/kolide_*.bats --filter-tags table:kolide_device
```

To run tests for a specific return type:

```bash
bats test/end-to-end/kolide_*.bats --filter-tags output:device
```

### Writing Tests

Writing tests against this plugin follows three broad patterns:

* Basic table queries using List semantics, such as `kolide_{table}`
* Filtered table queries using Get semantics, such as `kolide_{table}_by_id`
* Filtered table queries using List semantics, such as `kolide_{root-table}_{table}`

With our test structure, there is a reasonable degree of boilerplate in `setup_file`, `setup` and `teardown_file` functions. These can be borrowed from other tests of the same pattern

Each table typically had a `can_execute_query_via_steampipe` smoke test, a count-based test and then tests for the main fields being returned. We rely on structured idempotent queries that return consistent results, and a results file that articulates these.
