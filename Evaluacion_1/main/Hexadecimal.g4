
grammar Hexadecimal;

options {
    language='Go';
}

@parser::header {
    import (
        n "Evaluacion1/num"
    )
}

//Grammar
start 
    : total EOF {
        fmt.Println($total.res)
    }
    ;

total returns [float64 res]
    : l1=lista '.' l2=lista {
        $res = float64($l1.numero.GetValue1())+$l2.numero.GetValue2()
    }
    | l1=lista {
        $res = float64($l1.numero.GetValue1())
    }
    ;

lista returns [*n.Num numero]
    : x=lista y=digit {
        valor, _ := strconv.Atoi($y.s)
        valor1 := $x.numero.GetValue1()*16 + valor
        $numero = $x.numero
        $numero.SetAux($x.numero.GetAux()*0.0625)
        valor2 := $x.numero.GetValue2() + float64(valor)*$x.numero.GetAux()
        $numero.SetValue1(valor1)
        $numero.SetValue2(valor2)
    }
    | y=digit {
        valor, _ := strconv.Atoi($y.s)
        $numero = &n.Num{}
        $numero.SetValue1(valor)
        $numero.SetAux(0.0625)
        $numero.SetValue2(float64(valor)*0.0625)
    }
    ;

digit returns [string s]
    : '0' { $s = "0" } 
    | '1' { $s = "1" }
    | '2' { $s = "2" }
    | '3' { $s = "3" }
    | '4' { $s = "4" }
    | '5' { $s = "5" }
    | '6' { $s = "6" }
    | '7' { $s = "7" }
    | '8' { $s = "8" }
    | '9' { $s = "9" }
    | 'A' { $s = "10" }
    | 'B' { $s = "11" }
    | 'C' { $s = "12" }
    | 'D' { $s = "13" }
    | 'E' { $s = "14" }
    | 'F' { $s = "15" }    
    ;
    
// tokens
CERO: '0';
UNO: '1';
DOS: '2';
TRES: '3';
CUATRO: '4';
CINCO: '5';
SEIS: '6';
SIETE: '7';
OCHO: '8';
NUEVE: '9';
DIEZ: 'A';
ONCE: 'B';
DOCE: 'C';
TRECE: 'D';
CATORCE: 'E';
QUINCE: 'F';
PUNTO: '.';
WS: [ \r\n\t]+ -> skip;