/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package practica_2;

/**
 *
 * @author David Maldonado
 */
public class Num {
    private int valor1;
    private double valor2;
    private double aux;

    public Num() {
        this.valor1 = 0;
        this.valor2 = 0.0;
        this.aux = 0.0;
    }

    public int getValor1() {
        return valor1;
    }

    public void setValor1(int valor1) {
        this.valor1 = valor1;
    }

    public double getValor2() {
        return valor2;
    }

    public void setValor2(double valor2) {
        this.valor2 = valor2;
    }

    public double getAux() {
        return aux;
    }

    public void setAux(double aux) {
        this.aux = aux;
    }
   
}
