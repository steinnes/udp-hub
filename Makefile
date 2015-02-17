all:
	go build hub.go

install: all
	cp hub /usr/local/bin/hub

clean:
	rm -f hub
