# bats file_tags=table:kolide_admin_user, output:admin_user

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

@test "has_expected_first_name" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].first_name'"
    if [[ -z "$FIRST_NAME" ]]; then assert_output $FIRST_NAME ; else assert_success ; fi
}

@test "has_expected_last_name" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].last_name'"
    if [[ -z "$LAST_NAME" ]]; then assert_output $LAST_NAME ; else assert_success ; fi
}

#bats test_tags=exactness:fallback
@test "has_expected_email" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].email'"
    if [[ -z "$EMAIL" ]]; then assert_output $EMAIL ; else assert_output --partial $MY_DEFAULT_EMAIL_DOMAIN ; fi
}

#bats test_tags=exactness:default
@test "has_expected_access" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].access'"
    if [[ -z "$ACCESS" ]]; then assert_output $ACCESS ; else assert_output "full" ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_created_at" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].created_at'"
    if [[ -z "$CREATED_AT" ]]; then assert_output --partial $CREATED_AT ; else assert_success ; fi
}



teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
