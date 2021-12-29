package models

import (
	"encoding/json"
	"strconv"
	"time"
)

/******sql******
CREATE TABLE `alembic_version` (
  `version_num` varchar(32) NOT NULL,
  PRIMARY KEY (`version_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// AlembicVersion [...]
type AlembicVersion struct {
	VersionNum string `gorm:"primaryKey;column:version_num;type:varchar(32);not null"`
}

/******sql******
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `text` mediumtext NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `author` varchar(255) NOT NULL,
  `markdown` mediumtext,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `article_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `article_category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8
******sql******/
// Article [...]
type Article struct {
	ID              int             `gorm:"primaryKey;column:id;type:int(11);not null"`
	Title           string          `gorm:"column:title;type:varchar(255);not null"`
	Description     string          `gorm:"column:description;type:varchar(255)"`
	Tags            string          `gorm:"column:tags;type:varchar(255)"`
	Text            string          `gorm:"column:text;type:mediumtext;not null"`
	CreateTime      time.Time       `gorm:"column:create_time;type:datetime"`
	UpdateTime      time.Time       `gorm:"column:update_time;type:datetime"`
	CategoryID      int             `gorm:"index:category_id;column:category_id;type:int(11)"`
	ArticleCategory ArticleCategory `gorm:"joinForeignKey:category_id;foreignKey:id"`
	Author          string          `gorm:"column:author;type:varchar(255);not null"`
	Markdown        string          `gorm:"column:markdown;type:mediumtext"`
}

/******sql******
CREATE TABLE `article_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8
******sql******/
// ArticleCategory [...]
type ArticleCategory struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"column:name;type:varchar(255);not null"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `backup_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `service_id` int(11) DEFAULT NULL,
  `is_warning` tinyint(1) DEFAULT NULL,
  `backup_status` tinyint(1) DEFAULT NULL,
  `update_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `backup_log_ibfk_1` FOREIGN KEY (`service_id`) REFERENCES `service_base` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12334 DEFAULT CHARSET=utf8mb4
******sql******/
// BackupLog [...]
type BackupLog struct {
	ID           int         `gorm:"primaryKey;column:id;type:int(11);not null"`
	ServiceID    int         `gorm:"index:service_id;column:service_id;type:int(11)"`
	ServiceBase  ServiceBase `gorm:"joinForeignKey:service_id;foreignKey:id"`
	IsWarning    bool        `gorm:"column:is_warning;type:tinyint(1)"`
	BackupStatus bool        `gorm:"column:backup_status;type:tinyint(1)"`
	UpdateDate   time.Time   `gorm:"column:update_date;type:datetime"`
}

/******sql******
CREATE TABLE `cdn_domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `provider` varchar(64) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `belong_user` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`),
  CONSTRAINT `cdn_domain_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=444 DEFAULT CHARSET=utf8mb4
******sql******/
// CdnDomain [...]
type CdnDomain struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Provider    string    `gorm:"column:provider;type:varchar(64)"`
	Active      bool      `gorm:"column:active;type:tinyint(1)"`
	GameID      int       `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game        Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	BelongUser  string    `gorm:"column:belong_user;type:varchar(255)"`
}

/******sql******
CREATE TABLE `cdn_domain_spend_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) DEFAULT NULL,
  `record_month` varchar(32) DEFAULT NULL,
  `spend` decimal(10,2) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10527 DEFAULT CHARSET=utf8mb4
******sql******/
// CdnDomainSpendLog [...]
type CdnDomainSpendLog struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	DomainID    int       `gorm:"column:domain_id;type:int(11)"`
	RecordMonth string    `gorm:"column:record_month;type:varchar(32)"`
	Spend       float64   `gorm:"column:spend;type:decimal(10,2)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `cloud_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `room_id` int(11) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  `secret` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `room_id` (`room_id`),
  CONSTRAINT `cloud_account_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4
******sql******/
// CloudAccount [...]
type CloudAccount struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Account    string    `gorm:"column:account;type:varchar(64)"`
	Name       string    `gorm:"column:name;type:varchar(64)"`
	RoomID     int       `gorm:"index:room_id;column:room_id;type:int(11)"`
	Room       Room      `gorm:"joinForeignKey:room_id;foreignKey:id"`
	Key        string    `gorm:"column:key;type:varchar(255)"`
	Secret     string    `gorm:"column:secret;type:varchar(255)"`
	Version    string    `gorm:"column:version;type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
	Active     bool      `gorm:"column:active;type:tinyint(1)"`
}

/******sql******
CREATE TABLE `cloud_db_spend_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `db_type` varchar(255) DEFAULT NULL,
  `record_month` varchar(32) DEFAULT NULL,
  `spend` decimal(10,2) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `provider` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=711 DEFAULT CHARSET=utf8mb4
******sql******/
// CloudDbSpendLog [...]
type CloudDbSpendLog struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID      int       `gorm:"column:game_id;type:int(11)"`
	DbType      string    `gorm:"column:db_type;type:varchar(255)"`
	RecordMonth string    `gorm:"column:record_month;type:varchar(32)"`
	Spend       float64   `gorm:"column:spend;type:decimal(10,2)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	Provider    int       `gorm:"column:provider;type:int(11)"`
}

/******sql******
CREATE TABLE `cloud_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_type` varchar(255) DEFAULT NULL,
  `host_dead_time` varchar(255) DEFAULT NULL,
  `belong_user` varchar(255) DEFAULT NULL,
  `price_of_month` decimal(10,2) DEFAULT NULL,
  `project_id` varchar(255) DEFAULT NULL,
  `cpu_usage` decimal(10,2) DEFAULT NULL,
  `wan_usage` decimal(10,2) DEFAULT NULL,
  `cloud_account_id` int(11) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `pay_type` varchar(64) DEFAULT NULL,
  `resource_address` varchar(255) DEFAULT NULL,
  `resource_id` varchar(255) DEFAULT NULL,
  `resource_name` varchar(255) DEFAULT NULL,
  `resource_region_id` varchar(255) DEFAULT NULL,
  `status` varchar(64) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cloud_account_id` (`cloud_account_id`),
  CONSTRAINT `cloud_resource_ibfk_1` FOREIGN KEY (`cloud_account_id`) REFERENCES `cloud_account` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8mb4
******sql******/
// CloudResource [...]
type CloudResource struct {
	ID               int          `gorm:"primaryKey;column:id;type:int(11);not null"`
	ResourceType     string       `gorm:"column:resource_type;type:varchar(255)"`
	HostDeadTime     string       `gorm:"column:host_dead_time;type:varchar(255)"`
	BelongUser       string       `gorm:"column:belong_user;type:varchar(255)"`
	PriceOfMonth     float64      `gorm:"column:price_of_month;type:decimal(10,2)"`
	ProjectID        string       `gorm:"column:project_id;type:varchar(255)"`
	CPUUsage         float64      `gorm:"column:cpu_usage;type:decimal(10,2)"`
	WanUsage         float64      `gorm:"column:wan_usage;type:decimal(10,2)"`
	CloudAccountID   int          `gorm:"index:cloud_account_id;column:cloud_account_id;type:int(11)"`
	CloudAccount     CloudAccount `gorm:"joinForeignKey:cloud_account_id;foreignKey:id"`
	EndTime          time.Time    `gorm:"column:end_time;type:datetime"`
	GameID           int          `gorm:"column:game_id;type:int(11)"`
	PayType          string       `gorm:"column:pay_type;type:varchar(64)"`
	ResourceAddress  string       `gorm:"column:resource_address;type:varchar(255)"`
	ResourceID       string       `gorm:"column:resource_id;type:varchar(255)"`
	ResourceName     string       `gorm:"column:resource_name;type:varchar(255)"`
	ResourceRegionID string       `gorm:"column:resource_region_id;type:varchar(255)"`
	Status           string       `gorm:"column:status;type:varchar(64)"`
	IsDelete         bool         `gorm:"column:is_delete;type:tinyint(1)"`
}

/******sql******
CREATE TABLE `cloud_server_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) DEFAULT NULL,
  `host_id` varchar(255) DEFAULT NULL,
  `host_ip` varchar(255) DEFAULT NULL,
  `host_ip_id` varchar(255) DEFAULT NULL,
  `host_lip` varchar(255) DEFAULT NULL,
  `host_dead_time` varchar(255) DEFAULT NULL,
  `host_region_id` varchar(255) DEFAULT NULL,
  `belong_user` varchar(255) DEFAULT NULL,
  `project_id` varchar(255) DEFAULT NULL,
  `host_name` varchar(255) DEFAULT NULL,
  `price_of_month` decimal(10,2) DEFAULT NULL,
  `room_id` int(11) DEFAULT NULL,
  `mem` varchar(255) DEFAULT NULL,
  `cpu` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `server_id` (`server_id`),
  KEY `room_id` (`room_id`),
  CONSTRAINT `cloud_server_info_ibfk_1` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`),
  CONSTRAINT `cloud_server_info_ibfk_2` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1791 DEFAULT CHARSET=utf8mb4
******sql******/
// CloudServerInfo [...]
type CloudServerInfo struct {
	ID           int     `gorm:"primaryKey;column:id;type:int(11);not null"`
	ServerID     int     `gorm:"index:server_id;column:server_id;type:int(11)"`
	Server       Server  `gorm:"joinForeignKey:server_id;foreignKey:id"`
	HostID       string  `gorm:"column:host_id;type:varchar(255)"`
	HostIP       string  `gorm:"column:host_ip;type:varchar(255)"`
	HostIPID     string  `gorm:"column:host_ip_id;type:varchar(255)"`
	HostLip      string  `gorm:"column:host_lip;type:varchar(255)"`
	HostDeadTime string  `gorm:"column:host_dead_time;type:varchar(255)"`
	HostRegionID string  `gorm:"column:host_region_id;type:varchar(255)"`
	BelongUser   string  `gorm:"column:belong_user;type:varchar(255)"`
	ProjectID    string  `gorm:"column:project_id;type:varchar(255)"`
	HostName     string  `gorm:"column:host_name;type:varchar(255)"`
	PriceOfMonth float64 `gorm:"column:price_of_month;type:decimal(10,2)"`
	RoomID       int     `gorm:"index:room_id;column:room_id;type:int(11)"`
	Room         Room    `gorm:"joinForeignKey:room_id;foreignKey:id"`
	Mem          string  `gorm:"column:mem;type:varchar(255)"`
	CPU          string  `gorm:"column:cpu;type:varchar(32)"`
}

