package maven.jdbc;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.Properties;

public class FabricaConexao {
    private static Connection conexao;

    public static void conectar() {
        try {
            if (conexao == null) {
                String url = "jdbc:postgresql://localhost:5432/maven_jdbc";
                Properties props = new Properties();
                props.setProperty("user", "postgres");
                props.setProperty("password", "root");
                // props.setProperty("ssl", "true");
                conexao = DriverManager.getConnection(url, props);
            }

        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public static Connection getConexao() {
        return conexao;
    }
}
