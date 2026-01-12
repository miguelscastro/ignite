package br.com.miguelcastro.petland.atendimento.service;

import br.com.miguelcastro.petland.atendimento.model.dto.AtendimentoRequest;
import br.com.miguelcastro.petland.atendimento.model.entity.AtendimentoEntity;
import br.com.miguelcastro.petland.atendimento.repository.AtendimentoRepository;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AtendimentoService {

  @Autowired private AtendimentoRepository repository;

  public Integer gravar(AtendimentoRequest requisicao) {
    AtendimentoEntity entidade = new AtendimentoEntity();
    BeanUtils.copyProperties(requisicao, entidade);
    return repository.save(entidade).getId();
  }
}