/******sql******
CREATE TABLE `delete_service_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game` varchar(128) DEFAULT NULL,
  `platform` varchar(128) DEFAULT NULL,
  `service` varchar(128) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  `reason` varchar(256) DEFAULT NULL,
  `user_name` varchar(128) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2405 DEFAULT CHARSET=utf8mb4
******sql******/
// DeleteServiceLog [...]
type DeleteServiceLog struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Game       string    `gorm:"column:game;type:varchar(128)"`
	Platform   string    `gorm:"column:platform;type:varchar(128)"`
	Service    string    `gorm:"column:service;type:varchar(128)"`
	Number     int       `gorm:"column:number;type:int(11)"`
	Reason     string    `gorm:"column:reason;type:varchar(256)"`
	UserName   string    `gorm:"column:user_name;type:varchar(128)"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime"`
}

/******sql******
CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`),
  CONSTRAINT `department_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `department` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Department [...]
type Department struct {
	ID         int         `gorm:"primaryKey;column:id;type:int(11);not null"`
	ParentID   int         `gorm:"index:parent_id;column:parent_id;type:int(11)"`
	Department *Department `gorm:"joinForeignKey:parent_id;foreignKey:id"`
	Name       string      `gorm:"column:name;type:varchar(255)"`
	CreateDate time.Time   `gorm:"column:create_date;type:datetime"`
}

/******sql******
CREATE TABLE `ding_work_sheet` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` text,
  `level` smallint(6) DEFAULT NULL,
  `submit_user` varchar(255) DEFAULT NULL,
  `_submit_user_info` text,
  `create_time` datetime DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `work_sheet_id` int(11) DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `done_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1884 DEFAULT CHARSET=utf8mb4
******sql******/
// DingWorkSheet [...]
type DingWorkSheet struct {
	ID             int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Text           string    `gorm:"column:text;type:text"`
	Level          int16     `gorm:"column:level;type:smallint(6)"`
	SubmitUser     string    `gorm:"column:submit_user;type:varchar(255)"`
	SubmitUserInfo string    `gorm:"column:_submit_user_info;type:text"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`
	Status         int16     `gorm:"column:status;type:smallint(6)"`
	WorkSheetID    int       `gorm:"column:work_sheet_id;type:int(11)"`
	EndTime        time.Time `gorm:"column:end_time;type:datetime"`
	DoneTime       time.Time `gorm:"column:done_time;type:datetime"`
}

/******sql******
CREATE TABLE `dns` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(64) NOT NULL,
  `provider` varchar(255) DEFAULT NULL,
  `account` varchar(64) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `clound_account_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_dns_domain` (`domain`),
  KEY `clound_account_id` (`clound_account_id`),
  CONSTRAINT `dns_ibfk_1` FOREIGN KEY (`clound_account_id`) REFERENCES `cloud_account` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4
******sql******/
// DNS [...]
type DNS struct {
	ID              int          `gorm:"primaryKey;column:id;type:int(11);not null"`
	Domain          string       `gorm:"unique;column:domain;type:varchar(64);not null"`
	Provider        string       `gorm:"column:provider;type:varchar(255)"`
	Account         string       `gorm:"column:account;type:varchar(64)"`
	Active          bool         `gorm:"column:active;type:tinyint(1)"`
	CreateTime      time.Time    `gorm:"column:create_time;type:datetime"`
	UpdateTime      time.Time    `gorm:"column:update_time;type:datetime"`
	CloundAccountID int          `gorm:"index:clound_account_id;column:clound_account_id;type:int(11)"`
	CloudAccount    CloudAccount `gorm:"joinForeignKey:clound_account_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Domain [...]
type Domain struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Domain      string    `gorm:"column:domain;type:varchar(255)"`
	Active      bool      `gorm:"column:active;type:tinyint(1)"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `game` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `en_name` varchar(255) DEFAULT NULL,
  `game_type` smallint(6) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `max_service` smallint(6) DEFAULT NULL,
  `number_step` smallint(6) DEFAULT NULL,
  `is_manager` tinyint(1) DEFAULT NULL,
  `dns_level` smallint(6) DEFAULT NULL,
  `ali_accesskeyid` varchar(255) DEFAULT NULL,
  `ali_accesskeysecret` varchar(255) DEFAULT NULL,
  `dns_id` int(11) DEFAULT NULL,
  `salt_api_id` int(11) DEFAULT NULL,
  `pillar` text,
  `is_hunbu` tinyint(1) DEFAULT NULL,
  `manager_user` smallint(6) DEFAULT NULL,
  `merge_type` smallint(6) DEFAULT NULL,
  `business_user` varchar(255) DEFAULT NULL,
  `zabbix_group_id` int(11) DEFAULT NULL,
  `v1` tinyint(1) DEFAULT NULL,
  `step_info` mediumtext,
  `qq` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `dns_id` (`dns_id`),
  KEY `salt_api_id` (`salt_api_id`),
  CONSTRAINT `game_ibfk_1` FOREIGN KEY (`dns_id`) REFERENCES `dns` (`id`),
  CONSTRAINT `game_ibfk_2` FOREIGN KEY (`salt_api_id`) REFERENCES `salt_api` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=233 DEFAULT CHARSET=utf8mb4
******sql******/
// Game [...]
type Game struct {
	ID                 int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name               string    `gorm:"unique;column:name;type:varchar(32);not null"`
	EnName             string    `gorm:"column:en_name;type:varchar(255)"`
	GameType           int16     `gorm:"column:game_type;type:smallint(6)"`
	Active             bool      `gorm:"column:active;type:tinyint(1)"`
	CreateTime         time.Time `gorm:"column:create_time;type:datetime"`
	MaxService         int16     `gorm:"column:max_service;type:smallint(6)"`
	NumberStep         int16     `gorm:"column:number_step;type:smallint(6)"`
	IsManager          bool      `gorm:"column:is_manager;type:tinyint(1)"`
	DNSLevel           int16     `gorm:"column:dns_level;type:smallint(6)"`
	AliAccesskeyid     string    `gorm:"column:ali_accesskeyid;type:varchar(255)"`
	AliAccesskeysecret string    `gorm:"column:ali_accesskeysecret;type:varchar(255)"`
	DNSID              int       `gorm:"index:dns_id;column:dns_id;type:int(11)"`
	DNS                DNS       `gorm:"joinForeignKey:dns_id;foreignKey:id"`
	SaltAPIID          int       `gorm:"index:salt_api_id;column:salt_api_id;type:int(11)"`
	SaltAPI            SaltAPI   `gorm:"joinForeignKey:salt_api_id;foreignKey:id"`
	Pillar             string    `gorm:"column:pillar;type:text"`
	IsHunbu            bool      `gorm:"column:is_hunbu;type:tinyint(1)"`
	ManagerUser        int16     `gorm:"column:manager_user;type:smallint(6)"`
	MergeType          int16     `gorm:"column:merge_type;type:smallint(6)"`
	BusinessUser       string    `gorm:"column:business_user;type:varchar(255)"`
	ZabbixGroupID      int       `gorm:"column:zabbix_group_id;type:int(11)"`
	V1                 bool      `gorm:"column:v1;type:tinyint(1)"`
	StepInfo           string    `gorm:"column:step_info;type:mediumtext"`
	Qq                 string    `gorm:"column:qq;type:varchar(255)"`
}

/******sql******
CREATE TABLE `game_alias` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `name` varchar(32) NOT NULL,
  `en_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `game_id` (`game_id`),
  CONSTRAINT `game_alias_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4
******sql******/
// GameAlias [...]
type GameAlias struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID int    `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game   Game   `gorm:"joinForeignKey:game_id;foreignKey:id"`
	Name   string `gorm:"unique;column:name;type:varchar(32);not null"`
	EnName string `gorm:"column:en_name;type:varchar(255)"`
}

/******sql******
CREATE TABLE `game_app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `app_id` varchar(128) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_id` (`app_id`),
  KEY `game_id` (`game_id`),
  CONSTRAINT `game_app_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4
******sql******/
// GameApp [...]
type GameApp struct {
	ID     int    `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID int    `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game   Game   `gorm:"joinForeignKey:game_id;foreignKey:id"`
	AppID  string `gorm:"unique;column:app_id;type:varchar(128);not null"`
}

/******sql******
CREATE TABLE `game_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) NOT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `service_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `config` text NOT NULL,
  `path` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_game_config_game_id` (`game_id`),
  KEY `ix_game_config_platform_id` (`platform_id`),
  CONSTRAINT `game_config_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=205 DEFAULT CHARSET=utf8mb4
******sql******/
// GameConfig [...]
type GameConfig struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID     int       `gorm:"index:ix_game_config_game_id;column:game_id;type:int(11);not null"`
	Game       Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	PlatformID int       `gorm:"index:ix_game_config_platform_id;column:platform_id;type:int(11)"`
	ServiceID  int       `gorm:"column:service_id;type:int(11);not null"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	Config     string    `gorm:"column:config;type:text;not null"`
	Path       string    `gorm:"column:path;type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
}

