package main

import (
	"fmt"     //"format" - пакет для форматирования, ввода и вывода текста
	"strings" //этот пакет содержит функции для работы со строками
)

func main() {
    
    fmt.Println("Афанасий Капалкин. E-mail: kafriendly@yandex.ru")
    
    fmt.Println("Тестовое задание: «Калькулятор одной строкой».")
    
    fmt.Println("Подружить с римскими цифрами на данный момент не удалось, т.к. мне требуется еще время для изучения этого языка програмирования. \nНо есть идея - присвоить арабским цифрам римские символы и продолжить выполнение кода с арабскими цифрами:")
    fmt.Println("         C: 100, \n         XC: 90, \n         L: 50, \n         XL: 40, \n         X: 10, \n         IX: 9, \n         VIII: 8, \n         VII: 7, \n         VI: 6, \n         V: 5,\n         IV:  4, \n         III: 3, \n         II: 2, \n         I: 1")
    fmt.Println("\nИтак, собственно, сам консольный калькулятор: \n*********************************************")
    fmt.Println("")
	for { // этот цикл будет длиться вечно, пока не будет вызван оператор «break»
		//распечатать запрос
		fmt.Print("Введите математическое выражение (например, 3/3 или 2+7): ")

		//создаем строковую переменную
		var input string

		//заполняем переменную данными, введенными пользователем, и переходим на новую строку
		fmt.Scanln(&input)

		//ПРОВЕРЯЕМ ВЫВОД
		if strings.TrimSpace(strings.ToLower(input)) == "exit" { //«TrimSpace» возвращает входящую строку, в которой удалены все конечные и начальные пробелы.
			fmt.Println("Выход из программы...")
			break
			// ПЛАН ДЕЙСТВИЙ:
			//удаляем все пробельные символы, которые может ввести пользователь из поля ввода
			//строчные буквы во входящей строке
			//если ввод содержит "exit" - выход
		}

		//УДАЛЯЕМ ПРОБЕЛЫ. ОПЕРАТОРЫ
		// разделяем входные данные на два операнда и оператор
		expr := strings.ReplaceAll(input, " ", "") // заменяем все оставшиеся пробелы «ничем»
		opIndex := strings.IndexAny(expr, "+-*/")  // находим индекс оператора в строке
		if opIndex == -1 {                         // если оператор не найден - перезапускаем цикл
			fmt.Println("Ввод должен быть в формате: «цифра оператор (+-*/) цифра»")
			continue
		}
		op := expr[opIndex]                                                  //берем оператор из строки
		operands := strings.Split(expr[0:opIndex]+","+expr[opIndex+1:], ",") //отделяем операнды от строки и помещаем их как элементы внутри среза

		//ПРЕОБРАЗОВЫВАЕМ ОПЕРАНДЫ В ЧИСЛА
		var num1, num2 float64
		n, err := fmt.Sscanf(operands[0]+" "+operands[1], "%f %f", &num1, &num2)
		//Sscanf берет строку (operands[0]+" "operands[1]), форматируем строку (%f %f) для указания вывода из строки и указатели для сохранения вывода.
		//Sscanf изменяет num1 и num2 и возвращает количество выполненных назначений (n) и ошибку, если таковая имеется.
		if n != 2 || err != nil { //Если операндов больше двух или произошла ошибка, то цикл перезапускается
			fmt.Println("Ввод должен быть в формате: «цифра оператор (+-*/) цифра»")
			continue
		}
		if num1 == 0 {
            fmt.Print("Пожалуйста не используйте «0»! ")
            continue
		}
		
		if num2 == 0 {
            fmt.Print("Пожалуйста не используйте «0»! ")
            continue
		}
	
		//Создаем расчеты
		var result float64
		switch op {
		case '+':
			result = num1 + num2
		case '-':
			result = num1 - num2
		case '*':
			result = num1 * num2
		case '/':
			result = num1 / num2
		default:
			fmt.Println("Недопустимый оператор")
			continue
		}

		fmt.Printf("%g %c %g = %g\n", num1, op, num2, result) //последовательность вывода результата первая цифра, оператор и вторая цифра
		//
	}
}
