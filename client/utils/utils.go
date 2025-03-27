package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	config := &globals.Config{} // ✅ Inicializamos correctamente la estructura

	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error()) // Esto detendrá el programa si falla
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil { // Pasamos config sin el &
		log.Fatal("Error al decodificar config.json:", err)
	}

	return config
}


func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	text, _ := reader.ReadString('\n')
	log.Print(text)
}

func GenerarYEnviarPaquete() {
	// Crear un mensaje con los datos cargados de la configuración
	mensaje := Mensaje{
		Mensaje: globals.ClientConfig.Mensaje, // El mensaje se toma de la configuración
	}

	// Mostrar el paquete en los logs
	log.Printf("Paquete a enviar: %+v", mensaje)

	// Codificar el mensaje en JSON
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("Error al codificar el mensaje: %s", err.Error())
		return
	}

	// Enviar el mensaje al servidor
	url := fmt.Sprintf("http://%s:%d/mensaje", globals.ClientConfig.Ip, globals.ClientConfig.Puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error enviando el mensaje a %s:%d: %s", globals.ClientConfig.Ip, globals.ClientConfig.Puerto, err.Error())
		return
	}
	defer resp.Body.Close()

	// Registrar la respuesta del servidor
	log.Printf("Respuesta del servidor: %s", resp.Status)
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