/******sql******
CREATE TABLE `game_config_kwargs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) NOT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `service_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `value_type` varchar(20) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_game_config_kwargs_game_id` (`game_id`),
  KEY `ix_game_config_kwargs_platform_id` (`platform_id`),
  CONSTRAINT `game_config_kwargs_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=261 DEFAULT CHARSET=utf8mb4
******sql******/
// GameConfigKwargs [...]
type GameConfigKwargs struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID      int       `gorm:"index:ix_game_config_kwargs_game_id;column:game_id;type:int(11);not null"`
	Game        Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	PlatformID  int       `gorm:"index:ix_game_config_kwargs_platform_id;column:platform_id;type:int(11)"`
	ServiceID   int       `gorm:"column:service_id;type:int(11);not null"`
	Name        string    `gorm:"column:name;type:varchar(255)"`
	ValueType   string    `gorm:"column:value_type;type:varchar(20)"`
	Value       string    `gorm:"column:value;type:varchar(255)"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`
}

/******sql******
CREATE TABLE `game_key` (
  `game_id` int(11) NOT NULL,
  `key_id` int(11) NOT NULL,
  PRIMARY KEY (`game_id`,`key_id`),
  UNIQUE KEY `UC_game_id_key_id` (`game_id`,`key_id`),
  KEY `key_id` (`key_id`),
  CONSTRAINT `game_key_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `game_key_ibfk_2` FOREIGN KEY (`key_id`) REFERENCES `public_key` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// GameKey [...]
type GameKey struct {
	GameID    int       `gorm:"primaryKey;uniqueIndex:UC_game_id_key_id;column:game_id;type:int(11);not null"`
	Game      Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	KeyID     int       `gorm:"primaryKey;uniqueIndex:UC_game_id_key_id;index:key_id;column:key_id;type:int(11);not null"`
	PublicKey PublicKey `gorm:"joinForeignKey:key_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `game_platform` (
  `game_id` int(11) NOT NULL,
  `platform_id` int(11) NOT NULL,
  `platform_number` int(11) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `client_url` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `recharge_key` varchar(255) DEFAULT NULL,
  `recharge_secret` varchar(255) DEFAULT NULL,
  `choose_room` varchar(255) DEFAULT NULL,
  `platform_log_id` int(11) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `ext1` varchar(255) DEFAULT NULL,
  `ext2` varchar(255) DEFAULT NULL,
  `kwargs` text,
  PRIMARY KEY (`game_id`,`platform_id`),
  UNIQUE KEY `game_id` (`game_id`,`platform_id`),
  KEY `platform_id` (`platform_id`),
  CONSTRAINT `game_platform_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `game_platform_ibfk_2` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// GamePlatform [...]
type GamePlatform struct {
	GameID         int       `gorm:"primaryKey;uniqueIndex:game_id;column:game_id;type:int(11);not null"`
	Game           Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	PlatformID     int       `gorm:"primaryKey;uniqueIndex:game_id;index:platform_id;column:platform_id;type:int(11);not null"`
	Platform       Platform  `gorm:"joinForeignKey:platform_id;foreignKey:id"`
	PlatformNumber int       `gorm:"column:platform_number;type:int(11)"`
	Website        string    `gorm:"column:website;type:varchar(255)"`
	ClientURL      string    `gorm:"column:client_url;type:varchar(255)"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`
	RechargeKey    string    `gorm:"column:recharge_key;type:varchar(255)"`
	RechargeSecret string    `gorm:"column:recharge_secret;type:varchar(255)"`
	ChooseRoom     string    `gorm:"column:choose_room;type:varchar(255)"`
	PlatformLogID  int       `gorm:"column:platform_log_id;type:int(11)"`
	Active         bool      `gorm:"column:active;type:tinyint(1)"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime"`
	DeleteTime     time.Time `gorm:"column:delete_time;type:datetime"`
	Ext1           string    `gorm:"column:ext1;type:varchar(255)"`
	Ext2           string    `gorm:"column:ext2;type:varchar(255)"`
	Kwargs         string    `gorm:"column:kwargs;type:text"`
}

/******sql******
CREATE TABLE `game_services` (
  `game_id` int(11) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  UNIQUE KEY `UC_game_id_service_id` (`game_id`,`service_id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `game_services_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `game_services_ibfk_2` FOREIGN KEY (`service_id`) REFERENCES `service` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// GameServices [...]
type GameServices struct {
	GameID    int     `gorm:"uniqueIndex:UC_game_id_service_id;column:game_id;type:int(11)"`
	Game      Game    `gorm:"joinForeignKey:game_id;foreignKey:id"`
	ServiceID int     `gorm:"uniqueIndex:UC_game_id_service_id;index:service_id;column:service_id;type:int(11)"`
	Service   Service `gorm:"joinForeignKey:service_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `https_url` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) DEFAULT NULL,
  `not_after` datetime DEFAULT NULL,
  `cert` text,
  `key` text,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8mb4
******sql******/
// HTTPSURL [...]
type HTTPSURL struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	URL        string    `gorm:"column:url;type:varchar(255)"`
	NotAfter   time.Time `gorm:"column:not_after;type:datetime"`
	Cert       string    `gorm:"column:cert;type:text"`
	Key        string    `gorm:"column:key;type:text"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `inventory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room_id` int(11) DEFAULT NULL,
  `category` varchar(64) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `size` varchar(255) DEFAULT NULL,
  `numbers` smallint(6) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `update_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `room_id` (`room_id`),
  CONSTRAINT `inventory_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4
******sql******/
// Inventory [...]
type Inventory struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	RoomID      int       `gorm:"index:room_id;column:room_id;type:int(11)"`
	Room        Room      `gorm:"joinForeignKey:room_id;foreignKey:id"`
	Category    string    `gorm:"column:category;type:varchar(64)"`
	Brand       string    `gorm:"column:brand;type:varchar(255)"`
	Size        string    `gorm:"column:size;type:varchar(255)"`
	Numbers     int16     `gorm:"column:numbers;type:smallint(6)"`
	State       string    `gorm:"column:state;type:varchar(255)"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	CreateDate  time.Time `gorm:"column:create_date;type:datetime"`
	UpdateDate  time.Time `gorm:"column:update_date;type:datetime"`
}

/******sql******
CREATE TABLE `merge_manager_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `to_from` varchar(32) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `is_pass` tinyint(1) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `merge_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`),
  KEY `platform_id` (`platform_id`),
  CONSTRAINT `merge_manager_log_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `merge_manager_log_ibfk_2` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8616 DEFAULT CHARSET=utf8mb4
******sql******/
// MergeManagerLog [...]
type MergeManagerLog struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	ToFrom     string    `gorm:"column:to_from;type:varchar(32)"`
	GameID     int       `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game       Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	PlatformID int       `gorm:"index:platform_id;column:platform_id;type:int(11)"`
	Platform   Platform  `gorm:"joinForeignKey:platform_id;foreignKey:id"`
	IsPass     bool      `gorm:"column:is_pass;type:tinyint(1)"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime"`
	MergeTime  time.Time `gorm:"column:merge_time;type:datetime"`
}

