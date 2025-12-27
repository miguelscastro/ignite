package br.com.miguelcastro.gestaovagas.security;

import lombok.RequiredArgsConstructor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.method.configuration.EnableMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

@Configuration
@EnableMethodSecurity
@RequiredArgsConstructor
public class SecurityConfig {

  private final SecurityFilter securityFilter;
  private final SecurityCandidateFilter securityCandidateFilter;

  private static final String[] SWAGGER_LIST = {
    "/swagger-ui/**", "/v3/api-docs/**", "/swagger-resources/**", "/swagger-ui.html"
  };

  @Bean
  SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
    return http.csrf(csrf -> csrf.disable())
        .authorizeHttpRequests(
            auth -> {
              auth.requestMatchers("/candidate/")
                  .permitAll()
                  .requestMatchers("/company/")
                  .permitAll()
                  .requestMatchers("/company/auth")
                  .permitAll()
                  .requestMatchers("/candidate/auth")
                  .permitAll()
                  .requestMatchers(SWAGGER_LIST)
                  .permitAll();
              auth.anyRequest().authenticated();
            })
        .addFilterBefore(securityCandidateFilter, BasicAuthenticationFilter.class)
        .addFilterBefore(securityFilter, BasicAuthenticationFilter.class)
        .build();
  }

  @Bean
  PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder();
  }
}
