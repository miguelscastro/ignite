package br.com.miguelcastro.petland.cadastros.model.dto;

import br.com.miguelcastro.petland.cadastros.model.Endereco;
import br.com.miguelcastro.petland.cadastros.model.Perfil;
import lombok.Data;

@Data
public class CadastroRequest {
  private String nome;
  private Perfil perfil;
  private Endereco endereco;
}
