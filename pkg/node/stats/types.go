package stats

// Stats represents the stats.
type Stats struct {
	CPUs int   `json:"cpus"`
	MEM  int64 `json:"mem"`
	DISK int64 `json:"disk"`
}
