package equals;

import equals.carro.Carro;

public class Equals {
    public static void main(String[] args) {
        String string = "miguel";
        String string1 = "miguel";

        /*
         * vai ser true porque ambos estão na stack e a JVM entende que ambos
         * correspondem
         * ao mesmo objeto com uma única string de valor miguel
         */
        System.out.println(string1 == string);

        Integer i1 = 128;
        Integer i2 = Integer.parseInt("128");

        /*
         * compara se os objetos referenciam o mesmo valor, mas como fazem de locais
         * diferentes, da false
         */
        System.out.println(i1 == i2);

        /*
         * retorna true pois compara se os objetos e seus valores são iguais entre sí e
         * não se
         * referenciam o mesmo endereço
         */
        System.out.println(i1.equals(i2));

        Carro c1 = new Carro("palio");
        Carro c2 = new Carro("palio");

        /*
         * vai dar false já que embora tenham valores iguais são objetos em endereços de
         * memória distintos, sendo necessário comparar se os valores dos objetos são
         * iguais.
         */
        System.out.println(c1.equals(c2));

        /*
         * após sobrescrever o equals() e hashCode() retorna true, pois mesmo que sejam
         * objetos
         * armazenados em endereços de memória diferentes e apontem para atributos de
         * mesmo valor
         * porém armazenados em locais diferentes seus valores são iguais e é isso que o
         * equals()
         * sobrescrito faz
         */
        System.out.println(c1.equals(c2));
    }
}
