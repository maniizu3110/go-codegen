package usecase

type TestUsecase interface {
	CommonUsecase[domain.Test]
}

// implement test usecase
type testUsecaseImpl struct {
	CommonUsecase[domain.Test]
	repo repository.TestRepository
}

func NewTestUsecase(commonUsecase CommonUsecase[domain.Test], repo repository.TestRepository) TestUsecase {
	res := &testUsecaseImpl{}
	res.CommonUsecase = commonUsecase
	res.repo = repo
	return res
}
