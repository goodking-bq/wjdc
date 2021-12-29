package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ServerMgr struct {
	*_BaseMgr
}

// ServerMgr open func
func ServerMgr(db *gorm.DB) *_ServerMgr {
	if db == nil {
		panic(fmt.Errorf("ServerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ServerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("server"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ServerMgr) GetTableName() string {
	return "server"
}

// Reset 重置gorm会话
func (obj *_ServerMgr) Reset() *_ServerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ServerMgr) Get() (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ServerMgr) Gets() (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ServerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Server{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ServerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIP ip获取
func (obj *_ServerMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithIPInterfaceName ip_interface_name获取
func (obj *_ServerMgr) WithIPInterfaceName(ipInterfaceName string) Option {
	return optionFunc(func(o *options) { o.query["ip_interface_name"] = ipInterfaceName })
}

// WithUIp uip获取
func (obj *_ServerMgr) WithUIp(uip string) Option {
	return optionFunc(func(o *options) { o.query["uip"] = uip })
}

// WithTip tip获取
func (obj *_ServerMgr) WithTip(tip string) Option {
	return optionFunc(func(o *options) { o.query["tip"] = tip })
}

// WithLip lip获取
func (obj *_ServerMgr) WithLip(lip string) Option {
	return optionFunc(func(o *options) { o.query["lip"] = lip })
}

// WithLipInterfaceName lip_interface_name获取
func (obj *_ServerMgr) WithLipInterfaceName(lipInterfaceName string) Option {
	return optionFunc(func(o *options) { o.query["lip_interface_name"] = lipInterfaceName })
}

// WithPort port获取
func (obj *_ServerMgr) WithPort(port string) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithCPU cpu获取
func (obj *_ServerMgr) WithCPU(cpu string) Option {
	return optionFunc(func(o *options) { o.query["cpu"] = cpu })
}

// WithCPUCore cpu_core获取
func (obj *_ServerMgr) WithCPUCore(cpuCore string) Option {
	return optionFunc(func(o *options) { o.query["cpu_core"] = cpuCore })
}

// WithDisk disk获取
func (obj *_ServerMgr) WithDisk(disk string) Option {
	return optionFunc(func(o *options) { o.query["disk"] = disk })
}

// WithMemory memory获取
func (obj *_ServerMgr) WithMemory(memory string) Option {
	return optionFunc(func(o *options) { o.query["memory"] = memory })
}

// WithOs os获取
func (obj *_ServerMgr) WithOs(os string) Option {
	return optionFunc(func(o *options) { o.query["os"] = os })
}

// WithAssetTag asset_tag获取
func (obj *_ServerMgr) WithAssetTag(assetTag string) Option {
	return optionFunc(func(o *options) { o.query["asset_tag"] = assetTag })
}

// WithSerialNumber serial_number获取
func (obj *_ServerMgr) WithSerialNumber(serialNumber string) Option {
	return optionFunc(func(o *options) { o.query["serial_number"] = serialNumber })
}

// WithUnitNumber unit_number获取
func (obj *_ServerMgr) WithUnitNumber(unitNumber int) Option {
	return optionFunc(func(o *options) { o.query["unit_number"] = unitNumber })
}

// WithRackNumber rack_number获取
func (obj *_ServerMgr) WithRackNumber(rackNumber string) Option {
	return optionFunc(func(o *options) { o.query["rack_number"] = rackNumber })
}

// WithAddDate add_date获取
func (obj *_ServerMgr) WithAddDate(addDate string) Option {
	return optionFunc(func(o *options) { o.query["add_date"] = addDate })
}

// WithBrand brand获取
func (obj *_ServerMgr) WithBrand(brand string) Option {
	return optionFunc(func(o *options) { o.query["brand"] = brand })
}

// WithModeNumber mode_number获取
func (obj *_ServerMgr) WithModeNumber(modeNumber string) Option {
	return optionFunc(func(o *options) { o.query["mode_number"] = modeNumber })
}

// WithInUse in_use获取
func (obj *_ServerMgr) WithInUse(inUse bool) Option {
	return optionFunc(func(o *options) { o.query["in_use"] = inUse })
}

// WithSaltKey salt_key获取
func (obj *_ServerMgr) WithSaltKey(saltKey string) Option {
	return optionFunc(func(o *options) { o.query["salt_key"] = saltKey })
}

// WithCreateTime create_time获取
func (obj *_ServerMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUserID create_user_id获取
func (obj *_ServerMgr) WithCreateUserID(createUserID int) Option {
	return optionFunc(func(o *options) { o.query["create_user_id"] = createUserID })
}

// WithUpdateTime update_time获取
func (obj *_ServerMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUserID update_user_id获取
func (obj *_ServerMgr) WithUpdateUserID(updateUserID int) Option {
	return optionFunc(func(o *options) { o.query["update_user_id"] = updateUserID })
}

// WithInitUser init_user获取
func (obj *_ServerMgr) WithInitUser(initUser string) Option {
	return optionFunc(func(o *options) { o.query["init_user"] = initUser })
}

// WithInitPassword init_password获取
func (obj *_ServerMgr) WithInitPassword(initPassword string) Option {
	return optionFunc(func(o *options) { o.query["init_password"] = initPassword })
}

// WithSaltStatus salt_status获取
func (obj *_ServerMgr) WithSaltStatus(saltStatus int16) Option {
	return optionFunc(func(o *options) { o.query["salt_status"] = saltStatus })
}

// WithBaseStatus base_status获取
func (obj *_ServerMgr) WithBaseStatus(baseStatus bool) Option {
	return optionFunc(func(o *options) { o.query["base_status"] = baseStatus })
}

// WithPriceOfMonth price_of_month获取
func (obj *_ServerMgr) WithPriceOfMonth(priceOfMonth float64) Option {
	return optionFunc(func(o *options) { o.query["price_of_month"] = priceOfMonth })
}

// WithIsDeposit is_deposit获取
func (obj *_ServerMgr) WithIsDeposit(isDeposit bool) Option {
	return optionFunc(func(o *options) { o.query["is_deposit"] = isDeposit })
}

// WithDepositor depositor获取
func (obj *_ServerMgr) WithDepositor(depositor string) Option {
	return optionFunc(func(o *options) { o.query["depositor"] = depositor })
}

// WithGameID game_id获取
func (obj *_ServerMgr) WithGameID(gameID int) Option {
	return optionFunc(func(o *options) { o.query["game_id"] = gameID })
}

// WithRoomID room_id获取
func (obj *_ServerMgr) WithRoomID(roomID int) Option {
	return optionFunc(func(o *options) { o.query["room_id"] = roomID })
}

// WithHistoryGameID history_game_id获取
func (obj *_ServerMgr) WithHistoryGameID(historyGameID int) Option {
	return optionFunc(func(o *options) { o.query["history_game_id"] = historyGameID })
}

// WithIsFree is_free获取
func (obj *_ServerMgr) WithIsFree(isFree bool) Option {
	return optionFunc(func(o *options) { o.query["is_free"] = isFree })
}

// WithDeleteTime delete_time获取
func (obj *_ServerMgr) WithDeleteTime(deleteTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// WithDeleteUserID delete_user_id获取
func (obj *_ServerMgr) WithDeleteUserID(deleteUserID int) Option {
	return optionFunc(func(o *options) { o.query["delete_user_id"] = deleteUserID })
}

// WithIsDelete is_delete获取
func (obj *_ServerMgr) WithIsDelete(isDelete bool) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithPlatformID platform_id获取
func (obj *_ServerMgr) WithPlatformID(platformID int) Option {
	return optionFunc(func(o *options) { o.query["platform_id"] = platformID })
}

// WithIgnorePortScan ignore_port_scan获取
func (obj *_ServerMgr) WithIgnorePortScan(ignorePortScan bool) Option {
	return optionFunc(func(o *options) { o.query["ignore_port_scan"] = ignorePortScan })
}

// GetByOption 功能选项模式获取
func (obj *_ServerMgr) GetByOption(opts ...Option) (result Server, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ServerMgr) GetByOptions(opts ...Option) (results []*Server, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
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
func (obj *_ServerMgr) GetFromID(id int) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_ServerMgr) GetBatchFromID(ids []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIP 通过ip获取内容
func (obj *_ServerMgr) GetFromIP(ip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ip` = ?", ip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromIP 批量查找
func (obj *_ServerMgr) GetBatchFromIP(ips []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ip` IN (?)", ips).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIPInterfaceName 通过ip_interface_name获取内容
func (obj *_ServerMgr) GetFromIPInterfaceName(ipInterfaceName string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ip_interface_name` = ?", ipInterfaceName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIPInterfaceName 批量查找
func (obj *_ServerMgr) GetBatchFromIPInterfaceName(ipInterfaceNames []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ip_interface_name` IN (?)", ipInterfaceNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUIp 通过uip获取内容
func (obj *_ServerMgr) GetFromUIp(uip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`uip` = ?", uip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUIp 批量查找
func (obj *_ServerMgr) GetBatchFromUIp(uips []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`uip` IN (?)", uips).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTip 通过tip获取内容
func (obj *_ServerMgr) GetFromTip(tip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`tip` = ?", tip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromTip 批量查找
func (obj *_ServerMgr) GetBatchFromTip(tips []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`tip` IN (?)", tips).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLip 通过lip获取内容
func (obj *_ServerMgr) GetFromLip(lip string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`lip` = ?", lip).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLip 批量查找
func (obj *_ServerMgr) GetBatchFromLip(lips []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`lip` IN (?)", lips).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLipInterfaceName 通过lip_interface_name获取内容
func (obj *_ServerMgr) GetFromLipInterfaceName(lipInterfaceName string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`lip_interface_name` = ?", lipInterfaceName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLipInterfaceName 批量查找
func (obj *_ServerMgr) GetBatchFromLipInterfaceName(lipInterfaceNames []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`lip_interface_name` IN (?)", lipInterfaceNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPort 通过port获取内容
func (obj *_ServerMgr) GetFromPort(port string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`port` = ?", port).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPort 批量查找
func (obj *_ServerMgr) GetBatchFromPort(ports []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`port` IN (?)", ports).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCPU 通过cpu获取内容
func (obj *_ServerMgr) GetFromCPU(cpu string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`cpu` = ?", cpu).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCPU 批量查找
func (obj *_ServerMgr) GetBatchFromCPU(cpus []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`cpu` IN (?)", cpus).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCPUCore 通过cpu_core获取内容
func (obj *_ServerMgr) GetFromCPUCore(cpuCore string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`cpu_core` = ?", cpuCore).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCPUCore 批量查找
func (obj *_ServerMgr) GetBatchFromCPUCore(cpuCores []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`cpu_core` IN (?)", cpuCores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDisk 通过disk获取内容
func (obj *_ServerMgr) GetFromDisk(disk string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`disk` = ?", disk).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDisk 批量查找
func (obj *_ServerMgr) GetBatchFromDisk(disks []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`disk` IN (?)", disks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMemory 通过memory获取内容
func (obj *_ServerMgr) GetFromMemory(memory string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`memory` = ?", memory).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMemory 批量查找
func (obj *_ServerMgr) GetBatchFromMemory(memorys []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`memory` IN (?)", memorys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOs 通过os获取内容
func (obj *_ServerMgr) GetFromOs(os string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`os` = ?", os).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOs 批量查找
func (obj *_ServerMgr) GetBatchFromOs(oss []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`os` IN (?)", oss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAssetTag 通过asset_tag获取内容
func (obj *_ServerMgr) GetFromAssetTag(assetTag string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`asset_tag` = ?", assetTag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAssetTag 批量查找
func (obj *_ServerMgr) GetBatchFromAssetTag(assetTags []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`asset_tag` IN (?)", assetTags).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSerialNumber 通过serial_number获取内容
func (obj *_ServerMgr) GetFromSerialNumber(serialNumber string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`serial_number` = ?", serialNumber).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSerialNumber 批量查找
func (obj *_ServerMgr) GetBatchFromSerialNumber(serialNumbers []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`serial_number` IN (?)", serialNumbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUnitNumber 通过unit_number获取内容
func (obj *_ServerMgr) GetFromUnitNumber(unitNumber int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`unit_number` = ?", unitNumber).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUnitNumber 批量查找
func (obj *_ServerMgr) GetBatchFromUnitNumber(unitNumbers []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`unit_number` IN (?)", unitNumbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRackNumber 通过rack_number获取内容
func (obj *_ServerMgr) GetFromRackNumber(rackNumber string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`rack_number` = ?", rackNumber).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRackNumber 批量查找
func (obj *_ServerMgr) GetBatchFromRackNumber(rackNumbers []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`rack_number` IN (?)", rackNumbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAddDate 通过add_date获取内容
func (obj *_ServerMgr) GetFromAddDate(addDate string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`add_date` = ?", addDate).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAddDate 批量查找
func (obj *_ServerMgr) GetBatchFromAddDate(addDates []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`add_date` IN (?)", addDates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBrand 通过brand获取内容
func (obj *_ServerMgr) GetFromBrand(brand string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`brand` = ?", brand).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBrand 批量查找
func (obj *_ServerMgr) GetBatchFromBrand(brands []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`brand` IN (?)", brands).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromModeNumber 通过mode_number获取内容
func (obj *_ServerMgr) GetFromModeNumber(modeNumber string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`mode_number` = ?", modeNumber).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromModeNumber 批量查找
func (obj *_ServerMgr) GetBatchFromModeNumber(modeNumbers []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`mode_number` IN (?)", modeNumbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromInUse 通过in_use获取内容
func (obj *_ServerMgr) GetFromInUse(inUse bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`in_use` = ?", inUse).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromInUse 批量查找
func (obj *_ServerMgr) GetBatchFromInUse(inUses []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`in_use` IN (?)", inUses).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSaltKey 通过salt_key获取内容
func (obj *_ServerMgr) GetFromSaltKey(saltKey string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`salt_key` = ?", saltKey).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromSaltKey 批量查找
func (obj *_ServerMgr) GetBatchFromSaltKey(saltKeys []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`salt_key` IN (?)", saltKeys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_ServerMgr) GetFromCreateTime(createTime time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_ServerMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUserID 通过create_user_id获取内容
func (obj *_ServerMgr) GetFromCreateUserID(createUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`create_user_id` = ?", createUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUserID 批量查找
func (obj *_ServerMgr) GetBatchFromCreateUserID(createUserIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`create_user_id` IN (?)", createUserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_ServerMgr) GetFromUpdateTime(updateTime time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_ServerMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUserID 通过update_user_id获取内容
func (obj *_ServerMgr) GetFromUpdateUserID(updateUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`update_user_id` = ?", updateUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUserID 批量查找
func (obj *_ServerMgr) GetBatchFromUpdateUserID(updateUserIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`update_user_id` IN (?)", updateUserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromInitUser 通过init_user获取内容
func (obj *_ServerMgr) GetFromInitUser(initUser string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`init_user` = ?", initUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromInitUser 批量查找
func (obj *_ServerMgr) GetBatchFromInitUser(initUsers []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`init_user` IN (?)", initUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromInitPassword 通过init_password获取内容
func (obj *_ServerMgr) GetFromInitPassword(initPassword string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`init_password` = ?", initPassword).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromInitPassword 批量查找
func (obj *_ServerMgr) GetBatchFromInitPassword(initPasswords []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`init_password` IN (?)", initPasswords).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSaltStatus 通过salt_status获取内容
func (obj *_ServerMgr) GetFromSaltStatus(saltStatus int16) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`salt_status` = ?", saltStatus).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSaltStatus 批量查找
func (obj *_ServerMgr) GetBatchFromSaltStatus(saltStatuss []int16) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`salt_status` IN (?)", saltStatuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBaseStatus 通过base_status获取内容
func (obj *_ServerMgr) GetFromBaseStatus(baseStatus bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`base_status` = ?", baseStatus).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBaseStatus 批量查找
func (obj *_ServerMgr) GetBatchFromBaseStatus(baseStatuss []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`base_status` IN (?)", baseStatuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPriceOfMonth 通过price_of_month获取内容
func (obj *_ServerMgr) GetFromPriceOfMonth(priceOfMonth float64) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`price_of_month` = ?", priceOfMonth).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPriceOfMonth 批量查找
func (obj *_ServerMgr) GetBatchFromPriceOfMonth(priceOfMonths []float64) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`price_of_month` IN (?)", priceOfMonths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsDeposit 通过is_deposit获取内容
func (obj *_ServerMgr) GetFromIsDeposit(isDeposit bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_deposit` = ?", isDeposit).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsDeposit 批量查找
func (obj *_ServerMgr) GetBatchFromIsDeposit(isDeposits []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_deposit` IN (?)", isDeposits).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDepositor 通过depositor获取内容
func (obj *_ServerMgr) GetFromDepositor(depositor string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`depositor` = ?", depositor).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDepositor 批量查找
func (obj *_ServerMgr) GetBatchFromDepositor(depositors []string) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`depositor` IN (?)", depositors).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGameID 通过game_id获取内容
func (obj *_ServerMgr) GetFromGameID(gameID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`game_id` = ?", gameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGameID 批量查找
func (obj *_ServerMgr) GetBatchFromGameID(gameIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`game_id` IN (?)", gameIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoomID 通过room_id获取内容
func (obj *_ServerMgr) GetFromRoomID(roomID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`room_id` = ?", roomID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoomID 批量查找
func (obj *_ServerMgr) GetBatchFromRoomID(roomIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`room_id` IN (?)", roomIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromHistoryGameID 通过history_game_id获取内容
func (obj *_ServerMgr) GetFromHistoryGameID(historyGameID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`history_game_id` = ?", historyGameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromHistoryGameID 批量查找
func (obj *_ServerMgr) GetBatchFromHistoryGameID(historyGameIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`history_game_id` IN (?)", historyGameIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsFree 通过is_free获取内容
func (obj *_ServerMgr) GetFromIsFree(isFree bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_free` = ?", isFree).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsFree 批量查找
func (obj *_ServerMgr) GetBatchFromIsFree(isFrees []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_free` IN (?)", isFrees).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleteTime 通过delete_time获取内容
func (obj *_ServerMgr) GetFromDeleteTime(deleteTime time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`delete_time` = ?", deleteTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleteTime 批量查找
func (obj *_ServerMgr) GetBatchFromDeleteTime(deleteTimes []time.Time) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`delete_time` IN (?)", deleteTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleteUserID 通过delete_user_id获取内容
func (obj *_ServerMgr) GetFromDeleteUserID(deleteUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`delete_user_id` = ?", deleteUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleteUserID 批量查找
func (obj *_ServerMgr) GetBatchFromDeleteUserID(deleteUserIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`delete_user_id` IN (?)", deleteUserIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsDelete 通过is_delete获取内容
func (obj *_ServerMgr) GetFromIsDelete(isDelete bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_delete` = ?", isDelete).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsDelete 批量查找
func (obj *_ServerMgr) GetBatchFromIsDelete(isDeletes []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlatformID 通过platform_id获取内容
func (obj *_ServerMgr) GetFromPlatformID(platformID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`platform_id` = ?", platformID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlatformID 批量查找
func (obj *_ServerMgr) GetBatchFromPlatformID(platformIDs []int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`platform_id` IN (?)", platformIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIgnorePortScan 通过ignore_port_scan获取内容
func (obj *_ServerMgr) GetFromIgnorePortScan(ignorePortScan bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ignore_port_scan` = ?", ignorePortScan).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIgnorePortScan 批量查找
func (obj *_ServerMgr) GetBatchFromIgnorePortScan(ignorePortScans []bool) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ignore_port_scan` IN (?)", ignorePortScans).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
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
func (obj *_ServerMgr) FetchByPrimaryKey(id int) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByIxServerIP primary or index 获取唯一内容
func (obj *_ServerMgr) FetchUniqueByIxServerIP(ip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`ip` = ?", ip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByUIp primary or index 获取唯一内容
func (obj *_ServerMgr) FetchUniqueByUIp(uip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`uip` = ?", uip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByTip primary or index 获取唯一内容
func (obj *_ServerMgr) FetchUniqueByTip(tip string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`tip` = ?", tip).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByIxServerSaltKey primary or index 获取唯一内容
func (obj *_ServerMgr) FetchUniqueByIxServerSaltKey(saltKey string) (result Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`salt_key` = ?", saltKey).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.UpdateUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.GameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("room").Where("id = ?", result.RoomID).Find(&result.Room).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("game").Where("id = ?", result.HistoryGameID).Find(&result.Game).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user").Where("id = ?", result.DeleteUserID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("platform").Where("id = ?", result.PlatformID).Find(&result.Platform).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByCreateUserID  获取多个内容
func (obj *_ServerMgr) FetchIndexByCreateUserID(createUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`create_user_id` = ?", createUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByUpdateUserID  获取多个内容
func (obj *_ServerMgr) FetchIndexByUpdateUserID(updateUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`update_user_id` = ?", updateUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByGameID  获取多个内容
func (obj *_ServerMgr) FetchIndexByGameID(gameID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`game_id` = ?", gameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByRoomID  获取多个内容
func (obj *_ServerMgr) FetchIndexByRoomID(roomID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`room_id` = ?", roomID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByHistoryGameID  获取多个内容
func (obj *_ServerMgr) FetchIndexByHistoryGameID(historyGameID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`history_game_id` = ?", historyGameID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByDeleteUserID  获取多个内容
func (obj *_ServerMgr) FetchIndexByDeleteUserID(deleteUserID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`delete_user_id` = ?", deleteUserID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByPlatformID  获取多个内容
func (obj *_ServerMgr) FetchIndexByPlatformID(platformID int) (results []*Server, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Server{}).Where("`platform_id` = ?", platformID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].UpdateUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].GameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("room").Where("id = ?", results[i].RoomID).Find(&results[i].Room).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("game").Where("id = ?", results[i].HistoryGameID).Find(&results[i].Game).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].DeleteUserID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("platform").Where("id = ?", results[i].PlatformID).Find(&results[i].Platform).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
