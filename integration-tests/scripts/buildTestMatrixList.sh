#!/usr/bin/env bash

# requires a path to a test file to compare the test list against
# requires a matrix job name to be passed in, for example "automation"
# requires a node label to be passed in, for example "ubuntu-latest"

set -e

# get this scripts directory
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

cd "$SCRIPT_DIR"/../ || exit 1

FILENAME=$1
MATRIX_JOB_NAME=$2
NODE_LABEL=$3

# Get list of test names from JSON file
JSONFILE="${FILENAME}_test_list.json"
COUNTER=1

# Build a JSON object in the format expected by our integration-tests workflow matrix
matrix_output() {
  local counter=$1
  local job_name=$2
  local test_name=$3
  local node_label=$4
  local counter_out=$(printf "%02d\n" $counter)
  echo -n "{\"name\": \"${job_name}-${counter_out}\", \"file\": \"${job_name}\",\"nodes\": 1, \"os\": \"${node_label}\", \"pyroscope_env\": \"ci-smoke-${job_name}-evm-simulated\", \"run\": \"-run '^${test_name}$'\"}"
}

# Read the JSON file and loop through 'tests' and 'run'
jq -c '.tests[]' ${JSONFILE} | while read -r test; do
  testName=$(echo ${test} | jq -r '.name')
  subTests=$(echo ${test} | jq -r '.run[]?.name // empty')
  output=""
  
  # Loop through subtests, if any, and print in the desired format
  if [ -n "$subTests" ]; then
    for subTest in $subTests; do
      if [ $COUNTER -ne 1 ]; then
        echo -n ","
      fi
      matrix_output $COUNTER $MATRIX_JOB_NAME "${testName}/${subTest}" ${NODE_LABEL}
      ((COUNTER++))
    done
  else
    if [ $COUNTER -ne 1 ]; then
      echo -n ","
    fi
    matrix_output $COUNTER $MATRIX_JOB_NAME "${testName}" ${NODE_LABEL}
    ((COUNTER++))
  fi

done > "./tmpout.json"
OUTPUT=$(cat ./tmpout.json)
echo "[${OUTPUT}]"
rm ./tmpout.json
