# bats file_tags=table:kolide_auth_log, output:auth_log

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
    steampipe query $QUERY_UNDER_TEST --output json > $QUERY_RESULTS
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
@test "has_expected_person_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].person_name'"
    if [[ -z "$PERSON_NAME" ]]; then assert_output --partial $PERSON_NAME ; else assert_success ; fi
}

#bats test_tags=exactness:default
@test "has_expected_initial_status" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].initial_status'"
    if [[ -z "$INITIAL_STATUS" ]]; then assert_output $INITIAL_STATUS ; else assert_output "full" ; fi
}

#bats test_tags=exactness:default
@test "has_expected_result" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].result'"
    if [[ -z "$RESULT" ]]; then assert_output $RESULT ; else assert_output "full" ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
