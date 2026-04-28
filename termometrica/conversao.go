package main

import "fmt"

const ebulicaoKelvin = 373.15

func main() {
	temperaturaKelvin := ebulicaoKelvin
	temperaturaCelsius := temperaturaKelvin - 273.15
	fmt.Printf("A temperatura de ebulição da água é %.2f K ou %.2f °C\n", temperaturaKelvin, temperaturaCelsius)
}
