package main

import "os/exec"

func checkTerraformBinary() bool {
	_, err := exec.LookPath("terraform")
	if err != nil {
		_, err := exec.LookPath("tf")
		if err != nil {
			return false
		}
	}
	return true
}
