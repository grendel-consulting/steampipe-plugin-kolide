# bats file_tags=table:kolide_check, output:check

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
    assert_output $MY_CHECK_COUNT
}

@test "has_expected_name" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].name'"
    assert_output --partial $MY_TOP_CHECK_NAME
}

@test "has_expected_topics" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].topics'"
    assert_output --partial $MY_TOP_CHECK_TOPICS
}

@test "has_expected_compatible_platforms" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].compatible_platforms'"
    assert_output --partial $MY_TOP_CHECK_COMPATIBLE_PLATFORMS
}

@test "has_expected_targeted_groups" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].targeted_groups'"
    assert_output --partial $MY_TOP_CHECK_TARGETED_GROUPS
}

@test "has_expected_blocking_group_names" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].blocking_group_names'"
    assert_output --partial $MY_TOP_CHECK_BLOCKING_GROUP_NAMES
}

@test "has_expected_blocking_enabled" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].blocking_enabled'"
    assert_output --partial $MY_TOP_CHECK_BLOCKING_ENABLED
}

teardown_file(){
    rm -f $QUERY_RESULTS
}
