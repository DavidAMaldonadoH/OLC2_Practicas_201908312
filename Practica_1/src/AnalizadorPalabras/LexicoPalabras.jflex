package AnalizadorPalabras;

import java_cup.runtime.Symbol;

%% 
%class LexicoPalabras
%public
%line
%column
%cup
%unicode

%init{ 
    yyline = 1; 
    yycolumn = 1; 
%init} 

WhiteSpace = [ \t\r\f\n]+
Words = [\w]+

%%

"," { return new Symbol(sym.COMA, yyline, yycolumn, yytext()); }
"." { return new Symbol(sym.PUNTO, yyline, yycolumn, yytext()); }

\n { yycolumn = 1; }

{WhiteSpace} {}
{Words}      { return new Symbol(sym.WORD, yyline, yycolumn, yytext()); }

. {
    System.out.println("Este es un error lexico: "+yytext()+
    ", en la linea: "+yyline+", en la columna: "+yycolumn);
}
