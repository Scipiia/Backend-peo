package utils

import (
	"reflect"
	"vue-golang/internal/storage"
)

func MapToOrderDataGlyhari(result *storage.DemResultGlyhari) *storage.OrderData {
	operations := make([]storage.Operation, 0)

	// Добавляем только те, где значение > 0
	if result.NastrNapil > 0 {
		operations = append(operations, storage.Operation{"nast_napil", "Настройка оборудования для напиловки", result.NastrNapil})
	}
	if result.NapilKontr > 0 {
		operations = append(operations, storage.Operation{"napil_kontr", "Напиловка контура", result.NapilKontr})
	}
	if result.NapilKrishek > 0 {
		operations = append(operations, storage.Operation{"napil_krishek", "Напиловка крышек", result.NapilKrishek})
	}
	if result.NapilImpost > 0 {
		operations = append(operations, storage.Operation{"napil_impost", "Напиловка импоста", result.NapilImpost})
	}
	if result.Soedinitel > 0 {
		operations = append(operations, storage.Operation{"soedinitel", "Соединитель", result.Soedinitel})
	}
	if result.PromejSborka > 0 {
		operations = append(operations, storage.Operation{"promej_sborka", "Промежуточная сборка", result.PromejSborka})
	}
	if result.ImpostSverlovka > 0 {
		operations = append(operations, storage.Operation{"impost_sverlovka", "Импост (сверловка)", result.ImpostSverlovka})
	}
	if result.ImpostFrezerovka > 0 {
		operations = append(operations, storage.Operation{"impost_frezerovka", "Импост (фрезеровка)", result.ImpostFrezerovka})
	}
	if result.ImpostSborka > 0 {
		operations = append(operations, storage.Operation{"impost_sborka", "Импост (сборка)", result.ImpostSborka})
	}
	if result.OpresNastr > 0 {
		operations = append(operations, storage.Operation{"opres_nastr", "Опресовка (настройка)", result.OpresNastr})
	}
	if result.Opresovka > 0 {
		operations = append(operations, storage.Operation{"opresovka", "Опресовка", result.Opresovka})
	}
	if result.YstanYplotnitel > 0 {
		operations = append(operations, storage.Operation{"ystan_yplotnitel", "Установка уплотнителей, штифтовка", result.YstanYplotnitel})
	}
	if result.Zashivka > 0 {
		operations = append(operations, storage.Operation{"zashivka", "Зашивка", result.Zashivka})
	}
	if result.NapilStoikiDo3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_do3m", "Напиловка стойки до 3м", result.NapilStoikiDo3m})
	}
	if result.NapilStoikiBol3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_bol3m", "Напиловка стойки больше 3м", result.NapilStoikiBol3m})
	}
	if result.NapilRigelDo1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_do1m", "Напиловка ригеля до 1м", result.NapilRigelDo1m})
	}
	if result.NapilRigelBol1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_bol1m", "Напиловка ригеля больше 1м", result.NapilRigelBol1m})
	}
	if result.SverlRigelZamok > 0 {
		operations = append(operations, storage.Operation{"sverl_rigel_zamok", "Сверловка ригеля под замок", result.SverlRigelZamok})
	}
	if result.YstanZamkov > 0 {
		operations = append(operations, storage.Operation{"ystan_zamkov", "Установка замков", result.YstanZamkov})
	}
	if result.NapilShtapik > 0 {
		operations = append(operations, storage.Operation{"napil_shtapik", "Напиловка штапика", result.NapilShtapik})
	}
	if result.Ypakovka > 0 {
		operations = append(operations, storage.Operation{"ypakovka", "Упаковка", result.Ypakovka})
	}
	if result.FrezerRigel > 0 {
		operations = append(operations, storage.Operation{"frezer_rigel", "Фрезеровка ригеля", result.FrezerRigel})
	}
	if result.ObrabotRam > 0 {
		operations = append(operations, storage.Operation{"obrabot_ram", "Обработка рам", result.ObrabotRam})
	}
	if result.HandsSborka > 0 {
		operations = append(operations, storage.Operation{"hands_sborka", "Ручная сборка", result.HandsSborka})
	}
	if result.FrezerNastr > 0 {
		operations = append(operations, storage.Operation{"frezer_nastr", "Фрезеровка (настройка)", result.FrezerNastr})
	}
	if result.Shtiftovka > 0 {
		operations = append(operations, storage.Operation{"shtiftovka", "Штифтовка", result.Shtiftovka})
	}
	if result.YstanovkaZapoln > 0 {
		operations = append(operations, storage.Operation{"ystanovka_zapoln", "Установка заполнения", result.YstanovkaZapoln})
	}
	if result.NapilDonnik > 0 {
		operations = append(operations, storage.Operation{"napil_donnik", "Напиловка донника", result.NapilDonnik})
	}
	if result.AdapterNapil > 0 {
		operations = append(operations, storage.Operation{"adapter_napil", "Напиловка адаптера", result.AdapterNapil})
	}
	if result.AdapterYstan > 0 {
		operations = append(operations, storage.Operation{"adapter_ystan", "Установка адаптера", result.AdapterYstan})
	}
	if result.YstanYplotnFalc > 0 {
		operations = append(operations, storage.Operation{"ystan_yplotn_falc", "Установка уплотнителей фальца", result.YstanYplotnFalc})
	}
	if result.TotalTime > 0 {
		operations = append(operations, storage.Operation{"total_time", "Итого", result.TotalTime})
	}

	return &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Operations: operations,
		Profil:     result.Profil,
	}
}

