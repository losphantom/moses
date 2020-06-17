package main
import (
    "io/ioutil"
    "log"
    "fmt"
    "flag"
    yaml "gopkg.in/yaml.v2"
)
func main() {
    filePath := flag.String("yml", "./bootstrap.yml", "specify bootstrap.yml file path!")
    flag.Parse()

    resultMap := make(map[string]interface{})
    yamlFile, err := ioutil.ReadFile(*filePath)
   
    //log.Println("yamlFile:", yamlFile)
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, &resultMap)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    springMap := resultMap["spring"].(map[interface{}]interface{})
    profilesMap := springMap["profiles"].(map[interface{}]interface{})
    //log.Println(profilesMap["include"])
    profiles := profilesMap["include"]
    if profiles != nil {
        fmt.Print(profiles)
    }
}
