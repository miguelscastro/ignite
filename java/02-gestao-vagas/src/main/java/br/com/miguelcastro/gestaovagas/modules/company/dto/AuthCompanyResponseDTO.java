package br.com.miguelcastro.gestaovagas.modules.company.dto;

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
  private String accessToken;
  private LocalDateTime expiresIn;
}
