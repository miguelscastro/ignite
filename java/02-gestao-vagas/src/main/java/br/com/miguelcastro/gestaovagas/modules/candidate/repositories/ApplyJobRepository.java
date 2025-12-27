package br.com.miguelcastro.gestaovagas.modules.candidate.repositories;

import br.com.miguelcastro.gestaovagas.modules.candidate.entities.ApplyJobEntity;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ApplyJobRepository extends JpaRepository<ApplyJobEntity, UUID> {}
