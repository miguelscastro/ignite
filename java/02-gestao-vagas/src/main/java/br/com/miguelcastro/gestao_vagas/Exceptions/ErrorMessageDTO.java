package br.com.miguelcastro.gestao_vagas.Exceptions;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ErrorMessageDTO {
	private String message;
	private String field;
}
