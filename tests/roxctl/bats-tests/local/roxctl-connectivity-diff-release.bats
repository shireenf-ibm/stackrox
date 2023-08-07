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

@test "roxctl-release connectivity-diff no args" {
  run roxctl-release connectivity-diff
  assert_failure
  assert_line --partial "accepts 2 arg(s), received 0"
}

@test "roxctl-release connectivity-diff non existing dirs" {
  run roxctl-release connectivity-diff "$out_dir" "$out_dir"
  assert_failure
  assert_line --partial "error in connectivity diff analysis"
  assert_line --partial "no such file or directory"
}

@test "roxctl-release connectivity-diff stops on first error when run with --fail" {
  mkdir -p "$out_dir"
  write_yaml_to_file "$templated_fragment" "$(mktemp "$out_dir/templated-01-XXXXXX-file1.yaml")"
  write_yaml_to_file "$templated_fragment" "$(mktemp "$out_dir/templated-02-XXXXXX-file2.yaml")"

  run roxctl-release connectivity-diff "$out_dir/" "$out_dir/" --remove --output-file=/dev/null --fail
  assert_failure
  assert_output --partial 'YAML document is malformed'
  assert_output --partial 'file1.yaml'
  refute_output --partial 'file2.yaml'
}

@test "roxctl-release connectivity-diff produces no output when all yamls are templated" {
  mkdir -p "$out_dir"
  write_yaml_to_file "$templated_fragment" "$(mktemp "$out_dir/templated-XXXXXX.yaml")"

  echo "Analyzing a corrupted yaml file '$templatedYaml'" >&3
  run roxctl-release connectivity-diff "$out_dir/" "$out_dir/"
  assert_failure
  assert_output --partial 'YAML document is malformed'
  assert_output --partial 'no relevant Kubernetes resources found'
}

@test "roxctl-release connectivity-diff produces errors when some yamls are templated" {
  mkdir -p "$out_dir"
  write_yaml_to_file "$templated_fragment" "$(mktemp "$out_dir/templated-XXXXXX.yaml")"

  assert_file_exist "${test_data}/np-guard/scenario-minimal-service/frontend.yaml"
  assert_file_exist "${test_data}/np-guard/scenario-minimal-service/backend.yaml"
  cp "${test_data}/np-guard/scenario-minimal-service/frontend.yaml" "$out_dir/frontend.yaml"
  cp "${test_data}/np-guard/scenario-minimal-service/backend.yaml" "$out_dir/backend.yaml"

  echo "Analyzing a directory where 1/3 of yaml files are templated '$out_dir/'" >&3
  run roxctl-release connectivity-diff "$out_dir/" "$out_dir/" --remove --output-file=/dev/null
  assert_failure
  assert_output --partial 'YAML document is malformed'
  refute_output --partial 'no relevant Kubernetes resources found'
}

@test "roxctl-release connectivity-diff produces errors when all yamls are not K8s resources" {
  mkdir -p "$out_dir"
  assert_file_exist "${test_data}/np-guard/empty-yamls/empty.yaml"
  assert_file_exist "${test_data}/np-guard/empty-yamls/empty2.yaml"
  cp "${test_data}/np-guard/empty-yamls/empty.yaml" "$out_dir/empty.yaml"
  cp "${test_data}/np-guard/empty-yamls/empty2.yaml" "$out_dir/empty2.yaml"

  run roxctl-release connectivity-diff "$out_dir/" "$out_dir/" --remove --output-file=/dev/null
  assert_failure
  assert_output --partial 'Yaml document is not a K8s resource'
  assert_output --partial 'no relevant Kubernetes resources found'
  assert_output --partial 'ERROR:'
  assert_output --partial 'there were errors during execution'
}

diff_tests_dir="${BATS_TEST_DIRNAME}/../../../../roxctl/connectivity-diff/testdata/"

@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories default output format" {
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


@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories txt output" {
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
  run roxctl-release connectivity-diff "${dir1}" "${dir2}" --output-format=txt
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
    # partial is used to filter WARN and INFO messages
  assert_output --partial 'Connectivity diff:
source: payments/gateway[Deployment], destination: payments/visa-processor-v2[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload payments/visa-processor-v2[Deployment] added)
source: {ingress-controller}, destination: frontend/blog[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload frontend/blog[Deployment] added)
source: {ingress-controller}, destination: zeroday/zeroday[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload zeroday/zeroday[Deployment] added)'
}


@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories md output" {
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
  run roxctl-release connectivity-diff "${dir1}" "${dir2}" --output-format=md
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
  # partial is used to filter WARN and INFO messages
  assert_output --partial '| source | destination | dir1 | dir2 | diff-type |
|--------|-------------|------|------|-----------|
| payments/gateway[Deployment] | payments/visa-processor-v2[Deployment] | No Connections | TCP 8080 | added (workload payments/visa-processor-v2[Deployment] added) |
| {ingress-controller} | frontend/blog[Deployment] | No Connections | TCP 8080 | added (workload frontend/blog[Deployment] added) |
| {ingress-controller} | zeroday/zeroday[Deployment] | No Connections | TCP 8080 | added (workload zeroday/zeroday[Deployment] added) |'
}

@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories csv output" {
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
  run roxctl-release connectivity-diff "${dir1}" "${dir2}" --output-format=csv
  assert_success

  echo "$output" > "$ofile"
  assert_file_exist "$ofile"
    # partial is used to filter WARN and INFO messages
  assert_output --partial 'source,destination,dir1,dir2,diff-type
payments/gateway[Deployment],payments/visa-processor-v2[Deployment],No Connections,TCP 8080,added (workload payments/visa-processor-v2[Deployment] added)
{ingress-controller},frontend/blog[Deployment],No Connections,TCP 8080,added (workload frontend/blog[Deployment] added)
{ingress-controller},zeroday/zeroday[Deployment],No Connections,TCP 8080,added (workload zeroday/zeroday[Deployment] added)'
}

@test "roxctl-release connectivity-diff hould return error on not supported output format" {
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
  run roxctl-release connectivity-diff "${dir1}" "${dir2}" --output-format=png
  assert_failure

  assert_line --partial "error in formatting connectivity diff"
  assert_line --partial "png output format is not supported."
}

@test "roxctl-release connectivity-diff generates conns diff report between resources from two directories to output file" {
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
  run roxctl-release connectivity-diff "${dir1}" "${dir2}" --output-file="$out_dir/out.txt"
  assert_success

  assert_file_exist "$out_dir/out.txt"
  # partial is used to filter WARN and INFO messages
  assert_output --partial 'Connectivity diff:
source: payments/gateway[Deployment], destination: payments/visa-processor-v2[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload payments/visa-processor-v2[Deployment] added)
source: {ingress-controller}, destination: frontend/blog[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload frontend/blog[Deployment] added)
source: {ingress-controller}, destination: zeroday/zeroday[Deployment], dir1:  No Connections, dir2: TCP 8080, diff-type: added (workload zeroday/zeroday[Deployment] added)'
}

write_yaml_to_file() {
  image="${1}"
  templatedYaml="${2:-/dev/null}"
  cat >"$templatedYaml" <<-EOF
  cat $templatedYaml >&3
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
    spec:
      containers:
      - name: server
        image: $image
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        resources:
          requests:
            cpu: 100m
            memory: 64Mi
          limits:
            cpu: 200m
            memory: 128Mi
---
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  type: ClusterIP
  selector:
    app: frontend
  ports:
  - name: http
    port: 80
    targetPort: 8080
EOF
}
