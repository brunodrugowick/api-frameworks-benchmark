# Use official Maven image for building
FROM maven:3.9-eclipse-temurin-17 AS build

WORKDIR /app

# Copy Maven configuration first (for better caching)
COPY pom.xml ./
COPY src ./src

# Build the application
RUN mvn clean package -DskipTests

# Use a lightweight JDK runtime for the final image
FROM eclipse-temurin:17-jre

WORKDIR /app

# Copy the built JAR from the previous stage
COPY --from=build /app/target/*.jar app.jar

# Expose application port
# NOT USING BECAUSE I'M SETTING NETWORK TO HOST!
EXPOSE 8080

ENV SPRING_DATASOURCE_URL="jdbc:postgresql://localhost:1111/postgres"
ENV SERVER_PORT=9095
ENV SPRING_DATASOURCE_USERNAME=postgres
ENV SPRING_DATASOURCE_PASSWORD=postgres
ENV SPRING_DATASOURCE_DRIVERCLASSNAME=org.postgresql.Driver
ENV SPRING_JPA_HIBERNATE_DDLAUTO=create
ENV SPRING_JPA_PROPERTIES_HIBERNATE_DIALECT=org.hibernate.dialect.PostgreSQLDialect

# Run the application
ENTRYPOINT ["java", "-jar", "app.jar"]
