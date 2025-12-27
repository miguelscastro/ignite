package br.com.miguelcastro.gestao_vagas.modules.company.usecases;

import br.com.miguelcastro.gestao_vagas.modules.company.dto.AuthCompanyRequestDTO;
import br.com.miguelcastro.gestao_vagas.modules.company.dto.AuthCompanyResponseDTO;
import br.com.miguelcastro.gestao_vagas.modules.company.repositories.CompanyRepository;
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
public class AuthCompanyUseCase {

  @Value("${security.token.secret}")
  public String secretKey;

  private final PasswordEncoder passwordEncoder;

  private final CompanyRepository companyRepository;

  public AuthCompanyResponseDTO execute(AuthCompanyRequestDTO authCompanyRequestDTO) {
    var company =
        this.companyRepository
            .findByUsername(authCompanyRequestDTO.username())
            .orElseThrow(
                () -> {
                  throw new UsernameNotFoundException("User or password is invalid");
                });
    var passwordMatches =
        this.passwordEncoder.matches(authCompanyRequestDTO.password(), company.getPassword());

    if (!passwordMatches) {
      throw new BadCredentialsException("User or password is invalid");
    }

    Algorithm algorithm = Algorithm.HMAC256(secretKey);
    var expiresIn = Instant.now().plus(Duration.ofHours(2));
    var token =
        JWT.create()
            .withIssuer("javagas")
            .withSubject(company.getId().toString())
            .withClaim("roles", Arrays.asList("COMPANY"))
            .withExpiresAt(expiresIn)
            .sign(algorithm);

    var formattedExpireDate = LocalDateTime.ofInstant(expiresIn, ZoneId.systemDefault());

    return AuthCompanyResponseDTO.builder()
        .accessToken(token)
        .expiresIn(formattedExpireDate)
        .build();
  }
}
