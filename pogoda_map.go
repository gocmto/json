package main

func getConfig(key string) string {
	elements := map[string]string{
		"place":       "Место",
		"url":         "Подробный URL",
		"fact":        "Фактически",
		"forecast":    "Прогноз",
		"date":        "Дата",
		"now":         "Время замера",
		"info":        "Информация",
		"temp":        "Температура (°C)",
		"feels_like":  "Ощущаемая температура (°C)",
		"temp_water":  "Температура воды (°C)",
		"condition":   "Описание погоды",
		"wind_speed":  "Скорость ветра (в м/с)",
		"wind_gust":   "Скорость порывов ветра (в м/с)",
		"wind_dir":    "Направление ветра",
		"pressure_mm": "Давление (в мм рт. ст.)",
		"pressure_pa": "Давление (в гектопаскалях)",
		"humidity":    "Влажность воздуха (в процентах)",
		"daytime":     "Светлое или темное время суток",
		"polar":       "Признак полярного дня или ночи",
		"season":      "Время года в данном населенном пункте",
		"obs_time":    "Время замера погодных данных в формате Unixtime",
		"week":        "Порядковый номер недели",
		"sunrise":     "Время восхода Солнца",
		"sunset":      "Время заката Солнца",
		"moon_code":   "Код фазы Луны",
		"moon_text":   "Текстовый код для фазы Луны",
		"part_name":   "Название времени суток",
		"temp_min":    "Минимальная температура для времени суток (°C)",
		"temp_max":    "Максимальная температура для времени суток (°C)",
		"temp_avg":    "Средняя температура для времени суток (°C)",
		"prec_mm":     "Прогнозируемое количество осадков (в мм)",
		"prec_period": "Прогнозируемый период осадков (в минутах)",
		"prec_prob":   "Вероятность выпадения осадков",
	}
	return elements[key]
}

func getCondition(key string) string {
	elements := map[string]string{
		"clear":                            "ясно",
		"partly-cloudy":                    "малооблачно",
		"cloudy":                           "облачно с прояснениями",
		"overcast":                         "пасмурно",
		"partly-cloudy-and-light-rain":     "небольшой дождь",
		"partly-cloudy-and-rain":           "дождь",
		"overcast-and-rain":                "сильный дождь",
		"overcast-thunderstorms-with-rain": "сильный дождь, гроза",
		"cloudy-and-light-rain":            "небольшой дождь",
		"overcast-and-light-rain":          "небольшой дождь",
		"cloudy-and-rain":                  "дождь",
		"overcast-and-wet-snow":            "дождь со снегом",
		"partly-cloudy-and-light-snow":     "небольшой снег",
		"partly-cloudy-and-snow":           "снег",
		"overcast-and-snow":                "снегопад",
		"cloudy-and-light-snow":            "небольшой снег",
		"overcast-and-light-snow":          "небольшой снег",
		"cloudy-and-snow":                  "снег",
	}
	return elements[key]
}

func getWinDir(key string) string {
	elements := map[string]string{
		"nw": "северо-западное",
		"n":  "северное",
		"ne": "северо-восточное",
		"e":  "восточное",
		"se": "юго-восточное",
		"s":  "южное",
		"sw": "юго-западное",
		"w":  "западное",
		"с":  "штиль",
	}
	return elements[key]
}

func getD(key string) string {
	elements := map[string]string{
		"d": "светлое время суток",
		"n": "темное время суток",
	}
	return elements[key]
}

func getSeason(key string) string {
	elements := map[string]string{
		"summer": "лето",
		"autumn": "осень",
		"winter": "зима",
		"spring": "весна",
	}
	return elements[key]
}

func getMoonCode(key string) string {
	elements := map[string]string{
		"0":  "полнолуние",
		"1":  "убывающая Луна",
		"2":  "убывающая Луна",
		"3":  "убывающая Луна",
		"4":  "последняя четверть",
		"5":  "убывающая Луна",
		"6":  "убывающая Луна",
		"7":  "убывающая Луна",
		"8":  "новолуние",
		"9":  "растущая Луна",
		"10": "растущая Луна",
		"11": "растущая Луна",
		"12": "первая четверть",
		"13": "растущая Луна",
		"14": "растущая Луна",
		"15": "растущая Луна",
	}
	return elements[key]
}

func getPartName(key string) string {
	elements := map[string]string{
		"night":   "ночь",
		"morning": "утро",
		"day":     "день",
		"evening": "вечер",
	}
	return elements[key]
}
