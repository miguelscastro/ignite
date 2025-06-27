package expressoes;

public class ExpressoesSimples {
    public static void main(String[] args) {
        String nome = "Miguel";
        Integer matricula = 123;
        Double salario = 2544.6789;
        String textoImpresso = String.format(
                "Olá %-15.15s, de matrícula %05d e salário %,.2f",
                nome, matricula, salario);
        System.out.println(textoImpresso);
    }
}
