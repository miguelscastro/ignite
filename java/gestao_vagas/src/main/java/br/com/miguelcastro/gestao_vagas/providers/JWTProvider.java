package br.com.miguelcastro.gestao_vagas.providers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.exceptions.JWTVerificationException;

import br.com.miguelcastro.gestao_vagas.modules.company.useCases.AuthCompanyUseCase;

@Service
public class JWTProvider {

    @Autowired
    private AuthCompanyUseCase authCompanyUseCase;

    public String validateToken(String token) {
        token = token.replace("Bearer ", "");

        Algorithm algorithm = Algorithm.HMAC256(authCompanyUseCase.secretKey);

        try {
            var subject = JWT.require(algorithm)
                    .build()
                    .verify(token)
                    .getSubject();

            return subject;
        } catch (JWTVerificationException e) {
            e.printStackTrace();
            return "";
        }
    }
}
