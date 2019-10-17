package main

import (
	"log"
	"os/exec"
)

func runAdditionalTasks() error {
	cmd := exec.Command("systemctl", "restart", "nginx")
	err := cmd.Run()
	log.Printf("Command ran with error: %v", err)
	return err
}
