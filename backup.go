package main

import (  
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

var strings_list = make(map[int]string) //equivalen zu python dict
var ints_list = make(map[int]int) //equivalen zu python dict
var floats_list = make(map[int]float64) //equivalen zu python dict

func main() {  
    data, error := ioutil.ReadFile("test.lang") //Datei wird eingelesen
    if(error != nil) {
    	fmt.Println(error)
    }
    fmt.Println(string(data)) //datei wird ausgegeben
    FindAllNestedMarks(string(data)) //function wird aufgerufen
}

func FindAllNestedMarks(data string) {
	data = strings.ReplaceAll(data, "\\'", " $0$vesk$nested$type$0 ") //jedes \' wird durch .. ersetzt
	data = strings.ReplaceAll(data, "\\\"", " $0$vesk$nested$type$1 ") //jedes \" wird durch .. ersetzt
    data = strings.ReplaceAll(data, "\\/", " $0$vesk$nested$type$3 ") //jedes \/ wird durch .. ersetzt
    data = strings.ReplaceAll(data, "\\\\", " $0$vesk$nested$type$4 ") //jedes \\ wird durch .. ersetzt
	fmt.Println(data)
    FindAllStrings(string(data)) //function wird aufgerufen
}

func FindAllStrings(data string) {
    data_array := strings.Split(data, "\"") //data wird anahnd von " gesplittet
    fmt.Println(data_array) //data_array ist das ergebnis der letzen zeile, und ist eine art array
    data_counter := 1 //datacounter wird benutzt um jedes zweite in " stehende element zu finden
    dict_counter := 0 //wird einfach hals "hochzähler" benutzt
    for data_counter < len(data_array) { //solange data_counter kleiner ist als die anzahl an items in data_array
        fmt.Println(data_array[data_counter]) //jeder string wird ausgegeben
        strings_list[dict_counter] = data_array[data_counter] //jeder string wird zum dict hinzugefügt
        data = strings.Replace(data, "\""+strings_list[dict_counter]+"\"", "$1$vesk$strings$number$"+strconv.Itoa(dict_counter), 1) //die strings verden mit der vesk-id versehen
        data_counter = data_counter + 2 //counter +2
        dict_counter = dict_counter + 1 //couner +1
    }
    fmt.Println(data)
    AbstandHalten(string(data))

}

func AbstandHalten(data string) {
    data = strings.ReplaceAll(data, ",", " , ")
    data = strings.ReplaceAll(data, ";", " ; ")
    data = strings.ReplaceAll(data, "(", " ( ")
    data = strings.ReplaceAll(data, ")", " ) ")
    data = strings.ReplaceAll(data, "{", " { ")
    data = strings.ReplaceAll(data, "}", " } ")
    data = strings.ReplaceAll(data, "/", " / ")
    data = strings.ReplaceAll(data, "%", " % ")
    data = strings.ReplaceAll(data, "+", " + ")
    data = strings.ReplaceAll(data, "-", " - ")
    data = strings.ReplaceAll(data, "*", " * ")
    data = strings.ReplaceAll(data, "[", " [ ")
    data = strings.ReplaceAll(data, "]", " ] ")
    data = strings.ReplaceAll(data, ":", " : ")
    data = strings.ReplaceAll(data, "=", " = ")
    FindAllInts(string(data)) //nächste function
}

func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func FindAllInts(data string) {
    data_array := strings.Split(data, " ") //splitet an jedem leerzeichen
    hoch_counter := 0 //einfahcer counter
    dict_counter := 0
    float_counter := 0
    for hoch_counter < len(data_array) {
        if _, err := strconv.Atoi(data_array[hoch_counter]); err == nil { //check ob string nur zahlen von 1-9 enthält. bzw mit einem int hgleichzusetzen ist
            ints_list[dict_counter], err = strconv.Atoi(data_array[hoch_counter])
            data = strings.Replace(data, " "+data_array[hoch_counter], " $1$vesk$ints$number$"+strconv.Itoa(dict_counter)+" ", 1)
            dict_counter = dict_counter + 1
        } else if (isNumeric(data_array[hoch_counter]) == true) {
            floats_list[dict_counter], err = strconv.ParseFloat(data_array[hoch_counter], 64)
            data = strings.Replace(data, " "+data_array[hoch_counter], " $1$vesk$floats$number$"+strconv.Itoa(float_counter)+" ", 1)
            float_counter = float_counter + 1
        }
        hoch_counter = hoch_counter + 1

    }

    fmt.Println(data)
}

//maximale zahl: 1000000000000000000