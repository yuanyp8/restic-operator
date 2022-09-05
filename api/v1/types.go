package v1

// BackupInvokerRef contains information that points to the backup configuration or batch being used
type BackupInvokerRef struct {
	// APIGroup is the group for the resource being referenced
	APIGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

type TargetRef struct {
	APIVersion string `json:"apiVersion,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	//+optional
	Namespace string `json:"namespace,omitempty"`
}
