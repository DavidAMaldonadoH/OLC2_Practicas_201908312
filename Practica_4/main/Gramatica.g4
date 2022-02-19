grammar Gramatica;

options {
    language='Go';
}

@parser::header {
    import (
        m "practica4/Coordinate"
        arrayList "github.com/colegno/arraylist"
    )
}

//Grammar
start returns [*arrayList.List lista]
    : coordinates EOF {
        $lista = $coordinates.l
    }
    ;

coordinates returns [*arrayList.List l]
    @init {
       $l =  arrayList.New()
    }
    : e += coordinate*  {
        listCoord := localctx.(*CoordinatesContext).GetE()
            for _, e := range listCoord {
                $l.Add(e.GetC())
            }
    }
    ;

coordinate returns [m.Coordinate c]
    : '(' left=NUM ',' right=NUM ')' {
        $c = m.NewCoordinate($left.text, $right.text)
    }
    ;

// tokens
LPAR: '(';
RPAR: ')';
COMMA: ',';
NUM: '-'?[0-9]+;
WS: [ \r\n\t]+ -> skip;