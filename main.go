package main

import (
  "fmt"
  "net/http"
  "net"
	"os"
)

func main() {
  ipaddr := ""
  addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
        os.Stdout.WriteString(ipnet.IP.String() + "\n")
        ipaddr = ipnet.IP.String()
			}
		}
  }
  
  http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "Hello, world!\nIP Address ", ipaddr)
  })

  http.ListenAndServe(":8080", nil)
}
