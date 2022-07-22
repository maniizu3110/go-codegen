package usecase

type SampleUsecase interface {
	CommonUsecase[domain.Sample]
}

// implement sample usecase
type sampleUsecaseImpl struct {
	CommonUsecase[domain.Sample]
	repo repository.SampleRepository
}

func NewSampleUsecase(commonUsecase CommonUsecase[domain.Sample], repo repository.SampleRepository) SampleUsecase {
	res := &sampleUsecaseImpl{}
	res.CommonUsecase = commonUsecase
	res.repo = repo
	return res
}
