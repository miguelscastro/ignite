package excecoes;

public class Excecoes {
    public static void main(String[] args) {
        try {
            checkEstado("MA");
        } catch (EstadoValidateException e) {
            System.out.println(e.getMessage());
            System.out.println("Selecione outro estado");
        } catch (Exception e) {
            System.out.println(e.getMessage());
            System.out.println("Selecione um estado diferente");
        }
    }

    static void checkEstado(String nomeEstado) throws EstadoValidateException, Exception {
        if (nomeEstado.equalsIgnoreCase("FL")) {
            throw new Exception("Exceção");
        }
        if (!nomeEstado.equalsIgnoreCase("PI")) {
            throw new EstadoValidateException();
        }

        System.out.println("Bem vindo ao (a): " + nomeEstado.toUpperCase());

    }
}
