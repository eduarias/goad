package oslimits

import (
	"log"
	"syscall"
)

// GetOSSimultaneousFileDescriptors returns the maximum simultaneous request that OS can handle with current configuration
func GetOSSimultaneousFileDescriptors() (uint64, error) {
	fileDescriptors, err := getRlimitFileDescriptors()
	return fileDescriptors.Cur, err
}

// SetMaxOSSimultaneousFileDescriptors extend to hard limits the maximum simultaneous request that OS can handle and returns its value
func SetMaxOSSimultaneousFileDescriptors() error {
	var newLimit syscall.Rlimit
	newLimit, err := getRlimitFileDescriptors()
	if err != nil {
		return err
	}
	newLimit.Cur = newLimit.Max
	err = setRlimitFileDescriptors(newLimit)
	log.Printf("File descriptor limit set to: %d", newLimit.Max)
	return err
}

func getRlimitFileDescriptors() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return rLimit, err
}

func setRlimitFileDescriptors(newLimit syscall.Rlimit) error {
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &newLimit)
	return err
}
