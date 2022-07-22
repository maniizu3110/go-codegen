// Code generated; DO NOT EDIT.
package usecase

type CommonUsecase[T domain.Models] interface {
	Fetch(config domain.FetchConfig) ([]*T, int64, myerror.MyError)
	GetByID(id uint, expand ...string) (*T, myerror.MyError)
	Store(data *T) (*T, myerror.MyError)
	Update(id uint, data *T) (*T, myerror.MyError)
	SoftDelete(id uint) (*T, myerror.MyError)
	HardDelete(id uint) (*T, myerror.MyError)
	Restore(id uint) (*T, myerror.MyError)
}

type commonUsecase[T domain.Models] struct {
	commonRepo repository.CommonRepository[T]
}

func NewCommonUsecase[T domain.Models](r repository.CommonRepository[T]) CommonUsecase[T] {
	return &commonUsecase[T]{
		commonRepo: r,
	}
}

// Fetch: configの内容に従った情報を全てとってくる．
func (cu *commonUsecase[T]) Fetch(config domain.FetchConfig) ([]*T, int64, myerror.MyError) {
	data, cnt, err := cu.commonRepo.Fetch(config)
	if err != nil {
		return nil, 0, err
	}
	return data, cnt, nil
}

// GetByID: 指定IDを持つデータを取ってくる．
func (cu *commonUsecase[T]) GetByID(id uint, expand ...string) (*T, myerror.MyError) {
	data, err := cu.commonRepo.GetByID(id, expand...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Store: データの新規作成
func (cu *commonUsecase[T]) Store(data *T) (*T, myerror.MyError) {
	return cu.commonRepo.Store(data)
}

// Update: 既存データの内容の更新
func (cu *commonUsecase[T]) Update(id uint, data *T) (*T, myerror.MyError) {
	return cu.commonRepo.Update(id, data)
}

// SoftDelete:指定データを論理削除（復元可）する．
func (cu *commonUsecase[T]) SoftDelete(id uint) (*T, myerror.MyError) {
	_, err := cu.commonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return cu.commonRepo.SoftDelete(id)
}

// HardDelete:指定データを物理削除（復元不可）する．
func (cu *commonUsecase[T]) HardDelete(id uint) (*T, myerror.MyError) {
	_, err := cu.commonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return cu.commonRepo.HardDelete(id)
}

// Restore:論理削除されたデータを復元する．
func (cu *commonUsecase[T]) Restore(id uint) (*T, myerror.MyError) {
	_, err := cu.commonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return cu.commonRepo.Restore(id)
}
