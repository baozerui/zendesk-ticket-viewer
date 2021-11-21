all:
	go build
dev: 
	go run main.go network.go show.go const.go prompt.go ticket.go
clean:
	rm -f main