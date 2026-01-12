package br.com.miguelcastro.petland.cadastros.repository;

import br.com.miguelcastro.petland.cadastros.model.entity.AnimalEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AnimalRepository extends JpaRepository<AnimalEntity, Integer> {}
