package net_te

import "testing"

func TestTcp(t *testing.T) {
	scann := InitConns()
	scann.ScanningPorts("scanme.nmap.org", 100)
}
