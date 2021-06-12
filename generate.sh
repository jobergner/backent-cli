decltostring -input ./serverfactory/server_example/server/ -output ./serverfactory/stringified_server_decls.go -package serverfactory -only "gets_generated.go";
decltostring -input ./enginefactory/state_engine_example/ -output ./enginefactory/stringified_state_engine_decls.go -package enginefactory -exclude "test|easyjson";
easyjson -all -omit_empty -output_filename enginefactory/state_engine_example/tree_easyjson.go enginefactory/state_engine_example/tree.go;
easyjson -all -omit_empty -output_filename serverfactory/server_example/server/gets_generated_easyjson.go serverfactory/server_example/server/gets_generated.go;
go run . -engine_only -out serverfactory/server_example/server/;
