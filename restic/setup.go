package restic

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	RESTIC_REPOSOTORY    = "RESTIC_REPOSITORY"
	RESITC_PASSWORD      = "RESTIC_PASSWORD"
	RESTIC_PROGRESS_FPS  = "RESTIC_PROGRESS_FPS"
	TMPDIR               = "TMPDIR"
	REST_SERVER_USERNAME = "REST_SERVER_USERNAME"
	REST_SERVER_PASSWORD = "REST_SERVER_PASSWORD"

	// S3 protocol params
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_DEFAULT_REGION    = "AWS_DEFAULT_REGION"

	OS_STORAGE_URL = "OS_STORAGE_URL"
	OS_AUTH_TOKEN  = "OS_AUTH_TOKEN"

	// CA_CERT_DATA for using certs in Minio server or REST server
	CA_CERT_DATA   = "CA_CERT_DATA"
	resticCacheDir = "restic-cache"
)

func (w *ResticWrapper) setupEnv() error {
	// Set progress report frequency.
	// 0.016666 is for one report per minute.
	w.sh.SetEnv(RESTIC_PROGRESS_FPS, "0.016666")

	if w.config.StorageSecret == nil {
		return errors.New("missing storage Secret")
	}

	if err := w.


}



func (w *ResticWrapper) exportSecretKey(key string, required bool) error {
	if v, ok := w.config.StorageSecret.Data[key]; !ok {
		if required {
			return fmt.Errorf("storage Secret missing %s key", key)
		}
	} else {
		w.sh.SetEnv(key, string(v))
	}
	return nil
}