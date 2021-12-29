package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PlatformMgr struct {
	*_BaseMgr
}

// PlatformMgr open func
func PlatformMgr(db *gorm.DB) *_PlatformMgr {
	if db == nil {
		panic(fmt.Errorf("PlatformMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PlatformMgr{_BaseMgr: &_BaseMgr{DB: db.Table("platform"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PlatformMgr) GetTableName() string {
	return "platform"
}

// Reset 重置gorm会话
func (obj *_PlatformMgr) Reset() *_PlatformMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PlatformMgr) Get() (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PlatformMgr) Gets() (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PlatformMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Platform{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PlatformMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_PlatformMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEnName en_name获取
func (obj *_PlatformMgr) WithEnName(enName string) Option {
	return optionFunc(func(o *options) { o.query["en_name"] = enName })
}

// WithActive active获取
func (obj *_PlatformMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithCreateTime create_time获取
func (obj *_PlatformMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_PlatformMgr) GetByOption(opts ...Option) (result Platform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PlatformMgr) GetByOptions(opts ...Option) (results []*Platform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_PlatformMgr) GetFromID(id int) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PlatformMgr) GetBatchFromID(ids []int) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_PlatformMgr) GetFromName(name string) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`name` = ?", name).Find(&result).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_PlatformMgr) GetBatchFromName(names []string) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromEnName 通过en_name获取内容
func (obj *_PlatformMgr) GetFromEnName(enName string) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`en_name` = ?", enName).Find(&result).Error

	return
}

// GetBatchFromEnName 批量查找
func (obj *_PlatformMgr) GetBatchFromEnName(enNames []string) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`en_name` IN (?)", enNames).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_PlatformMgr) GetFromActive(active bool) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`active` = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量查找
func (obj *_PlatformMgr) GetBatchFromActive(actives []bool) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`active` IN (?)", actives).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_PlatformMgr) GetFromCreateTime(createTime time.Time) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_PlatformMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PlatformMgr) FetchByPrimaryKey(id int) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_PlatformMgr) FetchUniqueByName(name string) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`name` = ?", name).Find(&result).Error

	return
}

// FetchUniqueByEnName primary or index 获取唯一内容
func (obj *_PlatformMgr) FetchUniqueByEnName(enName string) (result Platform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Platform{}).Where("`en_name` = ?", enName).Find(&result).Error

	return
}
