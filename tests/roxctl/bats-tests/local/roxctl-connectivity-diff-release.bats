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
  dir2="${diff_tests_dir}/acs-security-demos-new/"
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
  assert_file_exist "${dir2}/backend/shipping/deployment.yaml"
  assert_file_exist "${dir2}/external/new_deployment.yaml"
  assert_file_exist "${dir2}/frontend/asset-cache/deployment.yaml"
  assert_file_exist "${dir2}/frontend/asset-cache/route.yaml"
  assert_file_exist "${dir2}/frontend/webapp/configmap.yaml"
  assert_file_exist "${dir2}/frontend/webapp/deployment.yaml"
  assert_file_exist "${dir2}/frontend/webapp/route.yaml"
  assert_file_exist "${dir2}/payments/gateway/deployment.yaml"
  assert_file_exist "${dir2}/payments/visa-processor/deployment.yaml"
  assert_file_exist "${dir2}/acs_netpols.yaml"
  echo "Writing diff report to ${ofile}" >&3
  run roxctl-release connectivity-diff "${dir1}" "${dir2}"
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
    # partial is used to filter WARN and INFO messages
  assert_output --partial 'source: backend/reports[Deployment], destination: backend/catalog[Deployment], dir1:  TCP 8080, dir2: TCP 9080, diff-type: changed
source: 0.0.0.0-255.255.255.255, destination: external/unicorn[Deployment], dir1:  No Connections, dir2: All Connections, diff-type: added (workload external/unicorn[Deployment] added)
source: backend/checkout[Deployment], destination: external/unicorn[Deployment], dir1:  No Connections, dir2: UDP 5353, diff-type: added (workload external/unicorn[Deployment] added)
source: backend/recommendation[Deployment], destination: external/unicorn[Deployment], dir1:  No Connections, dir2: UDP 5353, diff-type: added (workload external/unicorn[Deployment] added)
source: backend/reports[Deployment], destination: external/unicorn[Deployment], dir1:  No Connections, dir2: UDP 5353, diff-type: added (workload external/unicorn[Deployment] added)
source: external/unicorn[Deployment], destination: 0.0.0.0-255.255.255.255, dir1:  No Connections, dir2: All Connections, diff-type: added (workload external/unicorn[Deployment] added)
source: external/unicorn[Deployment], destination: frontend/webapp[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload external/unicorn[Deployment] added)
source: frontend/webapp[Deployment], destination: external/unicorn[Deployment], dir1:  No Connections, dir2: UDP 5353, diff-type: added (workload external/unicorn[Deployment] added)
source: payments/gateway[Deployment], destination: external/unicorn[Deployment], dir1:  No Connections, dir2: UDP 5353, diff-type: added (workload external/unicorn[Deployment] added)
source: frontend/webapp[Deployment], destination: backend/shipping[Deployment], dir1:  TCP 8080, dir2: No Connections, diff-type: removed
source: payments/gateway[Deployment], destination: payments/mastercard-processor[Deployment], dir1:  TCP 8080, dir2: No Connections, diff-type: removed (workload payments/mastercard-processor[Deployment] removed)
source: {ingress-controller}, destination: frontend/asset-cache[Deployment], dir1:  TCP 8080, dir2: No Connections, diff-type: removed'
  # a warn added by np-guard analysis
assert_output --partial 'Route resource frontend/asset-cache specified workload frontend/asset-cache[Deployment] as a backend, but network policies are blocking ingress connections from an arbitrary in-cluster source to this workload.Connectivity map will not include a possibly allowed connection between the ingress controller and this workload.'
}
