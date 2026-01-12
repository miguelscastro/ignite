package br.com.miguelcastro.petland.cadastros.controller;

import br.com.miguelcastro.petland.cadastros.model.dto.AnimalRequest;
import br.com.miguelcastro.petland.cadastros.model.dto.AnimalResponse;
import br.com.miguelcastro.petland.cadastros.service.AnimalService;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/animais")
public class AnimalController {

  @Autowired private AnimalService service;

  @GetMapping
  public List<AnimalResponse> get() {
    return service.listar();
  }

  @PostMapping()
  public Integer gravar(@RequestBody AnimalRequest requisicao) {
    return service.gravar(requisicao);
  }

  @PutMapping("/{id}")
  public Integer alterar(@PathVariable("id") Integer id, @RequestBody AnimalRequest requisicao) {
    return service.alterar(id, requisicao);
  }

  @DeleteMapping("/{id}")
  public void deletar(@PathVariable("id") Integer id) {
    service.excluir(id);
  }
}
