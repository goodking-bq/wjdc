package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _GameMgr struct {
	*_BaseMgr
}

// GameMgr open func
func GameMgr(db *gorm.DB) *_GameMgr {
	if db == nil {
		panic(fmt.Errorf("GameMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GameMgr{_BaseMgr: &_BaseMgr{DB: db.Table("game"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GameMgr) GetTableName() string {
	return "game"
}

// Reset 重置gorm会话
func (obj *_GameMgr) Reset() *_GameMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GameMgr) Get() (result Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_GameMgr) Gets() (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GameMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Game{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GameMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_GameMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEnName en_name获取
func (obj *_GameMgr) WithEnName(enName string) Option {
	return optionFunc(func(o *options) { o.query["en_name"] = enName })
}

// WithGameType game_type获取
func (obj *_GameMgr) WithGameType(gameType int16) Option {
	return optionFunc(func(o *options) { o.query["game_type"] = gameType })
}

// WithActive active获取
func (obj *_GameMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithCreateTime create_time获取
func (obj *_GameMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithMaxService max_service获取
func (obj *_GameMgr) WithMaxService(maxService int16) Option {
	return optionFunc(func(o *options) { o.query["max_service"] = maxService })
}

// WithNumberStep number_step获取
func (obj *_GameMgr) WithNumberStep(numberStep int16) Option {
	return optionFunc(func(o *options) { o.query["number_step"] = numberStep })
}

// WithIsManager is_manager获取
func (obj *_GameMgr) WithIsManager(isManager bool) Option {
	return optionFunc(func(o *options) { o.query["is_manager"] = isManager })
}

// WithDNSLevel dns_level获取
func (obj *_GameMgr) WithDNSLevel(dnsLevel int16) Option {
	return optionFunc(func(o *options) { o.query["dns_level"] = dnsLevel })
}

// WithAliAccesskeyid ali_accesskeyid获取
func (obj *_GameMgr) WithAliAccesskeyid(aliAccesskeyid string) Option {
	return optionFunc(func(o *options) { o.query["ali_accesskeyid"] = aliAccesskeyid })
}

// WithAliAccesskeysecret ali_accesskeysecret获取
func (obj *_GameMgr) WithAliAccesskeysecret(aliAccesskeysecret string) Option {
	return optionFunc(func(o *options) { o.query["ali_accesskeysecret"] = aliAccesskeysecret })
}

// WithDNSID dns_id获取
func (obj *_GameMgr) WithDNSID(dnsID int) Option {
	return optionFunc(func(o *options) { o.query["dns_id"] = dnsID })
}

// WithSaltAPIID salt_api_id获取
func (obj *_GameMgr) WithSaltAPIID(saltAPIID int) Option {
	return optionFunc(func(o *options) { o.query["salt_api_id"] = saltAPIID })
}

// WithPillar pillar获取
func (obj *_GameMgr) WithPillar(pillar string) Option {
	return optionFunc(func(o *options) { o.query["pillar"] = pillar })
}

// WithIsHunbu is_hunbu获取
func (obj *_GameMgr) WithIsHunbu(isHunbu bool) Option {
	return optionFunc(func(o *options) { o.query["is_hunbu"] = isHunbu })
}

// WithManagerUser manager_user获取
func (obj *_GameMgr) WithManagerUser(managerUser int16) Option {
	return optionFunc(func(o *options) { o.query["manager_user"] = managerUser })
}

// WithMergeType merge_type获取
func (obj *_GameMgr) WithMergeType(mergeType int16) Option {
	return optionFunc(func(o *options) { o.query["merge_type"] = mergeType })
}

// WithBusinessUser business_user获取
func (obj *_GameMgr) WithBusinessUser(businessUser string) Option {
	return optionFunc(func(o *options) { o.query["business_user"] = businessUser })
}

// WithZabbixGroupID zabbix_group_id获取
func (obj *_GameMgr) WithZabbixGroupID(zabbixGroupID int) Option {
	return optionFunc(func(o *options) { o.query["zabbix_group_id"] = zabbixGroupID })
}

// WithV1 v1获取
func (obj *_GameMgr) WithV1(v1 bool) Option {
	return optionFunc(func(o *options) { o.query["v1"] = v1 })
}

// WithStepInfo step_info获取
func (obj *_GameMgr) WithStepInfo(stepInfo string) Option {
	return optionFunc(func(o *options) { o.query["step_info"] = stepInfo })
}

// WithQq qq获取
func (obj *_GameMgr) WithQq(qq string) Option {
	return optionFunc(func(o *options) { o.query["qq"] = qq })
}

// GetByOption 功能选项模式获取
func (obj *_GameMgr) GetByOption(opts ...Option) (result Game, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GameMgr) GetByOptions(opts ...Option) (results []*Game, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
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
func (obj *_GameMgr) GetFromID(id int) (result Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_GameMgr) GetBatchFromID(ids []int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_GameMgr) GetFromName(name string) (result Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`name` = ?", name).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromName 批量查找
func (obj *_GameMgr) GetBatchFromName(names []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEnName 通过en_name获取内容
func (obj *_GameMgr) GetFromEnName(enName string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`en_name` = ?", enName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEnName 批量查找
func (obj *_GameMgr) GetBatchFromEnName(enNames []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`en_name` IN (?)", enNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGameType 通过game_type获取内容
func (obj *_GameMgr) GetFromGameType(gameType int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`game_type` = ?", gameType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGameType 批量查找
func (obj *_GameMgr) GetBatchFromGameType(gameTypes []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`game_type` IN (?)", gameTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActive 通过active获取内容
func (obj *_GameMgr) GetFromActive(active bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`active` = ?", active).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActive 批量查找
func (obj *_GameMgr) GetBatchFromActive(actives []bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`active` IN (?)", actives).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_GameMgr) GetFromCreateTime(createTime time.Time) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_GameMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMaxService 通过max_service获取内容
func (obj *_GameMgr) GetFromMaxService(maxService int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`max_service` = ?", maxService).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMaxService 批量查找
func (obj *_GameMgr) GetBatchFromMaxService(maxServices []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`max_service` IN (?)", maxServices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNumberStep 通过number_step获取内容
func (obj *_GameMgr) GetFromNumberStep(numberStep int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`number_step` = ?", numberStep).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNumberStep 批量查找
func (obj *_GameMgr) GetBatchFromNumberStep(numberSteps []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`number_step` IN (?)", numberSteps).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsManager 通过is_manager获取内容
func (obj *_GameMgr) GetFromIsManager(isManager bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`is_manager` = ?", isManager).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsManager 批量查找
func (obj *_GameMgr) GetBatchFromIsManager(isManagers []bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`is_manager` IN (?)", isManagers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDNSLevel 通过dns_level获取内容
func (obj *_GameMgr) GetFromDNSLevel(dnsLevel int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`dns_level` = ?", dnsLevel).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDNSLevel 批量查找
func (obj *_GameMgr) GetBatchFromDNSLevel(dnsLevels []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`dns_level` IN (?)", dnsLevels).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAliAccesskeyid 通过ali_accesskeyid获取内容
func (obj *_GameMgr) GetFromAliAccesskeyid(aliAccesskeyid string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`ali_accesskeyid` = ?", aliAccesskeyid).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAliAccesskeyid 批量查找
func (obj *_GameMgr) GetBatchFromAliAccesskeyid(aliAccesskeyids []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`ali_accesskeyid` IN (?)", aliAccesskeyids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAliAccesskeysecret 通过ali_accesskeysecret获取内容
func (obj *_GameMgr) GetFromAliAccesskeysecret(aliAccesskeysecret string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`ali_accesskeysecret` = ?", aliAccesskeysecret).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAliAccesskeysecret 批量查找
func (obj *_GameMgr) GetBatchFromAliAccesskeysecret(aliAccesskeysecrets []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`ali_accesskeysecret` IN (?)", aliAccesskeysecrets).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDNSID 通过dns_id获取内容
func (obj *_GameMgr) GetFromDNSID(dnsID int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`dns_id` = ?", dnsID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDNSID 批量查找
func (obj *_GameMgr) GetBatchFromDNSID(dnsIDs []int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`dns_id` IN (?)", dnsIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSaltAPIID 通过salt_api_id获取内容
func (obj *_GameMgr) GetFromSaltAPIID(saltAPIID int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`salt_api_id` = ?", saltAPIID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSaltAPIID 批量查找
func (obj *_GameMgr) GetBatchFromSaltAPIID(saltAPIIDs []int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`salt_api_id` IN (?)", saltAPIIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPillar 通过pillar获取内容
func (obj *_GameMgr) GetFromPillar(pillar string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`pillar` = ?", pillar).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPillar 批量查找
func (obj *_GameMgr) GetBatchFromPillar(pillars []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`pillar` IN (?)", pillars).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsHunbu 通过is_hunbu获取内容
func (obj *_GameMgr) GetFromIsHunbu(isHunbu bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`is_hunbu` = ?", isHunbu).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsHunbu 批量查找
func (obj *_GameMgr) GetBatchFromIsHunbu(isHunbus []bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`is_hunbu` IN (?)", isHunbus).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromManagerUser 通过manager_user获取内容
func (obj *_GameMgr) GetFromManagerUser(managerUser int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`manager_user` = ?", managerUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromManagerUser 批量查找
func (obj *_GameMgr) GetBatchFromManagerUser(managerUsers []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`manager_user` IN (?)", managerUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMergeType 通过merge_type获取内容
func (obj *_GameMgr) GetFromMergeType(mergeType int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`merge_type` = ?", mergeType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMergeType 批量查找
func (obj *_GameMgr) GetBatchFromMergeType(mergeTypes []int16) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`merge_type` IN (?)", mergeTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBusinessUser 通过business_user获取内容
func (obj *_GameMgr) GetFromBusinessUser(businessUser string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`business_user` = ?", businessUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBusinessUser 批量查找
func (obj *_GameMgr) GetBatchFromBusinessUser(businessUsers []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`business_user` IN (?)", businessUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromZabbixGroupID 通过zabbix_group_id获取内容
func (obj *_GameMgr) GetFromZabbixGroupID(zabbixGroupID int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`zabbix_group_id` = ?", zabbixGroupID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromZabbixGroupID 批量查找
func (obj *_GameMgr) GetBatchFromZabbixGroupID(zabbixGroupIDs []int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`zabbix_group_id` IN (?)", zabbixGroupIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromV1 通过v1获取内容
func (obj *_GameMgr) GetFromV1(v1 bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`v1` = ?", v1).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromV1 批量查找
func (obj *_GameMgr) GetBatchFromV1(v1s []bool) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`v1` IN (?)", v1s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStepInfo 通过step_info获取内容
func (obj *_GameMgr) GetFromStepInfo(stepInfo string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`step_info` = ?", stepInfo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStepInfo 批量查找
func (obj *_GameMgr) GetBatchFromStepInfo(stepInfos []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`step_info` IN (?)", stepInfos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQq 通过qq获取内容
func (obj *_GameMgr) GetFromQq(qq string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`qq` = ?", qq).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQq 批量查找
func (obj *_GameMgr) GetBatchFromQq(qqs []string) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`qq` IN (?)", qqs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
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
func (obj *_GameMgr) FetchByPrimaryKey(id int) (result Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_GameMgr) FetchUniqueByName(name string) (result Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`name` = ?", name).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("dns").Where("id = ?", result.DNSID).Find(&result.DNS).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("salt_api").Where("id = ?", result.SaltAPIID).Find(&result.SaltAPI).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByDNSID  获取多个内容
func (obj *_GameMgr) FetchIndexByDNSID(dnsID int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`dns_id` = ?", dnsID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexBySaltAPIID  获取多个内容
func (obj *_GameMgr) FetchIndexBySaltAPIID(saltAPIID int) (results []*Game, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Game{}).Where("`salt_api_id` = ?", saltAPIID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("dns").Where("id = ?", results[i].DNSID).Find(&results[i].DNS).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("salt_api").Where("id = ?", results[i].SaltAPIID).Find(&results[i].SaltAPI).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
