package java_nio;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

public class JavaNIO {
    public static void main(String[] args) {
        try {
            Path path = Paths.get("c:\\Code\\rocketseat\\one\\java\\teste\\teste.txt");

            // Lendo o path e convertendo todos os caracteres (bytes) de uma só vez
            byte[] bytesArquivo = Files.readAllBytes(path);

            // Como são bytes podemos criar uma String a partir de agora
            String imprimirConteudo = new String(bytesArquivo);
            System.out.println(imprimirConteudo);

            // lê o conteudo convertido para String
            String lerConteudo = Files.readString(path);
            System.out.println(lerConteudo);

            // decodifica os bytes do arquivos em caracteres usando UTF-8 e exibe no console
            List<String> arrayDeConteudos = Files.readAllLines(path);
            arrayDeConteudos.forEach(linha -> {
                System.out.println(linha);
            });

            // Como é adicionado byte a byte ele escreve 1 a 1 no arquivo passado no path
            String adicionarConteudo = "Lua";
            Files.write(path, adicionarConteudo.getBytes(StandardCharsets.UTF_8));

            List<String> nomes = new ArrayList<>();
            nomes.add("Pessoa1");
            nomes.add("Pessoa2");

            StringBuilder novoConteudo = new StringBuilder();
            nomes.forEach(n -> {
                novoConteudo.append(n + "\n");
            });
            Files.write(path, novoConteudo.toString().getBytes(StandardCharsets.UTF_8));
            System.out.println(novoConteudo);

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
