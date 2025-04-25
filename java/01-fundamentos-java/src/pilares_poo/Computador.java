package pilares_poo;

public class Computador {
    public static void main(String[] args) {

        SistemaDeMensagem smi = null;

        String appEscolhido = "???";

        if (appEscolhido.equals("msn"))
            smi = new MSNMessenger();
        else if (appEscolhido.equals("fbm"))
            smi = new FacebookMessenger();
        else if (appEscolhido.equals("tlg"))
            smi = new Telegram();

        smi.enviarMensagem();
        smi.receberMensagem();
    }
}
