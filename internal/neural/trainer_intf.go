// Representació general d'un tipus entrenador.
// CreatedAt: 2024/12/14 ds. JIQ

package neural

import (
	rF64 "github.com/jibort/ld_mcac/internal/core/rf64"
	dset "github.com/jibort/ld_mcac/internal/neural/dataset"
)

type TrainerIntf interface {
	LoadDataset(pDataset dset.Dataset)
	TrainEpochs(pEpochs int)
	EvaluateAccuracy(pTestSet dset.Dataset) (rAccuracy rF64.F64Range)
}
