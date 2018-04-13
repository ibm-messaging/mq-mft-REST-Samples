	/**
	 * <p>
	 * This sample applicaiton will take host name/IP, port number and the service (agent/transfer)
	 * </p>
	 * <dl>
	 * <dt>-host=name of the host</dt>
	 * <dd>Host name to which to connect - defaults to "localhost"</dd>
	 * <dt>-port=port</dt>
	 * <dd>Port number on which the api is exposed - defaults to "9080"</dd>
	 * <dt>service=name of the mft service. (for V9.0.5 support agent and transfer service)</dt>
	 * <dd>default service the application queries in agent list.</dd>
	 * <dd>
	 * 
	 * </dl>
	 */
	 

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"flag"
)

func main() {

var hostName string
flag.StringVar(&hostName, "host", "localhost", "host name")

var portNum string
flag.StringVar(&portNum, "port", "9080", "port number")

var serviceName string
flag.StringVar(&serviceName, "service", "agent", "name of the API service")

  flag.Parse()

	finalString := fmt.Sprintf("http://%s:%s/ibmmq/rest/v1/admin/mft/%s",hostName,portNum,serviceName)
	fmt.Println("Querying the service available at:")
	fmt.Println(finalString)


	response, err := http.Get(finalString)
	if err != nil {
	   fmt.Print(err.Error())
	   os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
