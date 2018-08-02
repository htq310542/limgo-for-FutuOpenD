
#!bin/sh
protoc -I . --go_out=import_path=futupb:. ./*.proto
