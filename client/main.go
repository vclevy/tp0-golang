package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Soy un Log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil { //TODO cuando ocurre?
		log.Fatalf("No se pudo cargar la configuración")
	}

	// loggeamos el valor de la config
	log.Println("Mensaje desde config.json:", globals.ClientConfig.Mensaje)

	// enviar un mensaje al servidor con el valor de la config
	log.Println("Enviando mensaje al servidor...")
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	paquete := utils.LeerConsola() //Genero un paquete que posteriormente le mando a GenerarYEnviarPaquete
	utils.GenerarYEnviarPaquete(paquete,globals.ClientConfig.Ip,globals.ClientConfig.Puerto)

}