# bats file_tags=table:kolide_device_group_device, output:device

setup() {
    load "${BATS_TEST_DIRNAME}/_support/extensions.bash"
    load_helpers
}

# With a Kolide K2 API key, this endpoint returns a 403
#bats test_tags=scope:smoke
@test "query_forbidden_under_billing_plan" {
    if ![ "$KOLIDE_PLAN" = "K2" ]; then skip; fi

    run ! steampipe query "select id, name from kolide_device_group order by name desc;"
    assert_output --partial "Kolide K2 API Error: 403 Forbidden"
}
