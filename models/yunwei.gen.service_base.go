package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ServiceBaseMgr struct {
	*_BaseMgr
}

// ServiceBaseMgr open func
func ServiceBaseMgr(db *gorm.DB) *_ServiceBaseMgr {
	if db == nil {
		panic(fmt.Errorf("ServiceBaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ServiceBaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("service_base"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ServiceBaseMgr) GetTableName() string {
	return "service_base"
}

// Reset 重置gorm会话
func (obj *_ServiceBaseMgr) Reset() *_ServiceBaseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ServiceBaseMgr) Get() (result ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("server").Where("id = ?", result.ServerID).Find(&result.Server).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeployUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("service").Where("id = ?", result.ServiceID).Find(&result.Service).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ServiceBaseMgr) Gets() (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ServiceBaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ServiceBaseMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServerID server_id获取
func (obj *_ServiceBaseMgr) WithServerID(serverID int) Option {
	return optionFunc(func(o *options) { o.query["server_id"] = serverID })
}

// WithGameID game_id获取
func (obj *_ServiceBaseMgr) WithGameID(gameID int) Option {
	return optionFunc(func(o *options) { o.query["game_id"] = gameID })
}

// WithPlatformID platform_id获取
func (obj *_ServiceBaseMgr) WithPlatformID(platformID int) Option {
	return optionFunc(func(o *options) { o.query["platform_id"] = platformID })
}

// WithDeployUserID deploy_user_id获取
func (obj *_ServiceBaseMgr) WithDeployUserID(deployUserID int) Option {
	return optionFunc(func(o *options) { o.query["deploy_user_id"] = deployUserID })
}

// WithServiceID service_id获取
func (obj *_ServiceBaseMgr) WithServiceID(serviceID int) Option {
	return optionFunc(func(o *options) { o.query["service_id"] = serviceID })
}

// WithConfigKwargs config_kwargs获取
func (obj *_ServiceBaseMgr) WithConfigKwargs(configKwargs string) Option {
	return optionFunc(func(o *options) { o.query["config_kwargs"] = configKwargs })
}

// WithDeployTime deploy_time获取
func (obj *_ServiceBaseMgr) WithDeployTime(deployTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["deploy_time"] = deployTime })
}

// WithUpdateTime update_time获取
func (obj *_ServiceBaseMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithNumber number获取
func (obj *_ServiceBaseMgr) WithNumber(number int64) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithOpenTime open_time获取
func (obj *_ServiceBaseMgr) WithOpenTime(openTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["open_time"] = openTime })
}

// WithMergeTime merge_time获取
func (obj *_ServiceBaseMgr) WithMergeTime(mergeTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["merge_time"] = mergeTime })
}

// WithCrossServiceID cross_service_id获取
func (obj *_ServiceBaseMgr) WithCrossServiceID(crossServiceID int) Option {
	return optionFunc(func(o *options) { o.query["cross_service_id"] = crossServiceID })
}

// WithMergedServices merged_services获取
func (obj *_ServiceBaseMgr) WithMergedServices(mergedServices string) Option {
	return optionFunc(func(o *options) { o.query["merged_services"] = mergedServices })
}

