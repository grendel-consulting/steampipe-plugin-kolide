# bats file_tags=table:kolide_person, output:person

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
@test "has_exactly_one_result" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi
    if [ "$MY_PERSON_COUNT" == "0" ]; then skip "no results"; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '. | length'"
    assert_output "1"
}

# Remaining functionality covered in kolide_person.bats

teardown_file(){
    rm -f $QUERY_RESULTS
}
