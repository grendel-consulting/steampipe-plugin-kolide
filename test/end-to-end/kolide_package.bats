# bats file_tags=table:kolide_package, output:package

setup_file() {
    load "${BATS_TEST_DIRNAME}/_support/globals.bash"
    define_file_globals
    define_test_results
}

setup() {
    load "${BATS_TEST_DIRNAME}/_support/extensions.bash"
    load_helpers
}

#bats test_tags=scope:smoke
@test "can_execute_query_via_steampipe" {
    steampipe query $QUERY_UNDER_TEST --output json > $QUERY_RESULTS
    assert_exists $QUERY_RESULTS
}

#bats test_tags=scope:smoke
@test "has_expected_number_of_results" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '. | length'"
    assert_output $MY_PACKAGE_COUNT
}

@test "has_expected_id" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].id'"
    assert_output --partial $MY_TOP_PACKAGE_ID
}

@test "has_expected_url" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].url'"
    assert_output --partial $MY_TOP_PACKAGE_URL
}

teardown_file(){
    rm -f $QUERY_RESULTS
}
