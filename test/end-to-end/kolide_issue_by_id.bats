# bats file_tags=table:kolide_issue, output:issue

setup_file() {
    load "${BATS_TEST_DIRNAME}/_support/globals.bash"
    define_file_globals
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

@test "has_no_more_than_one_result" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '. | length'"
    assert [ "$output" -le "1" ]
}

# Remaining functionality covered in kolide_issue.bats

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
