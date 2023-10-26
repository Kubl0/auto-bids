FROM eclipse-temurin:17-jdk-alpine as builder
WORKDIR /opt/app
COPY .mvn/ .mvn
COPY mvnw pom.xml ./
ARG SERVER_PORT=4000
RUN ./mvnw dependency:go-offline
COPY ./src ./src
RUN ./mvnw clean install

FROM eclipse-temurin:17-jre-alpine
WORKDIR /opt/app
EXPOSE 4000
COPY --from=builder /opt/app/target/*.jar /opt/app/app.jar
USER myuser
ENTRYPOINT ["java", "-jar", "/opt/app/app.jar"]
