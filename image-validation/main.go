package main 

import ( 
    "bufio"
    "fmt"
    "log"
	"os"
    "encoding/json"
    "net/http"
    "errors"
) 

var catalog = make(map[string][]string)
var cache = make(map[string]int)

func main() { 
  
    // os.Open() opens specific file in  
    // read-only mode and this return  
    // a pointer of type os. 
    file, err := os.Open("dump/input-dump") 
  
    if err != nil { 
        log.Fatalf("failed to open") 
  
    } 
  
    // The bufio.NewScanner() function is called in which the 
    // object os.File passed as its parameter and this returns a 
    // object bufio.Scanner which is further used on the 
    // bufio.Scanner.Split() method. 
    scanner := bufio.NewScanner(file) 
  
    // The bufio.ScanLines is used as an  
    // input to the method bufio.Scanner.Split() 
    // and then the scanning forwards to each 
    // new line using the bufio.Scanner.Scan() 
    // method. 
    scanner.Split(bufio.ScanLines) 


    for scanner.Scan() { 
		var line = scanner.Text()
        var data map[string]interface{}
        json.Unmarshal([]byte(line), &data)
        
        populateCatalog(data["productId"].(string), data["image"].(string))
    } 
    file.Close() 

    final()
}

func imageStatus(img string) (int, error) {
    resp, err := http.Get(img)    
    resp.Body.Close()

    if err != nil {
        return 0, errors.New("Server error")
    }

    cache[img] = resp.StatusCode
    return resp.StatusCode, nil

}

func populateCatalog(id string, img string) {
    catalog[id] = append(catalog[id], img)
}

func final() {

    for id, images := range catalog {
        var final_images []string

        for _,image := range images {
            
            if len(final_images) == 3{
                break

            } else if val, ok := cache[image]; ok {
                if val == 200 {
                    final_images = append(final_images, image)
                }
            } else {
                status, _ := imageStatus(image)

                if status == 200 {
                    final_images = append(final_images, image)
                }
            }   
        }

    type Product struct {
        ProductId string
        Images []string
    }
    res2B, _ := json.Marshal(Product{id, final_images})
    fmt.Println(string(res2B)) 
        
    }

}