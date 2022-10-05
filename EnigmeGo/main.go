package main



import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math/rand"
    "time"

)

func main() {
	readFile, err := os.Open("File.txt")
	if err != nil {
			log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lignes []string
	var result []string
	
	
	
	var tmp int
	var mot string
	for fileScanner.Scan(){
		lignes = append(lignes, fileScanner.Text())


	}
	readFile.Close()



	result = append(result, lignes[0])
	result = append(result, lignes[len(lignes)-1])
	for i, ligne := range lignes {
		
		if i > 1 {
			if lignes[i-1] == "before"{
				tmp, _ = strconv.Atoi(ligne)
				tmp = tmp - 1
				result = append(result, lignes[tmp])
				
			}
		}
		if i < len(lignes)-2 {
			if strings.Contains(lignes[i+1], "now"){
				mot = ligne
				for p, lettre := range mot {
					if p == 1 {

						result = append(result, lignes[int(lettre)/len(lignes)-1])
					}
					
				}
			}
				
			
		}
		
		


	}
	



fmt.Println(result)
PrintRandomNumber()
}



func PrintRandomNumber(){
	rand.Seed(time.Now().UnixNano())
    min := 0
    max := 100
    fmt.Println(rand.Intn(max - min + 1) + min)
}


