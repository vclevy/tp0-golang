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
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() Paquete { //LeerConsola no tiene parametros por eso vacio los () y devuelve Paquete por eso al lado eso
	// Crear el paquete para almacenar los mensajes
	paquete := Paquete{}

	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")

	for {
		text, _ := reader.ReadString('\n')
		log.Println(text)

		if text == "\r\n" { // Si el usuario presiona Enter sin escribir nada, termina. IMPORTANTE: PARA PROBAR EN WINDOWS, SACAR \r !!!
			break
		}

		paquete.Valores = append(paquete.Valores, text) // Guardar en el paquete
	}

	return paquete
}

 func GenerarYEnviarPaquete(paquete Paquete,ip string, puerto int) {

	if len(paquete.Valores) == 0 {
		log.Println("No se ingresaron mensajes para enviar.")
		return
	}

	log.Printf("Paquete a enviar: %+v", paquete)
	EnviarPaquete(ip, puerto, paquete) // Enviar el paquete al servidor
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