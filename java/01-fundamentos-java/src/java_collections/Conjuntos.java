package java_collections;

// import java.util.ArrayList;
import java.util.Collection;
// import java.util.HashSet;
// import java.util.List;
import java.util.Set;
import java.util.TreeSet;

public class Conjuntos {
    public static void main(String[] args) {
        Set<String> inscritos = new TreeSet<>();
        inscritos.add("Marcos");
        inscritos.add("Lucas");
        inscritos.add("Antonio");
        inscritos.add("Mirela");
        inscritos.add("Alessandra");
        inscritos.add("Felipe");
        inscritos.add("Sofia");

        imprimirInscritos(inscritos);
        /**
         * Alessandra
         * Felipe
         * Mirela
         * Sofia
         * Marcos
         * Lucas
         * Antonio
         */
    }

    // não se assuste com o nível de abstração, em breve você vai entender
    static void imprimirInscritos(Collection<String> inscritos) {
        System.out.println("***Listando os inscritos***");
        for (Object inscrito : inscritos) {
            System.out.println(inscrito);
        }
    }
}