func MapToOrderDataWindow(result *storage.DemResultWindow) *storage.OrderData {
	operations := make([]storage.Operation, 0)

	// Добавляем только те, где значение > 0
	if result.PodgotOboryd > 0 {
		operations = append(operations, storage.Operation{"podgot_oboryd", "Подготовка оборудования", result.PodgotOboryd})
	}
	if result.NapilRamStv > 0 {
		operations = append(operations, storage.Operation{"napil_ram_stv", "Напиловка рам и створок", result.NapilRamStv})
	}
	if result.NapilNramStv > 0 {
		operations = append(operations, storage.Operation{"napil_nram_stv", "Напиловка рам и створок нестандарт", result.NapilNramStv})
	}
	if result.NapilImposta > 0 {
		operations = append(operations, storage.Operation{"napil_imposta", "Напиловка импоста", result.NapilImposta})
	}
	if result.SverlovkaStoek > 0 {
		operations = append(operations, storage.Operation{"soedinitel", "Сверловка стоек", result.SverlovkaStoek})
	}
	if result.ObrabotRam > 0 {
		operations = append(operations, storage.Operation{"obrabot_ram", "Обработка рам", result.ObrabotRam})
	}
	if result.ObrabotkaStv > 0 {
		operations = append(operations, storage.Operation{"obrabotka_stv", "Обработка створки", result.ObrabotkaStv})
	}
	if result.PromejSborkaStv > 0 {
		operations = append(operations, storage.Operation{"promej_sborka_stv", "Промежуточная сборка створки", result.PromejSborkaStv})
	}
	if result.NapilAdaptera > 0 {
		operations = append(operations, storage.Operation{"napil_adaptera", "Напиловка адаптера", result.NapilAdaptera})
	}
	if result.PromejSborkaRam > 0 {
		operations = append(operations, storage.Operation{"promej_sborka_ram", "Промежуточная сборка рам", result.PromejSborkaRam})
	}
	if result.PromejSborkaGl > 0 {
		operations = append(operations, storage.Operation{"promej_sborka_gl", "Промежуточная сборка глухарей", result.PromejSborkaGl})
	}
	if result.OpresRam > 0 {
		operations = append(operations, storage.Operation{"opres_ram", "Опресовка рам", result.OpresRam})
	}
	if result.OpresGl > 0 {
		operations = append(operations, storage.Operation{"opres_gl", "Опрессовка глухаря", result.OpresGl})
	}
	if result.FrezerStv > 0 {
		operations = append(operations, storage.Operation{"frezer_stv", "Фрезеровка створок", result.FrezerStv})
	}
	if result.YstanYplRam > 0 {
		operations = append(operations, storage.Operation{"ystan_ypl_ram", "Установка уплотнителей рам", result.YstanYplRam})
	}
	if result.YstanYplStv > 0 {
		operations = append(operations, storage.Operation{"ystan_ypl_stv", "Установка уплотнителей створки", result.YstanYplStv})
	}
	if result.NapilTag > 0 {
		operations = append(operations, storage.Operation{"napil_tag", "Напиловка тяг", result.NapilTag})
	}
	if result.SverloTag > 0 {
		operations = append(operations, storage.Operation{"sverlo_tag", "Сверловка тяг", result.SverloTag})
	}
	if result.YstanFurn > 0 {
		operations = append(operations, storage.Operation{"ystan_furn", "Установка фурнитуры", result.YstanFurn})
	}
	if result.NaveshStv > 0 {
		operations = append(operations, storage.Operation{"navesh_stv", "Навешивание створок", result.NaveshStv})
	}
	if result.Zashivka > 0 {
		operations = append(operations, storage.Operation{"zashivka", "Зашивка", result.Zashivka})
	}
	if result.ObrabDopProfil > 0 {
		operations = append(operations, storage.Operation{"obrab_dop_profil", "Обработка доп профиля", result.ObrabDopProfil})
	}
	if result.YstanAdaptera > 0 {
		operations = append(operations, storage.Operation{"ystan_adaptera", "Установка адаптера", result.YstanAdaptera})
	}
	if result.FrezerImpostPilon > 0 {
		operations = append(operations, storage.Operation{"frezer_impost_pilon", "Фрезеровка импоста пилона", result.FrezerImpostPilon})
	}
	if result.KreplYglRam > 0 {
		operations = append(operations, storage.Operation{"krepl_ygl_ram", "Крепление углового соединителя к раме", result.KreplYglRam})
	}
	if result.GlyharDorab > 0 {
		operations = append(operations, storage.Operation{"glyhar_dorab", "Доработка глухаря", result.GlyharDorab})
	}
	if result.YplotGlyhar > 0 {
		operations = append(operations, storage.Operation{"yplot_glyhar", "Установка уплотнителя глухаря", result.YplotGlyhar})
	}
	if result.ZashivGlyhar > 0 {
		operations = append(operations, storage.Operation{"zashiv_glyhar", "Зашивка глухаря", result.ZashivGlyhar})
	}
	if result.OpresStv > 0 {
		operations = append(operations, storage.Operation{"opres_stv", "Опресовка створки", result.OpresStv})
	}
	if result.RazborErkera > 0 {
		operations = append(operations, storage.Operation{"razbor_erkera", "Разборка эркера", result.RazborErkera})
	}
	if result.GlWindow > 0 {
		operations = append(operations, storage.Operation{"gl_window", "Глухое окно", result.GlWindow})
	}
	if result.ObrabotShtylp > 0 {
		operations = append(operations, storage.Operation{"obrabot_shtylp", "Обработка штульпа", result.ObrabotShtylp})
	}
	if result.FrezerPodShtylp > 0 {
		operations = append(operations, storage.Operation{"frezer_pod_shtylp", "Фрезеровка под штульп", result.FrezerPodShtylp})
	}
	if result.YstanShtylp > 0 {
		operations = append(operations, storage.Operation{"ystan_shtylp", "Установка штульпа", result.YstanShtylp})
	}
	if result.YstanImpost > 0 {
		operations = append(operations, storage.Operation{"ystan_impost", "Установка импоста", result.YstanImpost})
	}
	if result.NastrForOpres > 0 {
		operations = append(operations, storage.Operation{"nastr_for_opres", "Настройка станка для опресовки", result.NastrForOpres})
	}
	if result.NapilRam > 0 {
		operations = append(operations, storage.Operation{"napil_ram", "Напиловка рам", result.NapilRam})
	}
	if result.NapilNram > 0 {
		operations = append(operations, storage.Operation{"napil_nram", "Напиловка рам нестандарт", result.NapilNram})
	}
	if result.NapilDopProfil > 0 {
		operations = append(operations, storage.Operation{"napil_dop_profil", "Напиловка доп профилей", result.NapilDopProfil})
	}
	if result.FrezerDopProfil > 0 {
		operations = append(operations, storage.Operation{"frezer_dop_profil", "Фрезеровка доп профилей", result.FrezerDopProfil})
	}
	if result.ObrabotkaStvRychka > 0 {
		operations = append(operations, storage.Operation{"obrabotka_stv_rychka", "Обработка створки под ручку", result.ObrabotkaStvRychka})
	}
	if result.ObrabotkaStvZamok > 0 {
		operations = append(operations, storage.Operation{"obrabotka_stv_zamok", "Обработка створки под замок", result.ObrabotkaStvZamok})
	}
	if result.NapilShtylp > 0 {
		operations = append(operations, storage.Operation{"napil_shtylp", "Напиловка штульпа", result.NapilShtylp})
	}
	if result.ImpostSbor > 0 {
		operations = append(operations, storage.Operation{"impost_sbor", "Импост (сборка)", result.ImpostSbor})
	}
	if result.RezinaRam > 0 {
		operations = append(operations, storage.Operation{"rezina_ram", "Обрезинивание рамы", result.RezinaRam})
	}
	if result.PodgotovFurn > 0 {
		operations = append(operations, storage.Operation{"podgotov_furn", "Подготовка фурнитуры", result.PodgotovFurn})
	}
	if result.PodgotovRam > 0 {
		operations = append(operations, storage.Operation{"podgotov_ram", "Подготовка рам", result.PodgotovRam})
	}
	if result.Razborka > 0 {
		operations = append(operations, storage.Operation{"razborka", "Разборка", result.Razborka})
	}
	if result.NapilZashitProf > 0 {
		operations = append(operations, storage.Operation{"napil_zashit_prof", "Напиловка защитного профиля", result.NapilZashitProf})
	}
	if result.ImpostFrezer > 0 {
		operations = append(operations, storage.Operation{"impost_frezer", "Импост фрезеровка", result.ImpostFrezer})
	}
	if result.YstanSteklaZaliv > 0 {
		operations = append(operations, storage.Operation{"ystan_stekla_zaliv", "Установка стекла и заливка", result.YstanSteklaZaliv})
	}
	if result.YstanKrishek > 0 {
		operations = append(operations, storage.Operation{"ystan_krishek", "Установка крышек", result.YstanKrishek})
	}
	if result.FrezerProfilZamazka > 0 {
		operations = append(operations, storage.Operation{"frezer_profil_zamazka", "Фрезеровка профиля и замазка герметиком", result.FrezerProfilZamazka})
	}
	if result.NapilStoikiDo3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_do3m", "Напиловка стойки до 3м", result.NapilStoikiDo3m})
	}
	if result.NapilStoikiBol3m > 0 {
		operations = append(operations, storage.Operation{"napil_stoiki_bol3m", "Напиловка стойки больше 3м", result.NapilStoikiBol3m})
	}
	if result.NapilRigelDo1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_do1m", "Напиловка ригеля до 1м", result.NapilRigelDo1m})
	}
	if result.NapilRigelBol1m > 0 {
		operations = append(operations, storage.Operation{"napil_rigel_bol1m", "Напиловка ригеля больше 1м", result.NapilRigelBol1m})
	}
	if result.SverloRigelZamok > 0 {
		operations = append(operations, storage.Operation{"sverlo_rigel_zamok", "Сверловка ригеля под замок", result.SverloRigelZamok})
	}
	if result.YstanZamok > 0 {
		operations = append(operations, storage.Operation{"ystan_zamok", "Установка замков", result.YstanZamok})
	}
	if result.Shtiftovka > 0 {
		operations = append(operations, storage.Operation{"shtiftovka", "Штифтовка", result.Shtiftovka})
	}
	if result.FrezerRigel > 0 {
		operations = append(operations, storage.Operation{"frezer_rigel", "Фрезеровка ригеля", result.FrezerRigel})
	}
	if result.PartSborka > 0 {
		operations = append(operations, storage.Operation{"part_sborka", "Частичная сборка", result.PartSborka})
	}
	if result.YstanRezYgl > 0 {
		operations = append(operations, storage.Operation{"ystan_rez_ygl", "Установка резиновых углов", result.YstanRezYgl})
	}
	if result.OpresYgl > 0 {
		operations = append(operations, storage.Operation{"opres_ygl", "Опрессовка углов", result.OpresYgl})
	}
	if result.YstanTermovst > 0 {
		operations = append(operations, storage.Operation{"ystan_termovst", "Установка термовставки", result.YstanTermovst})
	}
	if result.NapilShtapik > 0 {
		operations = append(operations, storage.Operation{"napil_shtapik", "Напиловка штапика", result.NapilShtapik})
	}
	if result.YpakSekcii > 0 {
		operations = append(operations, storage.Operation{"ypak_sekcii", "Упаковка секции", result.YpakSekcii})
	}
	if result.YpakRigel > 0 {
		operations = append(operations, storage.Operation{"ypak_rigel", "Упаковка ригелей", result.YpakRigel})
	}
	if result.YpakYplotn > 0 {
		operations = append(operations, storage.Operation{"ypak_yplotn", "Упаковка уплотнителей", result.YpakYplotn})
	}
	if result.YpakShtapik > 0 {
		operations = append(operations, storage.Operation{"ypak_shtapik", "Упаковка штапика", result.YpakShtapik})
	}
	if result.NapilDonnika > 0 {
		operations = append(operations, storage.Operation{"napil_donnika", "Упаковка штапика", result.NapilDonnika})
	}
	if result.NastrPbx > 0 {
		operations = append(operations, storage.Operation{"nastr_pbx", "Настройка РВХ", result.NastrPbx})
	}
	if result.MehObrabPzr > 0 {
		operations = append(operations, storage.Operation{"meh_obrab_pzr", "Мех обработка(ПЗР)", result.MehObrabPzr})
	}
	if result.RabotaPbx > 0 {
		operations = append(operations, storage.Operation{"rabota_pbx", "Работа станка (PBX)", result.RabotaPbx})
	}
	if result.SlesarObrFurn > 0 {
		operations = append(operations, storage.Operation{"slesar_obr_furn", "Слесарная обработка под установку фурнитуры", result.SlesarObrFurn})
	}
	if result.ImpostSverlo > 0 {
		operations = append(operations, storage.Operation{"impost_sverlo", "Импост (сверловка)", result.ImpostSverlo})
	}
	if result.Opresovka > 0 {
		operations = append(operations, storage.Operation{"opresovka", "Опресовка", result.Opresovka})
	}
	if result.SborRychka > 0 {
		operations = append(operations, storage.Operation{"sbor_rychka", "Сборка ручек", result.SborRychka})
	}
	if result.SborPetli > 0 {
		operations = append(operations, storage.Operation{"sbor_petli", "Сборка петли", result.SborPetli})
	}
	if result.PlastikYstnRam > 0 {
		operations = append(operations, storage.Operation{"plastik_ystn_ram", "Уст. Пластика на раму", result.PlastikYstnRam})
	}
	if result.NapilStv > 0 {
		operations = append(operations, storage.Operation{"napil_stv", "Напиловка створки", result.NapilStv})
	}
	if result.YstanFurnStv > 0 {
		operations = append(operations, storage.Operation{"ystan_furn_stv", "Установка фурнитуры створки", result.YstanFurnStv})
	}
	if result.YstanFurnRam > 0 {
		operations = append(operations, storage.Operation{"ystan_furn_ram", "Установка фурнитуры рам", result.YstanFurnRam})
	}
	if result.ComplOtg > 0 {
		operations = append(operations, storage.Operation{"compl_otg", "Комплектация в отгрузку", result.ComplOtg})
	}
	if result.YpakIzd > 0 {
		operations = append(operations, storage.Operation{"ypak_izd", "Упаковка изделия в сборе", result.YpakIzd})
	}
	if result.Otgryz > 0 {
		operations = append(operations, storage.Operation{"otgryz", "Вынос изделия", result.Otgryz})
	}
	if result.RaspFurn > 0 {
		operations = append(operations, storage.Operation{"rasp_furn", "Распаковка фурнитуры", result.RaspFurn})
	}
	if result.TotalTime > 0 {
		operations = append(operations, storage.Operation{"total_time", "Итого", result.TotalTime})
	}

	return &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Operations: operations,
		Profil:     result.Profil,
	}
}