/******sql******
CREATE TABLE `mergeservice_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(255) DEFAULT NULL,
  `game_en_name` varchar(255) DEFAULT NULL,
  `platform_name` varchar(255) DEFAULT NULL,
  `platform_en_name` varchar(255) DEFAULT NULL,
  `from_service_number` int(11) NOT NULL,
  `from_server_ip` varchar(15) DEFAULT NULL,
  `from_server_id` int(11) DEFAULT NULL,
  `from_db_ip` varchar(15) DEFAULT NULL,
  `to_service_number` int(11) NOT NULL,
  `to_server_id` int(11) DEFAULT NULL,
  `to_server_ip` varchar(15) DEFAULT NULL,
  `to_db_ip` varchar(15) DEFAULT NULL,
  `kwargs` text,
  `status` tinyint(1) DEFAULT NULL,
  `timestamp` int(11) DEFAULT NULL,
  `merge_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_mergeservice_log_from_service_number` (`from_service_number`),
  KEY `ix_mergeservice_log_game_name` (`game_name`(191)),
  KEY `ix_mergeservice_log_platform_en_name` (`platform_en_name`(191)),
  KEY `ix_mergeservice_log_to_service_number` (`to_service_number`)
) ENGINE=InnoDB AUTO_INCREMENT=8961 DEFAULT CHARSET=utf8mb4
******sql******/
// MergeserviceLog [...]
type MergeserviceLog struct {
	ID                int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameName          string    `gorm:"index:ix_mergeservice_log_game_name;column:game_name;type:varchar(255)"`
	GameEnName        string    `gorm:"column:game_en_name;type:varchar(255)"`
	PlatformName      string    `gorm:"column:platform_name;type:varchar(255)"`
	PlatformEnName    string    `gorm:"index:ix_mergeservice_log_platform_en_name;column:platform_en_name;type:varchar(255)"`
	FromServiceNumber int       `gorm:"index:ix_mergeservice_log_from_service_number;column:from_service_number;type:int(11);not null"`
	FromServerIP      string    `gorm:"column:from_server_ip;type:varchar(15)"`
	FromServerID      int       `gorm:"column:from_server_id;type:int(11)"`
	FromDbIP          string    `gorm:"column:from_db_ip;type:varchar(15)"`
	ToServiceNumber   int       `gorm:"index:ix_mergeservice_log_to_service_number;column:to_service_number;type:int(11);not null"`
	ToServerID        int       `gorm:"column:to_server_id;type:int(11)"`
	ToServerIP        string    `gorm:"column:to_server_ip;type:varchar(15)"`
	ToDbIP            string    `gorm:"column:to_db_ip;type:varchar(15)"`
	Kwargs            string    `gorm:"column:kwargs;type:text"`
	Status            bool      `gorm:"column:status;type:tinyint(1)"`
	Timestamp         int       `gorm:"column:timestamp;type:int(11)"`
	MergeTime         time.Time `gorm:"column:merge_time;type:datetime"`
}

/******sql******
CREATE TABLE `migrateservice_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_name` varchar(255) DEFAULT NULL,
  `game_en_name` varchar(255) DEFAULT NULL,
  `platform_name` varchar(255) DEFAULT NULL,
  `platform_en_name` varchar(255) DEFAULT NULL,
  `from_service_number` int(11) NOT NULL,
  `from_server_ip` varchar(15) DEFAULT NULL,
  `from_server_id` int(11) DEFAULT NULL,
  `to_server_id` int(11) DEFAULT NULL,
  `to_server_ip` varchar(15) DEFAULT NULL,
  `migrate_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_migrateservice_log_game_name` (`game_name`(191)),
  KEY `ix_migrateservice_log_platform_en_name` (`platform_en_name`(191))
) ENGINE=InnoDB AUTO_INCREMENT=1753 DEFAULT CHARSET=utf8mb4
******sql******/
// MigrateserviceLog [...]
type MigrateserviceLog struct {
	ID                int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameName          string    `gorm:"index:ix_migrateservice_log_game_name;column:game_name;type:varchar(255)"`
	GameEnName        string    `gorm:"column:game_en_name;type:varchar(255)"`
	PlatformName      string    `gorm:"column:platform_name;type:varchar(255)"`
	PlatformEnName    string    `gorm:"index:ix_migrateservice_log_platform_en_name;column:platform_en_name;type:varchar(255)"`
	FromServiceNumber int       `gorm:"column:from_service_number;type:int(11);not null"`
	FromServerIP      string    `gorm:"column:from_server_ip;type:varchar(15)"`
	FromServerID      int       `gorm:"column:from_server_id;type:int(11)"`
	ToServerID        int       `gorm:"column:to_server_id;type:int(11)"`
	ToServerIP        string    `gorm:"column:to_server_ip;type:varchar(15)"`
	MigrateTime       time.Time `gorm:"column:migrate_time;type:datetime"`
}

/******sql******
CREATE TABLE `open_time_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `platform_id` int(11) DEFAULT NULL,
  `platform_en_name` varchar(255) DEFAULT NULL,
  `platform_name` varchar(255) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `game_en_name` varchar(255) DEFAULT NULL,
  `game_name` varchar(255) DEFAULT NULL,
  `number` int(11) DEFAULT NULL,
  `open_time` datetime DEFAULT NULL,
  `result` smallint(6) DEFAULT NULL,
  `info` text,
  `create_time` datetime DEFAULT NULL,
  `create_user` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1036 DEFAULT CHARSET=utf8mb4
******sql******/
// OpenTimeLog [...]
type OpenTimeLog struct {
	ID             int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	PlatformID     int       `gorm:"column:platform_id;type:int(11)"`
	PlatformEnName string    `gorm:"column:platform_en_name;type:varchar(255)"`
	PlatformName   string    `gorm:"column:platform_name;type:varchar(255)"`
	GameID         int       `gorm:"column:game_id;type:int(11)"`
	GameEnName     string    `gorm:"column:game_en_name;type:varchar(255)"`
	GameName       string    `gorm:"column:game_name;type:varchar(255)"`
	Number         int       `gorm:"column:number;type:int(11)"`
	OpenTime       time.Time `gorm:"column:open_time;type:datetime"`
	Result         int16     `gorm:"column:result;type:smallint(6)"`
	Info           string    `gorm:"column:info;type:text"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`
	CreateUser     string    `gorm:"column:create_user;type:varchar(255)"`
}

