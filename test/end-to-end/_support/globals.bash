#!/usr/bin/env bash

define_file_globals(){
    export TABLE_UNDER_TEST=$(basename $BATS_TEST_FILENAME ".bats")
    export QUERY_UNDER_TEST="${BATS_TEST_DIRNAME}/_query/${TABLE_UNDER_TEST}.sql"
    export QUERY_RESULTS="${BATS_TEST_DIRNAME}/_output/${TABLE_UNDER_TEST}.json"
}

# Given we expect most tables to be relatively stable sets, we will
# make the first alphabetical entry our test results; where not, we will
# rely less on static test values and look for non-failing execution
#
# Unless otherwise required, tests will use the 'basic info' queries
# with an `order by` clause to ensure repeatability
define_test_results(){
    export KOLIDE_PLAN="K2"

    export MY_DEFAULT_EMAIL_DOMAIN="@grendel-consulting.com"

    export MY_ADMIN_USER_COUNT="1"
    export MY_TOP_ADMIN_USER_FIRST_NAME="James"
    export MY_TOP_ADMIN_USER_LAST_NAME="Ramirez"
    export MY_TOP_ADMIN_USER_ACCESS="full"

    # Given volatile table, no count based test
    export MY_TOP_AUDIT_LOG_TIMESTAMP="2020-01"
    export MY_TOP_AUDIT_LOG_DESCRIPTION="Live Query Campaign"
    export MY_TOP_AUDIT_LOG_ACTOR_NAME="James Ramirez"

    # Potentially volatile as Kolide, or tenant, may define and release new checks
    export MY_CHECK_COUNT="61"
    export MY_TOP_CHECK_NAME="macOS Sharing - Require Bluetooth Sharing to Be Disabled"
    export MY_TOP_CHECK_TOPICS="sharing-preferences"
    export MY_TOP_CHECK_COMPATIBLE_PLATFORMS="darwin"
    export MY_TOP_CHECK_TARGETED_GROUPS="macOS"
    export MY_TOP_CHECK_BLOCKING_GROUP_NAMES="macOS"
    export MY_TOP_CHECK_BLOCKING_ENABLED="false"

    export MY_DEPROVISIONED_PERSON_COUNT="2"
    export MY_TOP_DEPROVISIONED_USER_NAME="Cloud Manager"

    # Given 403, no expected results for device groups, or device group devices

    # Given volatile table, no count based test
    export MY_TOP_DEVICE_OPEN_ISSUE_TITLE="Device Battery Requires Servicing"

    export MY_DEVICE_COUNT="2"
    export MY_TOP_DEVICE_NAME="ikebana"
    export MY_TOP_DEVICE_HARDWARE_MODEL="MacBook Pro"
    export MY_TOP_DEVICE_SERIAL="C02V40F2HV2Q"

    # Given volatile table, no count based test
    export MY_TOP_ISSUE_TITLE="File Extensions not visible in Finder"

    export MY_PACKAGE_COUNT="4"
    export MY_TOP_PACKAGE_ID="darwin-launchd-pkg"
    export MY_TOP_PACKAGE_URL="https://api.kolide.com/package_downloads/darwin-launchd-pkg"

    # Given 403, no expected results for person groups

    # Given zero entries, minimal expected results for person
    export MY_PERSON_COUNT="0"
    export MY_TOP_PERSON_NAME="Xerxes"
    export MY_TOP_PERSON_REGISTERED_DEVICE="false"

    # Given zero entries, minimal expected results for person open issues
    export MY_PERSON_OPEN_ISSUE_COUNT="0"
    export MY_TOP_PERSON_OPEN_ISSUE_TITLE="Device Battery Requires Servicing"

    # Given zero entries, minimal expected results for person registered devices
    export MY_PERSON_REGISTERED_DEVICE_COUNT="0"
    export MY_TOP_PERSON_REGISTERED_DEVICE_NAME="ikebana"
    export MY_TOP_PERSON_REGISTERED_DEVICE_HARDWARE_MODEL="MacBook Pro"
    export MY_TOP_PERSON_REGISTERED_DEVICE_SERIAL="C02V40F2HV2Q"
}
