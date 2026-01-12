package br.com.miguelcastro.petland.atendimento.repository;

import br.com.miguelcastro.petland.atendimento.model.entity.AtendimentoEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AtendimentoRepository extends JpaRepository<AtendimentoEntity, Integer> {}
