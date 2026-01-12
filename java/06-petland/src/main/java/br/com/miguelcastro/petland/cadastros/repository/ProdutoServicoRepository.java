package br.com.miguelcastro.petland.cadastros.repository;

import br.com.miguelcastro.petland.cadastros.model.ProdutoServico;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProdutoServicoRepository extends JpaRepository<ProdutoServico, Integer> {}
