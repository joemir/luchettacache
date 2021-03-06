package main
 

import (
 "io/ioutil"
	"time"
	"io"
	"strings"
	"fmt"
	"net/http"
"net"
)



var (
cacheInstance Cache
)

func init(){

cacheInstance = NewRedisCache()

}



func createCacheKey(req *http.Request) (string){
	
	url := req.URL.String();

    host  :=req.Host;

    servletName := "/cache"

	return host +  url[strings.Index(url, servletName)+len(servletName):len(url)-1];
	}
	
	

func CacheHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Server", "luchettaCache BETA 0.0.1")
	
  
	key := createCacheKey(req)

    //cacheInstance.Put(key, []byte(key))
 
	contentCached := cacheInstance.Get(key);

	if(contentCached != nil){
		w.Header().Set("Content-Length", string(len(contentCached)))
		io.WriteString(w, string(contentCached))		
	}else
	{
		  newContent  := proxy(key)
		 cacheInstance.Put(key, newContent)
		 w.Header().Set("Content-Length", string(len(newContent)))
		io.WriteString(w, string(newContent))
	}


	
}

func proxy(url string) []byte {

          OKAddr := "200.221.2.45"
         OKAddress, _ := net.ResolveTCPAddr("tcp", OKAddr)

         transport := &http.Transport{
                 Proxy: http.ProxyFromEnvironment,
                 Dial: (&net.Dialer{
                         Timeout:   60 * time.Second,
                         KeepAlive: 60 * time.Second,
                         OKAddr: OKAddress}).Dial, TLSHandshakeTimeout: 10 * time.Second}

         client := &http.Client{
                 Transport: transport,
         }
     fmt.Println("proxy to "+url)
	resp, err := client.Get("http://"+url)
   if err != nil {
   	fmt.Println("Error on proxy to "+url)
	return nil
   }

defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

  return body
} 

func main() {

	serverTeste := &http.Server{
		Addr:           ":8082",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
	}

	http.HandleFunc("/cache/", CacheHandler)
	serverTeste.ListenAndServe()

	//log.Degub("Server started")

}


   