# bats file_tags=table:kolide_person, output:person

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

@test "has_expected_name" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].name'"
    if [[ -z "$NAME" ]]; then assert_output $NAME ; else assert_success ; fi
}

#bats test_tags=exactness:fallback
@test "has_expected_email" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].email'"
    if [[ -z "$EMAIL" ]]; then assert_output $EMAIL ; else assert_output --partial $MY_DEFAULT_EMAIL_DOMAIN ; fi
}

#bats test_tags=exactness:default
@test "has_expected_has_registered_device" {
    if [[ ! -e $QUERY_RESULTS ]]; then skip ; fi

    run bash -c "cat $QUERY_RESULTS | jq -r '.[0].has_registered_device'"
    if [[ -z "$HAS_REGISTERED_DEVICE" ]]; then assert_output $HAS_REGISTERED_DEVICE ; else assert_output "false" ; fi
}



teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
