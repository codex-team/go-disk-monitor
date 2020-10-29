// +build darwin

package main

import "syscall"

// Data amount constants
const (
	B  = 1         // Bytes
	KB = 1024 * B  // Kilobytes
	MB = 1024 * KB // Megabytes
	GB = 1024 * MB // Gigabytes
)

// DiskStatus - structure with information about disk usage
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// DiskUsage - disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}

	// calculate number of blocks that are reserved by the filesystem
	fsReservedBlocks := fs.Bfree - fs.Bavail

	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bavail * uint64(fs.Bsize)
	disk.Used = (fs.Blocks-fsReservedBlocks)*uint64(fs.Bsize) - disk.Free

	return
}
