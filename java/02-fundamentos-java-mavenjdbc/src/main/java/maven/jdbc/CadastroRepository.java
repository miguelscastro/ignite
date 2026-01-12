package maven.jdbc;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

public class CadastroRepository {
    private Connection conexao;

    public CadastroRepository() {
        conexao = FabricaConexao.getConexao();
    }

    public void salvar(Cadastro cadastro) throws SQLException {
        String sql = "INSERT INTO public.cadastro (nome, idade) VALUES (?, ?)";
        PreparedStatement pst = conexao.prepareStatement(sql);
        pst.setString(1, cadastro.getNome());
        pst.setInt(2, cadastro.getIdade());
        pst.executeUpdate();
    }

    public void alterar(Cadastro cadastro) throws SQLException {
        String sql = "UPDATE public.cadastro set nome = ?, idade = ? WHERE Id = ?";
        PreparedStatement pst = conexao.prepareStatement(sql);
        pst.setString(1, cadastro.getNome());
        pst.setInt(2, cadastro.getIdade());
        pst.setInt(3, cadastro.getId());
        pst.executeUpdate();
    }

    public void excluir(Integer id) throws SQLException {
        String sql = "DELETE from public.cadastro WHERE id = ?";
        PreparedStatement pst = conexao.prepareStatement(sql);
        pst.setInt(1, id);
        pst.executeUpdate();
    }

    public List<Cadastro> listar() throws SQLException {
        List<Cadastro> lista = new ArrayList<>();
        String sql = "SELECT * from public.cadastro";
        PreparedStatement pst = conexao.prepareStatement(sql);
        ResultSet cadastros = pst.executeQuery();
        while (cadastros.next()) {
            Integer id = cadastros.getInt("id");
            String nome = cadastros.getString("nome");
            Integer idade = cadastros.getInt("idade");

            Cadastro cadastro = new Cadastro();
            cadastro.setId(id);
            cadastro.setNome(nome);
            cadastro.setIdade(idade);

            lista.add(cadastro);
        }
        return lista;
    }

    public Cadastro buscar(Integer id) throws SQLException {
        String sql = "SELECT * from public.cadastro WHERE id = ?";
        PreparedStatement pst = conexao.prepareStatement(sql);
        pst.setInt(1, id);
        ResultSet result = pst.executeQuery();

        if (result.next()) {
            Integer idInt = result.getInt("id");
            String nome = result.getString("nome");
            Integer idade = result.getInt("idade");

            Cadastro cadastro = new Cadastro();
            cadastro.setId(idInt);
            cadastro.setNome(nome);
            cadastro.setIdade(idade);
            return cadastro;
        }
        return null;

    }
}
