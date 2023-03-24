package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	p := fmt.Println
	p("GoDotEnv package")

	p(os.Getenv("MODE"))
	p(os.Getenv("ENV_1"))
	p(os.Getenv("ENV_2"))

	// si quiero tomar las envs de un archivo en específico se debe poner la ruta del env caso contrario iría vacío
	if err := godotenv.Load("otherfolder/.var"); err != nil {
		p(err)
	}

	p(os.Getenv("MODE"))
	p(os.Getenv("ENV_1"))
	p(os.Getenv("ENV_2"))

	myEnvs, err := godotenv.Read("otherfolder/.var")
	if err != nil {
		p(err)
	}
	p(myEnvs)

	// Reemplaza los envs por los de default
	if err := godotenv.Overload(); err != nil {
		p(err)
	}

	p(os.Getenv("MODE"))
	p(os.Getenv("ENV_1"))
	p(os.Getenv("ENV_2"))

}
