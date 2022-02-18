grammar Gramatica;

options {
    language='Go';
}

@parser::header {
    import (
	    abs "practica3/Abstract"
	    exp "practica3/Expressions"
    )
}

// Gramatica
start returns [abs.Expression e]
    : init=exp EOF {
        $e = $init.e
        result := $e.Execute()
    	fmt.Println(result.Value)
    }
    ;

exp returns [abs.Expression e]
    : left=exp op=('*'|'/') right=exp {
        if $op.text == "*" {
            $e = exp.NewArithmetic(1, 1, $left.e, $right.e, exp.Multiplicacion)
        } else { 
            $e = exp.NewArithmetic(1, 1, $left.e, $right.e, exp.Division)
        }
    }
    | left=exp op=('+'|'-') right=exp {
        if $op.text == "+" {
            $e = exp.NewArithmetic(1, 1, $left.e, $right.e, exp.Suma)
        } else { 
            $e = exp.NewArithmetic(1, 1, $left.e, $right.e, exp.Resta)
        }
    }
    | ENTERO {
        $e = exp.NewLiteral(1, 1, abs.Int, $ENTERO.text)
    }
    | DECIMAL {
        $e = exp.NewLiteral(1, 1, abs.Double, $DECIMAL.text)
    }
    ;

// Tokens
PLUS: '+';
MINUS: '-';
TIMES: '*';
DIVISION: '/';
ENTERO: [0-9]+;
DECIMAL: [0-9]+('.'[0-9]+);
WS: [ \r\n\t]+ -> skip;