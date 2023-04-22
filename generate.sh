# required for running main (and examples, as main is used to generate code for examples)
go run ./generate/static_code/*

# required for running examples
easyjson -all -omit_empty ./examples/client/ &> /dev/null;
easyjson -all -omit_empty ./examples/message/ &> /dev/null;
easyjson -all -omit_empty ./examples/server/ &> /dev/null;
easyjson -all -omit_empty ./examples/state/ &> /dev/null;

# required for running unit tests
decltostring -input ./examples/server/ -output ./pkg/factory/server/stringified_server_decls.go -package server -include "\.generated";
decltostring -input ./examples/state/ -output ./pkg/factory/state/stringified_state_decls.go -package state -include "\.generated";
decltostring -input ./examples/message/ -output ./pkg/factory/message/stringified_message_decls.go -package message -include "\.generated";
decltostring -input ./examples/client/ -output ./pkg/factory/client/stringified_client_decls.go -package client -include "\.generated"; 
go run ./generate/typescript_test_decls/*

