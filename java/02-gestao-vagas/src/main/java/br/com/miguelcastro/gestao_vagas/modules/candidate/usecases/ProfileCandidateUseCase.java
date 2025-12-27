package br.com.miguelcastro.gestao_vagas.modules.candidate.usecases;

import br.com.miguelcastro.gestao_vagas.modules.candidate.CandidateRepository;
import br.com.miguelcastro.gestao_vagas.modules.candidate.dto.ProfileCandidateResponseDTO;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ProfileCandidateUseCase {

  private final CandidateRepository candidateRepository;

  public ProfileCandidateResponseDTO execute(UUID idCandidate) {
    var candidate =
        this.candidateRepository
            .findById(idCandidate)
            .orElseThrow(
                () -> {
                  throw new UsernameNotFoundException("User not found");
                });

    return ProfileCandidateResponseDTO.builder()
        .description(candidate.getDescription())
        .username(candidate.getUsername())
        .email(candidate.getEmail())
        .id(candidate.getId())
        .name(candidate.getName())
        .build();
  }
}
