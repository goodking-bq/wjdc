package web

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/goodking-bq/wjdc/models"
	"github.com/goodking-bq/wjdc/utils"
	"strconv"
	"strings"
)

type Handlers struct {
}

func (h *Handlers) Index(ctx *fiber.Ctx) error {
	q, _ := strconv.ParseUint(ctx.Query("q", "1"), 0, 64)
	questionnaire := models.Questionnaire{}
	utils.DB.Model(models.Questionnaire{}).Preload("Questions").Preload("Questions.Options").First(&questionnaire, q)
	return ctx.Render("ui/index", fiber.Map{"Q": questionnaire})
}

func (h *Handlers) SaveAnswer(ctx *fiber.Ctx) error {
	q, _ := strconv.ParseUint(ctx.Query("q", "1"), 0, 64)
	answer := models.Answer{QuestionnaireID: uint(q)}
	utils.DB.Save(&answer)
	if form, err := ctx.MultipartForm(); err == nil {

		fmt.Printf("%v", form)
		for k, v := range form.Value {
			if strings.Contains(k, "use") {
				continue
			}
			qid, _ := strconv.ParseUint(strings.Split(k, "_")[1], 0, 64)
			spend, _ := strconv.ParseUint(form.Value[k+"_use"][0], 0, 64)
			detail := &models.AnswerDetail{
				AnswerID:   answer.ID,
				QuestionID: uint(qid),
				Reply:      "",
				Spend:      uint(spend),
			}
			var question models.Question
			utils.DB.Model(models.Question{}).First(&question, qid)
			switch question.Type {
			case models.QuestionTypeText:
				detail.Reply = v[0]
			case models.QuestionTypeCheck, models.QuestionTypeRadio, models.QuestionTypeSelect:
				{
					var opts []*models.QuestionOption
					utils.DB.Model(models.QuestionOption{}).Where("id in ?", v).Find(&opts)
					var s []string
					for _, opt := range opts {
						s = append(s, opt.Content)
					}
					detail.Reply = strings.Join(s, ",")
				}
			}
			utils.DB.Save(detail)
		}
		fmt.Printf("%v", form.Value)
		return ctx.Render("ui/index", fiber.Map{"Message": "提交成功"})
	} else {
		return err
	}
}

func (h *Handlers) Csv(ctx *fiber.Ctx) error {
	bytesBuffer := &bytes.Buffer{}
	bytesBuffer.WriteString("xEFxBBxBF") // 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	writer := csv.NewWriter(bytesBuffer)
	q, _ := strconv.ParseUint(ctx.Query("q", "1"), 0, 64)
	var questions []*models.Question
	utils.DB.Model(models.Question{}).Where("questionnaire_id = ?", q).Order("id asc").Find(&questions)
	var header []string
	var order []uint
	for _, question := range questions {
		header = append(header, question.Content)
		order = append(order, question.ID)
	}
	_ = writer.Write(header)
	var answers []*models.Answer
	utils.DB.Model(models.Answer{}).Where("questionnaire_id = ?", q).Find(&answers)
	for _, answer := range answers {
		var tmp []string
		var results []map[string]interface{}
		utils.DB.Model(models.AnswerDetail{}).Where("answer_id = ?", answer.ID).Find(&results)
		for _, qid := range order {
			for _, d := range results {
				if d["question_id"].(uint) == qid {
					tmp = append(tmp, d["reply"].(string))
				}
			}
		}
		if len(tmp) > 0 {
			_ = writer.Write(tmp)
		}

	}
	writer.Flush()
	ctx.Append("Content-Disposition", "attachment;filename=wjdc.csv")
	ctx.Accepts("text/csv")
	_, err := ctx.Write(bytesBuffer.Bytes())
	return err
}