/******sql******
CREATE TABLE `operate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operate_group_id` int(11) DEFAULT NULL,
  `op_type` varchar(32) DEFAULT NULL,
  `conditions` text,
  `ignore_err` tinyint(1) DEFAULT NULL,
  `fun` varchar(255) DEFAULT NULL,
  `tgt` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `last_step_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `operate_group_id` (`operate_group_id`),
  KEY `last_step_id` (`last_step_id`),
  CONSTRAINT `operate_ibfk_2` FOREIGN KEY (`operate_group_id`) REFERENCES `operate_group` (`id`) ON DELETE CASCADE,
  CONSTRAINT `operate_ibfk_3` FOREIGN KEY (`last_step_id`) REFERENCES `operate` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1145 DEFAULT CHARSET=utf8mb4
******sql******/
// Operate [...]
type Operate struct {
	ID             int          `gorm:"primaryKey;column:id;type:int(11);not null"`
	OperateGroupID int          `gorm:"index:operate_group_id;column:operate_group_id;type:int(11)"`
	OperateGroup   OperateGroup `gorm:"joinForeignKey:operate_group_id;foreignKey:id"`
	OpType         string       `gorm:"column:op_type;type:varchar(32)"`
	Conditions     string       `gorm:"column:conditions;type:text"`
	IgnoreErr      bool         `gorm:"column:ignore_err;type:tinyint(1)"`
	Fun            string       `gorm:"column:fun;type:varchar(255)"`
	Tgt            string       `gorm:"column:tgt;type:varchar(255)"`
	CreateTime     time.Time    `gorm:"column:create_time;type:datetime"`
	Description    string       `gorm:"column:description;type:varchar(255)"`
	LastStepID     int          `gorm:"index:last_step_id;column:last_step_id;type:int(11)"`
	Operate        *Operate     `gorm:"joinForeignKey:last_step_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `operate_args` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operate_id` int(11) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `operate_id` (`operate_id`),
  CONSTRAINT `operate_args_ibfk_1` FOREIGN KEY (`operate_id`) REFERENCES `operate` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1897 DEFAULT CHARSET=utf8mb4
******sql******/
// OperateArgs [...]
type OperateArgs struct {
	ID        int     `gorm:"primaryKey;column:id;type:int(11);not null"`
	OperateID int     `gorm:"index:operate_id;column:operate_id;type:int(11)"`
	Operate   Operate `gorm:"joinForeignKey:operate_id;foreignKey:id"`
	Value     string  `gorm:"column:value;type:varchar(255)"`
}

/******sql******
CREATE TABLE `operate_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(255) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `operate_group_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `operate_group_ibfk_2` FOREIGN KEY (`service_id`) REFERENCES `service` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=328 DEFAULT CHARSET=utf8mb4
******sql******/
// OperateGroup [...]
type OperateGroup struct {
	ID          int     `gorm:"primaryKey;column:id;type:int(11);not null"`
	Type        string  `gorm:"column:type;type:varchar(255)"`
	GameID      int     `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game        Game    `gorm:"joinForeignKey:game_id;foreignKey:id"`
	ServiceID   int     `gorm:"index:service_id;column:service_id;type:int(11)"`
	Service     Service `gorm:"joinForeignKey:service_id;foreignKey:id"`
	Description string  `gorm:"column:description;type:varchar(255)"`
}

/******sql******
CREATE TABLE `permission` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `is_menu` tinyint(1) DEFAULT NULL,
  `can_get` tinyint(1) DEFAULT NULL,
  `can_post` tinyint(1) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Permission [...]
type Permission struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	IsMenu     bool      `gorm:"column:is_menu;type:tinyint(1)"`
	CanGet     bool      `gorm:"column:can_get;type:tinyint(1)"`
	CanPost    bool      `gorm:"column:can_post;type:tinyint(1)"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime"`
}

/******sql******
CREATE TABLE `platform` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `en_name` varchar(64) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `en_name` (`en_name`)
) ENGINE=InnoDB AUTO_INCREMENT=188 DEFAULT CHARSET=utf8mb4
******sql******/
// Platform [...]
type Platform struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Name       string    `gorm:"unique;column:name;type:varchar(64);not null" json:"name"`
	EnName     string    `gorm:"unique;column:en_name;type:varchar(64)" json:"en_name"`
	Active     bool      `gorm:"column:active;type:tinyint(1)" json:"active"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
}

/******sql******
CREATE TABLE `public_key` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(64) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `key` text,
  `create_date` datetime DEFAULT NULL,
  `is_default` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_public_key_user` (`user`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4
******sql******/
// PublicKey [...]
type PublicKey struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	User        string    `gorm:"unique;column:user;type:varchar(64);not null"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	Key         string    `gorm:"column:key;type:text"`
	CreateDate  time.Time `gorm:"column:create_date;type:datetime"`
	IsDefault   bool      `gorm:"column:is_default;type:tinyint(1)"`
}

/******sql******
CREATE TABLE `region` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `clound_account_id` int(11) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `code` varchar(64) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `clound_account_id` (`clound_account_id`),
  CONSTRAINT `region_ibfk_1` FOREIGN KEY (`clound_account_id`) REFERENCES `cloud_account` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Region [...]
type Region struct {
	ID              int          `gorm:"primaryKey;column:id;type:int(11);not null"`
	CloundAccountID int          `gorm:"index:clound_account_id;column:clound_account_id;type:int(11)"`
	CloudAccount    CloudAccount `gorm:"joinForeignKey:clound_account_id;foreignKey:id"`
	Name            string       `gorm:"column:name;type:varchar(64)"`
	Code            string       `gorm:"column:code;type:varchar(64)"`
	CreateTime      time.Time    `gorm:"column:create_time;type:datetime"`
	UpdateTime      time.Time    `gorm:"column:update_time;type:datetime"`
}

/******sql******
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// Role [...]
type Role struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime"`
}

/******sql******
CREATE TABLE `role_permission` (
  `permission_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  UNIQUE KEY `permission_id` (`permission_id`,`role_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `role_permission_ibfk_1` FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`),
  CONSTRAINT `role_permission_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// RolePermission [...]
type RolePermission struct {
	PermissionID int        `gorm:"uniqueIndex:permission_id;column:permission_id;type:int(11)"`
	Permission   Permission `gorm:"joinForeignKey:permission_id;foreignKey:id"`
	RoleID       int        `gorm:"uniqueIndex:permission_id;index:role_id;column:role_id;type:int(11)"`
	Role         Role       `gorm:"joinForeignKey:role_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `en_name` varchar(64) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `is_cloud` tinyint(1) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `salt_api_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `ix_room_en_name` (`en_name`),
  KEY `salt_api_id` (`salt_api_id`),
  CONSTRAINT `room_ibfk_1` FOREIGN KEY (`salt_api_id`) REFERENCES `salt_api` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4
******sql******/
// Room [...]
type Room struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(64);not null"`
	EnName     string    `gorm:"unique;column:en_name;type:varchar(64)"`
	Company    string    `gorm:"column:company;type:varchar(255)"`
	IsCloud    bool      `gorm:"column:is_cloud;type:tinyint(1)"`
	Address    string    `gorm:"column:address;type:varchar(255)"`
	Phone      string    `gorm:"column:phone;type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	SaltAPIID  int       `gorm:"index:salt_api_id;column:salt_api_id;type:int(11)"`
	SaltAPI    SaltAPI   `gorm:"joinForeignKey:salt_api_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `salt_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(255) DEFAULT NULL,
  `port` int(11) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `salt_key` varchar(255) DEFAULT NULL,
  `is_default` tinyint(1) DEFAULT NULL,
  `create_user_id` int(11) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `lhost` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4
******sql******/
// SaltAPI [...]
type SaltAPI struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Host         string    `gorm:"column:host;type:varchar(255)"`
	Port         int       `gorm:"column:port;type:int(11)"`
	Username     string    `gorm:"column:username;type:varchar(255)"`
	Password     string    `gorm:"column:password;type:varchar(255)"`
	SaltKey      string    `gorm:"column:salt_key;type:varchar(255)"`
	IsDefault    bool      `gorm:"column:is_default;type:tinyint(1)"`
	CreateUserID int       `gorm:"column:create_user_id;type:int(11)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	Lhost        string    `gorm:"column:lhost;type:varchar(15)"`
}

/******sql******
CREATE TABLE `server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(15) NOT NULL,
  `ip_interface_name` varchar(64) DEFAULT NULL,
  `uip` varchar(15) DEFAULT NULL,
  `tip` varchar(15) DEFAULT NULL,
  `lip` varchar(15) DEFAULT NULL,
  `lip_interface_name` varchar(64) DEFAULT NULL,
  `port` varchar(6) DEFAULT NULL,
  `cpu` varchar(255) DEFAULT NULL,
  `cpu_core` varchar(255) DEFAULT NULL,
  `disk` varchar(255) DEFAULT NULL,
  `memory` varchar(255) DEFAULT NULL,
  `os` varchar(255) DEFAULT NULL,
  `asset_tag` varchar(255) DEFAULT NULL,
  `serial_number` varchar(255) DEFAULT NULL,
  `unit_number` int(11) DEFAULT NULL,
  `rack_number` varchar(255) DEFAULT NULL,
  `add_date` varchar(10) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `mode_number` varchar(64) DEFAULT NULL,
  `in_use` tinyint(1) DEFAULT NULL,
  `salt_key` varchar(128) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `create_user_id` int(11) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `update_user_id` int(11) DEFAULT NULL,
  `init_user` varchar(255) DEFAULT NULL,
  `init_password` varchar(255) DEFAULT NULL,
  `salt_status` smallint(6) DEFAULT NULL,
  `base_status` tinyint(1) DEFAULT NULL,
  `price_of_month` decimal(10,2) DEFAULT NULL,
  `is_deposit` tinyint(1) DEFAULT NULL,
  `depositor` varchar(64) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `room_id` int(11) NOT NULL,
  `history_game_id` int(11) DEFAULT NULL,
  `is_free` tinyint(1) DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `delete_user_id` int(11) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `ignore_port_scan` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_server_ip` (`ip`),
  UNIQUE KEY `tip` (`tip`),
  UNIQUE KEY `uip` (`uip`),
  UNIQUE KEY `ix_server_salt_key` (`salt_key`),
  KEY `create_user_id` (`create_user_id`),
  KEY `game_id` (`game_id`),
  KEY `history_game_id` (`history_game_id`),
  KEY `room_id` (`room_id`),
  KEY `update_user_id` (`update_user_id`),
  KEY `delete_user_id` (`delete_user_id`),
  KEY `platform_id` (`platform_id`),
  CONSTRAINT `server_ibfk_1` FOREIGN KEY (`create_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `server_ibfk_2` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `server_ibfk_3` FOREIGN KEY (`history_game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `server_ibfk_4` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`),
  CONSTRAINT `server_ibfk_5` FOREIGN KEY (`update_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `server_ibfk_6` FOREIGN KEY (`delete_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `server_ibfk_7` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2715 DEFAULT CHARSET=utf8mb4
******sql******/
// Server [...]
type Server struct {
	ID               int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	IP               string    `gorm:"unique;column:ip;type:varchar(15);not null" json:"ip"`
	IPInterfaceName  string    `gorm:"column:ip_interface_name;type:varchar(64)" json:"ip_interface_name"`
	UIp              string    `gorm:"unique;column:uip;type:varchar(15)" json:"u_ip"`
	Tip              string    `gorm:"unique;column:tip;type:varchar(15)" json:"tip"`
	Lip              string    `gorm:"column:lip;type:varchar(15)" json:"lip"`
	LipInterfaceName string    `gorm:"column:lip_interface_name;type:varchar(64)" json:"lip_interface_name"`
	Port             string    `gorm:"column:port;type:varchar(6)" json:"port"`
	CPU              string    `gorm:"column:cpu;type:varchar(255)" json:"cpu"`
	CPUCore          string    `gorm:"column:cpu_core;type:varchar(255)" json:"cpu_core"`
	Disk             string    `gorm:"column:disk;type:varchar(255)" json:"disk"`
	Memory           string    `gorm:"column:memory;type:varchar(255)" json:"memory"`
	Os               string    `gorm:"column:os;type:varchar(255)" json:"os"`
	AssetTag         string    `gorm:"column:asset_tag;type:varchar(255)" json:"asset_tag"`
	SerialNumber     string    `gorm:"column:serial_number;type:varchar(255)" json:"serial_number"`
	UnitNumber       int       `gorm:"column:unit_number;type:int(11)" json:"unit_number"`
	RackNumber       string    `gorm:"column:rack_number;type:varchar(255)" json:"rack_number"`
	AddDate          string    `gorm:"column:add_date;type:varchar(10)" json:"add_date"`
	Brand            string    `gorm:"column:brand;type:varchar(255)" json:"brand"`
	ModeNumber       string    `gorm:"column:mode_number;type:varchar(64)" json:"mode_number"`
	InUse            bool      `gorm:"column:in_use;type:tinyint(1)" json:"in_use"`
	SaltKey          string    `gorm:"unique;column:salt_key;type:varchar(128)" json:"salt_key"`
	CreateTime       time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	CreateUserID     int       `gorm:"index:create_user_id;column:create_user_id;type:int(11)" json:"create_user_id"`
	User             User      `gorm:"joinForeignKey:create_user_id;foreignKey:id" json:"user"`
	UpdateTime       time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	UpdateUserID     int       `gorm:"index:update_user_id;column:update_user_id;type:int(11)" json:"update_user_id"`
	InitUser         string    `gorm:"column:init_user;type:varchar(255)" json:"init_user"`
	InitPassword     string    `gorm:"column:init_password;type:varchar(255)" json:"init_password"`
	SaltStatus       int16     `gorm:"column:salt_status;type:smallint(6)" json:"salt_status"`
	BaseStatus       bool      `gorm:"column:base_status;type:tinyint(1)" json:"base_status"`
	PriceOfMonth     float64   `gorm:"column:price_of_month;type:decimal(10,2)" json:"price_of_month"`
	IsDeposit        bool      `gorm:"column:is_deposit;type:tinyint(1)" json:"is_deposit"`
	Depositor        string    `gorm:"column:depositor;type:varchar(64)" json:"depositor"`
	GameID           int       `gorm:"index:game_id;column:game_id;type:int(11)" json:"game_id"`
	Game             Game      `gorm:"joinForeignKey:game_id;foreignKey:id" json:"game"`
	RoomID           int       `gorm:"index:room_id;column:room_id;type:int(11);not null" json:"room_id"`
	Room             Room      `gorm:"joinForeignKey:room_id;foreignKey:id" json:"room"`
	HistoryGameID    int       `gorm:"index:history_game_id;column:history_game_id;type:int(11)" json:"history_game_id"`
	IsFree           bool      `gorm:"column:is_free;type:tinyint(1)" json:"is_free"`
	DeleteTime       time.Time `gorm:"column:delete_time;type:datetime" json:"delete_time"`
	DeleteUserID     int       `gorm:"index:delete_user_id;column:delete_user_id;type:int(11)" json:"delete_user_id"`
	IsDelete         bool      `gorm:"column:is_delete;type:tinyint(1)" json:"is_delete"`
	PlatformID       int       `gorm:"index:platform_id;column:platform_id;type:int(11)" json:"platform_id"`
	Platform         Platform  `gorm:"joinForeignKey:platform_id;foreignKey:id" json:"platform"`
	IgnorePortScan   bool      `gorm:"column:ignore_port_scan;type:tinyint(1)" json:"ignore_port_scan"`
}

/******sql******
CREATE TABLE `server_config_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `room_id` int(11) DEFAULT NULL,
  `cpu` int(11) DEFAULT NULL,
  `memory` int(11) DEFAULT NULL,
  `diskspace` int(11) DEFAULT NULL,
  `disk_type` int(11) DEFAULT NULL,
  `ali_config_id` varchar(32) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `room_id` (`room_id`),
  CONSTRAINT `server_config_price_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4
******sql******/
// ServerConfigPrice [...]
type ServerConfigPrice struct {
	ID          int     `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name        string  `gorm:"column:name;type:varchar(255)"`
	RoomID      int     `gorm:"index:room_id;column:room_id;type:int(11)"`
	Room        Room    `gorm:"joinForeignKey:room_id;foreignKey:id"`
	CPU         int     `gorm:"column:cpu;type:int(11)"`
	Memory      int     `gorm:"column:memory;type:int(11)"`
	Diskspace   int     `gorm:"column:diskspace;type:int(11)"`
	DiskType    int     `gorm:"column:disk_type;type:int(11)"`
	AliConfigID string  `gorm:"column:ali_config_id;type:varchar(32)"`
	Price       float32 `gorm:"column:price;type:float"`
	Type        int     `gorm:"column:type;type:int(11)"`
}

/******sql******
CREATE TABLE `server_disk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) DEFAULT NULL,
  `slot` smallint(6) DEFAULT NULL,
  `brand` varchar(64) DEFAULT NULL,
  `disk_type` varchar(64) DEFAULT NULL,
  `disk_sn` varchar(255) DEFAULT NULL,
  `size` varchar(64) DEFAULT NULL,
  `state` varchar(64) DEFAULT NULL,
  `status` varchar(64) DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `is_ok` tinyint(1) DEFAULT NULL,
  `serial_number` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `server_id` (`server_id`),
  CONSTRAINT `server_disk_ibfk_1` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1255 DEFAULT CHARSET=utf8mb4
******sql******/
// ServerDisk [...]
type ServerDisk struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	ServerID     int       `gorm:"index:server_id;column:server_id;type:int(11)"`
	Server       Server    `gorm:"joinForeignKey:server_id;foreignKey:id"`
	Slot         int16     `gorm:"column:slot;type:smallint(6)"`
	Brand        string    `gorm:"column:brand;type:varchar(64)"`
	DiskType     string    `gorm:"column:disk_type;type:varchar(64)"`
	DiskSn       string    `gorm:"column:disk_sn;type:varchar(255)"`
	Size         string    `gorm:"column:size;type:varchar(64)"`
	State        string    `gorm:"column:state;type:varchar(64)"`
	Status       string    `gorm:"column:status;type:varchar(64)"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	IsOk         bool      `gorm:"column:is_ok;type:tinyint(1)"`
	SerialNumber string    `gorm:"column:serial_number;type:varchar(64)"`
}

/******sql******
CREATE TABLE `server_key` (
  `server_id` int(11) DEFAULT NULL,
  `key_id` int(11) DEFAULT NULL,
  UNIQUE KEY `server_id` (`server_id`,`key_id`),
  KEY `key_id` (`key_id`),
  CONSTRAINT `server_key_ibfk_1` FOREIGN KEY (`key_id`) REFERENCES `public_key` (`id`),
  CONSTRAINT `server_key_ibfk_2` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// ServerKey [...]
type ServerKey struct {
	ServerID  int       `gorm:"uniqueIndex:server_id;column:server_id;type:int(11)"`
	Server    Server    `gorm:"joinForeignKey:server_id;foreignKey:id"`
	KeyID     int       `gorm:"uniqueIndex:server_id;index:key_id;column:key_id;type:int(11)"`
	PublicKey PublicKey `gorm:"joinForeignKey:key_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `server_note` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) DEFAULT NULL,
  `level` int(11) DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `sheet_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `server_id` (`server_id`),
  KEY `ix_server_note_sheet_id` (`sheet_id`),
  CONSTRAINT `server_note_ibfk_1` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52871 DEFAULT CHARSET=utf8mb4
******sql******/
// ServerNote [...]
type ServerNote struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	ServerID   int       `gorm:"index:server_id;column:server_id;type:int(11)"`
	Server     Server    `gorm:"joinForeignKey:server_id;foreignKey:id"`
	Level      int       `gorm:"column:level;type:int(11)"`
	Note       string    `gorm:"column:note;type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
	SheetID    int       `gorm:"index:ix_server_note_sheet_id;column:sheet_id;type:int(11)"`
}

/******sql******
CREATE TABLE `server_service` (
  `server_id` int(11) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  UNIQUE KEY `server_id` (`server_id`,`service_id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `server_service_ibfk_1` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`),
  CONSTRAINT `server_service_ibfk_2` FOREIGN KEY (`service_id`) REFERENCES `service` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// ServerService [...]
type ServerService struct {
	ServerID  int     `gorm:"uniqueIndex:server_id;column:server_id;type:int(11)"`
	Server    Server  `gorm:"joinForeignKey:server_id;foreignKey:id"`
	ServiceID int     `gorm:"uniqueIndex:server_id;index:service_id;column:service_id;type:int(11)"`
	Service   Service `gorm:"joinForeignKey:service_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `server_zabbix_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) DEFAULT NULL,
  `hostid` varchar(255) DEFAULT NULL,
  `item_available_memory` varchar(255) DEFAULT NULL,
  `item_total_memory` varchar(255) DEFAULT NULL,
  `item_processor_load_1` varchar(255) DEFAULT NULL,
  `item_wan_intraffic` varchar(255) DEFAULT NULL,
  `item_wan_outtraffic` varchar(255) DEFAULT NULL,
  `max_processor_load_1` decimal(10,4) DEFAULT NULL,
  `min_processor_load_1` decimal(10,4) DEFAULT NULL,
  `avg_processor_load_1` decimal(10,4) DEFAULT NULL,
  `total_memory` decimal(10,2) DEFAULT NULL,
  `max_available_memory` decimal(10,2) DEFAULT NULL,
  `min_available_memory` decimal(10,2) DEFAULT NULL,
  `avg_available_memory` decimal(10,2) DEFAULT NULL,
  `max_usage_rate_memory` decimal(10,2) DEFAULT NULL,
  `min_usage_rate_memory` decimal(10,2) DEFAULT NULL,
  `avg_usage_rate_memory` decimal(10,2) DEFAULT NULL,
  `max_wan_intraffic` decimal(10,4) DEFAULT NULL,
  `min_wan_intraffic` decimal(10,4) DEFAULT NULL,
  `avg_wan_intraffic` decimal(10,4) DEFAULT NULL,
  `max_wan_outtraffic` decimal(10,4) DEFAULT NULL,
  `min_wan_outtraffic` decimal(10,4) DEFAULT NULL,
  `avg_wan_outtraffic` decimal(10,4) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `item_io_read` varchar(255) DEFAULT NULL,
  `item_io_write` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `server_id` (`server_id`),
  CONSTRAINT `server_zabbix_info_ibfk_1` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1847 DEFAULT CHARSET=utf8mb4
******sql******/
// ServerZabbixInfo [...]
type ServerZabbixInfo struct {
	ID                  int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	ServerID            int       `gorm:"index:server_id;column:server_id;type:int(11)"`
	Server              Server    `gorm:"joinForeignKey:server_id;foreignKey:id"`
	Hostid              string    `gorm:"column:hostid;type:varchar(255)"`
	ItemAvailableMemory string    `gorm:"column:item_available_memory;type:varchar(255)"`
	ItemTotalMemory     string    `gorm:"column:item_total_memory;type:varchar(255)"`
	ItemProcessorLoad1  string    `gorm:"column:item_processor_load_1;type:varchar(255)"`
	ItemWanIntraffic    string    `gorm:"column:item_wan_intraffic;type:varchar(255)"`
	ItemWanOuttraffic   string    `gorm:"column:item_wan_outtraffic;type:varchar(255)"`
	MaxProcessorLoad1   float64   `gorm:"column:max_processor_load_1;type:decimal(10,4)"`
	MinProcessorLoad1   float64   `gorm:"column:min_processor_load_1;type:decimal(10,4)"`
	AvgProcessorLoad1   float64   `gorm:"column:avg_processor_load_1;type:decimal(10,4)"`
	TotalMemory         float64   `gorm:"column:total_memory;type:decimal(10,2)"`
	MaxAvailableMemory  float64   `gorm:"column:max_available_memory;type:decimal(10,2)"`
	MinAvailableMemory  float64   `gorm:"column:min_available_memory;type:decimal(10,2)"`
	AvgAvailableMemory  float64   `gorm:"column:avg_available_memory;type:decimal(10,2)"`
	MaxUsageRateMemory  float64   `gorm:"column:max_usage_rate_memory;type:decimal(10,2)"`
	MinUsageRateMemory  float64   `gorm:"column:min_usage_rate_memory;type:decimal(10,2)"`
	AvgUsageRateMemory  float64   `gorm:"column:avg_usage_rate_memory;type:decimal(10,2)"`
	MaxWanIntraffic     float64   `gorm:"column:max_wan_intraffic;type:decimal(10,4)"`
	MinWanIntraffic     float64   `gorm:"column:min_wan_intraffic;type:decimal(10,4)"`
	AvgWanIntraffic     float64   `gorm:"column:avg_wan_intraffic;type:decimal(10,4)"`
	MaxWanOuttraffic    float64   `gorm:"column:max_wan_outtraffic;type:decimal(10,4)"`
	MinWanOuttraffic    float64   `gorm:"column:min_wan_outtraffic;type:decimal(10,4)"`
	AvgWanOuttraffic    float64   `gorm:"column:avg_wan_outtraffic;type:decimal(10,4)"`
	CreateTime          time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime          time.Time `gorm:"column:update_time;type:datetime"`
	ItemIoRead          string    `gorm:"column:item_io_read;type:varchar(255)"`
	ItemIoWrite         string    `gorm:"column:item_io_write;type:varchar(255)"`
}

/******sql******
CREATE TABLE `service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `en_name` varchar(64) NOT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_service_en_name` (`en_name`),
  UNIQUE KEY `ix_service_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4
******sql******/
// Service [...]
type Service struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(64);not null"`
	EnName     string    `gorm:"unique;column:en_name;type:varchar(64);not null"`
	Active     bool      `gorm:"column:active;type:tinyint(1)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `service_base` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) DEFAULT NULL,
  `game_id` int(11) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  `deploy_user_id` int(11) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  `config_kwargs` text,
  `deploy_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `number` smallint(6) DEFAULT NULL,
  `open_time` datetime DEFAULT NULL,
  `merge_time` datetime DEFAULT NULL,
  `cross_service_id` int(11) DEFAULT NULL,
  `merged_services` text,
  `version` varchar(64) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `no_backup_warning` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deploy_user_id` (`deploy_user_id`),
  KEY `game_id` (`game_id`),
  KEY `platform_id` (`platform_id`),
  KEY `server_id` (`server_id`),
  KEY `service_id` (`service_id`),
  KEY `cross_service_id` (`cross_service_id`),
  CONSTRAINT `service_base_ibfk_1` FOREIGN KEY (`deploy_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `service_base_ibfk_2` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `service_base_ibfk_3` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`),
  CONSTRAINT `service_base_ibfk_4` FOREIGN KEY (`server_id`) REFERENCES `server` (`id`),
  CONSTRAINT `service_base_ibfk_5` FOREIGN KEY (`service_id`) REFERENCES `service` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13878 DEFAULT CHARSET=utf8mb4
******sql******/
// ServiceBase [...]
type ServiceBase struct {
	ID              int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	ServerID        int       `gorm:"index:server_id;column:server_id;type:int(11)" json:"server_id"`
	Server          Server    `gorm:"joinForeignKey:server_id;foreignKey:id" json:"server"`
	GameID          int       `gorm:"index:game_id;column:game_id;type:int(11)" json:"game_id"`
	Game            Game      `gorm:"joinForeignKey:game_id;foreignKey:id" json:"game"`
	PlatformID      int       `gorm:"index:platform_id;column:platform_id;type:int(11)" json:"platform_id"`
	Platform        Platform  `gorm:"joinForeignKey:platform_id;foreignKey:id" json:"platform"`
	DeployUserID    int       `gorm:"index:deploy_user_id;column:deploy_user_id;type:int(11)" json:"deploy_user_id"`
	User            User      `gorm:"joinForeignKey:deploy_user_id;foreignKey:id" json:"user"`
	ServiceID       int       `gorm:"index:service_id;column:service_id;type:int(11)" json:"service_id"`
	Service         Service   `gorm:"joinForeignKey:service_id;foreignKey:id" json:"service"`
	ConfigKwargs    string    `gorm:"column:config_kwargs;type:text" json:"config_kwargs"`
	DeployTime      time.Time `gorm:"column:deploy_time;type:datetime" json:"deploy_time"`
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	Number          int64     `gorm:"column:number;type:int(11)" json:"number"`
	OpenTime        time.Time `gorm:"column:open_time;type:datetime" json:"open_time"`
	MergeTime       time.Time `gorm:"column:merge_time;type:datetime" json:"merge_time"`
	CrossServiceID  int       `gorm:"index:cross_service_id;column:cross_service_id;type:int(11)" json:"cross_service_id"`
	MergedServices  string    `gorm:"column:merged_services;type:text" json:"merged_services"`
	Version         string    `gorm:"column:version;type:varchar(64)" json:"version"`
	Status          string    `gorm:"column:status;type:varchar(255)" json:"status"`
	NoBackupWarning bool      `gorm:"column:no_backup_warning;type:tinyint(1)" json:"no_backup_warning"`
	Online          int       `gorm:"column:online;type:int(11)" json:"online"`
}

func (s *ServiceBase) PortID() string {
	var cfg map[string]interface{}
	_ = json.Unmarshal([]byte(s.ConfigKwargs), &cfg)
	return cfg["port_ID"].(string)
}

func (s *ServiceBase) Port() int64 {
	p := s.PortID()
	port, _ := strconv.ParseInt(p, 10, 64)
	return port + 5000
}

/******sql******
CREATE TABLE `service_handle` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `execute_type` varchar(32) DEFAULT NULL,
  `command` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`),
  KEY `service_id` (`service_id`),
  CONSTRAINT `service_handle_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`),
  CONSTRAINT `service_handle_ibfk_2` FOREIGN KEY (`service_id`) REFERENCES `service` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4
