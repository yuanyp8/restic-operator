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

type RetentionPolicy struct {
	Name        string   `json:"name"`
	KeepLast    int64    `json:"keepLast,omitempty"`
	KeepHourly  int64    `json:"keepHourly,omitempty"`
	KeepDaily   int64    `json:"keepDaily,omitempty"`
	KeepWeekly  int64    `json:"keepWeekly,omitempty"`
	KeepMonthly int64    `json:"keepMonthly,omitempty"`
	KeepYearly  int64    `json:"keepYearly,omitempty"`
	KeepTags    []string `json:"keepTags,omitempty"`
	Prune       bool     `json:"prune"`
	DryRun      bool     `json:"dryRun,omitempty"`
}
