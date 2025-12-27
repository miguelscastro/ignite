package br.com.miguelcastro.gestaovagas.modules.candidate.entities;

import br.com.miguelcastro.gestaovagas.modules.company.entities.JobEntity;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import java.time.LocalDateTime;
import java.util.UUID;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.CreationTimestamp;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity(name = "apply_jobs")
public class ApplyJobEntity {

  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private UUID id;

  @JoinColumn(name = "candidate_id", insertable = false, updatable = false)
  @ManyToOne
  private CandidateEntity candidateEntity;

  @JoinColumn(name = "job_id", insertable = false, updatable = false)
  @ManyToOne
  private JobEntity jobEntity;

  @Column(name = "candidate_id")
  private UUID candidateId;

  @Column(name = "job_id")
  private UUID jobId;

  @CreationTimestamp private LocalDateTime createdAt;
}
