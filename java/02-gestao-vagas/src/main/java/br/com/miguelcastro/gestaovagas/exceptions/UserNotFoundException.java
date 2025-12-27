package br.com.miguelcastro.gestaovagas.exceptions;

public class UserNotFoundException extends RuntimeException {
  public UserNotFoundException() {
    super("User not found");
  }
}
