package br.com.miguelcastro.gestaovagas.modules.company.usecases;

import br.com.miguelcastro.gestaovagas.exceptions.UserFoundException;
import br.com.miguelcastro.gestaovagas.modules.company.entities.CompanyEntity;
import br.com.miguelcastro.gestaovagas.modules.company.repositories.CompanyRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class CreateCompanyUseCase {

  private final CompanyRepository companyRepository;
  private final PasswordEncoder passwordEncoder;

  public CompanyEntity execute(CompanyEntity companyEntity) {
    this.companyRepository
        .findByUsernameOrEmail(companyEntity.getUsername(), companyEntity.getEmail())
        .ifPresent(
            user -> {
              throw new UserFoundException();
            });

    var password = passwordEncoder.encode(companyEntity.getPassword());
    companyEntity.setPassword(password);

    return this.companyRepository.save(companyEntity);
  }
}
