package java_time.local_time;

import java.time.LocalTime;
import java.time.format.DateTimeFormatter;

public class LocalTimeSample {
    public static void main(String[] args) {
        // Instanciação
        LocalTime horaAtual = LocalTime.now();
        LocalTime horaMinutoSegundoEspecifico = LocalTime.of(13, 21, 42); // 13:21:42
        LocalTime horaMinutoEspecifico = LocalTime.of(13, 21); // 13:21
        LocalTime horaDeUmaString = LocalTime.parse("13:21:42");
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("HHmmss");
        LocalTime horaDeUmaStringDespadronizada = LocalTime.parse("132142", formatter); // 13:21:42
        System.out.println(horaAtual);
        System.out.println(horaMinutoSegundoEspecifico);
        System.out.println(horaMinutoEspecifico);
        System.out.println(horaDeUmaString);
        System.out.println(horaDeUmaStringDespadronizada);
        System.out.println("-------");

        // Manipulação
        LocalTime horarioAtual = LocalTime.now();
        LocalTime horaAtualMais30Minutos = horarioAtual.plusMinutes(30);
        LocalTime horaAtualMenos2Horas = horarioAtual.minusHours(2);
        System.out.println(horarioAtual);
        System.out.println(horaAtualMais30Minutos);
        System.out.println(horaAtualMenos2Horas);
        System.out.println("-------");

        // Comparação
        LocalTime horaAtualmente = LocalTime.now();
        LocalTime horaAtualmenteMenos2Horas = horaAtual.minusHours(2);
        boolean passado = horaAtual.isAfter(horaAtualmenteMenos2Horas);
        System.out.println(horaAtualmente);
        System.out.println(passado);
        System.out.println("-------");

        // Formatação
        LocalTime hora = LocalTime.of(13, 21, 42);
        DateTimeFormatter formatterHora = DateTimeFormatter.ofPattern("ss-mm-HH");
        String horaFormatada = hora.format(formatterHora);
        System.out.println(horaFormatada); // 42-21-13

        // Dados específicos
        LocalTime horario = LocalTime.of(13, 21, 42);

        System.out.println(horario.getHour()); // 13
        System.out.println(horario.getMinute()); // 21
        System.out.println(horario.getSecond()); // 42
        System.out.println(horario.getNano()); // 0
    }
}
