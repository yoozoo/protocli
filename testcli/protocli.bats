#!/usr/bin/env bats

@test "packagetest.proto go output" {
  mkdir -p result/package/go

  ./testcli gen --lang=go result/package/go proto/package/common.proto
  ./testcli gen --lang=go result/package/go proto/package/gopackage_addReqFull.proto
  ./testcli gen --lang=go result/package/go proto/package/gopackage_addReq.proto
  ./testcli gen --lang=go result/package/go proto/package/gopackage_calcFull.proto
  ./testcli gen --lang=go result/package/go proto/package/gopackage_calc.proto
  ./testcli gen --lang=go result/package/go proto/package/gopackage_calc_warn.proto
  ./testcli gen --lang=go result/package/go proto/package/mixpackage_addReq.proto
  ./testcli gen --lang=go result/package/go proto/package/mixpackage_calc.proto
  ./testcli gen --lang=go result/package/go proto/package/nopackage_calc.proto
  ./testcli gen --lang=go result/package/go proto/package/nopackage_calc_warn.proto
  ./testcli gen --lang=go result/package/go proto/package/package_addReq.proto
  ./testcli gen --lang=go result/package/go proto/package/package_calc_commonerror.proto
  ./testcli gen --lang=go result/package/go proto/package/package_calc.proto
  ./testcli gen --lang=go result/package/go proto/package/package_calc._without_commonerror.proto

  diff -r result/package/go/ expected/package/go/
}
