package br.com.miguelcastro.gestaovagas.modules.candidate.usecases;

import br.com.miguelcastro.gestaovagas.modules.candidate.dto.AuthCandidateRequestDTO;
import br.com.miguelcastro.gestaovagas.modules.candidate.dto.AuthCandidateResponseDTO;
import br.com.miguelcastro.gestaovagas.modules.candidate.repositories.CandidateRepository;
import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import java.time.Duration;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.Arrays;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthCandidateUseCase {

  @Value("${security.token.secret.candidate}")
  public String secretKey;

  private final CandidateRepository candidateRepository;

  private final PasswordEncoder passwordEncoder;

  public AuthCandidateResponseDTO execute(AuthCandidateRequestDTO authCandidateRequestDTO) {
    var candidate =
        this.candidateRepository
            .findByUsername(authCandidateRequestDTO.username())
            .orElseThrow(
                () -> {
                  throw new UsernameNotFoundException("User or password is invalid");
                });
    var password =
        this.passwordEncoder.matches(authCandidateRequestDTO.password(), candidate.getPassword());

    if (!password) {
      throw new BadCredentialsException("User or password is invalid");
    }

    Algorithm algorithm = Algorithm.HMAC256(secretKey);

    var expiresIn = Instant.now().plus(Duration.ofMinutes(10));
    var token =
        JWT.create()
            .withIssuer("javagas")
            .withSubject(candidate.getId().toString())
            .withClaim("roles", Arrays.asList("CANDIDATE"))
            .withExpiresAt(expiresIn)
            .sign(algorithm);

    var formattedExpireDate = LocalDateTime.ofInstant(expiresIn, ZoneId.systemDefault());

    return AuthCandidateResponseDTO.builder()
        .accessToken(token)
        .expiresIn(formattedExpireDate)
        .build();
  }
}
