# Knife Programming Language Concepts
The Knife Programming Language is an Project by myself. The Language isnt developed for know, its planned to build an
Interpreter in V, Go or Rust.

## Comments
Comments in **Knife** are completly ignored by the Compiler/Interpreter. So you can write what ever you want.
```
//this is a single line comment

/*
This is a multiline 
comment!
*/


/*
[Description of a Feature]
This is the best practice example to write readable comments.
*/
```

## Datatypes
In **Knife** there few different Datatypes, wich are described as follow:

| Datatype | Definition                                        | Example                                        |
| :-----   | :----                                             | :----                                          |
| `num`    | Number, it combines _integer_ and _float_ in one. |  ``55``, `4.30`, `.25`                         |
| `string` | String, a sequence of given Characters.           |  `"Hello"`, `"Test Test Jes"`,  `'Hello'`      | 
| `bool`   | State, either `True` or `False`.                  |  `True`, `False`                               |
| `unset`  | Expremiental for non-defined Values.              |  `unset`                                       |
| `array`  | A group/list of multiple Values / Key:Values.     |  `[1,2,3]`, `["hello", 55.5, "sealion lel"]`   |
| `map`    | A group/list of same Datatypes Values.            |  `string["hello", "lol", "77xd"]`              |

### _Nums
Eevery `num` has two _Parts_: The NUMERIC and the DECIMAL area. The NUMERIC number represents the Full Number, while 
DECIMAL represents everything behind the `.` ("comma"). Every `num` in Knife is representet like that, also when you 
write a "non-comma"-Number.
```
25       //represents 25.0
1        //represents 1.0
0.5      //represents 0.5
.8       //represents 0.8
4.45563  //represents 4.45563
inf      //represents 0.             = Infinity
imf      //represents -0.            = -Infinity
nan      //represents .              = none
```





























d