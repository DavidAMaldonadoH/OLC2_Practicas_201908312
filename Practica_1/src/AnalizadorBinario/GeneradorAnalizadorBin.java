/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package AnalizadorBinario;

/**
 *
 * @author David Maldonado
 */
public class GeneradorAnalizadorBin {
    public static void main(String[] args) {
        generarAnalizadorBin();
    }
    
    public static void generarAnalizadorBin() {
        try {
            String ruta = "src/AnalizadorBinario/";
            String opcFlex[] = {ruta + "LexicoBinario.jflex", "-d", ruta};
            jflex.Main.generate(opcFlex);
           
            String opcCUP[] = {"-destdir", ruta, "-parser", "SintacticoBinario", ruta + "SintacticoBinario.cup"};
            java_cup.Main.main(opcCUP);
            
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
