.PHONY: clean copy gen

clean:
	rm -rf expected/go/*
	rm -rf expected/package/go/*

copy:
	rm -rf expected/*
	cp -r result/* expected

prepare:
	mkdir -p expected/package/go

pkg:
	./testcli gen --lang=go expected/package/go proto/package/common.proto
	./testcli gen --lang=go expected/package/go proto/package/gopackage_addReqFull.proto
	./testcli gen --lang=go expected/package/go proto/package/gopackage_addReq.proto
	./testcli gen --lang=go expected/package/go proto/package/gopackage_calcFull.proto
	./testcli gen --lang=go expected/package/go proto/package/gopackage_calc.proto
	./testcli gen --lang=go expected/package/go proto/package/gopackage_calc_warn.proto
	./testcli gen --lang=go expected/package/go proto/package/mixpackage_addReq.proto
	./testcli gen --lang=go expected/package/go proto/package/mixpackage_calc.proto
	./testcli gen --lang=go expected/package/go proto/package/nopackage_calc.proto
	./testcli gen --lang=go expected/package/go proto/package/nopackage_calc_warn.proto
	./testcli gen --lang=go expected/package/go proto/package/package_addReq.proto
	./testcli gen --lang=go expected/package/go proto/package/package_calc_commonerror.proto
	./testcli gen --lang=go expected/package/go proto/package/package_calc.proto
	./testcli gen --lang=go expected/package/go proto/package/package_calc._without_commonerror.proto
