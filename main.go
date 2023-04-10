package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var serviceName string = os.Getenv("SERVICE_NAME")

func main() {

	fmt.Printf("Stop systemd service %s.\n", serviceName)
	serviceStop, err := exec.Command("systemctl", "stop", serviceName).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(serviceStop))

	fmt.Printf("Clear docker cache: \"docker system prune\".\n")
	dockerSystem, err := exec.Command("docker", "system", "prune", "-f").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dockerSystem))

	fmt.Printf("Clear docker cache: \"docker volume rm\".\n")
	dockerVolumes, err := exec.Command("docker", "volume", "list", "-q").Output()
	words := strings.Split(string(dockerVolumes), "\n")
	for idx, volume := range words {
		exec.Command("docker", "volume", "rm", "-f", string(volume)).Output()
		fmt.Println(idx, volume)
	}

	fmt.Printf("Start systemd service %s.\n", serviceName)
	serviceStart, err := exec.Command("systemctl", "start", serviceName).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(serviceStart))

}
