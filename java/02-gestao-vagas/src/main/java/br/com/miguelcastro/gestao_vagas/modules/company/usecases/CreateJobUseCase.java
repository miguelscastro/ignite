package br.com.miguelcastro.gestao_vagas.modules.company.usecases;

import br.com.miguelcastro.gestao_vagas.modules.company.entities.JobEntity;
import br.com.miguelcastro.gestao_vagas.modules.company.repositories.JobRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class CreateJobUseCase {
  private final JobRepository jobRepository;

  public JobEntity execute(JobEntity jobEntity) {
    return this.jobRepository.save(jobEntity);
  }
}
