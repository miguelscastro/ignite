package excecoes;

public class EstadoValidateException extends RuntimeException {
    public EstadoValidateException() {
        super("O estado não foi localizado");
    }
}
