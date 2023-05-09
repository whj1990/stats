package repo

import (
	"context"
	"encoding/json"
	"mine/mine-grrpc/pbs"
	"time"

	"github.com/whj1990/app-mine/kitex_gen/api"
	"github.com/whj1990/go-core/handler"
	"github.com/whj1990/go-core/repository"
	"github.com/whj1990/go-core/util"
)

type ReviewProjectData struct {
	Id          int64  `form:"id" json:"id"` // 唯一标识
	Name        string //项目名称
	ModeCode    string //项目Mode code
	Description string //描述
	Status      int32  // 1-开启,2-关闭
	Deleted     int32  //1：删除
	CreatedUser string
	UpdatedUser string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*ReviewProjectData) TableName() string {
	return "rec_review_projects"
}

func (param *ReviewProjectData) getConditions() []*repository.Condition {
	return []*repository.Condition{
		repository.NewAndCondition(param),
	}
}

func (param *ReviewProjectData) listConditions() []*repository.Condition {
	var conditions []*repository.Condition
	return conditions
}

type ReviewProjectRepo interface {
	List(context.Context, *pbs.ReviewProjectListParams) (*[]ReviewProjectData, int64, error)
	Detail(ctx context.Context, req *api.IdsInt64Params) (*[]ReviewProjectData, error)
	TakeDetail(ctx context.Context, req *api.ReviewProjectData) (*ReviewProjectData, error)
	Add(context.Context, *api.ReviewProjectSaveData) (int64, error)
	Update(context.Context, *api.ReviewProjectSaveData) (int64, error)
	Delete(context.Context, *api.IdParam) (int64, error)
}

type reviewProjectRepo struct {
	repository.BaseRepo
}

func (r *reviewProjectRepo) List(ctx context.Context, req *pbs.ReviewProjectListParams) (*[]ReviewProjectData, int64, error) {
	var count int64
	var data []ReviewProjectData
	whereData := ReviewProjectData{
		Id:       req.Id,
		Name:     req.Name,
		ModeCode: req.ModeCode,
	}
	query := r.Db.Model(&ReviewProjectData{}).Where(whereData).Where("deleted = 0")
	if req.ShowStatus != 0 {
		query.Where("status = ? ", req.Status)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, count, handler.HandleError(err)
	}
	if err := query.Order("updated_at desc").Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&data).Error; err != nil {
		return nil, count, handler.HandleError(err)
	}
	return &data, count, nil
}

func (r *reviewProjectRepo) Detail(ctx context.Context, req *api.IdsInt64Params) (*[]ReviewProjectData, error) {
	data := make([]ReviewProjectData, 0)
	if len(req.Ids) == 0 {
		return nil, nil
	}
	err := r.Db.Model(&ReviewProjectData{}).Where("id in ?", req.Ids).Where("deleted = 0").Find(&data).Error
	if err != nil {
		return nil, handler.HandleError(err)
	}
	return &data, err
}

func (r *reviewProjectRepo) TakeDetail(ctx context.Context, req *api.ReviewProjectData) (*ReviewProjectData, error) {
	data := ReviewProjectData{}
	where := ReviewProjectData{}
	byteData, _ := json.Marshal(*req)
	err := json.Unmarshal(byteData, &where)
	if err != nil {
		return nil, handler.HandleError(err)
	}
	err = r.Db.Model(&ReviewProjectData{}).Where("deleted = 0").Take(&data, &where).Error
	if err != nil {
		return nil, handler.HandleError(err)
	}
	return &data, err
}

func (r *reviewProjectRepo) Add(ctx context.Context, req *api.ReviewProjectSaveData) (count int64, err error) {
	userName, err := util.GetMetaInfoCurrentUserName(ctx)
	if err != nil {
		return 0, err
	}
	data := ReviewProjectData{}
	dataByte, _ := json.Marshal(req)
	if err = json.Unmarshal(dataByte, &data); err != nil {
		return count, handler.HandleError(err)
	}
	data.CreatedUser = userName
	data.UpdatedUser = userName
	res := r.Db.Model(&ReviewProjectData{}).Create(&data)
	return res.RowsAffected, handler.HandleError(res.Error)
}

func (r *reviewProjectRepo) Update(ctx context.Context, req *api.ReviewProjectSaveData) (count int64, err error) {
	userName, err := util.GetMetaInfoCurrentUserName(ctx)
	if err != nil {
		return 0, err
	}
	data := ReviewProjectData{}
	dataByte, _ := json.Marshal(req)
	if err = json.Unmarshal(dataByte, &data); err != nil {
		return count, handler.HandleError(err)
	}
	data.UpdatedUser = userName
	res := r.Db.Model(&ReviewProjectData{}).Where("id = ? and deleted = 0", data.Id).Updates(&data)
	return res.RowsAffected, handler.HandleError(res.Error)
}
func (r *reviewProjectRepo) Delete(ctx context.Context, req *api.IdParam) (count int64, err error) {
	userName, err := util.GetMetaInfoCurrentUserName(ctx)
	if err != nil {
		return 0, err
	}
	res := r.Db.Model(&ReviewProjectData{}).Where("id = ? and deleted = 0", req.Id).Updates(&ReviewProjectData{
		Deleted:     1,
		UpdatedUser: userName,
	})
	return res.RowsAffected, handler.HandleError(res.Error)
}
