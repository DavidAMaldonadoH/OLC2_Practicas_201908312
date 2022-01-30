/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package practica_2;

import Analizador.Analizador;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Stream;

/**
 *
 * @author David Maldonado
 */
public class Practica_2 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        
        Analizador.getAnalizadorBin();
        Analizador.AnalizarCodigo(archivoToString("entrada.txt"));
    }
    
    public static String archivoToString(String filePath) {
        StringBuilder contentBuilder;
        String archivoString = "";
        try {
            Stream<String> stream = Files.lines(Paths.get(filePath), StandardCharsets.UTF_8);
            contentBuilder = new StringBuilder();
            stream.forEach(s -> contentBuilder.append(s));
            archivoString = contentBuilder.toString();
        } catch (IOException e) {
            System.err.println("Error: " + e.getMessage());
        }
        return archivoString;
    };
}
