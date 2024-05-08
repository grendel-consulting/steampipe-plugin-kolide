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
    steampipe query $QUERY_UNDER_TEST --output json > $QUERY_RESULTS
    assert_exists $QUERY_RESULTS
}

@test "has_expected_number_of_results" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '. | length'"

    if [[ -z "$EXPECTED_COUNT" ]]; then assert_output $EXPECTED_COUNT ; else assert [ "$output" -ge "1" ] ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_timestamp" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].timestamp'"
    if [[ -z "$TIMESTAMP" ]]; then assert_output --partial $TIMESTAMP ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_description" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].description'"
    if [[ -z "$DESCRIPTION" ]]; then assert_output --partial $DESCRIPTION ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_actor_name" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].actor_name'"
    if [[ -z "$ACTOR_NAME" ]]; then assert_output --partial $ACTOR_NAME ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
