package br.com.miguelcastro.petland.cadastros.model;

import jakarta.persistence.Embeddable;
import lombok.Data;

@Embeddable
@Data
public class Perfil {

  private boolean cliente;
  private boolean fornecedor;
  private boolean prestador;
}
