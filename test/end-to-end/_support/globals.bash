#!/usr/bin/env bash


define_file_globals(){
    export TABLE_UNDER_TEST=$(basename "$BATS_TEST_FILENAME" ".bats")
    export EXPECTED_RESULTS="${BATS_TEST_DIRNAME}/_results/${TABLE_UNDER_TEST}.bash"
    export QUERY_UNDER_TEST="${BATS_TEST_DIRNAME}/_query/${TABLE_UNDER_TEST}.sql"
    export QUERY_RESULTS="${BATS_TEST_DIRNAME}/_output/${TABLE_UNDER_TEST}.json"

    export MY_KOLIDE_PLAN="K2"
}

define_common_test_results(){
    export MY_DEFAULT_EMAIL_DOMAIN="@grendel-consulting.com"
}
