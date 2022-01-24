/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package AnalizadorPalabras;

/**
 *
 * @author David Maldonado
 */
public class GeneradorAnalizadorPas {
    public static void main(String[] args) {
        generarAnalizadorPas();
    }
    
    public static void generarAnalizadorPas() {
        try {
            String ruta = "src/AnalizadorPalabras/";
            String opcFlex[] = {ruta + "LexicoPalabras.jflex", "-d", ruta};
            jflex.Main.generate(opcFlex);
           
            String opcCUP[] = {"-destdir", ruta, "-parser", "SintacticoPalabras", ruta + "SintacticoPalabras.cup"};
            java_cup.Main.main(opcCUP);
            
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
