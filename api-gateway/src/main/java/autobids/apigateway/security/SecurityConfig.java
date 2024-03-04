package autobids.apigateway.security;

import autobids.apigateway.UriConstants;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.web.server.SecurityWebFilterChain;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsConfigurationSource;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;


@Configuration
@EnableWebFluxSecurity
public class SecurityConfig {

    @Autowired
    OAuth2LoginSuccessHandler oAuth2LoginSuccessHandler;

    @Autowired
    OAuth2LoginFailureHandler oAuth2LoginFailureHandler;

    @Bean
    public SecurityWebFilterChain securityWebFilterChain(ServerHttpSecurity http) {
        return http
                .cors(cors -> cors.configurationSource(corsConfigurationSource()))
                .authorizeExchange(exchanges -> exchanges
                        .pathMatchers(UriConstants.PROFILES_LOGIN).authenticated()
                        .pathMatchers(UriConstants.PROFILES_DELETE).authenticated()
                        .pathMatchers(UriConstants.PROFILES_EDIT).authenticated()
                        .pathMatchers(UriConstants.CARS_ADD).authenticated()
                        .pathMatchers(UriConstants.CARS_DELETE).authenticated()
                        .pathMatchers(UriConstants.CARS_DELETE_ALL).authenticated()
                        .pathMatchers(UriConstants.CARS_EDIT).authenticated()
                        .pathMatchers(UriConstants.CARS_SEARCH_ME).authenticated()
                        .anyExchange().permitAll()
                )
                .oauth2Login((oauth2Login) ->
                        oauth2Login
                                .authenticationSuccessHandler(oAuth2LoginSuccessHandler)
                                .authenticationFailureHandler(oAuth2LoginFailureHandler)
                )
                .csrf(ServerHttpSecurity.CsrfSpec::disable)
                .build();
    }

    @Bean
    CorsConfigurationSource corsConfigurationSource() {
        CorsConfiguration corsConfig = new CorsConfiguration();
        corsConfig.addAllowedOriginPattern("*");
        corsConfig.setMaxAge(3600L);
        corsConfig.setAllowCredentials(true);
        corsConfig.addAllowedMethod("*");
        corsConfig.addAllowedHeader("*");

        UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
        source.registerCorsConfiguration("/**", corsConfig);
        return source;
    }
}
