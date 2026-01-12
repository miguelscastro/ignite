package br.com.miguelcastro.gestaovagas.modules.candidate.dto;

import java.time.LocalDateTime;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class AuthCandidateResponseDTO {
  private String accessToken;
  private LocalDateTime expiresIn;
}
