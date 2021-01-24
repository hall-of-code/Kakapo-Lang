# Knife Concepts
This is an better Version of Knife - an overviewable Concept Description of an easy readable `not Indention-level` and `not only bracket-based`.

## Comments
Comments are complete Ignored by the Compiler/Interpreter.
```php
//this is a singleline Comment

/*
This is a multiline Comment.
This is a new Line in this Comment.
*/

/* Fixed singleline Comment */
```

## Datatypes
In **Knife** there few different Datatypes, wich are described as follow:

| Datatype | Definition                                        | Example                                        |
| :-----   | :----                                             | :----                                          |
| `multi`  | Everything, no specific Type defined.             |  Could be String, or Num or Bool ...           |
| `num`    | Number, it combines _integer_ and _float_ in one. |  ``55``, `4.30`, `.25`                         |
| `string` | String, a sequence of given Characters.           |  `"Hello"`, `"Test Test Jes"`,  `'Hello'`      | 
| `bool`   | State, either `True` or `False`.                  |  `True`, `False`                               |
| `array`  | A group/list of multiple Values / Key:Values.     |  `[1,2,3]`, `["hello", 55.5, "sealion lel"]`   |
| `map`    | A group/list of same Datatypes Values.            |  `["hello", "lol", "77xd"]string`              |

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
imf      //represents -0.0            = -Infinity
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
So you'll see, that there is no "null" or "undefined". Every undefined Object is `False`.

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
| `"hello", "world", "!" ]`| `[ "hello", "world", "!" ]string`       |  

A Maps basic definition is like that `[ <definition> ]<type>`.  
So when you want a small list, it's enought to define it so:  
```php
$myCoolMap = []string :map;
//and then store something in it:
$myCoolMap << "Hallo"; //same as [ 1 : "Hallo", ]string;
$myCoolMap << "Test";
$myCoolMap << "Lalalala";
// same as [1: "Hallo", 2: "Test", 3: "Lalalala"]string
```
And now one difficult example:  
```php
$myProfMap :map = [ []num ]string;
// Thats generates something like that:
    [
        <KEY string> : [<num>] :num
    ]
    
//or in filled form:
    [
        "example": [ 1, 2, 3, ],
        "some": [ 55.5, 25.25, 18, ],
    ]
    
!! Another Example : !!
$someMap :map = [ num ]string;

//is filled:
[
    "test": 44,
    "some": 5,
]

!! Another Example : !!
$someMap :map = [ [ num ]string ]string;

//is filled:
[
    "test": [
        "lel": 100,
    ],
    "test2": [
        "lel2": 200,
    ],
]

```

So now we have all basic Types explained :)  

## Variables
You could define two types of Variables: `Private` and `Global`, by default a Variable is `Private`. We recommended to define a Variables name with an `$`. But you also could name your Variable `test` instead of `$test`. 
#### _Private Variables
Private Variables are only avaible at the Level where its defined, and Lower. A private Variable can everytime changed to Global and back.
```php
$some   = "hello"; //unspecified Type, type: multi
$some2 := "hello"; //auto set type. type: string
$some3 :string = "hello"; //manual defined type. type: string
```
Also you could write `'private'` before the definition. Like that:
```php
private $some :string = "hello";
```

### _Global Variables
A Global Variable is avaible in the whole document from everywhere. A Global Variable could be overwritten with a Private Variable. Also you could move a `Global` to a `Private` Variable. Global Variables where defined with the `'global'` Keyword.
```php
global $g  = "hello world!"; //unspecified Type, type: multi
global $g := "hello world!"; //auto set Type, type: string
global $g :string = "hello world!"; //manual set Type, type: string
```
 ### _Extra Parameters
 There are also some extra Parameters wich could be assigned to a Variable. So when you would like to set a Variable once, and make it non-changeable - you could do this with `'once'`.
 ```php
 once $hello :string = "Hello"; //$hello has now the string Value "Hello"
      $hello = "Bye";           //this DOESNT WORK because you never can update an "once"-Variable any more.
      
 global once $example :num = 55; //you also could define multiple Paramters. So a Variable could be global - and once.
 ```

 There are a few more extra Params, wich can be assigned to a Variable:
 
 **The `freeze` Param:** That Param is simmilar to the `once` Param, but could be removed dynmaicly with `thaw`.  
 _Example_: 
 ```php
 $hello :string = "hello"; //define $hello
 $hello = "bye"; //update $hello
 freeze $hello; //freeze $hello at current state
 $hello = "lol"; //try to update it but it's not updating. So $hello is "bye" until you thaw it.
 
 //now we want to thaw it.
 thaw $hello;
 $hello = "nice";  //now we could update it again. $hello is now "nice".
 ```
 
--------
 **The `universe` Param:** That Param is a bit like the `global` Param, but it makes a Variable avaible over every Includet Module.   
 _Example_: 
 ```php
 File main.kf :
    import myModule as $module;
    universe $pro :string = "Hello World";  //Define Variable
    print($pro);                            //Prints out "Hello World!"
    
 File module.kf :
    load::universe;
    print($pro); //Prints out "Hello World!"
    
 ```
 
--------
 **The `unset` Param:** That Param undefines a Variable. So after that, the Variable is removed. But could defined after that again new whenever you would like to do that.     
 _Example_: 
 ```php
$hello := "Hello";
print($hello); //Prints out "Hello".

unset $hello;
print($hello); //Triggers an Error because the Variable '$hello' is not existing anymore.
```
 
