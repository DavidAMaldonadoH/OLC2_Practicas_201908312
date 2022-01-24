/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package AnalizadorPalabras;

import java.io.BufferedReader;
import java.io.StringReader;

/**
 *
 * @author David Maldonado
 */
public class AnalizadorPalabras {
    public static AnalizadorPalabras analizadorPas;
    
    public static AnalizadorPalabras getAnalizadorPas() {
        if (analizadorPas == null) {
            analizadorPas = new AnalizadorPalabras();
        }
        return analizadorPas;
    }
    
    public static void AnalizarCodigo(String entrada) {
        try {
            System.out.println("> Entrada: " + entrada);
            
            LexicoPalabras scanner = new LexicoPalabras(new BufferedReader(new StringReader(entrada)));
            SintacticoPalabras parser = new SintacticoPalabras(scanner);
            parser.parse();
            
            System.out.println("> Análisis realizado con exito!\n");
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
