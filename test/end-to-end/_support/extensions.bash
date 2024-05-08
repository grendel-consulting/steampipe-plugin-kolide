#!/usr/bin/env bash

load_helpers(){
    # We want to check return codes for some tests, hence run flags need to be enabled
    bats_require_minimum_version 1.5.0

    # For now, presume running locally on MacBook with BATS installed via Homebrew
    TEST_HELPER_INSTALL_ROOT="$(brew --prefix)"

    load "${TEST_HELPER_INSTALL_ROOT}/lib/bats-assert/load.bash"
    load "${TEST_HELPER_INSTALL_ROOT}/lib/bats-file/load.bash"
}
