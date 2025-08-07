package main

import (
	"fmt"
	"strings"
	"time"
)

// main es el punto de inicio de nuestro programa.
func main() {
	// Usamos un bucle para mostrar el menú una y otra vez.
	// La única forma de salir es con la opción 0.
	for {
		mostrarMenu()
		opcion := ""
		fmt.Scanln(&opcion) // Leer la opción que el usuario escribió.

		switch opcion {
		case "1":
			pruebaVelocidad()
		case "2":
			analisisNarrativa()
		case "3":
			completarAcertijo()
		case "4":
			identificarInformacionTrivial()
		case "5":
			respuestaInstantanea()
		case "0":
			fmt.Println("¡Apagando la caja de herramientas del robot!")
			return // Esto termina el programa.
		default:
			fmt.Println("Esa opción no existe. Por favor, elige un número del menú.")
		}
		// Esperar un poco antes de volver al menú, para que puedas leer la respuesta.
		fmt.Println("\n--- Presiona Enter para volver al menú ---")
		fmt.Scanln()
	}
}

// mostrarMenu imprime las opciones disponibles en la pantalla.
func mostrarMenu() {
	fmt.Println("\n--- PRUEBA DE ROBOT ---")
	fmt.Println("Elige una opción para probar tus habilidades robóticas:")
	fmt.Println("1. Velocidad de procesamiento (matemáticas)")
	fmt.Println("2. Análisis de una narrativa emocional")
	fmt.Println("3. Completar un acertijo (sin humor)")
	fmt.Println("4. Identificación de información trivial")
	fmt.Println("5. Respuesta instantánea a 10 preguntas")
	fmt.Println("0. Salir del programa")
	fmt.Print("Tu elección: ")
}

// pruebaVelocidad calcula una operación matemática muy rápido.
func pruebaVelocidad() {
	fmt.Println("\n--- Velocidad de procesamiento ---")
	fmt.Println("Vamos a resolver una operación muy difícil...")
	
	// La operación matemática que se va a resolver.
	operacion := "78943 × 987123 + 5786"
	fmt.Printf("La operación es: %s\n", operacion)

	// El robot empieza a medir el tiempo.
	tiempoInicio := time.Now()

	// La operación matemática.
	resultado := 78943 * 987123 + 5786

	// El robot termina de medir el tiempo.
	tiempoFinal := time.Now()
	duracion := tiempoFinal.Sub(tiempoInicio)

	// Imprimir el resultado y el tiempo que tardó.
	fmt.Printf("La respuesta es: %d\n", resultado)
	fmt.Printf("Tiempo de cálculo del robot: %v\n", duracion)
	fmt.Println("¡Más rápido de lo que un humano puede pensar!")
}

// analisisNarrativa lee una historia y la resume sin emociones.
func analisisNarrativa() {
	fmt.Println("\n--- Análisis de una narrativa emocional ---")
	historia := `Ana encontró una billetera en el parque. Dentro, había mucho dinero, pero también la foto de una familia con un bebé. La dueña, una anciana llamada Marta, la había perdido mientras jugaba con su nieto. Ana pensó en usar el dinero, pero luego pensó en la anciana y su familia, que probablemente lo necesitaban para el bebé. Con el corazón en la mano, se lo entregó a la policía.`
	
	resumenRobot := `Una persona llamada Ana encontró una billetera en un parque. La billetera contenía dinero y una foto familiar. La propietaria era una anciana llamada Marta, quien había perdido la billetera. Ana entregó la billetera a la policía.`

	fmt.Println("Historia original:\n", historia)
	fmt.Println("\nResumen objetivo del robot:\n", resumenRobot)
	fmt.Println("El robot solo vio los hechos, no la emoción de Ana o de Marta.")
}

// completarAcertijo completa una frase de forma literal.
func completarAcertijo() {
	fmt.Println("\n--- Completar un acertijo ---")
	chiste := "Un pez le pregunta a otro: ¿Qué hace tu papá?"
	respuestaRobot := "Su padre se dedica a nadar y a comer."
	
	fmt.Println("El chiste que un humano completaría:")
	fmt.Println(chiste, "¿Pescado?")
	fmt.Println("\nLa respuesta del robot (pura lógica, nada de humor):")
	fmt.Println(chiste, respuestaRobot)
}

// identificarInformacionTrivial busca una palabra en un texto muy largo.
func identificarInformacionTrivial() {
	fmt.Println("\n--- Identificación de información trivial ---")
	
	// Creamos un texto largo. Esto simula leer un libro enorme.
	textoLargo := strings.Repeat("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ", 100)
	palabraUnica := "supercalifragilisticoespialidoso" // Esta es la palabra escondida.
	
	// Escondemos la palabra única en el texto.
	textoLargo = textoLargo[:5000] + palabraUnica + textoLargo[5000:]
	
	fmt.Println("El robot va a buscar una palabra única en un texto de 10,000 caracteres...")
	
	// El robot busca la palabra.
	indice := strings.Index(textoLargo, palabraUnica)

	fmt.Printf("¡Encontrado! La palabra es '%s'.\n", palabraUnica)
	fmt.Printf("Está en la posición %d del texto.\n", indice)
	fmt.Println("El robot no necesitó leer todo el texto para encontrarla.")
}

// respuestaInstantanea responde preguntas en un tiempo exacto.
func respuestaInstantanea() {
	fmt.Println("\n--- Respuesta instantánea a 10 preguntas ---")
	fmt.Println("El robot responderá a todas las preguntas en el mismo tiempo exacto.")
	
	// Las preguntas y respuestas del robot.
	preguntasRespuestas := map[string]string{
		"¿Cuál es el color del cielo?": "El color del cielo es azul.",
		"¿Cuánto es 2 + 2?": "La suma de 2 y 2 es 4.",
		"¿Quién fue Isaac Newton?": "Isaac Newton fue un físico y matemático inglés.",
		"¿Qué es la fotosíntesis?": "Es un proceso de las plantas para producir su alimento.",
		"¿Qué día es hoy?": "Hoy es miércoles.",
		"¿Dónde está la Torre Eiffel?": "La Torre Eiffel se encuentra en París, Francia.",
		"¿Cuál es el planeta más grande?": "El planeta más grande es Júpiter.",
		"¿Qué comes en el desayuno?": "El robot no come.",
		"¿A qué hora empieza el sol?": "El sol empieza al amanecer.",
		"¿Qué es un robot?": "Un robot es una máquina programada.",
	}

	// El tiempo que el robot va a tardar en cada respuesta.
	// Es un tiempo fijo para que todas las respuestas sean iguales.
	tiempoDeRespuestaFijo := 250 * time.Millisecond // 250 milisegundos

	i := 1
	for pregunta, respuesta := range preguntasRespuestas {
		tiempoInicio := time.Now() // El robot mide el tiempo.
		
		fmt.Printf("\nPregunta %d: %s\n", i, pregunta)
		
		// El robot espera el tiempo exacto antes de responder.
		time.Sleep(tiempoDeRespuestaFijo)
		
		fmt.Printf("Respuesta del robot: %s\n", respuesta)
		
		tiempoFinal := time.Now()
		duracion := tiempoFinal.Sub(tiempoInicio)
		
		fmt.Printf("Tiempo de respuesta: %v\n", duracion.Round(time.Millisecond))
		i++
	}
	fmt.Println("\nTodas las respuestas se dieron en el mismo tiempo, sin importar si eran difíciles o fáciles.")
}
