# Fundamentos Spring Boot

Examples and exercises demonstrating Spring Boot fundamentals used during the Ignite formation.

---

## 🔧 Prerequisites
- Java 11+ (or project-targeted JDK)
- Maven (or use the included wrapper `./mvnw`)

---

## 🚀 Quick start (local)

```bash
cd java/03-fundamentos-springboot
# using the included maven wrapper
./mvnw spring-boot:run
```

Or build and run the jar:

```bash
./mvnw clean package
java -jar target/*.jar
```

---

## 🛠️ Notes
- Check `src/main/resources/application.properties` for environment configuration.
- Use your IDE to run and debug Spring Boot apps; main class is annotated with `@SpringBootApplication`.

Happy developing! ✅