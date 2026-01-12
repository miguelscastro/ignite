package br.com.miguelcastro.petland.cadastros.repository;

import br.com.miguelcastro.petland.cadastros.model.entity.CadastroEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CadastroRepository extends JpaRepository<CadastroEntity, Integer> {}
