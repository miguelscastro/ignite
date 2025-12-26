package br.com.miguelcastro.gestao_vagas.modules.company.dto;

import java.time.LocalDateTime;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor

public class AuthCompanyResponseDTO {
	private String access_token;
	private LocalDateTime expires_in;
}
