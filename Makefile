all: test

test: 
		go test -v --cover

cover:
		go test -v -covermode=count -coverprofile=count.out
		go tool cover -html=count.out