******sql******/
// ServiceHandle [...]
type ServiceHandle struct {
	ID          int     `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID      int     `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game        Game    `gorm:"joinForeignKey:game_id;foreignKey:id"`
	ServiceID   int     `gorm:"index:service_id;column:service_id;type:int(11)"`
	Service     Service `gorm:"joinForeignKey:service_id;foreignKey:id"`
	Type        string  `gorm:"column:type;type:varchar(255)"`
	ExecuteType string  `gorm:"column:execute_type;type:varchar(32)"`
	Command     string  `gorm:"column:command;type:varchar(255)"`
}

/******sql******
CREATE TABLE `update_release` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `restart_service_ids` varchar(255) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `file_name_list` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `publish_status` tinyint(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `service_ids` varchar(255) DEFAULT NULL,
  `user_name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_id` (`game_id`),
  CONSTRAINT `update_release_ibfk_1` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2660 DEFAULT CHARSET=utf8mb4
******sql******/
// UpdateRelease [...]
type UpdateRelease struct {
	ID                int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID            int       `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game              Game      `gorm:"joinForeignKey:game_id;foreignKey:id"`
	RestartServiceIDs string    `gorm:"column:restart_service_ids;type:varchar(255)"`
	FileName          string    `gorm:"column:file_name;type:varchar(255)"`
	FileNameList      string    `gorm:"column:file_name_list;type:varchar(255)"`
	Version           string    `gorm:"column:version;type:varchar(255)"`
	PublishStatus     bool      `gorm:"column:publish_status;type:tinyint(1)"`
	CreateTime        time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime        time.Time `gorm:"column:update_time;type:datetime"`
	UserID            int       `gorm:"column:user_id;type:int(11)"`
	ServiceIDs        string    `gorm:"column:service_ids;type:varchar(255)"`
	UserName          string    `gorm:"column:user_name;type:varchar(64)"`
}

/******sql******
CREATE TABLE `updaterelease_platform` (
  `update_release_id` int(11) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  UNIQUE KEY `UC_update_release_id_platform_id` (`update_release_id`,`platform_id`),
  KEY `platform_id` (`platform_id`),
  CONSTRAINT `updaterelease_platform_ibfk_1` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`),
  CONSTRAINT `updaterelease_platform_ibfk_2` FOREIGN KEY (`update_release_id`) REFERENCES `update_release` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// UpdatereleasePlatform [...]
