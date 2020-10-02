package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bufio"
)

func main() {
	domain := strings.Join(os.Args[1:], "")

	if len(os.Args) > 1 {
		fmt.Println(domain)
		cmd := "host -t ns " + domain + " | awk '{print $4}'"
		output, err := exec.Command("bash", "-c", cmd).CombinedOutput()

		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

 		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan(){
			fmt.Println("[*] Checking for Zone Transfer...  " + string(scanner.Text()) + " ==> " +  string(domain))
			cmd := "host -l " + string(domain) + " " + string(scanner.Text())
			output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			fmt.Println(string(output))
		}


	} else {
		fmt.Println("[*] Simple Zone Transfer")
		fmt.Println("[*] Usage: ./dns-axfr <domain name>\n")
	}

}
