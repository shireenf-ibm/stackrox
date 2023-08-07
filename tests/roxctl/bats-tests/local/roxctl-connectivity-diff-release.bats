#!/usr/bin/env bats

load "../helpers.bash"
out_dir=""
templated_fragment='"{{ printf "%s" ._thing.image }}"'

setup_file() {
    [[ -n "$NO_BATS_ROXCTL_REBUILD" ]] || rm -f "${tmp_roxctl}"/roxctl*
    echo "Testing roxctl version: '$(roxctl-release version)'" >&3
}

setup() {
  out_dir="$(mktemp -d -u)"
  ofile="$(mktemp)"
}

teardown() {
  rm -rf "$out_dir"
  rm -f "$ofile"
}

diff_tests_dir="${BATS_TEST_DIRNAME}/../../../../roxctl/connectivity-diff/testdata/"

@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories" {
  dir1="${diff_tests_dir}/acs-security-demos/"
  dir2="${diff_tests_dir}/acs-security-demos-new-version/"
  # assert files exist in dir1
  assert_file_exist "${dir1}/backend/catalog/deployment.yaml"
  assert_file_exist "${dir1}/backend/checkout/configmap.yaml"
  assert_file_exist "${dir1}/backend/checkout/deployment.yaml"
  assert_file_exist "${dir1}/backend/notification/deployment.yaml"
  assert_file_exist "${dir1}/backend/recommendation/configmap.yaml"
  assert_file_exist "${dir1}/backend/recommendation/deployment.yaml"
  assert_file_exist "${dir1}/backend/reports/configmap.yaml"
  assert_file_exist "${dir1}/backend/reports/deployment.yaml"
  assert_file_exist "${dir1}/backend/shipping/deployment.yaml"
  assert_file_exist "${dir1}/frontend/asset-cache/deployment.yaml"
  assert_file_exist "${dir1}/frontend/asset-cache/route.yaml"
  assert_file_exist "${dir1}/frontend/webapp/configmap.yaml"
  assert_file_exist "${dir1}/frontend/webapp/deployment.yaml"
  assert_file_exist "${dir1}/frontend/webapp/route.yaml"
  assert_file_exist "${dir1}/payments/gateway/deployment.yaml"
  assert_file_exist "${dir1}/payments/mastercard-processor/deployment.yaml"
  assert_file_exist "${dir1}/payments/visa-processor/deployment.yaml"
  assert_file_exist "${dir1}/acs_netpols.yaml"
  # assert files exist in dir2
  assert_file_exist "${dir2}/backend/catalog/deployment.yaml"
  assert_file_exist "${dir2}/backend/checkout/configmap.yaml"
  assert_file_exist "${dir2}/backend/checkout/deployment.yaml"
  assert_file_exist "${dir2}/backend/notification/deployment.yaml"
  assert_file_exist "${dir2}/backend/recommendation/configmap.yaml"
  assert_file_exist "${dir2}/backend/recommendation/deployment.yaml"
  assert_file_exist "${dir2}/backend/reports/configmap.yaml"
  assert_file_exist "${dir2}/backend/reports/deployment.yaml"
  assert_file_exist "${dir2}/backend/namespace.yaml"
  assert_file_exist "${dir2}/backend/shipping/deployment.yaml"
  assert_file_exist "${dir2}/frontend/asset-cache/deployment.yaml"
  assert_file_exist "${dir2}/frontend/asset-cache/route.yaml"
  assert_file_exist "${dir2}/frontend/blog/deployment.yaml"
  assert_file_exist "${dir2}/frontend/blog/route.yaml"
  assert_file_exist "${dir2}/frontend/namespace.yaml"
  assert_file_exist "${dir2}/frontend/webapp/configmap.yaml"
  assert_file_exist "${dir2}/frontend/webapp/deployment.yaml"
  assert_file_exist "${dir2}/frontend/webapp/route.yaml"
  assert_file_exist "${dir2}/payments/gateway/deployment.yaml"
  assert_file_exist "${dir2}/payments/mastercard-processor/deployment.yaml"
  assert_file_exist "${dir2}/payments/visa-processor/deployment.yaml"
  assert_file_exist "${dir2}/payments/visa-processor-v2/deployment.yaml"
  assert_file_exist "${dir2}/payments/namespace.yaml"
  assert_file_exist "${dir2}/zeroday/deployment.yaml"
  assert_file_exist "${dir2}/zeroday/namespace.yaml"
  assert_file_exist "${dir2}/zeroday/route.yaml"
  assert_file_exist "${dir2}/acs_netpols.yaml"
  echo "Writing diff report to ${ofile}" >&3
  run roxctl-release connectivity-diff "${dir1}" "${dir2}"
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
    # partial is used to filter WARN and INFO messages
  assert_output --partial 'Connectivity diff:
source: payments/gateway[Deployment], destination: payments/visa-processor-v2[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload payments/visa-processor-v2[Deployment] added)
source: {ingress-controller}, destination: frontend/blog[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload frontend/blog[Deployment] added)
source: {ingress-controller}, destination: zeroday/zeroday[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload zeroday/zeroday[Deployment] added)'
}
