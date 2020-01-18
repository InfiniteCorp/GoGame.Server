all: build run

lib:
	go build -buildmode=plugin -o gameslib/tictactoe.so games/tictactoe.go

build: lib
	go build -o GoGame

run:
	./GoGame

clean:
	rm -f GoGame