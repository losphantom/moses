package main

import (
    "net/http"
    "log"
    "encoding/json"
    "os/exec"
    "os"
    "fmt"
    "io"
)

func main() {
    http.HandleFunc("/relay", relayImage) 
    log.Println("server listening on 8087 ...")
    log.Fatal(http.ListenAndServe(":8087", nil))
}

/*
 {
   "image": "internet/xxx-api:branh-xxx",
   "relayReg": "10.1.1.55",
   "targetReg": "inner.harbor.com"
  }
*/
func relayImage(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var data map[string]string
    err := decoder.Decode(&data)
    if err != nil {
        log.Println(err)
        io.WriteString(rw, "error")
        return
    }
    image := data["image"]
    log.Println("-------------start for ", image, "---------------")
    fromImage := fmt.Sprintf("%s/%s", data["relayReg"], image)
    toImage := fmt.Sprintf("%s/%s", data["targetReg"], image)
    log.Println("*** 1. pull ", fromImage)
    cmd := exec.Command("docker", "pull", fromImage)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    log.Println("*** 2. tag ", fromImage, " ", toImage)
    cmd = exec.Command("docker", "tag", fromImage, toImage)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    log.Println("*** 3. push ", toImage)
    cmd = exec.Command("docker", "push", toImage)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    log.Println("-------------end of ", image, "---------------")
}
