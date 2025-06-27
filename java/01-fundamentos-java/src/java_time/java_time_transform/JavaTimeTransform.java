package java_time.java_time_transform;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class JavaTimeTransform {
    public static void main(String[] args) {
        // LocalDateTime to LocalDate
        LocalDateTime dataHora = LocalDateTime.of(2023, 1, 24, 16, 31, 27);
        LocalDate data = dataHora.toLocalDate();

        // LocalDateTime to LocalTime
        LocalDateTime dataHorario = LocalDateTime.of(2023, 1, 24, 16, 31, 27);
        LocalTime hora = dataHorario.toLocalTime();
        System.out.println(hora);

        // LocalDate to LocalDateTime
        LocalDate horarioData = LocalDate.of(2023, 1, 24);
        LocalDateTime horaData = horarioData.atTime(16, 31, 27); // 2023-01-24T16:31:27
        horaData = data.atTime(LocalTime.of(16, 31, 27)); // 2023-01-24T16:31:27
        horaData = data.atTime(LocalTime.of(16, 31)); // 2023-01-24T16:31
        horaData = data.atStartOfDay(); // 2023-01-24T00:00
        System.out.println(horaData);
    }
}
