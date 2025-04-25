package java_collections;

import java.util.ArrayList;
import java.util.List;

public class Listas {
    public static void main(String[] args) {
        // tipo da variável é mais comum ser a interface
        // mas o new será baseado na classe correspondente
        List<String> linguagens = new ArrayList<>();

        boolean adicionada = linguagens.add("Java");
        System.out.println("A linguagem Java foi adicionada na lista? " + adicionada);

        linguagens.add("CSharp");
        linguagens.add("Python");

        adicionada = linguagens.add("Java");
        System.out.println("A linguagem Java foi adicionada novamente na lista? " + adicionada);

        linguagens.add("JavaScript");
        linguagens.add("Go");
        linguagens.add("PHP");
        linguagens.add(4, "SQL");

        // pegadinha
        System.out.println("Qual o índice na lista da linguagem Java? " + linguagens.indexOf("Java"));

        System.out.println("Qual elemento está localizado no índice 3 (4-1)  " + linguagens.get(3));
        System.out.println("Qual elemento está localizado no índice 0 " + linguagens.get(0));

        boolean removida = linguagens.remove("Java");
        System.out.println("A linguagem Java foi removida? " + removida);

        // descomente a linha abaixo e execute mais uma vez
        // Collections.sort(linguagens);

        Object objetoRemovido = linguagens.remove(5);
        System.out.println("O objeto removido foi? " + objetoRemovido);

        System.out.println("***Listando as Linguagens***");
        for (Object linguagem : linguagens) {
            System.out.println(linguagem);
        }
    }
}
