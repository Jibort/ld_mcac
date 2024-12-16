// Representació general d'un tipus entrenador.
// CreatedAt: 2024/12/14 ds. JIQ

package neural

type TrainerIntf interface {
	LoadDataset(pDataset Dataset)
	TrainEpochs(pEpochs int)
	EvaluateAccuracy(pTestSet Dataset) (rAccuracy RangeF64)
}