--------
--------
So you'll see -> A Variable is really Dynamicly. So, except `once`, every Variable is updateable and Changeable.
Here a few Examples for more understanding:
```php
private $cool;
$yes = "yeess";

$cool :string, $more_cool :string = "Cool", "Cooler"; //define quick multiple Variables.
private $some :string, global $other :num, once $x :num = "hello", 5, 91.2;

/*
When you want to define multiple statements automaicly you could use the ':=' only for last one.
*/
$one, $two := "Yes"; //$one is defined as type: multi    | and $two is autodefined to type: string | Both has Value "Yes";
$one :auto, $two := "Yes";  //$one  is auto-defined as type: string | $two is also auto-defined to type: string | Both has value "Yes";

/*
When you want to define an Array with its items to multiple Variables, you could do that also like so:
*/
$first, $second, $third = ["item1", "item2", "item3"]; //thats maps $first to "item1", and $third to "item3" ... and so on.
$first, $second, $third = ["item1", "item2", "item3", "item4"]; //"item4" is never assigned in this example.

//but what - when you would like to assign the whole array to every of the Variables?
$first, $second, $third = (["item1", "item2", "item3"]) as ?instance; //now $first, $second and $third has the whole array each.


//The Or-expression:
private $some; //$some exists but not assigned.
$top = $some || "default Value"; //now when $some is not defined, $top gets the Value "Default Value";
//also you can use 'OR' instead of '||'. Further you could use so many || easypes es you want.
$top = $some || $another || $yetanother || "default value";

/*The Referencing System.
Basicly a Variable is storing the current state of its definion:
Lets take a look:
*/
$x = 25;
$y = $x;
$x = 62;
print($y); //This Prints out "25" because: at the time where $y was defined, $x has the Value 25.

/* Now we could also use References, thats only new Names for something.*/
$x = 25;
$y &= $x;                                 // or $y = &$x; 
$x = 62;
print($y); //now it prints out "62", because: $y references the $x Variable.
```
-----------

## Functions and Areas
In Knife Code-Blocks are defined as `<Area>` -> That means a Block **isn't** defined by `Indention Level` (like Python) - and also it **is optional** definded by `{}` like in JS, PHP or Golang. In Knife - we're using `<begin>` and `<end>` Keywords.
For don't scare off PHP and JS Developers to learn *Knife* - we've implemented `{}` as well. (Our Interpreter/Optimizer replaces `{` with "start:", and also `}` with ":end").
Because you can't define a Function in a Function (only unnamed and callback is allowed) -- you dont need to closeup every single Functions Tag when you're define the next Function.  

Functions are declared by `func` keyword - followed by Name, arguments and returntype. `func <name> (<params>) type: FUNCTION`.  
```php
func some($x :num, $y :num) num:
    return $x + $y;
:end

!! OR !!

func some($x :num, $y :num) :num {
    return $x + $y;
}

//Best Practice of multi define functions
func some($x :num, $y :num) num:
    return $x + $y;
    
func sub($x :num, $y :num) num:
    return $x - $y;
    
func mult($x :num, $y :num) num:
    return $x * $y;
    
:end
//now we have defined 3 Functions.
```
To call that function, you use: `<functionname>()`.
```php
$z = sub(19, 10); //$z is now 9.
```
**Functions** are private by default, to make it Avaible from other Files/Modules/Includes - we are using the `public` or shorter `pub` keyword.   
`pub func some() :string {...}`.

---
#### _Areas
An Area begins with `begin:` and ends up with `:end` **OR** `{` and `}`. We are using it in different ways.
For example in Loops ect.
---
## Loops
Loops are relativ the same as other Languages.
#### _For Loops
Basic Construction of a For Loop is this one.
`for( <define>, <run>, <target> ) ...`
So the Loop runs so long, as the Target is not reached, or sb call `break`.
```php
for( $f = 25, $f -= 1, $f == 0 ) {
    print("Lol"); 
}
```
This Loop, defines `$f = 25` once. Then It calls the `Loops Body` and the Loops `instruct` (`$f -= 1`) so long as the Expression `f == 0` is `false`. 

#### _Fixed Loops
Fixed Loops has a fixed number of Loop-rythms the have to make.
`loop(<count>, <optional expression>) ...`
```php
loop(10) {
    print("hello");
}
```
This example prints out "hello" ten times.  
Another Example:
```php
$number := 25;
$counter := 0;
loop($number, $counter += 1) {
    print($counter);
}
```
This Example prints out 1 then 2 then 3, then 4 ... 25 and stops execuing.  
Manipulating Fixed Loops? Yes thats also possible, the `<count>` is everytime overwriteable:
```php
private $data :array = [1, 2, 2, 2, 1, 2, 1];
private $c :num = $data.len;
loop ( $c ) {
    if( $data[0] == 1 ) {
        unset $data[0]; //remove current item from array ($data)
    }
    else
    {
        $data[0] -= 1;
        $c += 1;
    }
}
```
So that Loops runs 11 times. Because it runs so long as every Instance of $data is removed.

#### _While Loops
While Loops runs so long as it's Expression is `true`. `while ( <expression> ) ... `
```php
$a :num = 50;
while ( $a >= 1 ) {
    $a -= 1;
}
```
That Loops runs 50 Times, because then $a is 0 , and the Expression is `false`.

#### _Each Loops
Each Loops are implemented to Loop thought an Array/Map with nested Elements. `each (<array | map> as <reference> ) ...`
```php
$store :array = [
                    alfred: [first: "alfred", last: "gringer", age: 45],
                    gundel: [first: "gundel", last: "crisps", age: 59],
                    flash: [first: "flash", last: "grifindor", age: 17],
                ];

each ( $store as $item )
{
    print( _str($item.age) + $item.first);
}

```
      
      

