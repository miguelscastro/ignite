package java_collections;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Generics {
    public static void main(String[] args) {
        List<String> prateleira = new ArrayList<>();
        prateleira.add("Jeep");
        prateleira.add("Ford Ka");
        prateleira.add("Gol");
        prateleira.add("Wolkswagen");

        Collections.sort(prateleira);
        Collections.reverse(prateleira);
        Collections.shuffle(prateleira);

        for (String carrinho : prateleira) {
            System.out.println(carrinho);
        }
    }

}
