package br.com.miguelcastro.gestaovagas.modules.candidate.controllers;

import br.com.miguelcastro.gestaovagas.modules.candidate.dto.ProfileCandidateResponseDTO;
import br.com.miguelcastro.gestaovagas.modules.candidate.entities.CandidateEntity;
import br.com.miguelcastro.gestaovagas.modules.candidate.usecases.CreateCandidateUseCase;
import br.com.miguelcastro.gestaovagas.modules.candidate.usecases.ListAllJobsByFilterUseCase;
import br.com.miguelcastro.gestaovagas.modules.candidate.usecases.ProfileCandidateUseCase;
import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import java.util.List;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
@Tag(name = "Candidato", description = "Informações do candidato")
@RequestMapping("/candidate")
public class CandidateController {

  private final CreateCandidateUseCase createCandidateUseCase;
  private final ProfileCandidateUseCase profileCandidateUseCase;
  private final ListAllJobsByFilterUseCase listAllJobsByFilterUseCase;

  @PostMapping("/")
  @Operation(
      summary = "Cadastro de candidato",
      description = "Essa função é responsável por cadastrar um candidato")
  @ApiResponse(
      responseCode = "200",
      content = @Content(schema = @Schema(implementation = CandidateEntity.class)))
  @ApiResponse(responseCode = "400", description = "User not found")
  public ResponseEntity<Object> create(@Valid @RequestBody CandidateEntity candidateEntity) {
    try {
      var result = this.createCandidateUseCase.execute(candidateEntity);
      return ResponseEntity.ok().body(result);
    } catch (Exception e) {
      return ResponseEntity.badRequest().body(e.getMessage());
    }
  }

  @PreAuthorize("hasRole('CANDIDATE')")
  @GetMapping("/")
  @Operation(
      summary = "Perfil do candidato",
      description = "Essa função é responsável por buscar as informações do candidato")
  @ApiResponse(
      responseCode = "200",
      content = {@Content(schema = @Schema(implementation = ProfileCandidateResponseDTO.class))})
  @ApiResponse(responseCode = "400", description = "User not found")
  @SecurityRequirement(name = "jwt_auth")
  public ResponseEntity<Object> get(HttpServletRequest request) {
    var candidateId = request.getAttribute("candidate_id");
    try {
      var profile = this.profileCandidateUseCase.execute(UUID.fromString(candidateId.toString()));
      return ResponseEntity.ok().body(profile);

    } catch (Exception e) {
      return ResponseEntity.badRequest().body(e.getMessage());
    }
  }

  @PreAuthorize("hasRole('CANDIDATE')")
  @GetMapping("/job")
  @Operation(
      summary = "Listagem de vagas disponível para o candidato",
      description =
          "Essa função é responsável por listar todas as vagas disponíveis baseado no filtro")
  @ApiResponse(
      responseCode = "200",
      content = {
        @Content(array = @ArraySchema(schema = @Schema(implementation = JobEntity.class)))
      })
  @SecurityRequirement(name = "jwt_auth")
  public List<JobEntity> findJobByFilter(@RequestParam String filter) {
    return this.listAllJobsByFilterUseCase.execute(filter);
  }
}
