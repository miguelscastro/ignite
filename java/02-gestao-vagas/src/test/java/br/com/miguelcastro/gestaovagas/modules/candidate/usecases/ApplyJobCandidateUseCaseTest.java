package br.com.miguelcastro.gestaovagas.modules.candidate.usecases;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.when;

import br.com.miguelcastro.gestaovagas.exceptions.JobNotFoundException;
import br.com.miguelcastro.gestaovagas.exceptions.UserNotFoundException;
import br.com.miguelcastro.gestaovagas.modules.candidate.entities.ApplyJobEntity;
import br.com.miguelcastro.gestaovagas.modules.candidate.entities.CandidateEntity;
import br.com.miguelcastro.gestaovagas.modules.candidate.repositories.ApplyJobRepository;
import br.com.miguelcastro.gestaovagas.modules.candidate.repositories.CandidateRepository;
import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import br.com.miguelcastro.gestaovagas.modules.company.repositories.JobRepository;
import java.util.Optional;
import java.util.UUID;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

@ExtendWith(MockitoExtension.class)
class ApplyJobCandidateUseCaseTest {

  @InjectMocks private ApplyJobCandidateUseCase applyJobCandidateUseCase;

  @Mock private CandidateRepository candidateRepository;
  @Mock private JobRepository jobRepository;
  @Mock private ApplyJobRepository applyJobRepository;

  @Test
  @DisplayName("Should not be able to apply to job when candidate is not found")
  void should_not_be_able_to_apply_job_when_candidate_not_found() {
    assertThrows(
        UserNotFoundException.class,
        () -> {
          applyJobCandidateUseCase.execute(null, null);
        });
  }

  @Test
  @DisplayName("Should not be able to apply to job when job is not found")
  void should_not_be_able_to_apply_job_when_job_not_found() {
    var idCandidate = UUID.randomUUID();
    var candidate = new CandidateEntity();
    candidate.setId(idCandidate);

    when(candidateRepository.findById(idCandidate)).thenReturn(Optional.of(candidate));

    assertThrows(
        JobNotFoundException.class,
        () -> {
          applyJobCandidateUseCase.execute(idCandidate, null);
        });
  }

  @Test
  void should_be_able_to_create_a_new_apply_job() {
    var idCandidate = UUID.randomUUID();
    var idJob = UUID.randomUUID();

    var applyJob = ApplyJobEntity.builder().candidateId(idCandidate).jobId(idJob).build();

    var applyJobCreated = ApplyJobEntity.builder().id(UUID.randomUUID()).build();

    when(candidateRepository.findById(idCandidate)).thenReturn(Optional.of(new CandidateEntity()));
    when(jobRepository.findById(idJob)).thenReturn(Optional.of(new JobEntity()));
    when(applyJobRepository.save(applyJob)).thenReturn(applyJobCreated);

    var result = applyJobCandidateUseCase.execute(idCandidate, idJob);
    assertNotNull(result.getId());
    assertThat(result).hasFieldOrProperty("id");
  }
}
