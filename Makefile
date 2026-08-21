test-lexer:
	clear && go test ./frontend/lexer -v

test-parser:
	clear && go test ./frontend/parser -v

test-ir-pipeline:
	clear && go test ./middleware/... -v

run:
	go run main.go source/lexer.gob