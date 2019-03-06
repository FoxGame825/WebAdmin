#!/bin/bash
protoc --go_out=./ s2s.proto
cp s2s.pb.go ../../wuxia_server/wxserver/src/shinerjoy.com/sspb
