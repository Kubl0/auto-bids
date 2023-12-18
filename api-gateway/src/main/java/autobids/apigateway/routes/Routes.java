package autobids.apigateway.routes;

import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.security.core.context.ReactiveSecurityContextHolder;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.oauth2.core.user.OAuth2User;
import org.springframework.stereotype.Component;


@Component
public class Routes {


    @Bean
    public RouteLocator profileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/user/**", r -> r
                        .path("/profiles/user/**")
                        .uri("http://localhost:4100"))
                .build();
    }

    @Bean
    public RouteLocator securedProfileRoutes(RouteLocatorBuilder builder) {
        return builder.routes()
                .route("/profiles/me", r -> r
                        .path("/profiles/me")
                        .filters(f -> f.rewritePath("/profiles/me", "/profiles/user/example@example.com")
                                .filter((exchange, chain) ->
                                        ReactiveSecurityContextHolder.getContext()
                                                .map(SecurityContext::getAuthentication)
                                                .doOnNext(authentication -> {
                                                    OAuth2User user = (OAuth2User) authentication.getPrincipal();
                                                    System.out.println("Token: " + user.getAttribute("email"));
                                                })
                                                .then(chain.filter(exchange))
                                )
                        )
                        .uri("http://localhost:4100"))
                .build();
    }
}
