package main

import (
	"log"

        "encoding/json"
        "io/ioutil"

        )


type Configuration struct {

        API_Key string

        Http_Listen string
        Http_Cert   string
        Http_Key    string
        Http_Mode   string

        Maxmind_Username        string
        Maxmind_Password        string
        Maxmind_Url             string

        Redis_Enabled           bool
        Redis_Host              string
        Redis_Port              int
        Redis_Password          string
        Redis_Database          string
        Redis_Cache_Time        int

}

var Config *Configuration

func Load( ConfigFile string ) {

        json_file, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
        log.Fatalln("Cannot open %s.\n", ConfigFile)
	}

        err = json.Unmarshal(json_file, &Config)

	if err != nil { 
	log.Fatalln("Cannot parse configuration file %s", ConfigFile)
	}

	/* Sanity Check */

        if Config.API_Key == "" {
                log.Fatalln("Cannot find 'api_key' in %s.", ConfigFile)
        }

        if Config.Http_Listen == "" {
                log.Fatalln("Cannot find 'http_listen' in %s.", ConfigFile)
        }

        if Config.Http_Cert == "" {
                log.Fatalln("Cannot find 'http_cert' in %s.", ConfigFile)
        }

        if Config.Http_Key == "" {
                log.Fatalln("Cannot find 'http_key' in %s.", ConfigFile)
        }

        if Config.Http_Mode == "" {
                log.Fatalln("Cannot find 'http_mode' in %s.", ConfigFile)
        }

        if Config.Http_Mode != "release" && Config.Http_Mode != "debug" && Config.Http_Mode != "test" {
        log.Fatalf("Invalid 'http_mode' :  %s.  Valid 'http_modes' are 'release', 'debug' and 'test'\n", Config.Http_Mode)
        }

	if Config.Maxmind_Username == "" {
		log.Fatalln("Cannot find 'maxmind_username' in %s.", ConfigFile)
		}

	if Config.Maxmind_Password == "" {
		log.Fatalln("Cannot find 'maxmind_password' in %s.", ConfigFile)
		}

	if Config.Maxmind_Url == "" {
		log.Fatalln("Cannot find 'maxmind_url' in %s.", ConfigFile)
		}

	if Config.Redis_Host == "" {
		log.Fatalln("Cannot find 'redis_host' in %s.", ConfigFile)
		}

	if Config.Redis_Port == 0 {
		log.Fatalln("Cannot find 'redis_port' in %s.", ConfigFile)
		}

	if Config.Redis_Password == "" {
		log.Fatalln("Cannot find 'redis_password' in %s.", ConfigFile)
		}

	if Config.Redis_Database == "" {
		log.Fatalln("Cannot find 'redis_database' in %s.", ConfigFile)
		}

	/* Cache_Time is in hours */

	if Config.Redis_Cache_Time == 0 {
		log.Fatalln("Cannot find 'redis_cache_time' in %s.", ConfigFile)
		}





}

