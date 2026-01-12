package br.com.miguelcastro.petland.cadastros.model.dto;

import br.com.miguelcastro.petland.cadastros.model.AnimalEspecie;
import java.time.LocalDate;
import lombok.Data;

@Data
public class AnimalRequest {
  private String nome;
  private LocalDate aniversário;
  private AnimalEspecie especie;
}
