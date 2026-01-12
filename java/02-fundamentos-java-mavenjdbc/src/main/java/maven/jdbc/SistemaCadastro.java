package maven.jdbc;

import java.sql.SQLException;

public class SistemaCadastro {

    public static void main(String[] args) throws SQLException {
        FabricaConexao.conectar();
        CadastroRepository repository = new CadastroRepository();

        Cadastro cadastro = repository.buscar(13);
        System.out.println(cadastro.getId() + cadastro.getNome() + cadastro.getIdade());
    }
}