package main

import (  
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "regexp"
    //"reflect"
)

var strings_list = make(map[int]string) //equivalen zu python dict
var ints_list = make(map[int]int) //equivalen zu python dict
var floats_list = make(map[int]float64) //equivalen zu python dict

var global_result = ""

var dict_counter = 0 //wird einfach hals "hochzähler" benutzt

func main() {  
    data, error := ioutil.ReadFile("test.lang") //Datei wird eingelesen
    if(error != nil) {
    	fmt.Println(error)
    }
    //fmt.Println(string(data)) //datei wird ausgegeben
    FindAllNestedMarks(string(data)) //function wird aufgerufen
}

func FindAllNestedMarks(data string) {
	data = strings.ReplaceAll(data, "\\'", " $0$vesk$nested$type$0 ") //jedes \' wird durch .. ersetzt
	data = strings.ReplaceAll(data, "\\\"", " $0$vesk$nested$type$1 ") //jedes \" wird durch .. ersetzt
    data = strings.ReplaceAll(data, "\\/", " $0$vesk$nested$type$3 ") //jedes \/ wird durch .. ersetzt
    data = strings.ReplaceAll(data, "\\\\", " $0$vesk$nested$type$4 ") //jedes \\ wird durch .. ersetzt
    FindAllStrings(string(data)) //function wird aufgerufen
}

func FindAllStrings(data string) {
    data_array := strings.Split(data, "\"") //data wird anahnd von " gesplittet
    //fmt.Println(data_array) //data_array ist das ergebnis der letzen zeile, und ist eine art array
    var data_counter = 1 //datacounter wird benutzt um jedes zweite in " stehende element zu finden
    for data_counter < len(data_array) { //solange data_counter kleiner ist als die anzahl an items in data_array
     //   fmt.Println(data_array[data_counter]) //jeder string wird ausgegeben
        strings_list[dict_counter] = data_array[data_counter] //jeder string wird zum dict hinzugefügt
        data = strings.Replace(data, "\""+strings_list[dict_counter]+"\"", "$1$vesk$strings$number$"+strconv.Itoa(dict_counter), 1) //die strings verden mit der vesk-id versehen
        data_counter = data_counter + 2 //counter +2
        dict_counter = dict_counter + 1 //couner +1
    }
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

    //fmt.Println(data)
    FindAllPrimaryBrackets(string(data))
}

func FindAllPrimaryBrackets(data string) {
    this := strings.Split(data, "(");
   // fmt.Println(this);

    for item, _ := range this {
        item = item
        re := regexp.MustCompile(`(\((?:\(??[^\(]*?\)))`);
        result := re.FindString(data);
     //   fmt.Println("RESULT:"+result)
        if(len(result) > 0) {
            target := WOBcompiler(string(result))
            data = strings.Replace(data, result, target, 1);
        }
    }
    //fmt.Println(data);
}

func WOBcompiler(data string) string {
    data = strings.ReplaceAll(data, "(", ""); //alle Klammern werden entfehrnt
    data = strings.ReplaceAll(data, ")", "");
    data_array := strings.Split(data, " ");
    this_ints_counter := 0
    this_floats_counter := 0
    this_strings_counter := 0
   // this_vars_counter := 0
    for item, _ := range data_array { 
        data_array[item] = strings.ReplaceAll(data_array[item], " ", ""); // leerzeichen werden ersetzt
        if(strings.Contains(data_array[item], "$1$vesk$ints") == true) {
            this_ints_counter += 1
        } else if(strings.Contains(data_array[item], "$1$vesk$floats") == true) {
            this_floats_counter += 1
        } else if(strings.Contains(data_array[item], "$1$vesk$strings") == true) {
            this_strings_counter += 1
        } else if(data_array[item] == "+" || data_array[item] == "-" || data_array[item] == "*" || data_array[item] == "/") {
            _ = 0
        }
    }
    //fmt.Println("Ergebniss:")
    //fmt.Println(data_array);

    //mathematische operartionen
    if(this_strings_counter > 0) { // wenn mindestens ein String in der Klammer vorgekommen ist
        for item, _ := range data_array { // für jedes item im data array
            if(data_array[item] != "+" && data_array[item] != "-" && data_array[item] != "*" && data_array[item] != "/" && data_array[item] != "") { //solanfge es kein + - usw ist
                if(strings.Contains(data_array[item], "$1$vesk$ints") == true) { // wenn es ein Int ist 
                    x := strings.Split(data_array[item], "$") // wird es hier 
                    ex, _ := strconv.Atoi(x[5])               // zu einem String
                    data_array[item] = strconv.Itoa(ints_list[ex]) // umgewandelt
    //                fmt.Println(data_array[item]) // debug
                } else if(strings.Contains(data_array[item], "$1$vesk$floats") == true) {
                    this_floats_counter += 1 //float soll hier zum string werden
                } else if(strings.Contains(data_array[item], "$1$vesk$strings") == true) {
                    x := strings.Split(data_array[item], "$") // wird es hier 
                    ex, _ := strconv.Atoi(x[5])               // bleibt string
                    data_array[item] = strings_list[ex] // umgewandelt
      //              fmt.Println(data_array[item]) // debug // string bleibt string
                }
                global_result += data_array[item]
            }    
        }

        strings_list[dict_counter] = global_result 
    }// else if(this_floats_counter > 0) {

    //} else if(this_ints_counter > 0) {

    //}

    result := "$1$vesk$string$number$"+strconv.Itoa(dict_counter) // nur für strings
    dict_counter += 1 // nur für strings
    fmt.Println("ENDE:")//debug
    fmt.Println(global_result)//debug
    global_result = "";
    return result
}

//maximale zahl: 1000000000000000000