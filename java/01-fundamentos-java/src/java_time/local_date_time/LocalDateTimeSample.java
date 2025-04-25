package java_time.local_date_time;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class LocalDateTimeSample {
    public static void main(String[] args) {
        // instanciação
        // 2023-06-25T16:25 -> data e hora sem segundo
        LocalDateTime dataHora = LocalDateTime.of(2023, 6, 25, 16, 25);
        // 2023-06-25T16:25:33 -> data e hora com segundo
        dataHora = LocalDateTime.of(2023, 6, 25, 16, 25, 33);
        System.out.println(dataHora);
        System.out.println("-------");

        // Manipulação
        LocalDateTime dataHorario = LocalDateTime.of(2023, 6, 25, 16, 25);
        LocalDateTime proximoMes = dataHorario.plusMonths(1);
        LocalDateTime dataHorarioComMenos15Minutos = dataHorario.minusMinutes(15);
        LocalDateTime dataHorarioDefinindoSegundoZero = dataHorario.withSecond(33);
        System.out.println(proximoMes);
        System.out.println(dataHorarioComMenos15Minutos);
        System.out.println(dataHorarioDefinindoSegundoZero);
        System.out.println("-------");

        // Comparação
        LocalDateTime dataHora1 = LocalDateTime.of(2023, 6, 25, 16, 25);
        LocalDateTime dataHora2 = LocalDateTime.of(2023, 6, 25, 16, 25, 17);
        // comparando dois objetos data/hora em sua totalidade
        boolean dataHoraIgual = dataHora1.equals(dataHora2); // false
        System.out.println(dataHoraIgual);
        // comparando dígito específico de uma data e hora
        boolean dataHoraDigitoIgual = dataHora1.getDayOfMonth() == dataHora2.getDayOfMonth(); // true
        System.out.println(dataHoraDigitoIgual);
        System.out.println("-------");

        // Formatação
        LocalDateTime horaData = LocalDateTime.of(2023, 7, 22, 17, 33);
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("dd/MM/yy HH:mm:ss");
        String horaDataFormatada = horaData.format(formatter);
        System.out.println(horaDataFormatada); // 22/07/23 17:33:00
    }
}
