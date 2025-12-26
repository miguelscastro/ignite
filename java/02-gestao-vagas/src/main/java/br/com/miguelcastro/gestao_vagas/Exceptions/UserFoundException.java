package br.com.miguelcastro.gestao_vagas.Exceptions;

public class UserFoundException extends RuntimeException {
	public UserFoundException() {
		super("User already exists");
	}
}
