# bats file_tags=table:kolide_deprovisioned_person, output:deprovisioned_person

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
    assert_output $MY_DEPROVISIONED_PERSON_COUNT
}

@test "has_expected_name" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].name'"
    assert_output --partial $MY_TOP_DEPROVISIONED_PERSON_NAME
}

@test "has_expected_email_domain" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].email'"
    assert_output --partial $MY_DEFAULT_EMAIL_DOMAIN
}

teardown_file(){
    rm -f $QUERY_RESULTS
}
