# bats file_tags=table:kolide_package, output:package

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

@test "has_expected_id" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].id'"
    if [[ -z "$ID" ]]; then assert_output $ID ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_built_at" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].built_at'"
    if [[ -z "$BUILT_AT" ]]; then assert_output --partial $BUILT_AT ; else assert_success ; fi
}

@test "has_expected_version" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].version'"
    if [[ -z "$VERSION" ]]; then assert_output $VERSION ; else assert_success ; fi
}


@test "has_expected_url" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].url'"
    if [[ -z "$URL" ]]; then assert_output $URL ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
