package service

import (
	"context"
	"log"
	"strings"
	"vue-golang/internal/storage"
	"vue-golang/internal/storage/mysql"
)

type NormService struct {
	storage *mysql.Storage
}

func NewNormService(storage *mysql.Storage) *NormService {
	return &NormService{storage: storage}
}

func (s *NormService) CalculateNorm(ctx context.Context, orderID, pos int, typeIzd string, templateCode string) ([]storage.Operation, error) {
	// 1. Получаем материалы
	materials, err := s.storage.GetOrderMaterials(ctx, orderID, pos)
	if err != nil {
		return nil, err
	}

	//TODO кастыль
	//typeIzd := "glyhar"

	// 2. Строим контекст — ВОТ ОНО МЕСТО!
	ctxData := BuildContext(materials, typeIzd)

	// 3. Определяем код шаблона (пример: жёстко задан или по логике)
	//templateCode := "55" // ← позже сделаем умнее

	// 4. Получаем шаблон (операции + правила)
	template, err := s.storage.GetTemplateByCode(ctx, templateCode)
	if err != nil {
		return nil, err
	}

	// 5. Применяем правила
	result := ApplyRules(template.Operations, template.Rules, ctxData)

	return result, nil
}

type Context struct {
	Type string

	HasImpost   bool
	ImpostCount int
	// Добавишь больше признаков позже: тип профиля, площадь, кол-во камер и т.д.
}

func BuildContextGlyhar(materials []*storage.KlaesMaterials) Context {
	ctx := Context{Type: "glyhar"}

	for _, m := range materials {
		name := strings.TrimSpace(strings.ToLower(m.NameMat))
		//art := strings.ToLower(m.ArticulMat)

		//if strings.Contains(name, "импост") {
		//	ctx.HasImpost = true
		//	ctx.ImpostCount++
		//}

		if name == "импост" || name == "impost" || name == "доп. импост" {
			ctx.HasImpost = true
			ctx.ImpostCount++
		}
	}

	log.Printf("Смотрим материалы: HasImpost=%v, ImpostCount=%d", ctx.HasImpost, ctx.ImpostCount)

	return ctx
}

func BuildContextWindow(materials []*storage.KlaesMaterials) Context {
	ctx := Context{Type: "window"}

	for _, m := range materials {
		name := strings.ToLower(m.NameMat)
		//art := strings.ToLower(m.ArticulMat)

		if strings.Contains(name, "рама") {
			ctx.HasImpost = true
			ctx.ImpostCount++
		}
	}

	log.Printf("Смотрим материалы: HasImpost=%v, ImpostCount=%d", ctx.HasImpost, ctx.ImpostCount)

	return ctx
}

func BuildContextDoor(materials []*storage.KlaesMaterials) Context {
	ctx := Context{Type: "door"}

	for _, m := range materials {
		name := strings.TrimSpace(strings.ToLower(m.NameMat))
		//art := strings.ToLower(m.ArticulMat)

		//if strings.Contains(name, "импост") {
		//	ctx.HasImpost = true
		//	ctx.ImpostCount++
		//}

		if name == "импост" || name == "impost" || name == "доп. импост" {
			ctx.HasImpost = true
			ctx.ImpostCount++
		}
	}

	log.Printf("Смотрим материалы: HasImpost=%v, ImpostCount=%d", ctx.HasImpost, ctx.ImpostCount)

	return ctx
}

func BuildContext(materials []*storage.KlaesMaterials, typeIzd string) Context {
	switch typeIzd {
	case "glyhar":
		return BuildContextGlyhar(materials)
	case "window":
		return BuildContextWindow(materials)
	case "door":
		return BuildContextDoor(materials)
	default:
		return Context{Type: "неизвестный тип изделия"}
	}

}

func ApplyRules(operations []storage.Operation, rules []storage.Rule, ctx Context) []storage.Operation {
	result := make([]storage.Operation, len(operations))
	copy(result, operations)

	log.Printf("Загружено правил: %d", len(rules))
	for i, r := range rules {
		log.Printf("Правило %d: op=%s, cond=%v", i, r.Operation, r.Condition)
	}

	for i := range result {
		for _, rule := range rules {
			if rule.Operation != result[i].Name {
				continue
			}

			if !MatchesCondition(rule.Condition, ctx) {
				continue
			}

			switch rule.Mode {
			case "set":
				result[i].Value = rule.SetValue
				result[i].Minutes = rule.SetMinutes
			case "multiplied":
				result[i].Value = rule.ValuePerUnit * float64(ctx.ImpostCount)
				result[i].Minutes = rule.MinutesPerUnit * float64(ctx.ImpostCount)
				result[i].Count = float64(ctx.ImpostCount)
			case "additive":
				result[i].Value += rule.ValuePerUnit * float64(ctx.ImpostCount)
				result[i].Minutes += rule.MinutesPerUnit * float64(ctx.ImpostCount)
			default:
				// По умолчанию — просто замена
				result[i].Value = rule.SetValue
				result[i].Minutes = rule.SetMinutes
			}
			break // применили первое подходящее правило
		}
	}

	return result
}

func MatchesCondition(condition map[string]interface{}, ctx Context) bool {
	for key, expected := range condition {
		if !fieldMatches(key, expected, ctx) {
			return false
		}
	}
	return true
}

func fieldMatches(key string, expected interface{}, ctx Context) bool {
	switch key {
	case "HasImpost":
		if val, ok := expected.(bool); ok {
			return ctx.HasImpost == val
		}
	case "ImpostCount":
		// Поддержка: {"min": 2} или просто 2
		return compareIntField(ctx.ImpostCount, expected)
	default:
		return false
	}
	return false
}

func compareIntField(actual int, expected interface{}) bool {
	// Вариант 1: просто число → точное совпадение
	if val, ok := expected.(float64); ok { // JSON числа — float64
		return actual == int(val)
	}

	// Вариант 2: объект { "min": 2 }
	if obj, ok := expected.(map[string]interface{}); ok {
		if minVal, hasMin := obj["min"].(float64); hasMin {
			if actual < int(minVal) {
				return false
			}
		}
		if maxVal, hasMax := obj["max"].(float64); hasMax {
			if actual > int(maxVal) {
				return false
			}
		}
		return true
	}

	return false
}
