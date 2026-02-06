//go:build windows

package common

import "C"

import (
	"log"
	"syscall"
)

const (
	ES_CONTINUOUS       = 0x80000000
	ES_SYSTEM_REQUIRED  = 0x00000001
	ES_DISPLAY_REQUIRED = 0x00000002
)

var (
	kernel32                = syscall.NewLazyDLL("kernel32.dll")
	setThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")
)

func PreventSleep() {
	log.Println("Disabling sleep mode")
	setThreadExecutionState.Call(ES_CONTINUOUS | ES_SYSTEM_REQUIRED)
}

func AllowSleep() {
	log.Println("Restoring sleep mode")
	setThreadExecutionState.Call(ES_CONTINUOUS)
}
