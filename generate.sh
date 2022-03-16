# required for running main (and examples, as main is used to generate code for examples)
# go run generate/*;

# required for running examples. initially generatin marshallers causes errors which we swallow
# easyjson -all -output_filename examples/application/server/gets_generated_easyjson.go examples/application/server/gets_generated.go &> /dev/null;
# easyjson -all -output_filename examples/application/server/message_easyjson.go examples/application/server/message.go &> /dev/null;
# go run . -engine_only -out=examples/application/server/ generate;
# easyjson -all -output_filename examples/engine/tree_easyjson.go examples/engine/tree.go;

# required for running unit tests
decltostring -input ./examples/server/ -output ./pkg/factory/server/stringified_server_decls.go -package server -exclude "test|easyjson";
decltostring -input ./examples/state/ -output ./pkg/factory/state/stringified_state_decls.go -package state -exclude "test|easyjson";
decltostring -input ./examples/message/ -output ./pkg/factory/message/stringified_message_decls.go -package message -exclude "test|easyjson";
decltostring -input ./examples/client/ -output ./pkg/factory/client/stringified_client_decls.go -package client -exclude "test|easyjson"; 

# required for running integration tests
# go run . -out=integrationtest/state/ generate;

