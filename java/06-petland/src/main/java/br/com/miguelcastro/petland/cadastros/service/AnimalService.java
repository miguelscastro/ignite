package br.com.miguelcastro.petland.cadastros.service;

import br.com.miguelcastro.petland.cadastros.model.dto.AnimalRequest;
import br.com.miguelcastro.petland.cadastros.model.dto.AnimalResponse;
import br.com.miguelcastro.petland.cadastros.model.entity.AnimalEntity;
import br.com.miguelcastro.petland.cadastros.repository.AnimalRepository;
import java.util.ArrayList;
import java.util.List;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AnimalService {
  @Autowired private AnimalRepository repository;

  public Integer gravar(AnimalRequest requisicao) {
    AnimalEntity entidade = new AnimalEntity();
    BeanUtils.copyProperties(requisicao, entidade);
    return repository.save(entidade).getId();
  }

  public Integer alterar(Integer id, AnimalRequest requisicao) {
    AnimalEntity entidade = repository.findById(id).orElse(null);
    if (entidade != null) {
      BeanUtils.copyProperties(requisicao, entidade);
      return repository.save(entidade).getId();
    }
    return null;
  }

  public List<AnimalResponse> listar() {
    List<AnimalEntity> entities = repository.findAll();
    List<AnimalResponse> response = new ArrayList<>();
    for (AnimalEntity e : entities) {
      AnimalResponse res = new AnimalResponse();
      BeanUtils.copyProperties(e, res);
      response.add(res);
    }
    return response;
  }

  public void excluir(Integer id) {
    repository.deleteById(id);
  }
}
