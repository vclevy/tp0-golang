package main

import (
	"bufio"
	"client/globals"
	"client/utils"
	"fmt"
	"log"
	"os"
	"time"
	"net"
)

// VerificarServidor intenta conectarse al servidor y devuelve true si está disponible.
func VerificarServidor(ip string, puerto int) bool {
	address := fmt.Sprintf("%s:%d", ip, puerto)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second) // Intentamos conectar con timeout de 2s
	if err != nil {
		return false // El servidor no está disponible
	}
	conn.Close() // Cerramos la conexión
	return true
}

func main() {
	utils.ConfigurarLogger()
	log.Println("Soy un Log")

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	if !VerificarServidor(globals.ClientConfig.Ip, globals.ClientConfig.Puerto) {
		log.Fatalf("Error: El servidor no está disponible en %s:%d", globals.ClientConfig.Ip, globals.ClientConfig.Puerto)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escribe un mensaje para enviar al servidor (presiona Enter sin escribir nada para salir):")

	for {
		fmt.Print("> ")
		scanner.Scan()
		mensaje := scanner.Text()

		// Si el usuario presiona Enter sin texto, salimos del bucle
		if mensaje == "" {
			log.Println("Finalizando envío de mensajes...")
			break
		}

		// Actualizamos la configuración global con el nuevo mensaje
		globals.ClientConfig.Mensaje = mensaje

		// Enviar el mensaje al servidor con el valor de la configuración actual
		utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

		// Generamos y enviamos el paquete con el mensaje
		utils.GenerarYEnviarPaquete()
	}

	log.Println("Programa finalizado.")
}
