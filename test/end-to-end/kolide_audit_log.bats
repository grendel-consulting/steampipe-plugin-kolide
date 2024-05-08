# bats file_tags=table:kolide_audit_log, output:audit_log

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
    skip "too volatile"
}

@test "has_expected_timestamp" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].timestamp'"
    assert_output --partial $MY_TOP_AUDIT_LOG_TIMESTAMP
}

@test "has_expected_description" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].description'"
    assert_output --partial $MY_TOP_AUDIT_LOG_DESCRIPTION
}

@test "has_expected_actor_name" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].actor_name'"
    assert_output --partial $MY_TOP_AUDIT_LOG_ACTOR_NAME
}

teardown_file(){
    rm -f $QUERY_RESULTS
}
