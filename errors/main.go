package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Errors in Golang")
	func() {
		f, err := os.Open("/values.json")
		if err != nil {

			if pErr, ok := err.(*os.PathError); ok {
				fmt.Println("Failed to open the file at path", pErr.Path)
				return
			}
			fmt.Println("Generic error message", err)
			return
		}
		fmt.Println(f.Name(), "opened successfully")
	}()

	func() {
		addr, netErr := net.LookupHost("gwen.stacy.com")
		if netErr != nil {
			if dnsErr, ok := netErr.(*net.DNSError); ok {
				if dnsErr.Timeout() {
					fmt.Println("operation timed out", netErr)
					return
				}
				if dnsErr.Temporary() {
					fmt.Println("temporary error", netErr)
					return
				}
				fmt.Println("generic DNS error", netErr)
				return
			}
			fmt.Println("generic error", netErr)
			return
		}
		fmt.Println(addr)
	}()

	func() {
		files, err := filepath.Glob("\\")
		if err != nil {
			if err == filepath.ErrBadPattern {
				fmt.Println("bad pattern error:", err)
				return
			}
			fmt.Println("generic error:", err)
			return
		}
		fmt.Println("matched file", files)
	}()
}