func MapToOrderDataDoor(result *storage.DemResultDoor) *storage.OrderData {
	operations := make([]storage.Operation, 0)

	// Добавляем только те, где значение > 0
	if result.NapilRamStv > 0 {
		operations = append(operations, storage.Operation{"napil_ram_stv", "Напиловка рамы и створки", result.NapilRamStv})
	}
	if result.PetliObr > 0 {
		operations = append(operations, storage.Operation{"petli_obr", "Обработка под петли", result.PetliObr})
	}
	if result.ZamokObr > 0 {
		operations = append(operations, storage.Operation{"zamok_obr", "Обработка под замок; ручку", result.ZamokObr})
	}
	if result.ShpingObrStv > 0 {
		operations = append(operations, storage.Operation{"shping_obr_stv", "Обработка створки под шпингалет; ответку", result.ShpingObrStv})
	}
	if result.ShpingObrRam > 0 {
		operations = append(operations, storage.Operation{"shping_obr_ram", "Обработка рамы под шпингалет", result.ShpingObrRam})
	}
	if result.YstanZakld > 0 {
		operations = append(operations, storage.Operation{"ystan_zakld", "Установка закладных в раму; створку", result.YstanZakld})
	}
	if result.SverlOtvShtift > 0 {
		operations = append(operations, storage.Operation{"sverl_otv_shtift", "Сверловка отверстий под штифты на створке и раме", result.SverlOtvShtift})
	}
	if result.FrezerStoekRigel > 0 {
		operations = append(operations, storage.Operation{"frezer_stoek_rigel", "Фрезеровка стойки, ригелей ленивой створки", result.FrezerStoekRigel})
	}
	if result.SborkaRam > 0 {
		operations = append(operations, storage.Operation{"sborka_ram", "Сборка рамы", result.SborkaRam})
	}
	if result.ShitfRam > 0 {
		operations = append(operations, storage.Operation{"shitf_ram", "Штифтовка рамы", result.ShitfRam})
	}
	if result.OpresStv > 0 {
		operations = append(operations, storage.Operation{"opres_stv", "Опрессовка створок", result.OpresStv})
	}
	if result.SborStv > 0 {
		operations = append(operations, storage.Operation{"sbor_stv", "Сборка створок", result.SborStv})
	}
	if result.ShiftStv > 0 {
		operations = append(operations, storage.Operation{"shift_stv", "Штифтовка створок", result.ShiftStv})
	}
	if result.NanesKlei > 0 {
		operations = append(operations, storage.Operation{"nanes_klei", "Нанесение клея", result.NanesKlei})
	}
	if result.YstanYplRam > 0 {
		operations = append(operations, storage.Operation{"ystan_ypl_ram", "Установка уплотнителей рам", result.YstanYplRam})
	}
	if result.YstanYplStv > 0 {
		operations = append(operations, storage.Operation{"ystan_ypl_stv", "Установка уплотнителей створки", result.YstanYplStv})
	}
	if result.YstZamokNakl > 0 {
		operations = append(operations, storage.Operation{"yst_zamok_nakl", "Установка замка; накладок; примыкающего профиля", result.YstZamokNakl})
	}
	if result.YstShpingOtv > 0 {
		operations = append(operations, storage.Operation{"yst_shping_otv", "Установка шпингалета; ответки; примыкающего профиля", result.YstShpingOtv})
	}
	if result.SborPetliRam > 0 {
		operations = append(operations, storage.Operation{"sbor_petli_ram", "Сборка петель рама", result.SborPetliRam})
	}
	if result.SborYstPorog > 0 {
		operations = append(operations, storage.Operation{"sbor_yst_porog", "Сборка и установка порога", result.SborYstPorog})
	}
	if result.Naveshiv > 0 {
		operations = append(operations, storage.Operation{"naveshiv", "Навешивание", result.Naveshiv})
	}
	if result.Zashiv > 0 {
		operations = append(operations, storage.Operation{"zashiv", "Зашивка", result.Zashiv})
	}
	if result.OpresRam > 0 {
		operations = append(operations, storage.Operation{"opres_ram", "Опрессовка рам", result.OpresRam})
	}
	if result.SborPetliStv > 0 {
		operations = append(operations, storage.Operation{"sbor_petli_stv", "Сборка петель створки", result.SborPetliStv})
	}
	if result.NastrStanok > 0 {
		operations = append(operations, storage.Operation{"nastr_stanok", "Настройка станка для напиловки", result.NastrStanok})
	}
	if result.MehObrabPzr > 0 {
		operations = append(operations, storage.Operation{"meh_obrab_pzr", "ПЗР", result.MehObrabPzr})
	}
	if result.RabotaPbx > 0 {
		operations = append(operations, storage.Operation{"rabota_pbx", "Работа станка РВХ", result.RabotaPbx})
	}
	if result.FrezerNastr > 0 {
		operations = append(operations, storage.Operation{"frezer_nastr", "Фрезеровка (Настройка)", result.FrezerNastr})
	}
	if result.FrezerPorogSborka > 0 {
		operations = append(operations, storage.Operation{"frezer_porog_sborka", "Фрезеровка (порогов),промежуточная сборка", result.FrezerPorogSborka})
	}
	if result.FrezerYstShtyp > 0 {
		operations = append(operations, storage.Operation{"frezer_yst_shtyp", " Фрезеровка и установка штульпов", result.FrezerYstShtyp})
	}
	if result.OpresNastr > 0 {
		operations = append(operations, storage.Operation{"opres_nastr", "Опрессовка (настройка)", result.OpresNastr})
	}
	if result.Opres > 0 {
		operations = append(operations, storage.Operation{"opres", "Опрессовка", result.Opres})
	}
	if result.PodgDerjShetki > 0 {
		operations = append(operations, storage.Operation{"podg_derj_shetki", "Подготовка держателя щетки", result.PodgDerjShetki})
	}
	if result.YstPorogYplDr > 0 {
		operations = append(operations, storage.Operation{"yst_porog_ypl_dr", "Установка порогов, уплотнителей, держателя, щетки, штифтовка, удаление излишков герметика после застывания", result.YstPorogYplDr})
	}
	if result.NaveshivStv > 0 {
		operations = append(operations, storage.Operation{"naveshiv_stv", "Навешивание створки, установка отв. Планок", result.NaveshivStv})
	}
	if result.YstZapoln > 0 {
		operations = append(operations, storage.Operation{"nastr_for_opres", "Установка заполнения", result.YstZapoln})
	}
	if result.ImpostNapil > 0 {
		operations = append(operations, storage.Operation{"impost_napil", "Доп. импост (напиловка)", result.ImpostNapil})
	}
	if result.ImpostFrezer > 0 {
		operations = append(operations, storage.Operation{"impost_frezer", "Доп. импост (Фрезеровка)", result.ImpostFrezer})
	}
	if result.ImpostSverlo > 0 {
		operations = append(operations, storage.Operation{"impost_sverlo", "Доп. импост (Сверловка)", result.ImpostSverlo})
	}
	if result.ImpostYst > 0 {
		operations = append(operations, storage.Operation{"impost_yst", "Доп. импост (установка)", result.ImpostYst})
	}
	if result.ImpostShtift > 0 {
		operations = append(operations, storage.Operation{"impost_shtift", "Доп. импост (штифтовка)", result.ImpostShtift})
	}
	if result.YplFalc > 0 {
		operations = append(operations, storage.Operation{"ypl_falc", "Уплотнитель фальца", result.YplFalc})
	}
	if result.NapilNalich > 0 {
		operations = append(operations, storage.Operation{"napil_nalich", "Напиловка наличника", result.NapilNalich})
	}
	if result.NapilRam > 0 {
		operations = append(operations, storage.Operation{"napil_ram", "Напиловка рамы", result.NapilRam})
	}
	if result.NapilStv > 0 {
		operations = append(operations, storage.Operation{"napil_stv", "Напиловка створки", result.NapilStv})
	}
	if result.KontrSbork > 0 {
		operations = append(operations, storage.Operation{"kontr_sbork", "Контрольная сборка и разборка стоек", result.KontrSbork})
	}
	if result.SverlRam > 0 {
		operations = append(operations, storage.Operation{"sverl_ram", "Сверловка рамы для монтажа; соед", result.SverlRam})
	}
	if result.SverlZink > 0 {
		operations = append(operations, storage.Operation{"sverl_zink", "Сверловка, зинковка и установка клёпок", result.SverlZink})
	}
	if result.ZashitPl > 0 {
		operations = append(operations, storage.Operation{"zashit_pl", "Наклейка защитной плёнки", result.ZashitPl})
	}
	if result.SborRam > 0 {
		operations = append(operations, storage.Operation{"sbor_ram", "Сборка рамы", result.SborRam})
	}
	if result.NapilYstKrishStv > 0 {
		operations = append(operations, storage.Operation{"napil_yst_krish_stv", "Напиловка и установка крышек на ств", result.NapilYstKrishStv})
	}
	if result.NapilYstKrishRam > 0 {
		operations = append(operations, storage.Operation{"napil_yst_krish_ram", "Напиловка и установка крышек на рам", result.NapilYstKrishRam})
	}
	if result.SborPetli > 0 {
		operations = append(operations, storage.Operation{"sbor_petli", "Сборка петель", result.SborPetli})
	}
	if result.YstPtliRamStv > 0 {
		operations = append(operations, storage.Operation{"yst_ptli_ram_stv", "Установка петель на раму и створку", result.YstPtliRamStv})
	}
	if result.RezkaPlast > 0 {
		operations = append(operations, storage.Operation{"rezka_plast", "Нарезать расклинивающий пластик", result.RezkaPlast})
	}
	if result.Brysok > 0 {
		operations = append(operations, storage.Operation{"brysok", "Прикрепить брусок (фальшпорог)", result.Brysok})
	}
	if result.IzgPritv > 0 {
		operations = append(operations, storage.Operation{"izg_pritv", "Изготовление притвора и настройка", result.IzgPritv})
	}
	if result.ObrPritv > 0 {
		operations = append(operations, storage.Operation{"obr_pritv", "Обработка притвора под замок и ответку", result.ObrPritv})
	}
	if result.YstPritv > 0 {
		operations = append(operations, storage.Operation{"yst_pritv", "Установка притвора", result.YstPritv})
	}
	if result.ObrabotkaAll > 0 {
		operations = append(operations, storage.Operation{"obrabotka_all", "Обработка (замок; ручка; ответка; шпингалет)", result.ObrabotkaAll})
	}
	if result.YstanPlnPetli > 0 {
		operations = append(operations, storage.Operation{"ystan_pln_petli", "Установка пластин под петли", result.YstanPlnPetli})
	}
	if result.YstFetr > 0 {
		operations = append(operations, storage.Operation{"yst_fetr", "Установка фетра", result.YstFetr})
	}
	if result.Rezina > 0 {
		operations = append(operations, storage.Operation{"rezina", "Обрезинивание", result.Rezina})
	}
	if result.FrezerShping > 0 {
		operations = append(operations, storage.Operation{"frezer_shping", "Фрезеровка под шпингалет", result.FrezerShping})
	}
	if result.Gl > 0 {
		operations = append(operations, storage.Operation{"gl", "Глухое окно", result.Gl})
	}
	if result.Fortochka > 0 {
		operations = append(operations, storage.Operation{"fortochka", "Форточка", result.Fortochka})
	}
	if result.Upak > 0 {
		operations = append(operations, storage.Operation{"ypak", "Упаковка", result.Upak})
	}
	if result.TotalTime > 0 {
		operations = append(operations, storage.Operation{"total_time", "Итого", result.TotalTime})
	}

	return &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Operations: operations,
		Profil:     result.Profil,
	}
}

