package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"reno/reno"
)

func renoManual() {
	fmt.Println(
		`
	Reno1 is a hashing-based system configured specifically for web devs.

	Usage:
		renoi
		renoi [help]

	Description:
		Shows basic configuration and purpose of Reno1. 
		Renoi helps with restarting the dev server anytime changes are made to the the 'views' and 'locales' folder and inside sub-apps. 
		All other files like .js .ts are all monitored by nodemon.

		Without you having to restart casa's server again, renoi just does that for you like your new autosave 
		with no configuration or complexity involved. Your server still works fine, so no worries
		
	Examples:
		renoi
			Reno1 starts watching all 'views' and 'locales' folders, and restarts server if any change is made.
			To quit, simply hit Ctrl-C or what you configured in your terminal as Reno1 exits gloriously.
	
	If you'd like to suggest more features, bug reports or new ideas, feel free to let us know or open up a PR
		`)

}

// checkCWD is just a safety gaurd to ensure that devs don't run renoi outside projects that are not node_based 
//
func checkCWD() {
	nodeEnv := []string{"node_modules", "package.json", "package-lock.json", "views"}

	for _, file := range nodeEnv {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			fmt.Printf("Invalid enviroment, %s does not exist, make sure you're in the right folder or make %s available\n", file, file)
			os.Exit(1)
		}
	}
}

func main() {

	if len(os.Args) > 1 && os.Args[1] == "help" {
		renoManual()
		os.Exit(0)
	} else if len(os.Args) > 1 && os.Args[1] != "help" {
		fmt.Print("fatal: Invalid usage or unsupported feature\n\n")
		renoManual()
		os.Exit(1)
	}

	checkCWD()
	cmd := exec.Command("bash", "-c", "npm run dev") // Go can only run Unix based binaries, so we have to tell bash to resolve the commands

	// Attaches the new shell to Go's stdin and stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		trackedFiles := reno.TrackFiles()
		initialHash := reno.ComputeFinalHash(trackedFiles)

		time.Sleep(1 * time.Second)

		secondTrack := reno.TrackFiles()
		newHash := reno.ComputeFinalHash(secondTrack)

		if initialHash != newHash {
			fmt.Println("[renoi]: detected change...restarting server")
			if cmd != nil && cmd.Process != nil {
				cmd.Process.Kill()
			}
			// starting up a new server
			cmd = exec.Command("bash", "-c", "npm run dev")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start()
		}
	}

}
