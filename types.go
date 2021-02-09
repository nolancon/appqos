package appqos

// Pool - AppQoS pool; equates to a single node representation of a
// K8s PowerWorkload object. PowerWorkload will contain data such as
// Nodes and/or NodeSelector labels.
type Pool struct {
	Name         *string       `json:"name,omitempty"`
	ID           *int          `json:"id,omitempty"`
	Apps         *[]App        `json:"apps,omitempty"`
	Cbm          *int          `json:"cbm,omitempty"`
	Mba          *int          `json:"mba,omitempty"`
	MbaBw        *int          `json:"mba_bw,omitempty"`
	Cores        *[]int        `json:"cores,omitempty"`
	PowerProfile *PowerProfile `json:"power_profile,omitempty"`
}

// Power Profile - AppQoS power_profile; equates to a K8s PowerProfile object
type PowerProfile struct {
	ID      *int    `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	MinFreq *int    `json:"min_freq,omitempty"`
	MaxFreq *int    `json:"max_freq,omitempty"`
	Epp     *string `json:"epp,omitempty"`
}

// App - Not necessary for power operator, added for completeness
type App struct {
	PoolID *int    `json:"pool_id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Cores  *[]int  `json:"cores,omitempty"`
	Pids   *[]int  `json:"pids,omitempty"`
}
