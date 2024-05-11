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
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows | length'"

    if [[ -z "$EXPECTED_COUNT" ]]; then assert_output $EXPECTED_COUNT ; else assert [ "$output" -ge "1" ] ; fi
}

@test "has_expected_first_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].first_name'"
    if [[ -z "$FIRST_NAME" ]]; then assert_output $FIRST_NAME ; else assert_success ; fi
}

@test "has_expected_last_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].last_name'"
    if [[ -z "$LAST_NAME" ]]; then assert_output $LAST_NAME ; else assert_success ; fi
}

#bats test_tags=exactness:fallback
@test "has_expected_email" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].email'"
    if [[ -z "$EMAIL" ]]; then assert_output $EMAIL ; else assert_output --partial $MY_DEFAULT_EMAIL_DOMAIN ; fi
}

#bats test_tags=exactness:default
@test "has_expected_access" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].access'"
    if [[ -z "$ACCESS" ]]; then assert_output $ACCESS ; else assert_output "full" ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_created_at" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].created_at'"
    if [[ -z "$CREATED_AT" ]]; then assert_output --partial $CREATED_AT ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
