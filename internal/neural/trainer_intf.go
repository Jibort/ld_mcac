// Representació general d'un tipus entrenador.
// CreatedAt: 2024/12/14 ds. JIQ

package neural

import (
	"github.com/jibort/ld_mcac/internal/core"
)

type TrainerIntf interface {
	LoadDataset(pDataset Dataset)
	TrainEpochs(pEpochs int)
	EvaluateAccuracy(pTestSet Dataset) (rAccuracy core.RangeF64)
}
