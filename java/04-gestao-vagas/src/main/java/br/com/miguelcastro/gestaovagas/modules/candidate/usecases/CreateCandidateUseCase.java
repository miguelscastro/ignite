package br.com.miguelcastro.gestaovagas.modules.candidate.usecases;

import br.com.miguelcastro.gestaovagas.exceptions.UserFoundException;
import br.com.miguelcastro.gestaovagas.modules.candidate.entities.CandidateEntity;
import br.com.miguelcastro.gestaovagas.modules.candidate.repositories.CandidateRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class CreateCandidateUseCase {

  private final CandidateRepository candidateRepository;

  private final PasswordEncoder passwordEncoder;

  public CandidateEntity execute(CandidateEntity candidateEntity) {
    this.candidateRepository
        .findByUsernameOrEmail(candidateEntity.getUsername(), candidateEntity.getEmail())
        .ifPresent(
            user -> {
              throw new UserFoundException();
            });
    var password = passwordEncoder.encode(candidateEntity.getPassword());
    candidateEntity.setPassword(password);

    return this.candidateRepository.save(candidateEntity);
  }
}