func MapToOrderDataVitraj(result *storage.DemResultVitraj) *storage.OrderData {
	operations := make([]storage.Operation, 0)

	// Словарь "поле -> человекочитаемое название"
	operationNames := map[string]string{
		"PodgotOboryd":        "Подготовка оборудования",
		"NapilStkDo3m":        "Напиловка стойки до 3м",
		"NapilStkBol3m":       "Напиловка стойки свыше 3м",
		"NapilStkBol5m":       "Напиловка стойки свыше 5м",
		"PZR":                 "ПЗР для PBX",
		"NastrPbx":            "Настройка PBX",
		"DrenajOtv":           "Выполнить дренажные отверстия (работа станка)",
		"ZashelkObr":          "Обработка под защёлки (работа станка)",
		"PrisetPrObr":         "Выполнить обработку пристеночного профиля",
		"YplYst":              "Установить уплотнитель (по кол-ву панелей)",
		"FetrYst":             "Установить фетр (см. комплектац. щёточный уплотнитель) пгм",
		"StekloYst":           "Установка стеклопакета",
		"RazmetOtv":           "Разметка отверстий",
		"SverloOtv":           "Просверлить отверстия для роликов и накладок (по 2 на ролик. По 2 ролика на панель)",
		"NakladYst":           "Установка накладок",
		"Panel":               "Упаковать панели",
		"Naprav":              "Упаковать направляющие",
		"Ydlin":               "Упаковать удлинители",
		"Upak":                "Упаковка",
		"NapilRigelDo1m":      "Напиливание ригеля до 1м",
		"NapilRigelBol1m":     "Напиливание ригеля более 1м",
		"NakldNapil":          "Напиливание накладок",
		"NapilKriskhSt":       "Напиливание крышки стойки",
		"NapilKriskhRg":       "Напиливание крышки ригеля",
		"KomplKriskh":         "Комплектация крышек",
		"NapilAdapt":          "Напиливание адаптеров",
		"FrezerRigel":         "Фрезеровка ригеля",
		"FukelYst":            "Установка фукелей",
		"StoikiPbx":           "PBX стойки",
		"ZamokYst":            "Установка замка",
		"ZamokYstBolt":        "Установка замка под болты",
		"RigelSverloZamok":    "Сверление ригеля под замок",
		"NakladSverlo":        "Сверление накладок",
		"YplYstRigel":         "Уплотнение на ригель",
		"YplYstNakld":         "Уплотнение на накладку",
		"YplYstStoik":         "Уплотнение на стойки",
		"AdaptYstStoik":       "Установка адаптеров на стойки",
		"AdaptYstRigel":       "Установка адаптеров на ригель",
		"TermovstYstStoiki":   "Установка термовставок на стойки",
		"TermovstYstRigel":    "Установка термовставок на ригель",
		"SborNog":             "Сборка ножек",
		"Birki":               "Печать бирок по местам",
		"UpakStoik":           "Упаковка стоек",
		"UpakPet":             "Упаковка петель",
		"UpakRigel":           "Упаковка ригелей",
		"UpakRigel2m":         "Упаковка ригелей >2м",
		"UpakKriskh":          "Упаковка крышек",
		"UpakNakld":           "Упаковка накладок",
		"UpakYplNog":          "Упаковка уплотнителей и ножек",
		"UpakKronsht":         "Упаковка кронштейнов",
		"VinosGotovIzd":       "Вынос готового изделия",
		"YstStikZakld":        "Установка стыковочной закладной",
		"ObnYsovDo4m":         "Обнижение усов до 4м",
		"ObnYsovBol4m":        "Обнижение усов более 4м",
		"OtmRezin":            "Отматывание резины в отгруз",
		"ObnNastr":            "Настройка оборудования для обнижения",
		"TrybaProf":           "Труба профильная",
		"KomplShtapik":        "Комплектация штапиков",
		"OtvVO":               "Штамповка отверстий под водоотлив",
		"NastrStanokRigel":    "Настройка станка для ригеля",
		"FrezerStoikiPr":      "Фрезеровка стоек (прям)",
		"NastrStanokStoiki1":  "Настройка станка для стоек",
		"FrezerStoikiYgl":     "Фрезеровка углов стоек",
		"NastrStanokStoiki2":  "Настройка станка для стоек №2",
		"ZashitPl":            "Установка защитных плёнок",
		"YstPritv":            "Установка притвора",
		"YstKapel":            "Установка капельников",
		"SborSekci":           "Сборка секций",
		"GermetYpl":           "Герметизация и уст. уплотнения",
		"UpakSoed":            "Упаковка соединителей",
		"UpakShtapik":         "Упаковка штапиков",
		"UpakSekcii":          "Упаковка секций",
		"Obrezanie":           "Обрезание",
		"VinosSekcii":         "Вынос секций",
		"YgolKorob":           "Уголок и коробка",
		"PbxVO":               "PBX В/О",
		"ClearVO":             "Зачистка В/О",
		"YstKomptrSt":         "Установка компенсатора,пгм стойки",
		"YstKomptrRg":         "Установка компенсатора,пгм ригель",
		"UpakYgl":             "Упаковка углов",
		"NapilYgl":            "Напиловка уголка",
		"NarezTermovst":       "Нарезка термовставок",
		"NarezKompensr":       "Нарезка компенсаторов",
		"UpakTermovst":        "Упаковка термовставок",
		"RigelSverloZamok3m":  "Сверление ригеля под замок выше 3м",
		"RigelNsverloZamok3m": "Сверление нестандарт ригеля под замок",
		"ObjYsovTermvst":      "Обжим усов для термовставки",
		"YstKPS":              "Установка КПС",
		"NapilStkBol6m":       "Напиловка стойки свыше 6м",
		"NapilShtapik":        "Напиливание штапика",
		"RezProfil":           "Резина пгм в профиль",
		"RezSekcii":           "Резина пгм в секцию",
		"PartSbSeck":          "Частичная сборка секций",
		"PartSbMarkSeck":      "Частичная сборка маркировка секции",
		"PartSbStoek":         "Частичная сборка стойки",
		"YstFormr":            "Установка формирователя",
		"TotalTime":           "Итого",
	}

	// Используем рефлексию для итерации по всем float64 полям
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Пропускаем не float64 поля
		if field.Kind() != reflect.Float64 {
			continue
		}

		value := field.Float()
		// Добавляем операцию только если значение > 0
		if value > 0 {
			// Получаем человекочитаемое имя из словаря
			displayName, exists := operationNames[fieldName]
			if !exists {
				displayName = fieldName // fallback
			}
			operations = append(operations, storage.Operation{
				ID:    fieldName,
				Name:  displayName,
				Value: value,
			})
		}
	}

	return &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Operations: operations,
		Profil:     result.Profil,
	}
}

