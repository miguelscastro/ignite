package br.com.miguelcastro.gestaovagas.modules.candidate.usecases;

import br.com.miguelcastro.gestaovagas.exceptions.JobNotFoundException;
import br.com.miguelcastro.gestaovagas.exceptions.UserNotFoundException;
import br.com.miguelcastro.gestaovagas.modules.candidate.CandidateRepository;
import br.com.miguelcastro.gestaovagas.modules.company.repositories.JobRepository;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ApplyJobCandidateUseCase {

  private final CandidateRepository candidateRepository;
  private final JobRepository jobRepository;

  public void execute(UUID idCandidate, UUID idJob) {
    this.candidateRepository
        .findById(idCandidate)
        .orElseThrow(
            () -> {
              throw new UserNotFoundException();
            });

    this.jobRepository
        .findById(idJob)
        .orElseThrow(
            () -> {
              throw new JobNotFoundException();
            });
  }
}
