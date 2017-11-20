package main

import "regexp"
import "strings"


/* TASK 1
 *
 * Релизовать функцию, которая принимает на вход slice из чисел и возвращает
 * slice в котором удалены все четные числа.
 *
 ******************************************************************************/

func Filter(filtered_slice []int, predicate func(int) bool) []int {
    result_slice := make([]int, 0)
    for _, element := range filtered_slice {
        if predicate(element) {
            result_slice = append(result_slice, element)
        }
    }
    return result_slice
}

func Odd(number int) bool {
	return number % 2 == 1
}

func RemoveEven(filtered_slice []int) []int {
	return Filter(filtered_slice, Odd)
}

/******************************************************************************/

/* TASK 2
 *
 * Написать генератор, который будет принимать число и при вызове возвращать
 * очередную степень этого часла.
 ******************************************************************************/

func PowerGenerator(seed int) func() int {
	current_value := 1
	return func() int {
		current_value *= seed
		return current_value
	}
}

/******************************************************************************/

/* TASK 3
 *
 * Реализуйте функцию, которая подсчитывает число различных слов в тексте
 * (строчка). “Словом” считается непустая последовательность букв. Слова,
 * отличающиеся только регистром символов, считаются одинаковыми.
 ******************************************************************************/

func PrepareString(prepared_string string) string {
	non_letters := regexp.MustCompile(`\P{L}`)
	result := non_letters.ReplaceAllString(prepared_string, " ")
	multiple_spaces := regexp.MustCompile(`\s+`)
	result = multiple_spaces.ReplaceAllString(result, " ")
	spaces_on_sides := regexp.MustCompile(`(^\s|\s$)`)
	result = spaces_on_sides.ReplaceAllString(result, "")
	return strings.ToLower(result)
}

func GetWords(extracted_string string) []string {
    words:= regexp.MustCompile(`\p{L}+`)
    return words.FindAllString(extracted_string, -1)
}

func DifferentWordsCountFromSlice(words []string) int {
    word_counts := make(map[string]int)
    for _, word :=range words {
        word_counts[word]++
    }
    return len(word_counts);
}

func DifferentWordsCount(text string) int {
	prepared_string := PrepareString(text)
	words_slice := GetWords(prepared_string)
	return DifferentWordsCountFromSlice(words_slice)
}
