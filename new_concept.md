Die Programmiersprache `Knife` baut auf einen Low_Level Data Interpreter auf - ist jedoch dynamisch ausführbar, und verfügt über eine Statisch-erweiterbare Typisierung.

### _Datentypen

| Typ  | Erklärung                                   | Beispiel                                                     |
| ---  | ---                                         | ---                                                          |
|`num` | Räpresentiert eine Gerade oder Kommazahl.   | `5`, `1.7644`, `.5`                                          |
|`int` | Räpresentiert eine Gerade nicht-komma Zahl. | `6`, `59`, `624320402`                                       |
|`dub` | Räpresentiert eine Kommazahl.               | `.76`, `1.7`, `56.34667`                                     |
|`str` | Räpresentiert einen String.                 | `"Hallo"`, `"Toll wie Toll"`, `"Tach ich bin ein String"`    |
|`boo` | Räpresentiert einen Boolean.                | `true`, `false`                                              |
|`cha` | Räpresentiert einen Char.                   | `'A'`, `'2'`                                                 |

Das waren jetzt die Normalen Datentypen. Natürlich gibt es auch noch sogenannte `List Types`, diese werden benutzt um die oben geannten Datentypen vereinfacht in einer Liste zu speichern und zu strukturieren.

| Typ  | Erklärung                                   | Beispiel                                                     |
| ---  | ---                                         | ---                                                          |
|`arr` | Räpresentiert ein Array - also eine Liste.  | `[4, 26, 91, 8, 44]`, `["Hallo", "Welt"]`                    |
|`vec` | Räpresentiert einen Vector - mehr-d-liste.  | `<'a': 1, 2, 3 | 'b': 3, 2, 1 | 'c': 2, 1, 3 >`              |
|`map` | Räpresentiert eine Map -                 .  | `{ }`                                                        |

#### _Arrays [arr]
Ein Array (`arr`) ist eine Aufzählung mehrerer Elemente. Ein Array umfasst meistens Werte des selben Typen.
```
$mein_array :arr = ["Hallo", "Welt"];
```

#### _Eigene Typen
```
type :looooong [":i32 + :i32"];
$zahl :num = 10.00;
$lang :looooong = $zahl .parseNum(1, 100*100*100, 0);
```

### _Parameter
```
f myFunction ( $one :num, $two :str, $thr :arr ) ... //normale Parameter
f myFunction ( $one :num = 10, $two :str = "Hallo", $thr :arr = [1, 2, 3] ) ... //default Werte
f myFunction ( ?$one :num, ?$two :str, ?$thr :arr) ... //optional Parameters
f myFunction ( $:num ) ... //unnamed Parameter
f myFunction ( $:num .. $:num) ... //Undefined nuimber of Parameters
f myFunction ( ? ) :auto {
    return _self[0] ?? "Not Found"; //returns first Parameter or Error-string
}
```