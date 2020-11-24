# Kakapo-Lang
The Kakapo Scripting/Programming Languge.

This Language is a Hobby Project by myself and its not working yet ^^

## Info
The Syntax of Kakapo is a bit Javascript-like and inspired. It completly compiles to fast Golang, wich makes Kakapo verry performant. 

```
//This is a Comment in Kakapo;

//
This is a
Multiline Comment 
in Kakapo;

// *Author: [...];


//Data Types:
in Kakapo there are 3 different Base-Types:
String: String
Num: Number/Int or Float
Boolean: True/Flase
Null: Null/Nil/undefined/nothing
;


//Defining Variables;
var $test = "Hello, World!"; // this variable is privat and only in that case avaible, where its defined;
$test = "Hello, World!"; // this variable is privat and only in that case avaible, where its defined;
global $test = "Hello, World!"; // this ist a global variable, wich is avaible from everywhere;
private $test = "Only here"; // this variable is privat and only in that case avaible, where its defined;
once $test = "Nice"; // this variable can't be Updated or Modifyed after once defined;

$string = "Hello World!";
$string = "Hello "+"World";
$num = 16 + 5;
$num = 14.5 + .5;
$array = ["Arg1", 55.5, "test"+"lol", 33];
$object = {"Lol": "test"};
$construct = construct({"data" => type.int, "alter" => type.int, "name" => type.str});

//defining operations;
+=     //add VAL2 to VAL1;
-=     //substract VAL2 form VAL1;
.=     //synchronize Floating Value from VAL2;
/=     //VAL1 / VAL2;
*=     //VAL1 * VAL2;

//functions;
func myFunction($arg)
{
   $arg = $arg.split(":")[0];
   return ($arg);
};

function my_Cool_Box($x, $y)
{
   $this.x = $x;
   $this.y = $y;
   $this.constructor = { //the $this.constructor represents a Small function-based Function-constrcutor;
     if([$this.x, $this.y] != type.int) return;
   };
   $string = "Width: "+$this.x+" and Hiehgt is "+$this.y;
   return ($string);
};

//unnamed functions;

```
