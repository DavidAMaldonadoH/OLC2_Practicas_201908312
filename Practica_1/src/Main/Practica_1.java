/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package Main;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Stream;

import AnalizadorBinario.AnalizadorBinario;
import AnalizadorPalabras.AnalizadorPalabras;

/**
 *
 * @author David Maldonado
 */
public class Practica_1 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        AnalizadorPalabras.getAnalizadorPas();
        AnalizadorPalabras.AnalizarCodigo(archivoToString("entradaPalabras.txt"));
        
        AnalizadorBinario.getAnalizadorBin();
        AnalizadorBinario.AnalizarCodigo(archivoToString("entradaBinario.txt"));
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
