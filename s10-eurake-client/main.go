package main

import (
	"fmt"
	eureka "github.com/xuanbo/eureka-client"
	"net/http"
)

func main() {
	// create eureka client
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://127.0.0.1:8502/eureka/",
		App:                   "gomoudle",
		Port:                  8506,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
	// start client, register、heartbeat、refresh
	client.Start()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// full applications from eureka server
		//apps := client.Applications
		//b, _ := json.Marshal(apps)
		name := "hello welcome go moudle"
		_, _ = writer.Write([]byte(name))
	})

	// start http server
	if err := http.ListenAndServe(":8506", nil); err != nil {
		fmt.Println(err)
	}
}
