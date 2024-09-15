package main

import (
    "fmt"
    "regexp"
    "strconv"
)

func main() {
    var input string
    fmt.Println("Введите выражение (для арабских или римских):")
    _, err := fmt.Scanln(&input)
    if err != nil {
        panic("Ошибка ввода: " + err.Error())
    }

    // Регулярное выражение для проверки формата ввода
    re := regexp.MustCompile(`^(\d|[IVXLCDM]{1,4})\s*([-+*/])\s*(\d|[IVXLCDM]{1,4})$`)
    matches := re.FindStringSubmatch(input)

    if matches == nil {
        panic("Неверный формат ввода")
    }

    a, operator, b := matches[1], matches[2], matches[3]

    // Проверка, является ли ввод арабскими числами
    var isArabic bool
    if isArabicNumber(a) && isArabicNumber(b) {
        isArabic = true
    } else if isRomanNumber(a) && isRomanNumber(b) {
        // Проверка, является ли ввод римскими числами
    } else {
        panic("Соблюдайте правила ввода")
    }

    var num1, num2 int
    if isArabic {
        num1 = toArabic(a)
        num2 = toArabic(b)
    } else {
        num1 = toRoman(a)
        num2 = toRoman(b)
    }

    var result int
    switch operator {
    case "+":
        result = num1 + num2
    case "-":
        if isArabic || num1-num2 > 0 {
            result = num1 - num2
        } else {
            panic("Результат не может быть отрицательным для римских чисел")
        }
    case "*":
        result = num1 * num2
    case "/":
        if num2 == 0 {
            panic("Ошибка деления на ноль")
        }
        result = num1 / num2
    default:
        panic("Неверный оператор")
    }

    if isArabic {
        fmt.Println(result)
    } else {
        // Для римских чисел, если результат < 1, вызывать панику
        if result < 1 {
            panic("Результат меньше 1, что недопустимо для римских чисел")
        }
        fmt.Println(toRomanOutput(result))
    }
}

// Проверка, является ли строка арабским числом
func isArabicNumber(s string) bool {
    num, err := strconv.Atoi(s)
    return err == nil && num >= 1 && num <= 10
}

// Проверка, является ли строка римским числом
func isRomanNumber(s string) bool {
    return regexp.MustCompile(`^(I|V|X|L|C|D|M)+$`).MatchString(s)
}

// Преобразование римского числа в арабское
func toRoman(s string) int {
    roman := map[rune]int{
        'I': 1, 'V': 5, 'X': 10,
        'L': 50, 'C': 100, 'D': 500, 'M': 1000,
    }
    sum, prev := 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        current := roman[rune(s[i])]
        if current < prev {
            sum -= current
        } else {
            sum += current
        }
        prev = current
    }
    return sum
}

// Преобразование арабского числа в римское
func toRomanOutput(num int) string {
    vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    result := ""
    for i, v := range vals {
        for num >= v {
            num -= v
            result += romans[i]
        }
    }
    return result
}

// Преобразование арабского числа (строка) в целое значение
func toArabic(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil || num < 1 || num > 10 {
        panic("Недопустимые арабские числа. Должно быть от 1 до 10.")
    }
    return num
}

