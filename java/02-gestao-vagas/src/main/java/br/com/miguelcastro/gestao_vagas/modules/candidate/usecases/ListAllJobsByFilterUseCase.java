package br.com.miguelcastro.gestao_vagas.modules.candidate.usecases;

import br.com.miguelcastro.gestao_vagas.modules.company.entities.JobEntity;
import br.com.miguelcastro.gestao_vagas.modules.company.repositories.JobRepository;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ListAllJobsByFilterUseCase {

  private final JobRepository jobRepository;

  public List<JobEntity> execute(String filter) {
    return this.jobRepository.findByDescriptionContainingIgnoreCase(filter);
  }
}
