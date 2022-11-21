package main

import (


    "github.com/gocolly/colly"
	"encoding/csv"
    "log"
    "os"
)



func main() {
    fName := "data.csv"
    file, err := os.Create(fName)
    if err != nil {
        log.Fatalf("Could not create file, err: %q", err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()
	c := colly.NewCollector()
    c.OnHTML("table.wikitable", func(e *colly.HTMLElement) {
        e.ForEach("tr", func(index int, el *colly.HTMLElement) {
            if (index == 0){
                writer.Write([]string{
                    el.ChildText("th:nth-child(1)"),
                    
                    el.ChildText("th:nth-child(2)"),
                    el.ChildText("th:nth-child(3)"),		
                })
            } else {
                writer.Write([]string{
                    el.ChildText("th:nth-child(1)"),
                    
                    el.ChildText("td:nth-child(2)"),
                    el.ChildText("td:nth-child(3)"),		
                })
            }
            
        })

    })
    c.Visit("https://es.wikipedia.org/wiki/Copa_Mundial_de_F%C3%BAtbol#Resultados_y_estad%C3%ADsticas")
}