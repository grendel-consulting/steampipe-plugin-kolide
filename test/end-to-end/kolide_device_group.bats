# bats file_tags=table:kolide_device_group, output:device_group

setup_file() {
    load "${BATS_TEST_DIRNAME}/_support/globals.bash"
    define_common_test_results
}

setup() {
    load "${BATS_TEST_DIRNAME}/_support/extensions.bash"
    load_helpers
}

# With a Kolide K2 API key, this endpoint returns a 403
#bats test_tags=scope:smoke, failing:billing
@test "query_forbidden_under_billing_plan" {
    if [[ "$MY_KOLIDE_PLAN" != "K2" ]]; then skip; fi

    run ! steampipe query $QUERY_UNDER_TEST
    assert_output --partial "403"
}
