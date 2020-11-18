/*Author: Alessandro Festa
Infra package allow a user to install a local cluster based on:
K3s
K0s
Kind
*/
package infra

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// K0s check the OS flavor and provide an input to the subsequent functions
func K0s(osFlavor string, infraSelection string) {
	// where are we? If windows we have to call wsl function if not proceed
	switch osFlavor {
	case "windows":
		infraK0sWSL()
	case "linux":
		infraK0sDefault()
	case "arm":
		infraK0sARM()
	default:
		infraK0sDefault()
	}
	// check if K3s is in the path
	// if k3s is not in the path download it
	// now we do install with the default flags

}

func infraK0sWSL() {
	// we are in WSL so we cannot use the default installer
	cmd := exec.Command("bash", "-c", "curl -Lo ./k0s https://github.com/k0sproject/k0s/releases/download/v0.7.0/k0s-v0.7.0-amd64; chmod +x ./k0s; sudo mv ./k0s /usr/local/bin;mkdir -p ${HOME}/.k0s")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func infraK0sDefault() {
	// Let's download and install K3s the usual way
	cmd := exec.Command("/bin/sh", "-c", "curl -Lo ./k0s https://github.com/k0sproject/k0s/releases/download/v0.7.0/k0s-v0.7.0-amd64; chmod +x ./k0s; sudo mv ./k0s /usr/local/bin;mkdir -p ${HOME}/.k0s")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	launchK0sFile()
}

func infraK0sARM() {
	//ARM installation
}

func runK0sDefault() {
	fmt.Println("file written successfully")
	cmd := exec.Command("/bin/bash", "-c", "sudo mv ./start.sh ${HOME}/.k0s/; chmod +x ${HOME}/.k0s/start.sh ;. /${HOME}/.k0s/start.sh")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func launchK0sFile() {
	err := ioutil.WriteFile("start.sh", []byte("#!/bin/bash\n"), 0644)
	f, err := os.Create("start.sh")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"#!/bin/bash", "echo 'Installing K0s...'", "sleep 5", "sudo k0s server -c ${HOME}/.k0s/k0s.yaml --enable-worker > /dev/null 2>&1 &", "echo 'Configuring last steps...'", "sleep 5", "echo 'Copy the following line and paste it to your session to use K0s'", "echo 'export KUBECONFIG=/var/lib/k0s/pki/admin.conf'"}

	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	runK0sDefault()
}
