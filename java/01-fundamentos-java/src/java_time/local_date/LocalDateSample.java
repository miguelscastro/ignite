package java_time.local_date;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.time.format.FormatStyle;
import java.util.Locale;

public class LocalDateSample {
    public static void main(String[] args) {
        // LocalDate
        // retorna a data atual aaaa/mm/dd
        LocalDate dataAtual = LocalDate.now();
        // permite determinar a data
        LocalDate meuAniversario = LocalDate.of(2005, 5, 14);
        System.out.println(dataAtual);
        System.out.println(meuAniversario);
        System.out.println("-------");

        // DateTimeFormatter
        String stringDataBr = "17/01/2023";
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("dd/MM/yyyy");
        LocalDate dataConcreta = LocalDate.parse(stringDataBr, formatter);
        System.out.println(dataConcreta);
        System.out.println("-------");

        // Manipulação de datas
        LocalDate aniversarioIzabelly = LocalDate.of(2023, 5, 3);
        System.out.println(aniversarioIzabelly);
        LocalDate novaData = aniversarioIzabelly.minusDays(7);
        System.out.println(novaData);
        LocalDate dataQueTera15Anos = aniversarioIzabelly.plusYears(15);
        System.out.println(dataQueTera15Anos);
        System.out.println("-------");

        // Comparação de datas
        LocalDate data1 = LocalDate.of(2023, 3, 4);
        LocalDate data2 = LocalDate.of(2024, 4, 3);
        // data2 está após a data 1? true
        System.out.println(data2.isAfter(data1));
        // data2 está antes que data 1? flase
        System.out.println(data2.isBefore(data1));
        // data2 é igual a data 1? false
        System.out.println(data2.isEqual(data1));
        System.out.println("-------");

        // Formatação de datas
        LocalDate data = LocalDate.of(2023, 1, 13);
        DateTimeFormatter dateFormatter = DateTimeFormatter.ofPattern("dd/MM/yyyy");
        System.out.println(data); // 2023-01-13
        // formatando a data criada acima
        System.out.println(dateFormatter.format(data)); // 13/01/2023
        System.out.println("-------");

        LocalDate dataPorLocal = LocalDate.of(2023, 1, 13);
        // 13 de janeiro de 2023
        String dataFormatada = dataPorLocal.format(DateTimeFormatter.ofLocalizedDate(FormatStyle.LONG));
        System.out.println(dataFormatada);
        // 13 de jan. de 2023
        dataFormatada = data.format(DateTimeFormatter.ofLocalizedDate(FormatStyle.MEDIUM));
        System.out.println(dataFormatada);
        // 13/01/2023
        dataFormatada = data.format(DateTimeFormatter.ofLocalizedDate(FormatStyle.SHORT));
        System.out.println(dataFormatada);
        // sexta-feira, 13 de janeiro de 2023
        dataFormatada = data.format(DateTimeFormatter.ofLocalizedDate(FormatStyle.FULL));
        System.out.println(dataFormatada);
        System.out.println("-------");
        // formatação de datas por local
        LocalDate dataPais = LocalDate.of(2021, 8, 23);
        Locale[] locales = { Locale.CANADA, Locale.US, Locale.of("pt", "BR"), Locale.UK };
        for (Locale locale : locales) {
            String novaDataPais = dataPais
                    .format(DateTimeFormatter.ofLocalizedDate(FormatStyle.SHORT).withLocale(locale));
            System.out.println(novaDataPais);
        }
        // 2021-08-23
        // 8/23/21
        // 23.8.2021
        // 23/08/2021
        System.out.println("-------");

        // Dados específicos
        LocalDate ExactDate = LocalDate.of(2021, 8, 23);

        System.out.println(ExactDate.getDayOfMonth()); // retorna o dia do mês -> 23
        System.out.println(ExactDate.getYear()); // retorna o ano -> 2023
        System.out.println(ExactDate.getMonth()); // retorna o elemento de enum java.time.Month -> AUGUST
        System.out.println(ExactDate.getMonthValue()); // retorna o número do mês entre 1-12 -> 8
    }
}