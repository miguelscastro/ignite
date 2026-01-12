package br.com.miguelcastro.gestaovagas.modules.company.controllers;

import br.com.miguelcastro.gestaovagas.modules.company.dto.CreateJobDTO;
import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import br.com.miguelcastro.gestaovagas.modules.company.usecases.CreateJobUseCase;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.validation.Valid;
import java.util.UUID;
import lombok.RequiredArgsConstructor;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequiredArgsConstructor
@RequestMapping("/company/job")
public class JobController {

  private final CreateJobUseCase createJobUseCase;

  @PostMapping("/")
  @PreAuthorize("hasRole('COMPANY')")
  @Tag(name = "Vagas")
  @Operation(
      summary = "Cadastro de vagas",
      description = "Essa função é responsável por cadastrar as vagas dentro da empresa")
  @ApiResponse(
      responseCode = "200",
      content = {@Content(schema = @Schema(implementation = JobEntity.class))})
  @SecurityRequirement(name = "jwt_auth")
  public JobEntity create(
      @Valid @RequestBody CreateJobDTO createJobDTO, HttpServletRequest request) {
    var companyId = request.getAttribute("company_id");

    var jobEntity =
        JobEntity.builder()
            .benefits(createJobDTO.getBenefits())
            .companyId(UUID.fromString(companyId.toString()))
            .description(createJobDTO.getDescription())
            .level(createJobDTO.getLevel())
            .build();

    return this.createJobUseCase.execute(jobEntity);
  }
}
