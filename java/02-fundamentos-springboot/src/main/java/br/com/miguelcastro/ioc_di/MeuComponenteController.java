package br.com.miguelcastro.ioc_di;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/component")
public class MeuComponenteController {

    @Autowired
    MeuComponente meuComponente;

    @GetMapping("/")
    public String chamandoComponente() {
        var meuComponent = new MeuComponente();
        var resultado = meuComponent.chamarMeuComponente();
        return resultado;
    }

}
