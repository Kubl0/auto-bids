package autobids.apigateway.security;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.core.context.ReactiveSecurityContextHolder;
import org.springframework.security.web.server.SecurityWebFilterChain;

import static org.springframework.security.config.Customizer.withDefaults;

@EnableWebFluxSecurity
public class SecurityConfig {

    @Bean
    public SecurityWebFilterChain securityWebFilterChain(ServerHttpSecurity http) {
        return http
                .authorizeExchange(exchanges -> exchanges
                        .pathMatchers(HttpMethod.POST, "/profiles/user_profile/**").authenticated()
                        .pathMatchers(HttpMethod.PUT, "/profiles/user_profile/**").authenticated()
                        .pathMatchers("/profiles/me").authenticated()
                        .anyExchange().permitAll()
                )
                .oauth2Login(withDefaults())
                .csrf(ServerHttpSecurity.CsrfSpec::disable)
                .build();
    }
}
