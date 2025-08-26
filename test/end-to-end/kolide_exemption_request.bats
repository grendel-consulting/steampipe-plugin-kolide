# bats file_tags=table:kolide_exemption_request, output:exemption_request

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

@test "has_expected_id" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].id'"
    if [[ -z "$ID" ]]; then assert_output $ID ; else assert_success ; fi
}

@test "has_expected_status" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].status'"
    if [[ -z "$STATUS" ]]; then assert_output --partial $STATUS ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_requested_at" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].requested_at'"
    if [[ -z "$REQUESTED_AT" ]]; then assert_output $REQUESTED_AT ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_requester_message" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].requester_message'"
    if [[ -z "$REQUESTER_MESSAGE" ]]; then assert_output --partial $REQUESTER_MESSAGE ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_requester_id" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].requester_id'"
    if [[ -z "$REQUESTER_ID" ]]; then assert_output --partial $REQUESTER_ID ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_device_id" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].device_id'"
    if [[ -z "$DEVICE_ID" ]]; then assert_output --partial $DEVICE_ID ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
