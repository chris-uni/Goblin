test-lexer:
	clear && go test ./frontend/lexer -v

run:
	go run main.go source/lexer.gob