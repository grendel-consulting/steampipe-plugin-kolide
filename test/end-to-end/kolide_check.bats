# bats file_tags=table:kolide_check, output:check

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

@test "has_expected_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].name'"
    if [[ -z "$NAME" ]]; then assert_output $NAME ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_topics" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].topics'"
    if [[ -z "$TOPICS" ]]; then assert_output --partial $TOPICS ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_compatible_platforms" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].compatible_platforms'"
    if [[ -z "$COMPATIBLE_PLATFORMS" ]]; then assert_output --partial $COMPATIBLE_PLATFORMS ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_targeted_groups" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].targeted_groups'"
    if [[ -z "$TARGETED_GROUPS" ]]; then assert_output --partial $TARGETED_GROUPS ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_blocking_group_names" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].blocking_group_names'"
    if [[ -z "$BLOCKING_GROUP_NAMES" ]]; then assert_output --partial $BLOCKING_GROUP_NAMES ; else assert_success ; fi
}

#bats test_tags=exactness:default
@test "has_expected_blocking_enabled" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].blocking_enabled'"
    if [[ -z "$BLOCKING_ENABLED" ]]; then assert_output $BLOCKING_ENABLED ; else assert_output "false" ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
