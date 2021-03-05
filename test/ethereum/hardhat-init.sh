#!/bin/bash

set -e

cd "${0%/*}" # cd to current script dir
yarn
yarn hardhat accounts