type UpdatereleasePlatform struct {
	UpdateReleaseID int           `gorm:"uniqueIndex:UC_update_release_id_platform_id;column:update_release_id;type:int(11)"`
	UpdateRelease   UpdateRelease `gorm:"joinForeignKey:update_release_id;foreignKey:id"`
	PlatformID      int           `gorm:"uniqueIndex:UC_update_release_id_platform_id;index:platform_id;column:platform_id;type:int(11)"`
	Platform        Platform      `gorm:"joinForeignKey:platform_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login` varchar(20) DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  `name` varchar(20) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  `qq` int(11) DEFAULT NULL,
  `weixin_id` varchar(255) DEFAULT NULL,
  `last_login` datetime DEFAULT NULL,
  `last_online` datetime DEFAULT NULL,
  `department_id` int(11) DEFAULT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `is_super` tinyint(1) DEFAULT NULL,
  `can_terminal` tinyint(1) DEFAULT NULL,
  `public_key` text,
  `session` text,
  `game_id` int(11) DEFAULT NULL,
  `ding_id` varchar(255) DEFAULT NULL,
  `on_duty` tinyint(1) DEFAULT NULL,
  `on_duty_change_time` datetime DEFAULT NULL,
  `work_state` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ix_user_login` (`login`),
  KEY `department_id` (`department_id`),
  KEY `game_id` (`game_id`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`department_id`) REFERENCES `department` (`id`),
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`game_id`) REFERENCES `game` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4
******sql******/
// User [...]
type User struct {
	ID               int        `gorm:"primaryKey;column:id;type:int(11);not null"`
	Login            string     `gorm:"unique;column:login;type:varchar(20)"`
	Password         string     `gorm:"column:password;type:varchar(100);not null"`
	Name             string     `gorm:"column:name;type:varchar(20)"`
	Phone            string     `gorm:"column:phone;type:varchar(255)"`
	Email            string     `gorm:"column:email;type:varchar(30)"`
	Qq               int        `gorm:"column:qq;type:int(11)"`
	WeixinID         string     `gorm:"column:weixin_id;type:varchar(255)"`
	LastLogin        time.Time  `gorm:"column:last_login;type:datetime"`
	LastOnline       time.Time  `gorm:"column:last_online;type:datetime"`
	DepartmentID     int        `gorm:"index:department_id;column:department_id;type:int(11)"`
	Department       Department `gorm:"joinForeignKey:department_id;foreignKey:id"`
	Active           bool       `gorm:"column:active;type:tinyint(1)"`
	IsSuper          bool       `gorm:"column:is_super;type:tinyint(1)"`
	CanTerminal      bool       `gorm:"column:can_terminal;type:tinyint(1)"`
	PublicKey        string     `gorm:"column:public_key;type:text"`
	Session          string     `gorm:"column:session;type:text"`
	GameID           int        `gorm:"index:game_id;column:game_id;type:int(11)"`
	Game             Game       `gorm:"joinForeignKey:game_id;foreignKey:id"`
	DingID           string     `gorm:"column:ding_id;type:varchar(255)"`
	OnDuty           bool       `gorm:"column:on_duty;type:tinyint(1)"`
	OnDutyChangeTime time.Time  `gorm:"column:on_duty_change_time;type:datetime"`
	WorkState        bool       `gorm:"column:work_state;type:tinyint(1)"`
}

/******sql******
CREATE TABLE `user_domain` (
  `user_id` int(11) DEFAULT NULL,
  `domain_id` int(11) DEFAULT NULL,
  UNIQUE KEY `user_id` (`user_id`,`domain_id`),
  KEY `domain_id` (`domain_id`),
  CONSTRAINT `user_domain_ibfk_1` FOREIGN KEY (`domain_id`) REFERENCES `domain` (`id`),
  CONSTRAINT `user_domain_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// UserDomain [...]
type UserDomain struct {
	UserID   int    `gorm:"uniqueIndex:user_id;column:user_id;type:int(11)"`
	User     User   `gorm:"joinForeignKey:user_id;foreignKey:id"`
	DomainID int    `gorm:"uniqueIndex:user_id;index:domain_id;column:domain_id;type:int(11)"`
	Domain   Domain `gorm:"joinForeignKey:domain_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `user_platform` (
  `user_id` int(11) DEFAULT NULL,
  `platform_id` int(11) DEFAULT NULL,
  UNIQUE KEY `user_id` (`user_id`,`platform_id`),
  KEY `platform_id` (`platform_id`),
  CONSTRAINT `user_platform_ibfk_1` FOREIGN KEY (`platform_id`) REFERENCES `platform` (`id`),
  CONSTRAINT `user_platform_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// UserPlatform [...]
type UserPlatform struct {
	UserID     int      `gorm:"uniqueIndex:user_id;column:user_id;type:int(11)"`
	User       User     `gorm:"joinForeignKey:user_id;foreignKey:id"`
	PlatformID int      `gorm:"uniqueIndex:user_id;index:platform_id;column:platform_id;type:int(11)"`
	Platform   Platform `gorm:"joinForeignKey:platform_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `user_role` (
  `user_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  UNIQUE KEY `user_id` (`user_id`,`role_id`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `user_role_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `user_role_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
******sql******/
// UserRole [...]
type UserRole struct {
	UserID int  `gorm:"uniqueIndex:user_id;column:user_id;type:int(11)"`
	User   User `gorm:"joinForeignKey:user_id;foreignKey:id"`
	RoleID int  `gorm:"uniqueIndex:user_id;index:role_id;column:role_id;type:int(11)"`
	Role   Role `gorm:"joinForeignKey:role_id;foreignKey:id"`
}

/******sql******
CREATE TABLE `user_work_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `log_type` smallint(6) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4
******sql******/
// UserWorkLog [...]
type UserWorkLog struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	UserID     int       `gorm:"column:user_id;type:int(11)"`
	LogType    int16     `gorm:"column:log_type;type:smallint(6)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `web_payment_import_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) DEFAULT NULL,
  `record_month` varchar(32) DEFAULT NULL,
  `xiyou_payment` decimal(10,2) DEFAULT NULL,
  `tencent_payment` decimal(10,2) DEFAULT NULL,
  `lianyun_payment` decimal(10,2) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=581 DEFAULT CHARSET=utf8mb4
******sql******/
// WebPaymentImportLog [...]
type WebPaymentImportLog struct {
	ID             int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	GameID         int       `gorm:"column:game_id;type:int(11)"`
	RecordMonth    string    `gorm:"column:record_month;type:varchar(32)"`
	XiyouPayment   float64   `gorm:"column:xiyou_payment;type:decimal(10,2)"`
	TencentPayment float64   `gorm:"column:tencent_payment;type:decimal(10,2)"`
	LianyunPayment float64   `gorm:"column:lianyun_payment;type:decimal(10,2)"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `work_sheet` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source_type` varchar(32) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `img_url` text,
  `info` text,
  `warning_level` smallint(6) DEFAULT NULL,
  `working_user` smallint(6) DEFAULT NULL,
  `solution` text,
  `is_automate` smallint(6) DEFAULT NULL,
  `event_type` varchar(32) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `other_info` text,
  `submit_user` varchar(32) DEFAULT NULL,
  `event_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20030 DEFAULT CHARSET=utf8mb4
******sql******/
// WorkSheet [...]
type WorkSheet struct {
	ID           int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	SourceType   string    `gorm:"column:source_type;type:varchar(32)"`
	Title        string    `gorm:"column:title;type:varchar(255)"`
	ImgURL       string    `gorm:"column:img_url;type:text"`
	Info         string    `gorm:"column:info;type:text"`
	WarningLevel int16     `gorm:"column:warning_level;type:smallint(6)"`
	WorkingUser  int16     `gorm:"column:working_user;type:smallint(6)"`
	Solution     string    `gorm:"column:solution;type:text"`
	IsAutomate   int16     `gorm:"column:is_automate;type:smallint(6)"`
	EventType    string    `gorm:"column:event_type;type:varchar(32)"`
	Status       int16     `gorm:"column:status;type:smallint(6)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime"`
	OtherInfo    string    `gorm:"column:other_info;type:text"`
	SubmitUser   string    `gorm:"column:submit_user;type:varchar(32)"`
	EventID      int       `gorm:"column:event_id;type:int(11)"`
}

/******sql******
CREATE TABLE `work_sheet_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sheet_id` int(11) DEFAULT NULL,
  `info` text,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sheet_id` (`sheet_id`),
  CONSTRAINT `work_sheet_log_ibfk_1` FOREIGN KEY (`sheet_id`) REFERENCES `work_sheet` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11201 DEFAULT CHARSET=utf8mb4
******sql******/
// WorkSheetLog [...]
type WorkSheetLog struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	SheetID    int       `gorm:"index:sheet_id;column:sheet_id;type:int(11)"`
	WorkSheet  WorkSheet `gorm:"joinForeignKey:sheet_id;foreignKey:id"`
	Info       string    `gorm:"column:info;type:text"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

/******sql******
CREATE TABLE `work_sheet_task_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4
******sql******/
// WorkSheetTaskType [...]
type WorkSheetTaskType struct {
	ID         int       `gorm:"primaryKey;column:id;type:int(11);not null"`
	Name       string    `gorm:"column:name;type:varchar(255)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}
