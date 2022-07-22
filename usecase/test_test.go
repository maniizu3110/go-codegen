package usecase

type TestTestUsecase interface {
	CommonUsecase[domain.TestTest]
}

// implement testTest usecase
type testTestUsecaseImpl struct {
	CommonUsecase[domain.TestTest]
	repo repository.TestTestRepository
}

func NewTestTestUsecase(commonUsecase CommonUsecase[domain.TestTest], repo repository.TestTestRepository) TestTestUsecase {
	res := &testTestUsecaseImpl{}
	res.CommonUsecase = commonUsecase
	res.repo = repo
	return res
}
