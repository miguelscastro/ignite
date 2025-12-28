package br.com.miguelcastro.gestaovagas.exceptions;

public class CompanyNotFoundException extends RuntimeException {
  public CompanyNotFoundException() {
    super("Company not found");
  }
}