// WithVersion version获取
func (obj *_ServiceBaseMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithStatus status获取
func (obj *_ServiceBaseMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithNoBackupWarning no_backup_warning获取
func (obj *_ServiceBaseMgr) WithNoBackupWarning(noBackupWarning bool) Option {
	return optionFunc(func(o *options) { o.query["no_backup_warning"] = noBackupWarning })
}

// GetByOption 功能选项模式获取
func (obj *_ServiceBaseMgr) GetByOption(opts ...Option) (result ServiceBase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("server").Where("id = ?", result.ServerID).Find(&result.Server).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeployUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("service").Where("id = ?", result.ServiceID).Find(&result.Service).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ServiceBaseMgr) GetByOptions(opts ...Option) (results []*ServiceBase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ServiceBaseMgr) GetFromID(id int) (result ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("server").Where("id = ?", result.ServerID).Find(&result.Server).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeployUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("service").Where("id = ?", result.ServiceID).Find(&result.Service).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromID(ids []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromServerID 通过server_id获取内容
func (obj *_ServiceBaseMgr) GetFromServerID(serverID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`server_id` = ?", serverID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromServerID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromServerID(serverIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`server_id` IN (?)", serverIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGameID 通过game_id获取内容
func (obj *_ServiceBaseMgr) GetFromGameID(gameID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`game_id` = ?", gameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGameID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromGameID(gameIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`game_id` IN (?)", gameIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlatformID 通过platform_id获取内容
func (obj *_ServiceBaseMgr) GetFromPlatformID(platformID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`platform_id` = ?", platformID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlatformID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromPlatformID(platformIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`platform_id` IN (?)", platformIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeployUserID 通过deploy_user_id获取内容
func (obj *_ServiceBaseMgr) GetFromDeployUserID(deployUserID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`deploy_user_id` = ?", deployUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeployUserID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromDeployUserID(deployUserIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`deploy_user_id` IN (?)", deployUserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromServiceID 通过service_id获取内容
func (obj *_ServiceBaseMgr) GetFromServiceID(serviceID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`service_id` = ?", serviceID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromServiceID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromServiceID(serviceIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`service_id` IN (?)", serviceIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromConfigKwargs 通过config_kwargs获取内容
func (obj *_ServiceBaseMgr) GetFromConfigKwargs(configKwargs string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`config_kwargs` = ?", configKwargs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromConfigKwargs 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromConfigKwargs(configKwargss []string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`config_kwargs` IN (?)", configKwargss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeployTime 通过deploy_time获取内容
func (obj *_ServiceBaseMgr) GetFromDeployTime(deployTime time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`deploy_time` = ?", deployTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeployTime 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromDeployTime(deployTimes []time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`deploy_time` IN (?)", deployTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ServiceBaseMgr) GetFromUpdateTime(updateTime time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNumber 通过number获取内容
func (obj *_ServiceBaseMgr) GetFromNumber(number int16) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`number` = ?", number).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNumber 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromNumber(numbers []int16) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`number` IN (?)", numbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOpenTime 通过open_time获取内容
func (obj *_ServiceBaseMgr) GetFromOpenTime(openTime time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`open_time` = ?", openTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOpenTime 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromOpenTime(openTimes []time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`open_time` IN (?)", openTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMergeTime 通过merge_time获取内容
func (obj *_ServiceBaseMgr) GetFromMergeTime(mergeTime time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`merge_time` = ?", mergeTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMergeTime 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromMergeTime(mergeTimes []time.Time) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`merge_time` IN (?)", mergeTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCrossServiceID 通过cross_service_id获取内容
func (obj *_ServiceBaseMgr) GetFromCrossServiceID(crossServiceID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`cross_service_id` = ?", crossServiceID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCrossServiceID 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromCrossServiceID(crossServiceIDs []int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`cross_service_id` IN (?)", crossServiceIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMergedServices 通过merged_services获取内容
func (obj *_ServiceBaseMgr) GetFromMergedServices(mergedServices string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`merged_services` = ?", mergedServices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMergedServices 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromMergedServices(mergedServicess []string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`merged_services` IN (?)", mergedServicess).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容
func (obj *_ServiceBaseMgr) GetFromVersion(version string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromVersion(versions []string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_ServiceBaseMgr) GetFromStatus(status string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromStatus(statuss []string) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNoBackupWarning 通过no_backup_warning获取内容
func (obj *_ServiceBaseMgr) GetFromNoBackupWarning(noBackupWarning bool) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`no_backup_warning` = ?", noBackupWarning).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNoBackupWarning 批量查找
func (obj *_ServiceBaseMgr) GetBatchFromNoBackupWarning(noBackupWarnings []bool) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`no_backup_warning` IN (?)", noBackupWarnings).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ServiceBaseMgr) FetchByPrimaryKey(id int) (result ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("server").Where("id = ?", result.ServerID).Find(&result.Server).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeployUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("service").Where("id = ?", result.ServiceID).Find(&result.Service).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByServerID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByServerID(serverID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`server_id` = ?", serverID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByGameID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByGameID(gameID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`game_id` = ?", gameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByPlatformID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByPlatformID(platformID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`platform_id` = ?", platformID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByDeployUserID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByDeployUserID(deployUserID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`deploy_user_id` = ?", deployUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByServiceID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByServiceID(serviceID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`service_id` = ?", serviceID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByCrossServiceID  获取多个内容
func (obj *_ServiceBaseMgr) FetchIndexByCrossServiceID(crossServiceID int) (results []*ServiceBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ServiceBase{}).Where("`cross_service_id` = ?", crossServiceID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("server").Where("id = ?", results[i].ServerID).Find(&results[i].Server).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeployUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("service").Where("id = ?", results[i].ServiceID).Find(&results[i].Service).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
