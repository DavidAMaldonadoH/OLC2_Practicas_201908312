/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package Analizador;

import java.io.BufferedReader;
import java.io.StringReader;

/**
 *
 * @author David Maldonado
 */
public class Analizador {
    public static Analizador analizador;
    
    public static Analizador getAnalizadorBin() {
        if (analizador == null) {
            analizador = new Analizador();
        }
        return analizador;
    }
    
    public static void AnalizarCodigo(String entrada) {
        try {
            System.out.println("> Entrada: " + entrada);
            
            Lexico scanner = new Lexico(new BufferedReader(new StringReader(entrada)));
            Sintactico parser = new Sintactico(scanner);
            parser.parse();
            
            System.out.println("> Análisis realizado con exito!\n");
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
