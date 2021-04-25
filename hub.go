package main;

import (
	"fmt";
	"os";
	"net";
	"time";
	"encoding/json";
	"io/ioutil";
);


type Address struct {
	Host string
	Port int16
};

type Config struct {
	Maps []ProxyMap
}

type ProxyMap struct {
	SrcPort int16
	DstAddr []Address
};


func sendBuf(dst Address, buf []byte, ret chan int) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", dst.Host, dst.Port));
	check(err);

	_, err = conn.Write(buf)
	check(err);

	ret <- 1
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	var cfg Config;
	count := 0;

	file, err := ioutil.ReadFile("./config.json");
	check(err);

	err = json.Unmarshal(file, &cfg);
	check(err);

	trigger := time.Tick(10 * time.Second);
	counter := make(chan int);

	for i := range(cfg.Maps) {
		proxymap := cfg.Maps[i];
		go proxy(proxymap.SrcPort, proxymap.DstAddr, counter);
	}

	for {
		select {
			case <- counter:
				count++;
			case <- trigger:
				fmt.Printf("count=%d\n", count);
		}
	}
}


func proxy(port int16, destinations []Address, count chan int) {
	sock, err := net.ListenPacket("udp", fmt.Sprintf(":%d", port));
	check(err);

	// select on socket
	for {
		var buf [65535]byte;
		rlen, _, err := sock.ReadFrom(buf[0:]);
		if err != nil {
			fmt.Println(err);
			os.Exit(1);
		}
		if rlen == 0 {
			continue;
		}
		// send to destinations
		for k := range(destinations) {
			dst := destinations[k];
			go sendBuf(dst, buf[:rlen], count);
		}
	}
}
