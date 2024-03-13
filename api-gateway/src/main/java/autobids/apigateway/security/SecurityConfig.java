package autobids.apigateway.security;

import autobids.apigateway.UriConstants;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.authorization.AuthorizationDecision;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.authority.mapping.GrantedAuthoritiesMapper;
import org.springframework.security.oauth2.core.oidc.user.OidcUserAuthority;
import org.springframework.security.oauth2.core.user.OAuth2User;
import org.springframework.security.oauth2.core.user.OAuth2UserAuthority;
import org.springframework.security.web.server.SecurityWebFilterChain;
import org.springframework.security.web.server.authorization.AuthorizationContext;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsConfigurationSource;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;
import reactor.core.publisher.Mono;

import java.util.*;


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
                .oauth2Login(oauth2 -> oauth2
                        .authenticationSuccessHandler(oAuth2LoginSuccessHandler)
                        .authenticationFailureHandler(oAuth2LoginFailureHandler)
//                        .userInfoEndpoint(userInfo -> userInfo
//                                .userAuthoritiesMapper(userAuthoritiesMapper())
//                        )
                )
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
                        .pathMatchers(UriConstants.PROFILES_USER_ROUTES).access(this::isAdmin)
                        .anyExchange().permitAll()
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

    private Mono<AuthorizationDecision> isAdmin(Mono<Authentication> authentication, AuthorizationContext context) {
        return authentication.map(auth -> {
                ArrayList<String> roles = ((OAuth2User) auth.getPrincipal()).getAttribute("userRoles");
                return roles != null && roles.contains("ADMIN") ? new AuthorizationDecision(true) : new AuthorizationDecision(false);
        });
    }

    // Niby to jest cos co ma se z Principal zmapowac do Authorities jesli sie nie przesyla, ale i tak nic sie nie da bo nie da sie tego wrzucic wyzej jako mapper
    // Sam oauth nie przekazuje nic w authorities pomimo ze niby powinien, a ustawienia sa tak jak byc powinny niby na auth0
    @Bean
    GrantedAuthoritiesMapper userAuthoritiesMapper() {
        return (authorities) -> {
            Set<GrantedAuthority> mappedAuthorities = new HashSet<>();

            authorities.forEach(authority -> {
                if (authority instanceof OidcUserAuthority) {

                    if (authority.getAuthority().equals("ROLE_USER")) {
                        mappedAuthorities.add(new SimpleGrantedAuthority("ROLE_USER"));
                    } else if (authority.getAuthority().equals("ROLE_ADMIN")) {
                        mappedAuthorities.add(new SimpleGrantedAuthority("ROLE_ADMIN"));
                    }

                } else if (authority instanceof OAuth2UserAuthority oauth2UserAuthority) {
                    // Handling OAuth2UserAuthority
                    Map<String, Object> userAttributes = oauth2UserAuthority.getAttributes();
                    if (userAttributes.containsKey("userRole")) {
                        Object rolesObj = userAttributes.get("userRole");
                        if (rolesObj instanceof List) {
                            List<String> roles = (List<String>) rolesObj;
                            roles.forEach(role -> mappedAuthorities.add(new SimpleGrantedAuthority("ROLE_" + role.toUpperCase())));
                        } else if (rolesObj instanceof String) {
                            // Corrected syntax for handling String
                            mappedAuthorities.add(new SimpleGrantedAuthority("ROLE_" + ((String) rolesObj).toUpperCase()));
                        }
                    }


                }
            });
            return mappedAuthorities;
        };
    }

}
