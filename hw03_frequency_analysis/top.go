package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func Top10(s string) []string {
	res := make([]string, 0, 10)

	// Если пришла пустая строка
	if s == "" {
		return res
	}

	// Разделяем строку на слова при этом удаляем все лишние пробелы по решулярке
	r := regexp.MustCompile(`\s+`)
	words := strings.Split(r.ReplaceAllString(s, " "), " ")

	// Мапа для подсчета кол-ва каждого слова
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// Слайс пар слово-количество для удобной работы
	pairList := make(PairList, 0, len(counts))
	for k, v := range counts {
		pairList = append(pairList, Pair{k, v})
	}

	// Сортировка слайса по убыванию кол-ва вхождения слова
	sort.SliceStable(pairList, func(i, j int) bool {
		return pairList[i].Value > pairList[j].Value
	})

	// Текущее кол-во повторений (при первом проходе максимальное)
	currentRepeat := pairList[0].Value

	// Слайс для слов с одинаковым кол-вом повторений
	keys := make([]string, 0, 10)

	for j, pair := range pairList {
		if j >= 10 {
			sort.Strings(keys)         // Сортируем слайс лексикографически
			res = append(res, keys...) // Добавляем в результат
			break
		}

		if pair.Value == currentRepeat {
			keys = append(keys, pair.Key)
		} else {
			currentRepeat = pair.Value    // Изменяем текущее кол-во повторений
			sort.Strings(keys)            // Сортируем слайс лексикографически
			res = append(res, keys...)    // Добавляем в результат
			keys = []string{}             // Очищаем слайс
			keys = append(keys, pair.Key) // Добавляем это слово в слайс
		}
	}

	return res
}
