package web

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.kube.2xi.com/golden/ynds/models"
	"gitlab.kube.2xi.com/golden/ynds/utils"
	"strconv"
	"time"
)

type Handlers struct {
}

func (h *Handlers) ServiceList(c *fiber.Ctx) error {
	_sort := c.Query("sort", "number")
	order := c.Query("order", "asc")
	number := c.Query("number")
	version := c.Query("version")
	status := c.Query("status")
	platformID := c.Query("platform")
	ip := c.Query("ip")
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	where := utils.DB.Where("service_id = ?", 1)
	if ip != "" {
		var ss []struct {
			ID int
		}
		utils.DB.Model(models.Server{}).Where("ip like ?", ip+"%").Find(&ss)
		var serverIDs []int
		for _, s := range ss {
			serverIDs = append(serverIDs, s.ID)
		}
		where = where.Where("server_id in ?", serverIDs)
	}
	if number != "" {
		where = where.Where("number = ?", number)
	}
	if version != "" {
		where = where.Where("version like ?", "%"+version+"%")
	}
	if status != "" {
		where = where.Where("status = ?", status)
	}
	if platformID != "" {
		where = where.Where("platform_id = ?", platformID)
	}
	var count int64
	models.ServiceBaseMgr(where).Count(&count)
	tx := where.Offset(offset)
	if c.Query("limit") != "" {
		limit, _ := strconv.Atoi(c.Query("limit", "20"))
		tx = where.Limit(limit)
	}
	if _sort != "" {
		tx = tx.Order(_sort + " " + order)
	}
	serviceBaseMgr := models.ServiceBaseMgr(tx)
	gameServices, err := serviceBaseMgr.Gets()
	if err != nil {
		return err
	}
	var total int64
	models.ServiceBaseMgr(utils.DB.Where("service_id = ?", 1)).Count(&total)
	return c.JSON(fiber.Map{
		"total":            count,
		"totalNotFiltered": total - count,
		"rows":             gameServices,
	})
}

func (h *Handlers) ServerList(c *fiber.Ctx) error {
	ip := c.Query("ip")
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	tx := utils.DB.Offset(offset)
	if c.Query("limit") != "" {
		limit, _ := strconv.Atoi(c.Query("limit"))
		tx = tx.Limit(limit)
	}
	query := map[string]interface{}{}
	if ip != "" {
		query["ip"] = ip
	}
	tx = tx.Where(query)
	serverMgr := models.ServerMgr(tx)
	servers, err := serverMgr.Gets()
	if err != nil {
		return err
	}
	var total int64
	models.ServerMgr(utils.DB).Count(&total)
	var count int64
	models.ServerMgr(utils.DB.Where(query)).Count(&count)
	type newServers struct {
		ID           int       `json:"id"`
		IP           string    `json:"ip"`
		Lip          string    `json:"lip"`
		SaltKey      string    `json:"salt_key"`
		ServiceCount int64     `json:"service_count"`
		CreateTime   time.Time `json:"create_time"`
	}
	var s []newServers
	for _, server := range servers {
		var c int64
		models.ServiceBaseMgr(utils.DB.Where("server_id = ?", server.ID)).Count(&c)
		s = append(s, newServers{
			ID:           server.ID,
			IP:           server.IP,
			Lip:          server.Lip,
			SaltKey:      server.SaltKey,
			ServiceCount: c,
			CreateTime:   server.CreateTime,
		})
	}
	return c.JSON(fiber.Map{
		"total":            total,
		"totalNotFiltered": total - count,
		"rows":             s,
	})
}

func (h *Handlers) UpdateService(c *fiber.Ctx) error {
	saltKey := c.Query("salt_key")
	platformEnName := c.Query("platform_en_name")
	number := c.Query("number")
	gameEnName := c.Query("game_en_name")
	status := c.Query("status")
	version := c.Query("version")
	online := c.Query("online")
	openTime := c.Query("opentime")
	server, err := models.ServerMgr(utils.DB).GetFromSaltKey(saltKey)
	if err != nil {
		return c.JSON(fiber.Map{"status": 10000, "msg": "salt key  error"})
	}
	game, err := models.GameMgr(utils.DB).GetFromEnName(gameEnName)
	if err != nil || len(game) != 1 {
		return c.JSON(fiber.Map{"status": 10000, "msg": "game_en_name error"})
	}
	platform, err := models.PlatformMgr(utils.DB).GetFromEnName(platformEnName)
	if err != nil {
		return c.JSON(fiber.Map{"status": 10000, "msg": "platform_en_name error"})
	}
	gameServiceMgr := models.ServiceBaseMgr(utils.DB)
	n, _ := strconv.ParseInt(number, 10, 64)
	gameServices, err := gameServiceMgr.GetByOptions(
		gameServiceMgr.WithServiceID(1),
		gameServiceMgr.WithPlatformID(platform.ID),
		gameServiceMgr.WithGameID(game[0].ID),
		gameServiceMgr.WithServerID(server.ID),
		gameServiceMgr.WithNumber(n),
	)
	if err != nil || len(gameServices) != 1 {
		return c.JSON(fiber.Map{"status": 10000, "msg": "游戏服不存在，检查number"})
	}
	gameService := gameServices[0]
	gameService.Status = status
	gameService.Version = version
	i, _ := strconv.Atoi(online)
	gameService.Online = i
	var op time.Time
	if openTime != "" {
		if tm, err := strconv.ParseInt(openTime, 10, 64); err == nil {
			op = time.Unix(tm, 0)
		}
	}
	err = utils.DB.Model(models.ServiceBase{}).Where("id = ?", gameService.ID).Updates(map[string]interface{}{"status": status,
		"online": online, "Version": version,
		"open_time": op}).Error
	return err
}

func (h *Handlers) PlatformList(c *fiber.Ctx) error {
	platformMgr := models.PlatformMgr(utils.DB)
	platforms, _ := platformMgr.Gets()
	return c.JSON(platforms)
}
