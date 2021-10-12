package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func esPrimo(numero int) bool {
	if numero == 0 || numero == 1 {
		return false
	}
	contador := 0
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			contador++
		}
	}
	return !(contador > 2)
}

func dePalabrasAFrase(palabras []string) string {
	return strings.Join(palabras, " ")
}

func arrayIncludes(array []int, elemento int) bool {
	for _, id := range array {
		if id == elemento {
			return true
		}
	}
	return false
}

func promedio(resultados []int) int {
	suma := 0
	for _, id := range resultados {
		suma += id
	}
	return suma / len(resultados)
}

func todosIguales(arreglo []int) bool {
	for i := 0; i < len(arreglo)-1; i++ {
		if arreglo[i] != arreglo[i+1] {
			return false
		}
	}
	return true
}

func mesesDelAnio(array []string) []string {
	result := []string{}
	contador := map[string]int{}
	for _, id := range array {
		if id == "Enero" {
			result = append(result, "Enero")
			contador["Enero"]++
		} else if id == "Marzo" {
			result = append(result, "Marzo")
			contador["Marzo"]++
		} else if id == "Noviembre" {
			result = append(result, "Noviembre")
			contador["Noviembre"]++
		}
	}

	if len(contador) == 3 {
		return result
	} else {
		return []string{"No se encontraron los meses pedidos"}
	}
}

type Gato struct {
	nombre string
	edad   int
}

func (this Gato) meow() string {
	return this.nombre + "says Meow!"
}

func crearGato(nombre string, edad int) {
	gato := Gato{nombre, edad}
	fmt.Println(gato)
}

func operacionMatematica(n1, n2 int, cb func(i, j int) int) {
	fmt.Println(cb(n1, n2))
}

func capicua(numero int) bool {
	strNum := []rune(strconv.Itoa(numero))
	aux := []string{}
	for i := len(strNum) - 1; i >= 0; i-- {
		aux = append(aux, string(strNum[i]))
	}
	return string(strNum) == strings.Join(aux, "")
}

func deleteAbc(cadena string) string {
	rune := []rune(cadena)
	aux := []string{}
	for i := 0; i < len(rune); i++ {
		if string(rune[i]) != "a" && string(rune[i]) != "b" && string(rune[i]) != "c" {
			aux = append(aux, string(rune[i]))
		}
	}
	return string(strings.Join(aux, ""))
}

func sortArray(arr []string) []string {
	aux := map[int][]string{}
	result := []string{}
	for _, value := range arr {
		if len(aux[len(value)]) == 0 {
			aux[len(value)] = []string{}
		}
		aux[len(value)] = append(aux[len(value)], value)
	}
	//sort map by key
	keys := make([]int, 0, len(aux))
	for k := range aux {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, value := range keys {
		result = append(result, aux[value]...)
	}
	fmt.Println(aux)
	return result
	//probar implementar con quick sort adaptado
}

func buscoInterseccion(arr1, arr2 []int) []int {
	result := []int{}
	aux := map[int]bool{}
	longest := []int{}
	shortest := []int{}

	if len(arr1) > len(arr2) {
		longest = arr1
		shortest = arr2
	} else {
		longest = arr2
		shortest = arr1
	}
	for _, value := range longest {
		aux[value] = true
	}
	for _, value := range shortest {
		if aux[value] {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println("EsPrimo 7?: ", esPrimo(7))
	fmt.Println("EsPrimo 97?: ", esPrimo(97))
	fmt.Println("EsPrimo 10?: ", esPrimo(10))
	fmt.Println("EsPrimo 100?: ", esPrimo(100))
	fmt.Println("EsPrimo 0?: ", esPrimo(0))
	fmt.Println("EsPrimo 1?: ", esPrimo(1))
	fmt.Println("De palabras a frase ['Henry', 'JavaScript', 'Class']", dePalabrasAFrase([]string{"Henry", "JavaScript", "Class"}))
	fmt.Println("De palabras a frase ['Henry']", dePalabrasAFrase([]string{"Henry"}))
	fmt.Println("arrayContiene(arr, item)", arrayIncludes([]int{10, 10, 16, 12}, 10))
	fmt.Println("arrayContiene(arr, item)", arrayIncludes([]int{10, 10, 16, 12}, 12))
	fmt.Println("arrayContiene(arr, item)", arrayIncludes([]int{10, 10, 16, 12}, 22))
	fmt.Println("promedio(resultados", promedio([]int{10, 10, 16, 12}))
	fmt.Println("todosIguales(arreglo)", todosIguales([]int{20, 20, 20, 20}))
	fmt.Println("todosIguales(arreglo)", todosIguales([]int{97, 100, 190, 9}))
	fmt.Println("mesesDelAño(array)", mesesDelAnio([]string{"Marzo", "Diciembre", "Abril", "Junio", "Julio", "Noviembre", "Enero", "Mayo", "Febrero"}))
	fmt.Println("mesesDelAño(array)", mesesDelAnio([]string{"Marzo", "Marzo", "Diciembre", "Julio", "Noviembre"}))
	crearGato("golang", 3)
	operacionMatematica(2, 3, func(i, j int) int {
		return i + j
	})
	fmt.Println("Es capicua 12321 ? ", capicua(12321))
	fmt.Println("Es capicua 105217 ? ", capicua(105217))
	fmt.Println("deleteAbc(abcefgh)", deleteAbc("abcefgh"))
	fmt.Println("deleteAbc(abc)", deleteAbc("abc"))
	fmt.Println("deleteAbc(plural)", deleteAbc("plural"))
	fmt.Println("deleteAbc(limon)", deleteAbc("limon"))
	fmt.Println("sortArray(array)", sortArray([]string{"You", "are", "beautiful", "looking"}))
	fmt.Println("sortArray(array)", sortArray([]string{"pera", "manzana", "alcaucil", "papa"}))
	fmt.Println("buscoInterseccion(arreglo1, arreglo2)", buscoInterseccion([]int{1, 2, 3}, []int{1, 5, 8, 3}))
	fmt.Println("buscoInterseccion(arreglo1, arreglo2)", buscoInterseccion([]int{7, 23, 4}, []int{23, 70}))
	fmt.Println("buscoInterseccion(arreglo1, arreglo2)", buscoInterseccion([]int{1, 20, 3}, []int{22, 5, 7}))
}
