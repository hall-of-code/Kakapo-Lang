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
" //thats a multiline string, wich complete ignores newline. You only can write multiline strings in " not in '.
```
When you'll need newline, you could use `_generateSynExpr(<multilineString>, LN_BREAK, "\n")`.
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
So you see: With the function `_stateExt(<object>)` you could get the status of an Objects "main definition space" - and not it's default Referenced Expression State.

### _Arrays and Lists
In Knife there is no explecit difference between an List and an Dict/map/hash/object - you could use every array as an Dict also as an simple List.
**[Understand why]:**
Imagine you have a list of the Age of some anonymous People:
`[25, 71, 28, 65, 43]`
Internal this Array is describeable like that: 
```php
    [
        1 : 25,
        2 : 71,
        3 : 28,
        4 : 65,
        5 : 43,
    ]
```
When you now want to access one of this `Items` you've to do: `$variable[2]` to access the second Key in the Array.  
Next we want to store the Ages with the People Names. For that we have to set unique `<Keys>` wich is bindet to the Value.
```php
    [
        Piet : 25,
        Gundula: 71,
        Bob : 28,
        Berbel : 65,
        Xatar : 43,
    ]
```
When we now want to access one of the Values in the Array, we could do it like this: `$variable['Berbel']` (= 65).  
Also you could mix it, in that case you have to know: `num` based keys are automated in the correct direction - that means as follows: 
```php
    !! INCORRECT !!
    [
        Piet : 25,
        Gundula : 71,
        1 : 28,
        Berbel : 65,
        2 : 43,
    ]
    
    The right way: (!)
    [
        Piet : 25,
        Gundula : 71,
        3 : 28,
        Berbel : 65,
        5 : 43,
    ]
```
Also Arrays supporting nesting (of course :D).  
This was a little example:  
```php
    [
        abc : (4 + 25),
        efg : "Hello",
        hij : [1, 2, 44, 82],
        lel : [
            x1: "hello",
            x2: "world",
            x3: "!",
        ],
    ]
```

Also Knife has an JSON inline support, but its recommendet to format it into Native-array because the JSON parser's   performance isnt nice :) Read more later...  


### _Maps
`Maps` are the "Professional Version of Arrays" -> when you understand what I mean :)  
While an `Array` has optional Static Types -> a `Map` is based on Static Types.  
| Array | Map                                                        |  
| :----  | ----:                                                     |  
| `"hello", "world", "!" ]`| `string[ "hello", "world", "!" ]`       |  

A Maps basic definition is like that `[ <storage> ]<content>`.  
So when you want a small list, it's enought to define it so:  
```php
$myCoolMap = []string :map;
//and then store something in it:
$myCoolMap << "Hallo"; //same as [ 1 : "Hallo", ];
$myCoolMap << "Test";
$myCoolMap << "Lalalala";
// same as [1: "Hallo", 2: "Test", 3: "Lalalala"]
```
And now one difficult example:  
```php
$myProfMap = [string][string][int] :map;
// Thats generates something like that:
    [
        <KEY string> : [<key string> : <int> ]
    ]
    
//or in filled form:
    [
        example: [
            a: 7,
            b: 24,
            c: 15,
        ],
    ]
```

So now we have all basic Types explained :)  

## Variables
A Variable is a small storage to store some data in it, especially give data a name.  
In Knife there are a few stack of different Variable Types:    
[ **Normal** | Avaible where its defined and lower level. ]  
```php
$x = "Hello"; //variable $x has now the value "Hello".

//recommendet, static typed.
$x :string = "Hello";

//automatic detect datatype
$x := "Hello"; //string
```
A defined variable can get changed, but not the Type! So a `string` cant switched to `num` or `array`. To do that, you could
use `?`, example: `$x :string = "Hey;     $x ?= 25;` OR simply define your own dynamic Type, described later...  
[ **Global** | Avaible everywhere (in its own module). ]  
```php
global $g = "Heyy"; //variable $g has now the value "Heyy".

//recommendet
global $g :string = "Heyy"; //string
```

[ **Private** | Avaible where its defined and lower level. (like default) ]  
```php
private $p = "Hi"; //variable $p has now the value "Hi".

//recommendet
global $p :string = "Hi"; //string
```

[ **Once** | Only onetime defineable - and never changeable  ]
```php
once $o = "Allo"; //variable $o is now "Allo", and cant updated any more.  
global once $go = "Höllö"; //global once
//recommendet
once $o :string = "Allo"; //string
```
(Equivalent to `const in JS`)
**Now the Special Variable Types:**
Some Variable Types are really rare used, so they are only for special cases.  
[ **freeze** | Freeze a Variable at it's current point ]  
```php
$f = "hi";
$f = "bye";
freeze $f; 
$f = "waaaa"; //$f is now "bye" because its not possible to update it, after freeze  
```
You also can unfreeze a freezed variable, this you could do with `thaw`. `thaw $f = "nice!";  //$f is now "nice!"`.  

[ **universe** | Universe make a Variable accessable from everywhere in the current runtime ]  
```php
universe $u = "Im am here!";  

//you now can call it from every includet file or module. But it could overwritten easily.  
```

`unset` a Variable is also possible, just use `unset $variable;` to undefine it.  

## Functions




























d