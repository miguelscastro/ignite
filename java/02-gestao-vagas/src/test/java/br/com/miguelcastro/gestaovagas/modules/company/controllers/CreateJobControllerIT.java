package br.com.miguelcastro.gestaovagas.modules.company.controllers;

import br.com.miguelcastro.gestaovagas.modules.company.dto.CreateJobDTO;
import br.com.miguelcastro.gestaovagas.utils.TestUtils;
import java.util.UUID;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.junit.jupiter.api.TestInstance.Lifecycle;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.http.MediaType;
import org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
@TestInstance(Lifecycle.PER_CLASS)
class CreateJobControllerIT {

  private MockMvc mvc;

  @Autowired private WebApplicationContext context;

  @BeforeAll
  void setup() {
    mvc =
        MockMvcBuilders.webAppContextSetup(context)
            .apply(SecurityMockMvcConfigurers.springSecurity())
            .build();
  }

  @Test
  void should_be_able_to_create_a_new_job() throws Exception {
    var createJobDTO =
        CreateJobDTO.builder()
            .benefits("BENEFITS_TEST")
            .description("DESCRIPTION_TEST")
            .level("LEVEL_TEST")
            .build();

    var result =
        mvc.perform(
                MockMvcRequestBuilders.post("/company/job/")
                    .contentType(MediaType.APPLICATION_JSON)
                    .content(TestUtils.objectToJson(createJobDTO))
                    .header(
                        "Authorization",
                        TestUtils.generateToken(UUID.randomUUID(), "JAVAGAS_@123#")))
            .andExpect(MockMvcResultMatchers.status().isOk());

    System.out.println(result);
  }
}
