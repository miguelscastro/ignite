package br.com.miguelcastro.gestaovagas.modules.company.usecases;

import br.com.miguelcastro.gestaovagas.exceptions.CompanyNotFoundException;
import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import br.com.miguelcastro.gestaovagas.modules.company.repositories.CompanyRepository;
import br.com.miguelcastro.gestaovagas.modules.company.repositories.JobRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class CreateJobUseCase {
  private final JobRepository jobRepository;
  private final CompanyRepository companyRepository;

  public JobEntity execute(JobEntity jobEntity) {
    companyRepository
        .findById(jobEntity.getCompanyId())
        .orElseThrow(
            () -> {
              throw new CompanyNotFoundException();
            });

    return this.jobRepository.save(jobEntity);
  }
}
