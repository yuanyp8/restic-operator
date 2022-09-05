package restic

import (
	v1 "github.com/yuanyp8/restic-operator/api/v1"
	shell "gomodules.xyz/go-sh"
	core "k8s.io/api/core/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

const (
	DefaultOutputFileName = "output.json"
	DefaultScratchDir     = "/tmp"
	DefaultHost           = "host-0"
)

type ResticWrapper struct {
	sh     *shell.Session
	config SetupOptions
}

type SetupOptions struct {
	Provider       string // define the backup items, e.g. S3/Minio/OSS/Local Storage
	Bucket         string // S3 buckets
	Endpoint       string // S3 Endpoints
	Region         string // S3 Region
	Path           string
	CacertFile     string
	ScratchDir     string // restic binary file directory
	EnableCache    bool
	MaxConnections int64
	StorageSecret  *core.Secret
	Nice           *ofst.NiceSettings
	IONice         *ofst.IONiceSettings
}

type Command struct {
	Name string
	Args []string
}

type BackupOptions struct {
	Host              string
	BackupPaths       []string
	StdinPipeCommands []Command
	StdinFileName     string // default as "stdin"
	RetentionPolicy   v1.RetentionPolicy
	Exclude           []string
	Args              []string
}

type RestoreOptions struct {
	Host         string
	SourceHost   string
	RestorePaths []string
	Snapshots    []string
	Destination  string
	Exclude      []string
	Include      []string
	Args         []string
}

type DumpOptions struct {
	Host               string
	SourceHost         string
	Snapshot           string
	Path               string
	FileName           string
	StdoutPipeCommands []Command
}

func NewResticWrapper(options SetupOptions) (*ResticWrapper, error) {
	wrapper := &ResticWrapper{
		sh:     shell.NewSession(),
		config: options,
	}

	if err := wrapper.configure(); err != nil {
		return nil, err
	}

	return wrapper, nil
}

func (w *ResticWrapper) configure() error {
	w.sh.SetDir(w.config.ScratchDir)
	// enable for debug
	w.sh.ShowCMD = true
	w.sh.PipeFail = true
	w.sh.PipeStdErrors = true

	// setup restic environment
	return w.setupEnv()
}
