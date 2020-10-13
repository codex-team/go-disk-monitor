// +build darwin

package main

import "syscall"

const (
	// Bytes
	B = 1
	// Kilobytes
	KB = 1024 * B
	// Megabytes
	MB = 1024 * KB
	// Gigabytes
	GB = 1024 * MB
)

// structure with information about disk usage
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bavail * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
