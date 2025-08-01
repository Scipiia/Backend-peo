package utils

import "vue-golang/internal/storage"

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
