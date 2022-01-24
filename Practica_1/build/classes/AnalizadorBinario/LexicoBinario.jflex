package AnalizadorBinario;

import java_cup.runtime.Symbol;

%% 
%class LexicoBinario
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

%%

"1" { return new Symbol(sym.UNO, yyline, yycolumn, yytext()); }
"0" { return new Symbol(sym.CERO, yyline, yycolumn, yytext()); }

\n { yycolumn = 1; }

{WhiteSpace} {}

. {
    System.out.println("Este es un error lexico: "+yytext()+
    ", en la linea: "+yyline+", en la columna: "+yycolumn);
}