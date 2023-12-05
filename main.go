package main

import (

        "os"
	"log"

	"github.com/gin-gonic/gin"
	)


func main() {


        if len(os.Args) == 1 {
                log.Fatalln("No configuration file specified.")
        }

        log.Println("Firing up Maxmind-API-Proxy!")

	 /* Load configuration into "global" memory. */

	 Load(os.Args[1])

	 Redis_Init()

	 log.Println("Connect to Redis.")

	 log.Printf("Setting gin to \"%s\" mode.", Config.Http_Mode)

	 gin.SetMode(Config.Http_Mode)


	 router := gin.Default()

	 /* Always force API authentication! */
	 
	 router.Use(Authenticate_API())

	 router.GET("/:ip_address", Maxmind_Query_IP)
	 
	 log.Printf("Listening for TLS traffic on %s.", Config.Http_Listen)

	 err := router.RunTLS(Config.Http_Listen, Config.Http_Cert, Config.Http_Key)

        if err != nil {
                log.Fatalf("Cannot bind it %s or cannot open %s or %s.\n", Config.Http_Listen, Config.Http_Cert, Config.Http_Key)
        }

}
