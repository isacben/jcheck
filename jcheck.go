package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
    jsonData := []byte(`{
    "hello":"world",
    "name": "isaac",
    "age": 23,
    "object": {
        "one": 1,
        "two": 2,
        "bool": false
    }
    "type": "some data",
    "array": [1, 2, 3, 4]
}`)
    
    var data interface{}
    decoder := json.NewDecoder(bytes.NewReader(jsonData))
    err := decoder.Decode(&data)
    if err != nil {
        if syntaxErr, ok := err.(*json.SyntaxError); ok {
            line := strings.Count(string(jsonData[:syntaxErr.Offset]), "\n") + 1
            fmt.Printf("syntax error at line %d: %s\n", line, err)
        } else if err == io.ErrUnexpectedEOF {
            line := strings.Count(string(jsonData[:]), "\n") + 1
            fmt.Printf("syntax error at line %d: %s\n", line, err)
        } else {
            fmt.Println("error:", err)
        }
        os.Exit(1)
    }

    fmt.Printf("jcheck: valid!\n%+v\n", data)
}
