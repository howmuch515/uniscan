package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func file_exists_check(path string) bool {
	// if file exists, return true.
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	target_domain_name := os.Args[1]
	subfinder_filename := "subfinder_result"
	subjack_filename := "subjack_result"
	aquatone_dirname := "aquatone_result"
	chromium_path := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	var err error

	// subfinder
	fmt.Println("[+] subfinder scan")
	if file_exists_check(subfinder_filename) {
		fmt.Println("[!] result file is exists! if you rescan, please remove the result file.")
		os.Exit(1)
	} else {
		err = exec.Command("subfinder", "-o", subfinder_filename, "-d", target_domain_name).Run()

		if err != nil {
			log.Println("Error in exec subfinder")
			log.Fatal(err)
		}
	}

	// subjack
	fmt.Println("[+] subjack scan")
	os.Stat(subfinder_filename)
	err = exec.Command("subjack", "-w", subfinder_filename, "-o", subjack_filename, "-timeout", "30", "-v", "-ssl").Run()

	if err != nil {
		log.Println("Error in exec subjack")
		log.Fatal(err)
	}

	// aquaton
	fmt.Println("[+] aquatone scan")
	aquatone_cmd := fmt.Sprintf("cat %s | aquatone -chrome-path '%s' -out %s", subfinder_filename, chromium_path, aquatone_dirname)

	err = exec.Command("sh", "-c", aquatone_cmd).Run()

	if err != nil {
		log.Println("Error in exec aquatone")
		log.Fatal(err)
	}
}
