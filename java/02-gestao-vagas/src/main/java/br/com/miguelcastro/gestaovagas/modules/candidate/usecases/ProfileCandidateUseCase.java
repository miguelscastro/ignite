package br.com.miguelcastro.gestaovagas.modules.candidate.usecases;

import br.com.miguelcastro.gestaovagas.modules.candidate.dto.ProfileCandidateResponseDTO;
import br.com.miguelcastro.gestaovagas.modules.candidate.repositories.CandidateRepository;
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
