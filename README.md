## Krakend-Generator
A tool for generating krakend configuration file with ease.

as we all know it is a good practice to generate the krakend configuration in a container and make it immune to further changes.
so instead of writing the json file by hand (which easily you can get lost between definitions), or using krakend-designer (which shows a lot of useless settings) maybe the best choice is to write some codes which generate the config.json file for us.
so you can simply generate the `krakend.json` file by using the defined structs. 
#### Installation
we support go modules so for installing you can run:
```
go get -u github.com/p3ym4n/krakend-generator
```
#### Usage
a simple route definition using this tool is like below:
```go
import kgen "github.com/p3ym4n/krakend-generator"

func main() {
    app := kgen.Default()
    
    app.AddEndpoints(
        kgen.NOOPEndpoint(kgen.GET, "/hello",
            kgen.NOOPBackend(kgen.GET, "github.com", "/hello"),
        ),
    )
    
    // config.json file is the path you want for the file
    if err := app.Generate("config.json"); err != nil {
        log.Fatalf("error occurred: %s",err)
    } else {
        log.Println("config.json generated.")
    }
}
```

for more information please check the documentation on [go.dev](https://pkg.go.dev/github.com/p3ym4n/krakend-generator) 

 

