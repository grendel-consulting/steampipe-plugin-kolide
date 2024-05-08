# bats file_tags=table:kolide_admin_user, output:admin_user

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
    assert_output $MY_ADMIN_USER_COUNT
}

@test "has_expected_first_name" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].first_name'"
    assert_output $MY_TOP_ADMIN_USER_FIRST_NAME
}

@test "has_expected_last_name" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].last_name'"
    assert_output $MY_TOP_ADMIN_USER_LAST_NAME
}

@test "has_expected_email_domain" {
    if ![[ -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].email'"
    assert_output --partial $MY_DEFAULT_EMAIL_DOMAIN
}

@test "has_expected_access" {
    if ![[ -e $QUERY_RESULTS ]]; then skip; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].access'"
    assert_output $MY_TOP_ADMIN_USER_ACCESS
}

teardown_file(){
    rm -f $QUERY_RESULTS
}
