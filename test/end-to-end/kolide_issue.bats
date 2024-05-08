# bats file_tags=table:kolide_issue, output:issue

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
@test "has_expected_title" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].title'"
    if [[ -z "$TITLE" ]]; then assert_output --partial $TITLE ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_detected_at" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].detected_at'"
    if [[ -z "$DETECTED_AT" ]]; then assert_output --partial $DETECTED_AT ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_blocks_device_at" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].blocks_device_at'"
    if [[ -z "$BLOCKS_DEVICE_AT" ]]; then assert_output --partial $BLOCKS_DEVICE_AT ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_resolved_at" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].resolved_at'"
    if [[ -z "$RESOLVED_AT" ]]; then assert_output --partial $RESOLVED_AT ; else assert_success ; fi
}

#bats test_tags=exactness:default
@test "has_expected_exempted" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].exempted'"
    if [[ -z "$EXEMPTED" ]]; then assert_output $EXEMPTED ; else assert_output "false" ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
