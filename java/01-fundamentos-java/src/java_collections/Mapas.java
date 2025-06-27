package java_collections;

import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.Map;

public class Mapas {
    public static void main(String[] args) {
        Map<String, String> estadosBrasileiro = new HashMap<>();
        estadosBrasileiro.put("SP", "SAO PAULO");
        estadosBrasileiro.put("PI", "PIAUI");
        estadosBrasileiro.put("MA", "MARANHAO");
        estadosBrasileiro.put("AM", "AMAZONAS");
        estadosBrasileiro.put("BA", "BAHIA");
        estadosBrasileiro.put("GO", "GOIAS");

        Iterator<String> ufsIterator = estadosBrasileiro.keySet().iterator();
        while (ufsIterator.hasNext()) {
            Object key = ufsIterator.next();

            Object value = estadosBrasileiro.get(key);

            System.out.println("Sigla: " + key + " Nome: " + value);
        }

        // outro exemplo
        // Map e Enum são contexto diferentes
        // Avalie cada contexto e usabilidade

        Map<String, String> estadoCivil = new LinkedHashMap<>();
        estadoCivil.put("S", "SOLTEIRO(A)");
        estadoCivil.put("C", "CASADO(A)");
        estadoCivil.put("D", "DIVORCIADO(A)");
        estadoCivil.put("V", "VIUVO(A)");

        Iterator<String> estadoIterator = estadoCivil.keySet().iterator();
        while (estadoIterator.hasNext()) {
            Object key = estadoIterator.next();

            Object value = estadoCivil.get(key);

            System.out.println("Sigla: " + key + " Nome: " + value);
        }
    }
}