/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package AnalizadorBinario;

import java.io.BufferedReader;
import java.io.StringReader;
/**
 *
 * @author David Maldonado
 */
public class AnalizadorBinario {
    public static AnalizadorBinario analizadorBin;
    
    public static AnalizadorBinario getAnalizadorBin() {
        if (analizadorBin == null) {
            analizadorBin = new AnalizadorBinario();
        }
        return analizadorBin;
    }
    
    public static void AnalizarCodigo(String entrada) {
        try {
            System.out.println("> Entrada: " + entrada);
            
            LexicoBinario scanner = new LexicoBinario(new BufferedReader(new StringReader(entrada)));
            SintacticoBinario parser = new SintacticoBinario(scanner);
            parser.parse();
            
            System.out.println("> Análisis realizado con exito!\n");
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
    
}
