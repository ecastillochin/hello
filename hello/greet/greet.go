package greet

//variable por package, puede ser usada dentro de todo el package
//Primera letra en mayuscula se puede exportar
//Primera letra en minuscula no se puede exportar
var emoji = "emoji"

//English return greet in English
func English() string {
	return "Hi " + emoji
}

//Italian return greet in Italian
func Italian() string {
	return "Ciao " + emoji
}
