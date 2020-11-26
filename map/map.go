package main

import "fmt"

func main() {
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[789456123] = "Pedro"
	aprovados[121513551] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 789456123)

	fmt.Println(aprovados)

	funcsESalarios := map[string]float64{
		"José":  11231.23,
		"Maria": 11542.36,
		"Pedro": 3525.22,
	}

	funcsESalarios["João"] = 2500.00

	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s %.2f\n", nome, salario)
	}

	fmt.Println("##############################")

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12345.00,
			"Gustavo":  1020.30,
		},
		"J": {
			"Julio": 20000.00,
		},
		"P": {
			"Pedro": 3000.00,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")

	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}

}
