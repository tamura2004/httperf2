package infra

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Netstat() {
	ff := NewFileFactory()
	file := ff.CreateTeeFile("NETSTAT", "20060102150405", "csv")
	fmt.Fprintln(file, "DATE,TIME,TYPE,HOSTNAME,ESTABLISHED,CLOSE_WAIT,SYN_SENT")
	hostname := GetHostname()

	for {
		dic := numEstablished()
		fmt.Fprintf(
			file,
			"%s,NETSTAT,%s,%d,%d,%d\n",
			time.Now().Format("2006-01-02,15:04:05"),
			hostname,
			dic["ESTABLISHED"],
			dic["CLOSE_WAIT"],
			dic["SYN_SENT"],
		)
		time.Sleep(10 * time.Second)
	}
}

var status []string = []string{
	"ESTABLISHED",
	"CLOSE_WAIT",
	"SYN_SENT",
}

func numEstablished() map[string]int {
	out, err := exec.Command("netstat", "-na").Output()
	if err != nil {
		log.Fatal(err)
	}
	r := strings.NewReader(string(out))
	s := bufio.NewScanner(r)

	dic := make(map[string]int)
	for s.Scan() {
		for _, st := range status {
			if strings.Index(s.Text(), st) != -1 {
				dic[st]++
			}
		}
	}
	return dic
}
