// 1. get the user setup from yaml file
// 2. set the git user in the repo
// 2a. check where the command has been invoked
// 2b. run setGit()

// 3. create new repo

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Main struct {
		Name  string
		Email string
	}
	Work struct {
		Name  string
		Email string
	}
}

func main() {
	wd := checkWhereIAm()
	fmt.Printf("setting user in %s\n", wd)
	users := readYaml()

	if len(os.Args) == 1 {
		fmt.Println("Provide user for which I have to set up credentials")
		os.Exit(1)
	}

	args := os.Args[1]

	switch args {
	case "main":
		setGit(users.Main.Name, users.Main.Email)
		fmt.Printf("Git config set for %s\n", users.Main.Name)
	case "work":
		setGit(users.Work.Name, users.Work.Email)
		fmt.Printf("Git config set for %s\n", users.Work.Name)
	default:
		fmt.Printf("No setup for this user")
		os.Exit(1)
	}

}

func setGit(user, email string) {
	exec.Command("git", "config", "user.name", user).Run()
	exec.Command("git", "config", "user.email", email).Run()
}

func checkWhereIAm() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func readYaml() Yaml {
	users := Yaml{}

	usr, _ := user.Current()
	filePath := usr.HomeDir + "/setup.yaml"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error reading file: %s\n", err)
	}

	err = yaml.Unmarshal(data, &users)
	if err != nil {
		log.Fatalf("error unmarshalling: %s\n", err)
	}
	return users
}
