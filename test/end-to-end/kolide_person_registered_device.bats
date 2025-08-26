# bats file_tags=table:kolide_registered_device, output:device

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

@test "has_expected_name" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].name'"
    if [[ -z "$NAME" ]]; then assert_output $NAME ; else assert_success ; fi
}

#bats test_tags=exactness:fuzzy
@test "has_expected_hardware_model" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].hardware_model'"
    if [[ -z "$HARDWARE_MODEL" ]]; then assert_output --partial $HARDWARE_MODEL ; else assert_success ; fi
}

@test "has_expected_serial" {
    run bash -c "cat $QUERY_RESULTS | jq -r '.rows.[0].serial'"
    if [[ -z "$SERIAL" ]]; then assert_output $SERIAL ; else assert_success ; fi
}

teardown_file(){
    if [[ -f $QUERY_RESULTS ]]; then
        rm -f $QUERY_RESULTS
    fi
}
