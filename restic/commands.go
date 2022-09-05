package restic

import (
	"time"
)

const ResticCMD = "/bin/restic"

type Snapshot struct {
	ID       string    `json:"id"`
	Time     time.Time `json:"time"`
	Tree     string    `json:"tree"`
	Paths    []string  `json:"paths"`
	Hostname string    `json:"hostname"`
	Username string    `json:"username"`
	UID      int       `json:"uid"`
	Gid      int       `json:"gid"`
	Tags     []string  `json:"tags"`
}

type backupParams struct {
	path        string
	host        string
	snapshotId  string
	destination string
	excludes    []string
	includes    []string
	args        []string
}

func (w *ResticWrapper) listSnapshots(snapshotIDs []string) ([]Snapshot, error) {
	
}
