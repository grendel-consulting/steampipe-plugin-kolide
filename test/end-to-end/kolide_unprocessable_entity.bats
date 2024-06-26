# bats file_tags=table:kolide_person_open_issue, output:issue

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

# Regression test for https://github.com/grendel-consulting/steampipe-plugin-kolide/issues/129
#bats test_tags=scope:regression
@test "can_execute_query_via_steampipe" {
    steampipe query $QUERY_UNDER_TEST --output json > $QUERY_RESULTS
    assert_exists $QUERY_RESULTS
}

@test "has_expected_number_of_results" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows | length'"

    if [[ -z "$EXPECTED_COUNT" ]]; then assert_output $EXPECTED_COUNT ; else assert [ "$output" -ge "1" ] ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_title" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].title'"
    if [[ -z "$TITLE" ]]; then assert_output --partial $TITLE ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_person_id" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].person_id'"
    if [[ -z "$PERSON_ID" ]]; then assert_output --partial $PERSON_ID ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
