package java_nio;

import java.io.File;
import java.io.IOException;

public class JavaFileIO {
    public static void main(String[] args) throws IOException {
        File diretorio = new File("c:\\Code\\rocketseat\\one\\java\\teste");
        System.out.println(diretorio.exists());
        if (!diretorio.exists()) {
            diretorio.mkdir();
            System.out.println("criado");
        }
        File arquivo = new File(diretorio, "teste.txt");
        System.out.println(arquivo.exists());

        arquivo.createNewFile();
        System.out.println(arquivo.exists());
    }
}
