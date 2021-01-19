# Knife Programming Language Concepts
The Knife Programming Language is an Project by myself. The Language isnt developed for know, its planned to build an
Interpreter in V, Go or Rust.

## Comments
Comments in **Knife** are completly ignored by the Compiler/Interpreter. So you can write what ever you want.
```php
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
Every `num` has two _Parts_: The NUMERIC and the DECIMAL area. The NUMERIC number represents the Full Number, while 
DECIMAL represents everything behind the `.` ("comma"). Every `num` in Knife is representet like that, also when you 
write a "non-comma"-Number.
```php
25       //represents 25.0
1        //represents 1.0
-18      //represents -18.0
0.5      //represents 0.5
.8       //represents 0.8
4.45563  //represents 4.45563
inf      //represents --0.0             = Infinity
imf      //represents -0.            = -Infinity
nan      //represents .              = non Math
```

### _Strings
Strings in Knife can written in `''` or `""` - Even when last one is recommendet. (There are two differences between them):
```php
"Hello, Wolrd!" //a normal string
'Hello, Wolrd!' //also a normal string

"
Hello
World
!
" //thats a multiline string, wich replaces newline in \n. You only can write multiline strings in " not in '.
```
Also its supportet to escape some Strings. Based on a 'Backslash' :
```php
'Hello " ' //no escape needet.
"ha ' hu '" //also no escape needet.

"Hello \" lalala" //escape needet. 
```

### _Bools and Unsets
A Bool represents the current state of an `OBJECT` (`OBJCET` = Every Code Argument). The Bool could have the States `True` or `False`
(case sensitiv). By default every `Object` is `False`, until its defined. Lets imagine this case:
```php
output( _state($x) ); //the native function _state() returns the current state of an object. This returns FALSE

$x = "Hello"; //lets define $x as the String "Hello".
output( _state($x) ); //now the functions returns us True because $x is *not* undefined.
```
So you'll see, that there is no "null" or "undefined". Only there is an "extra Param" wich is assigned to an `Object` when it's not defined. This "extra Param" called `unset` -> But where we need this case?
```php
//lets imagine we have an expression wich is assigned to a variable:
$expr = 25 > 26; //because 25 is not greater than 26, this expression is false. So $expr is `false` even it's defined.

/*
Whats the solution? -> Yes! The "unset Param".
*/
output( _stateExt($expr) ); // returns True. because it's defined.
```






























d