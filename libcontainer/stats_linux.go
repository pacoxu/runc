package libcontainer

import (
	"github.com/pacoxu/runc/libcontainer/cgroups"
	"github.com/pacoxu/runc/libcontainer/intelrdt"
	"github.com/pacoxu/runc/types"
)

type Stats struct {
	Interfaces    []*types.NetworkInterface
	CgroupStats   *cgroups.Stats
	IntelRdtStats *intelrdt.Stats
}
