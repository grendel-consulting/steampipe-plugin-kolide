# bats file_tags=table:kolide_audit_log, output:audit_log

setup_file() {
    load "${BATS_TEST_DIRNAME}/_support/globals.bash"
    define_file_globals

    define_common_test_results

    if [[ -f $EXPECTED_RESULTS ]]; then
        load $EXPECTED_RESULTS
    fi
}

setup() {
    load "${BATS_TEST_DIRNAME}/_support/extensions.bash"
    load_helpers
}

#bats test_tags=scope:smoke
@test "can_execute_query_via_steampipe" {
    steampipe query $QUERY_UNDER_TEST --output json > $QUERY_RESULTS 3>&-
    assert_exists $QUERY_RESULTS
}

@test "has_expected_number_of_results" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows | length'"

    if [[ -z "$EXPECTED_COUNT" ]]; then assert_output $EXPECTED_COUNT ; else assert [ "$output" -ge "1" ] ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_timestamp" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].timestamp'"
    if [[ -z "$TIMESTAMP" ]]; then assert_output --partial $TIMESTAMP ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_description" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].description'"
    if [[ -z "$DESCRIPTION" ]]; then assert_output --partial $DESCRIPTION ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_actor_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].actor_name'"
    if [[ -z "$ACTOR_NAME" ]]; then assert_output --partial $ACTOR_NAME ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
