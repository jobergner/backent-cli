# required for running main (and examples, as main is used to generate code for examples)
go run generate/*;

# required for running examples. initially generatin marshallers causes errors which we swallow
easyjson -all -output_filename examples/application/server/gets_generated_easyjson.go examples/application/server/gets_generated.go &> /dev/null;
easyjson -all -output_filename examples/application/server/message_easyjson.go examples/application/server/message.go &> /dev/null;
go run . -engine_only -out=examples/application/server/ generate;
easyjson -all -output_filename examples/engine/tree_easyjson.go examples/engine/tree.go;

# required for running unit tests
decltostring -input ./examples/application/server/ -output ./pkg/factory/server/stringified_server_decls.go -package serverfactory -only "gets_generated.go";
decltostring -input ./examples/engine/ -output ./pkg/factory/engine/stringified_state_engine_decls.go -package engine -exclude "test|easyjson";

# required for running integration tests
go run . -out=integrationtest/state/ generate;

