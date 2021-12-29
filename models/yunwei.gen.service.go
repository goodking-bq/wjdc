package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ServiceMgr struct {
	*_BaseMgr
}

// ServiceMgr open func
func ServiceMgr(db *gorm.DB) *_ServiceMgr {
	if db == nil {
		panic(fmt.Errorf("ServiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ServiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("service"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ServiceMgr) GetTableName() string {
	return "service"
}

// Reset 重置gorm会话
func (obj *_ServiceMgr) Reset() *_ServiceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ServiceMgr) Get() (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ServiceMgr) Gets() (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ServiceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Service{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ServiceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_ServiceMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEnName en_name获取
func (obj *_ServiceMgr) WithEnName(enName string) Option {
	return optionFunc(func(o *options) { o.query["en_name"] = enName })
}

// WithActive active获取
func (obj *_ServiceMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithCreateTime create_time获取
func (obj *_ServiceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_ServiceMgr) GetByOption(opts ...Option) (result Service, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ServiceMgr) GetByOptions(opts ...Option) (results []*Service, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ServiceMgr) GetFromID(id int) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ServiceMgr) GetBatchFromID(ids []int) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ServiceMgr) GetFromName(name string) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`name` = ?", name).Find(&result).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_ServiceMgr) GetBatchFromName(names []string) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromEnName 通过en_name获取内容
func (obj *_ServiceMgr) GetFromEnName(enName string) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`en_name` = ?", enName).Find(&result).Error

	return
}

// GetBatchFromEnName 批量查找
func (obj *_ServiceMgr) GetBatchFromEnName(enNames []string) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`en_name` IN (?)", enNames).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_ServiceMgr) GetFromActive(active bool) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_ServiceMgr) GetBatchFromActive(actives []bool) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ServiceMgr) GetFromCreateTime(createTime time.Time) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ServiceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ServiceMgr) FetchByPrimaryKey(id int) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByIxServiceName primary or index 获取唯一内容
func (obj *_ServiceMgr) FetchUniqueByIxServiceName(name string) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`name` = ?", name).Find(&result).Error

	return
}

// FetchUniqueByIxServiceEnName primary or index 获取唯一内容
func (obj *_ServiceMgr) FetchUniqueByIxServiceEnName(enName string) (result Service, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Service{}).Where("`en_name` = ?", enName).Find(&result).Error

	return
}