func MapToOrderDataLoggia(result *storage.DemResultLoggia) *storage.OrderData {
	operations := make([]storage.Operation, 0)

	// Словарь: поле Go -> человекочитаемое название операции
	operationNames := map[string]string{
		"PodgotOboryd":      "Подготовка оборудования",
		"NapilStkDo3m":      "Напиловка стойки до 3м",
		"NapilRigelDo1m":    "Напиливание ригеля до 1м",
		"NapilRigelBol1m":   "Напиливание ригеля более 1м",
		"NapilShtapik":      "Напиливание штапика",
		"KomplShtapik":      "Комплектация штапика по местам",
		"NapilAdaptTr":      "Напиливание адаптеров и труб",
		"RigelFrezer":       "Фрезеровка ригеля",
		"FrezerRigelZamok":  "Фрезеровка ригеля под замок",
		"ZamokYst":          "Установка замка",
		"RezPgmSt":          "Резина пгм в профиль стойки",
		"RezPgmRg":          "Резина пгм в профиль ригель",
		"FrezVo":            "Фрезеровка водоотливных отверстий (2шт) + настрой",
		"PartSbSekci":       "Частичная сборка секций",
		"YstPritv":          "Установка притвора",
		"YstFormir":         "Установка формирователя",
		"Birki":             "Печать бирок по местам",
		"UpakRam":           "Упаковка рам",
		"UpakStoik":         "Упаковка стоек",
		"UpakRigel":         "Упаковка ригелей",
		"UpakRigel2m":       "Упаковка ригелей выше 2м",
		"UpakShtapik":       "Упаковка штапиков",
		"UpakAdaptTr":       "Упаковка адаптеров и труб",
		"UpakYplNog":        "Упаковка уплотнителей и ножек",
		"VinosGotovIzd":     "Вынос готового изделия",
		"NapilPrVirav":      "Напиловка профиль выравнивающий",
		"NapilRam":          "Напиливание рамы",
		"NapilStv":          "Напиливание створки",
		"NapilPritv":        "Напиливание притвора",
		"NapilSoed":         "Напиливание соединителей",
		"FrezRam":           "Фрезеровка рамы",
		"FrezStv":           "Фрезеровка створки",
		"FrezPritv":         "Фрезеровка притвора",
		"ClearSverlStv":     "Зачистить после фрез. Сверлить отв. По шаблону",
		"ClearSverlYsilStv": "Зачистить после фрез. Сверлить отв. По шаблону. Усиленная стойка",
		"Kraska":            "Подкраска торцов",
		"YstRolik":          "Установка роликов, фетра (риг), удаление пленки",
		"YstZashel":         "Установка защёлок, фетра",
		"Rezina":            "Обрезинивание",
		"SborStv":           "Сборка створки",
		"SborPritv":         "Сборка притвора",
		"PodgKompl":         "Подготовка комплектующих",
		"ShtampStStv":       "Штамповка стоек створки",
		"PodgRezin":         "Подготовка резины",
		"SborYpl":           "Сборка уплотнителей",
		"YstZaklep":         "Установка заклёпок",
		"TotalTime":         "Итого",
	}

	// Используем рефлексию для доступа к полям структуры
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		// Обрабатываем только float64 поля
		if field.Kind() != reflect.Float64 {
			continue
		}

		value := field.Float()

		// Добавляем только если значение > 0
		if value > 0 {
			displayName, exists := operationNames[fieldName]
			if !exists {
				displayName = fieldName // fallback, если нет в словаре
			}
			operations = append(operations, storage.Operation{
				ID:    fieldName,
				Name:  displayName,
				Value: value,
			})
		}
	}

	return &storage.OrderData{
		ID:         result.ID,
		OrderNum:   result.OrderNum,
		Name:       result.Name,
		Count:      result.Count,
		Profil:     result.Profil,
		Operations: operations,
	}
}
