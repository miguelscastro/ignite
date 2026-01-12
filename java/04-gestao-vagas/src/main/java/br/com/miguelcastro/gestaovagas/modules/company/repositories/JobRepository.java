package br.com.miguelcastro.gestaovagas.modules.company.repositories;

import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import java.util.List;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;

public interface JobRepository extends JpaRepository<JobEntity, UUID> {
  List<JobEntity> findByDescriptionContainingIgnoreCase(String filter);
}
