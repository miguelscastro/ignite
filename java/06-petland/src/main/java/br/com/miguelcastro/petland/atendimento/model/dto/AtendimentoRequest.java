package br.com.miguelcastro.petland.atendimento.model.dto;

import br.com.miguelcastro.petland.atendimento.model.AtendimentoStatus;
import br.com.miguelcastro.petland.atendimento.model.AtendimentoTipo;
import java.time.LocalDate;
import lombok.Data;

@Data
public class AtendimentoRequest {
  private String descricao;
  private LocalDate data;
  private Double valor;
  private boolean emergencia;
  private AtendimentoTipo tipo;
  private AtendimentoStatus status;
  private Integer animal;
  private Integer produtoServico;
  private Integer cadastro;
}
