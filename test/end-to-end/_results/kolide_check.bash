#!/usr/bin/env bash

define_test_results(){
    # export EXPECTED_COUNT=""
    # export ID=""
    export NAME="macOS Sharing - Require Bluetooth Sharing to Be Disabled"
    export TOPICS="sharing-preferences"
    export COMPATIBLE_PLATFORMS="darwin"
    export TARGETED_GROUPS="macOS"
    export BLOCKING_GROUP_NAMES="macOS"
    export BLOCKING_ENABLED="false"
}
